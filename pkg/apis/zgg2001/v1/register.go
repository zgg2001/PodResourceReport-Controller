package v1

import (
	"github.com/zgg2001/PodResourceReport-Controller/pkg/apis/zgg2001"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: zgg2001.GroupName, Version: "v1"}

var (
	// SchemeBuilder localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	// AddToScheme localSchemeBuilder AddToScheme
	Scheme      *runtime.Scheme
	AddToScheme = localSchemeBuilder.AddToScheme
)

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// addKnownTypes adds our types to the API scheme by registering
func addKnownTypes(scheme *runtime.Scheme) error {
	Scheme = scheme
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&NamespaceResourceReport{},
		&NamespaceResourceReportList{},
	)
	// register the type in the scheme
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
