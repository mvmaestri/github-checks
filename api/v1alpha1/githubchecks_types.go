/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GithubChecksSpec defines the desired state of GithubChecks
type GithubChecksSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of GithubChecks. Edit GithubChecks_types.go to remove/update
	Foo  string `json:"foo,omitempty"`
	Size int32  `json:"size"`
}

// GithubChecksStatus defines the observed state of GithubChecks
type GithubChecksStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GithubChecks is the Schema for the githubchecks API
// +kubebuilder:subresource:status
type GithubChecks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GithubChecksSpec   `json:"spec,omitempty"`
	Status GithubChecksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GithubChecksList contains a list of GithubChecks
type GithubChecksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GithubChecks `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GithubChecks{}, &GithubChecksList{})
}
