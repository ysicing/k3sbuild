// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=k3s.cattle.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddonList is a list of Addon resources
type AddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Addon `json:"items"`
}

func NewAddon(namespace, name string, obj Addon) *Addon {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Addon").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}