// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	buildv1 "github.com/openshift/api/build/v1"
)

// SourceRevisionApplyConfiguration represents a declarative configuration of the SourceRevision type for use
// with apply.
type SourceRevisionApplyConfiguration struct {
	Type *buildv1.BuildSourceType             `json:"type,omitempty"`
	Git  *GitSourceRevisionApplyConfiguration `json:"git,omitempty"`
}

// SourceRevisionApplyConfiguration constructs a declarative configuration of the SourceRevision type for use with
// apply.
func SourceRevision() *SourceRevisionApplyConfiguration {
	return &SourceRevisionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *SourceRevisionApplyConfiguration) WithType(value buildv1.BuildSourceType) *SourceRevisionApplyConfiguration {
	b.Type = &value
	return b
}

// WithGit sets the Git field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Git field is set to the value of the last call.
func (b *SourceRevisionApplyConfiguration) WithGit(value *GitSourceRevisionApplyConfiguration) *SourceRevisionApplyConfiguration {
	b.Git = value
	return b
}
