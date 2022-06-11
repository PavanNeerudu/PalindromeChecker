/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PalindromeCheckerSpec defines the desired state of PalindromeChecker
type PalindromeCheckerSpec struct {
	Input string `json:"input,omitempty"`
}

// PalindromeCheckerStatus defines the observed state of PalindromeChecker
type PalindromeCheckerStatus struct {
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:printcolumn:name="Initial_Input",type=string,JSONPath=`.spec.input`
//+kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// PalindromeChecker is the Schema for the palindromecheckers API
type PalindromeChecker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PalindromeCheckerSpec   `json:"spec,omitempty"`
	Status PalindromeCheckerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// PalindromeCheckerList contains a list of PalindromeChecker
type PalindromeCheckerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PalindromeChecker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PalindromeChecker{}, &PalindromeCheckerList{})
}
