// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ClusterVersionOperatorStatusApplyConfiguration represents a declarative configuration of the ClusterVersionOperatorStatus type for use
// with apply.
type ClusterVersionOperatorStatusApplyConfiguration struct {
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// ClusterVersionOperatorStatusApplyConfiguration constructs a declarative configuration of the ClusterVersionOperatorStatus type for use with
// apply.
func ClusterVersionOperatorStatus() *ClusterVersionOperatorStatusApplyConfiguration {
	return &ClusterVersionOperatorStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ClusterVersionOperatorStatusApplyConfiguration) WithObservedGeneration(value int64) *ClusterVersionOperatorStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}
