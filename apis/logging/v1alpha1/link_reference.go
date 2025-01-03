// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &LoggingLinkRef{}

// LoggingLinkRef defines the resource reference to LoggingLink, which "External" field
// holds the GCP identifier for the KRM object.
type LoggingLinkRef struct {
	// A reference to an externally managed LoggingLink resource.
	// Should be in the format "projects/<projectID>/locations/<location>/buckets/<bucketID>/links/<linkID>".
	External string `json:"external,omitempty"`

	// The name of a LoggingLink resource.
	Name string `json:"name,omitempty"`

	// The namespace of a LoggingLink resource.
	Namespace string `json:"namespace,omitempty"`

	parent *LoggingLinkParent
}

// NormalizedExternal provision the "External" value for other resource that depends on LoggingLink.
// If the "External" is given in the other resource's spec.LoggingLinkRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual LoggingLink object from the cluster.
func (r *LoggingLinkRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", LoggingLinkGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseLoggingLinkExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(LoggingLinkGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", LoggingLinkGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

// New builds a LoggingLinkRef from the Config Connector LoggingLink object.
func NewLoggingLinkRef(ctx context.Context, reader client.Reader, obj *LoggingLink) (*LoggingLinkRef, error) {
	id := &LoggingLinkRef{}

	// Get Parent
	loggingBucketRef, err := refsv1beta1.ResolveLoggingLogBucketRef(ctx, reader, obj, obj.Spec.LoggingLogBucketRef)
	if err != nil {
		return nil, err
	}
	projectID := loggingBucketRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := loggingBucketRef.Location
	bucketID := loggingBucketRef.LoggingLogBucketID
	id.parent = &LoggingLinkParent{ProjectID: projectID, Location: location, LogBucket: bucketID}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = asLoggingLinkExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseLoggingLinkExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("loggingBucketRef.projectID changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualParent.Location != location {
		return nil, fmt.Errorf("loggingBucketRef.location changed, expect %s, got %s", actualParent.Location, location)
	}
	if actualParent.LogBucket != bucketID {
		return nil, fmt.Errorf("loggingBucketRef.LoggingLogBucketID changed, expect %s, got %s", actualParent.Location, bucketID)
	}
	if actualResourceID != resourceID {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			resourceID, actualResourceID)
	}
	id.External = externalRef
	id.parent = &LoggingLinkParent{ProjectID: projectID, Location: location, LogBucket: bucketID}
	return id, nil
}

func (r *LoggingLinkRef) Parent() (*LoggingLinkParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := parseLoggingLinkExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("LoggingLinkRef not initialized from `NewLoggingLinkRef` or `NormalizedExternal`")
}

type LoggingLinkParent struct {
	ProjectID string
	Location  string
	LogBucket string
}

func (p *LoggingLinkParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/buckets/" + p.LogBucket
}

func asLoggingLinkExternal(parent *LoggingLinkParent, resourceID string) (external string) {
	return parent.String() + "/links/" + resourceID
}

func parseLoggingLinkExternal(external string) (parent *LoggingLinkParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "bucket" || tokens[6] != "link" {
		return nil, "", fmt.Errorf("format of LoggingLink external=%q was not known (use projects/<projectId>/locations/<location>/buckets/<bucketID>/links/<linkID>)", external)
	}
	parent = &LoggingLinkParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		LogBucket: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
