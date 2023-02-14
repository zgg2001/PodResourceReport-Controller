/*
Copyright The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	zgg2001v1 "github.com/zgg2001/PodResourceReport-Controller/pkg/apis/zgg2001/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNamespaceResourceReports implements NamespaceResourceReportInterface
type FakeNamespaceResourceReports struct {
	Fake *FakeZgg2001V1
	ns   string
}

var namespaceresourcereportsResource = schema.GroupVersionResource{Group: "zgg2001.github.com", Version: "v1", Resource: "namespaceresourcereports"}

var namespaceresourcereportsKind = schema.GroupVersionKind{Group: "zgg2001.github.com", Version: "v1", Kind: "NamespaceResourceReport"}

// Get takes name of the namespaceResourceReport, and returns the corresponding namespaceResourceReport object, and an error if there is any.
func (c *FakeNamespaceResourceReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *zgg2001v1.NamespaceResourceReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(namespaceresourcereportsResource, c.ns, name), &zgg2001v1.NamespaceResourceReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zgg2001v1.NamespaceResourceReport), err
}

// List takes label and field selectors, and returns the list of NamespaceResourceReports that match those selectors.
func (c *FakeNamespaceResourceReports) List(ctx context.Context, opts v1.ListOptions) (result *zgg2001v1.NamespaceResourceReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(namespaceresourcereportsResource, namespaceresourcereportsKind, c.ns, opts), &zgg2001v1.NamespaceResourceReportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &zgg2001v1.NamespaceResourceReportList{ListMeta: obj.(*zgg2001v1.NamespaceResourceReportList).ListMeta}
	for _, item := range obj.(*zgg2001v1.NamespaceResourceReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaceResourceReports.
func (c *FakeNamespaceResourceReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(namespaceresourcereportsResource, c.ns, opts))

}

// Create takes the representation of a namespaceResourceReport and creates it.  Returns the server's representation of the namespaceResourceReport, and an error, if there is any.
func (c *FakeNamespaceResourceReports) Create(ctx context.Context, namespaceResourceReport *zgg2001v1.NamespaceResourceReport, opts v1.CreateOptions) (result *zgg2001v1.NamespaceResourceReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(namespaceresourcereportsResource, c.ns, namespaceResourceReport), &zgg2001v1.NamespaceResourceReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zgg2001v1.NamespaceResourceReport), err
}

// Update takes the representation of a namespaceResourceReport and updates it. Returns the server's representation of the namespaceResourceReport, and an error, if there is any.
func (c *FakeNamespaceResourceReports) Update(ctx context.Context, namespaceResourceReport *zgg2001v1.NamespaceResourceReport, opts v1.UpdateOptions) (result *zgg2001v1.NamespaceResourceReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(namespaceresourcereportsResource, c.ns, namespaceResourceReport), &zgg2001v1.NamespaceResourceReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zgg2001v1.NamespaceResourceReport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNamespaceResourceReports) UpdateStatus(ctx context.Context, namespaceResourceReport *zgg2001v1.NamespaceResourceReport, opts v1.UpdateOptions) (*zgg2001v1.NamespaceResourceReport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(namespaceresourcereportsResource, "status", c.ns, namespaceResourceReport), &zgg2001v1.NamespaceResourceReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zgg2001v1.NamespaceResourceReport), err
}

// Delete takes name of the namespaceResourceReport and deletes it. Returns an error if one occurs.
func (c *FakeNamespaceResourceReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(namespaceresourcereportsResource, c.ns, name, opts), &zgg2001v1.NamespaceResourceReport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNamespaceResourceReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(namespaceresourcereportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &zgg2001v1.NamespaceResourceReportList{})
	return err
}

// Patch applies the patch and returns the patched namespaceResourceReport.
func (c *FakeNamespaceResourceReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *zgg2001v1.NamespaceResourceReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(namespaceresourcereportsResource, c.ns, name, pt, data, subresources...), &zgg2001v1.NamespaceResourceReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*zgg2001v1.NamespaceResourceReport), err
}