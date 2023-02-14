package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=namespaceresourcereports,scope=Namespaced
type NamespaceResourceReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// Spec defines the behavior of a NamespaceResourceReport
	Spec NamespaceResourceReportSpec `json:"spec"`

	// Most recently observed status of the NamespaceResourceReport~
	Status NamespaceResourceReportStatus `json:"status"`
}

// NamespaceResourceReportList is NamespaceResourceReport list
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NamespaceResourceReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NamespaceResourceReport `json:"items"`
}

type NamespaceResourceReportSpec struct {
	Namespace string `json:"namespace"`
}

type NamespaceResourceReportStatus struct {
	CpuUsed string `json:"cpuused,omitempty"`
	MemUsed string `json:"memused,omitempty"`
}
