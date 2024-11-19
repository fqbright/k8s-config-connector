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

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func LoggingLinkSpec_LoggingLogBucketRef_FromProto(mapCtx *direct.MapContext, in string) *refs.LoggingLogBucketRef {
	if in == "" {
		return nil
	}
	return &refs.LoggingLogBucketRef{
		External: in,
	}
}

func LoggingLinkSpec_LoggingLogBucketRef_ToProto(mapCtx *direct.MapContext, in *refs.LoggingLogBucketRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}


func BigQueryDataset_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDataset) *krm.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataset{}
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	return out
}
func BigQueryDataset_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataset) *pb.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDataset{}
	out.DatasetId = direct.ValueOf(in.DatasetID)
	return out
}
func LoggingLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Link) *krm.LoggingLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLinkSpec{}
	out.ResourceID = direct.LazyPtr(in.GetName()) // this needs to be link ID, which is just the unique name of the link (strip the prefix)
	out.Description = direct.LazyPtr(in.GetDescription())
	// Build from proto and to proto for Log Bucket Ref
	out.LoggingLogBucketRef =  LoggingLinkSpec_LoggingLogBucketRef_FromProto(mapCtx, in.GetName) //but strip the suffix of links/linkID
	return out
}
func LoggingLinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLinkSpec) *pb.Link {
	if in == nil {
		return nil
	}
	out := &pb.Link{}
	// out.Name = direct.ValueOf(in.Name) - keep in mind, this is set, but in the caller by the controller
	out.Description = direct.ValueOf(in.Description)

	return out
}
func LoggingLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Link) *krm.LoggingLinkObservedState {
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// This is the first lifecycle state return by a direct controller, so this is a guess based on other enums
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifeCycleState())
	out.BigqueryDataset = BigQueryDataset_FromProto(mapCtx, in.BigQueryDataset)
}
func LoggingLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLinkObservedState) *pb.Link {
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// This is the first lifecycle state return by a direct controller, so this is a guess based on other enums
	out.LifecycleState = direct.Enum_ToProto(mapCtx, in.GetLifeCycleState())
	out.BigqueryDataset = BigQueryDataset_ToProto(mapCtx, in.BigQueryDataset)
}
