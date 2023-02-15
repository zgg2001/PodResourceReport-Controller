/*
Copyright 2017 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"reflect"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	v1 "github.com/zgg2001/PodResourceReport-Controller/pkg/apis/zgg2001/v1"
	clientset "github.com/zgg2001/PodResourceReport-Controller/pkg/client/clientset/versioned"
	samplescheme "github.com/zgg2001/PodResourceReport-Controller/pkg/client/clientset/versioned/scheme"
	informers "github.com/zgg2001/PodResourceReport-Controller/pkg/client/informers/externalversions/zgg2001/v1"
	listers "github.com/zgg2001/PodResourceReport-Controller/pkg/client/listers/zgg2001/v1"
)

const controllerAgentName = "PodResourceReport-Controller"

const (
	SuccessSynced         = "Synced"
	ErrResourceExists     = "ErrResourceExists"
	MessageResourceExists = "Resource %q already exists and is not managed by NamespaceResourceReport"
	MessageResourceSynced = "NamespaceResourceReport synced successfully"
)

// Controller is the controller implementation for NamespaceResourceReport resources
type Controller struct {
	kubeclientset   kubernetes.Interface
	sampleclientset clientset.Interface

	nsrrLister listers.NamespaceResourceReportLister
	nsrrSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface
	recorder  record.EventRecorder
}

func NewController(
	kubeclientset kubernetes.Interface,
	sampleclientset clientset.Interface,
	deploymentInformer appsinformers.DeploymentInformer,
	namespaceResourceReportInformer informers.NamespaceResourceReportInformer) *Controller {

	utilruntime.Must(samplescheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:   kubeclientset,
		sampleclientset: sampleclientset,
		nsrrLister:      namespaceResourceReportInformer.Lister(),
		nsrrSynced:      namespaceResourceReportInformer.Informer().HasSynced,
		workqueue:       workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "NamespaceResourceReports"),
		recorder:        recorder,
	}

	// Set up an event handler for when NamespaceResourceReport resources change
	klog.Info("Setting up event handlers")
	namespaceResourceReportInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueNamespaceResourceReport,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueNamespaceResourceReport(new)
		},
	})

	return controller
}

func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {

	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()
	klog.Info("Starting NamespaceResourceReport controller")

	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.nsrrSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {

	obj, shutdown := c.workqueue.Get()
	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		if key, ok = obj.(string); !ok {
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		if err := c.syncHandler(key); err != nil {
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

// reconcile
func (c *Controller) syncHandler(key string) error {

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	nrr, err := c.nsrrLister.NamespaceResourceReports(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("NamespaceResourceReport '%s' in work queue no longer exists", key))
			return nil
		}
		return err
	}
	nrrCopy := nrr.DeepCopy()

	nrrCopy.Status.CpuUsed = "123"
	nrrCopy.Status.MemUsed = "456"
	if !reflect.DeepEqual(nrrCopy.Status, nrr.Status) {
		err = c.updateNamespaceResourceReportStatus(nrrCopy)
		if err != nil {
			return err
		}
	}
	c.recorder.Event(nrr, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateNamespaceResourceReportStatus(nrr *v1.NamespaceResourceReport) error {
	nrrCopy := nrr.DeepCopy()
	var updateErr error
	_ = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, updateErr = c.sampleclientset.Zgg2001V1().NamespaceResourceReports(nrrCopy.Namespace).UpdateStatus(context.TODO(), nrrCopy, metav1.UpdateOptions{})
		return updateErr
	})
	return updateErr
}

func (c *Controller) enqueueNamespaceResourceReport(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}
