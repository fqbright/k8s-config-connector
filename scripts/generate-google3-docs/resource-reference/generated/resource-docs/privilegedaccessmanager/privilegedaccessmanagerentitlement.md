{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}PrivilegedAccessManagerEntitlement{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>PrivilegedAccessManager</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/iam/docs/pam-overview">/iam/docs/pam-overview</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>
<pre>v1.folders.locations.entitlements</pre>
<pre>v1.organizations.locations.entitlements</pre>
<pre>v1.projects.locations.entitlements</pre>
</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td>
<pre><a href="/iam/docs/reference/pam/rest/v1/folders.locations.entitlements">/logging/docs/reference/v2/rest/v2/folders.exclusions</a></pre>
<pre><a href="/iam/docs/reference/pam/rest/v1/organizations.locations.entitlements">/logging/docs/reference/v2/rest/v2/organizations.exclusions</a></pre>
<pre><a href="/iam/docs/reference/pam/rest/v1/projects.locations.entitlements">/logging/docs/reference/v2/rest/v2/projects.exclusions</a></pre>
</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpprivilegedaccessmanagerentitlement<br>gcpprivilegedaccessmanagerentitlements<br>privilegedaccessmanagerentitlement</td>
</tr>
<td>{{product_name_short}} Service Name</td>
<td>privilegedaccessmanager.googleapis.com</td>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>privilegedaccessmanagerentitlements.privilegedaccessmanager.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


<tr>
<td>{{product_name_short}} Default Average Reconcile Interval In Seconds</td>
<td>600</td>
</tr>
</tbody>
</table>

## Custom Resource Definition Properties



### Spec
#### Schema
```yaml
additionalNotificationTargets:
  adminEmailRecipients:
  - string
  requesterEmailRecipients:
  - string
approvalWorkflow:
  manualApprovals:
    requireApproverJustification: boolean
    steps:
    - approvalsNeeded: integer
      approverEmailRecipients:
      - string
      approvers:
      - principals:
        - string
eligibleUsers:
- principals:
  - string
folderRef:
  external: string
  name: string
  namespace: string
location: string
maxRequestDuration: string
organizationRef:
  external: string
privilegedAccess:
  gcpIAMAccess:
    roleBindings:
    - conditionExpression: string
      role: string
projectRef:
  external: string
  kind: string
  name: string
  namespace: string
requesterJustificationConfig:
  notMandatory: {}
  unstructured: {}
resourceID: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>additionalNotificationTargets</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Additional email addresses to be notified based on actions taken.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>additionalNotificationTargets.adminEmailRecipients</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. Additional email addresses to be notified when a principal (requester) is granted access.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>additionalNotificationTargets.adminEmailRecipients[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>additionalNotificationTargets.requesterEmailRecipients</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. Additional email address to be notified about an eligible entitlement.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>additionalNotificationTargets.requesterEmailRecipients[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. The approvals needed before access are granted to a requester. No approvals are needed if this field is null.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}An approval workflow where users designated as approvers review and act on the grants.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.requireApproverJustification</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Optional. Whether the approvers need to provide a justification for their actions.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Optional. List of approval steps in this workflow. These steps are followed in the specified order sequentially. Only 1 step is supported.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Step represents a logical step in a manual approval workflow.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[].approvalsNeeded</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Required. How many users from the above list need to approve. If there aren't enough distinct users in the list, then the workflow indefinitely blocks. Should always be greater than 0. 1 is the only supported value.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[].approverEmailRecipients</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. Additional email addresses to be notified when a grant is pending approval.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[].approverEmailRecipients[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[].approvers</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Optional. The potential set of approvers in this step. This list must contain at most one entry.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[].approvers[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}AccessControlEntry is used to control who can do some operation.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[].approvers[].principals</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. Users who are allowed for the operation. Each entry should be a valid v1 IAM principal identifier. The format for these is documented at: https://cloud.google.com/iam/docs/principal-identifiers#v1{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>approvalWorkflow.manualApprovals.steps[].approvers[].principals[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>eligibleUsers</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Who can create grants using this entitlement. This list should contain at most one entry.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>eligibleUsers[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}AccessControlEntry is used to control who can do some operation.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>eligibleUsers[].principals</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. Users who are allowed for the operation. Each entry should be a valid v1 IAM principal identifier. The format for these is documented at: https://cloud.google.com/iam/docs/principal-identifiers#v1{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>eligibleUsers[].principals[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>folderRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. The Folder that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>folderRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The 'name' field of a folder, when not managed by Config Connector. This field must be set when 'name' field is not set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>folderRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The 'name' field of a 'Folder' resource. This field must be set when 'external' field is not set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>folderRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The 'namespace' field of a 'Folder' resource. If unset, the namespace is defaulted to the namespace of the referencer resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>location</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Location of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>maxRequestDuration</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Required. The maximum amount of time that access is granted for a request. A requester can ask for a duration less than this, but never more.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>organizationRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. The Organization that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>organizationRef.external</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The 'name' field of an organization, when not managed by Config Connector.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>privilegedAccess</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The access granted to a requester on successful approval.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>privilegedAccess.gcpIAMAccess</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Access to a Google Cloud resource through IAM.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>privilegedAccess.gcpIAMAccess.roleBindings</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Required. Role bindings that are created on successful grant.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>privilegedAccess.gcpIAMAccess.roleBindings[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}RoleBinding represents IAM role bindings that are created after a successful grant.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>privilegedAccess.gcpIAMAccess.roleBindings[].conditionExpression</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. The expression field of the IAM condition to be associated
with the role. If specified, a user with an active grant for this
entitlement is able to access the resource only if this condition
evaluates to true for their request.

This field uses the same CEL format as IAM and supports all attributes
that IAM supports, except tags. More details can be found at
https://cloud.google.com/iam/docs/conditions-overview#attributes.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>privilegedAccess.gcpIAMAccess.roleBindings[].role</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Required. IAM role to be granted. More details can be found at https://cloud.google.com/iam/docs/roles-overview.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. The Project that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `projectID` field of a project, when not managed by Config Connector.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.kind</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The kind of the Project resource; optional but must be `Project` if provided.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `name` field of a `Project` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `namespace` field of a `Project` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>requesterJustificationConfig</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Required. The manner in which the requester should provide a justification for requesting access.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>requesterJustificationConfig.notMandatory</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}NotMandatory justification type means the justification isn't required and can be provided in any of the supported formats. The user must explicitly opt out using this field if a justification from the requester isn't mandatory. The only accepted value is `{}` (empty struct). Either 'notMandatory' or 'unstructured' field must be set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>requesterJustificationConfig.unstructured</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Unstructured justification type means the justification is in the format of a string. If this is set, the server allows the requester to provide a justification but doesn't validate it. The only accepted value is `{}` (empty struct). Either 'notMandatory' or 'unstructured' field must be set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The PrivilegedAccessManagerEntitlement name. If not given, the 'metadata.name' will be used.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>* Field is required when parent field is specified</p>


### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
externalRef: string
observedGeneration: integer
observedState:
  createTime: string
  etag: string
  state: string
  updateTime: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observations of the object's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>externalRef</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A unique specifier for the PrivilegedAccessManagerEntitlement resource in GCP.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to 'metadata.generation', then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedState</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}ObservedState is the state of the resource as most recently observed in GCP.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedState.createTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Create time stamp.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedState.etag</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}An 'etag' is used for optimistic concurrency control as a way to prevent simultaneous updates to the same entitlement. An 'etag' is returned in the response to 'GetEntitlement' and the caller should put the 'etag' in the request to 'UpdateEntitlement' so that their change is applied on the same version. If this field is omitted or if there is a mismatch while updating an entitlement, then the server rejects the request.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedState.state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Current state of this entitlement.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedState.updateTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Update time stamp.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Folder Level Entitlement
```yaml
# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: privilegedaccessmanager.cnrm.cloud.google.com/v1beta1
kind: PrivilegedAccessManagerEntitlement
metadata:
  name: privilegedaccessmanagerentitlement-sample-folder
spec:
  folderRef:
    # Replace ${FOLDER_ID?} with your folder ID.
    external: folders/${FOLDER_ID?}
  location: global
  maxRequestDuration: 1800s
  privilegedAccess:
    gcpIAMAccess:
      roleBindings:
        - role: roles/pubsub.viewer
          conditionExpression: "request.time > timestamp(\"2019-12-31T12:00:00.000Z\")"
  requesterJustificationConfig:
    notMandatory: {}
  eligibleUsers:
    - principals:
        # Replace ${PROJECT_ID?} with your project ID.
        - serviceAccount:pame-dep1-folder@${PROJECT_ID?}.iam.gserviceaccount.com
  additionalNotificationTargets:
    adminEmailRecipients:
      # Replace ${PROJECT_ID?} with your project ID.
      - pame-dep1-folder@${PROJECT_ID?}.iam.gserviceaccount.com
    requesterEmailRecipients:
      # Replace ${PROJECT_ID?} with your project ID.
      - pame-dep1-folder@${PROJECT_ID?}.iam.gserviceaccount.com
      - pame-dep2-folder@${PROJECT_ID?}.iam.gserviceaccount.com
  approvalWorkflow:
    manualApprovals:
      requireApproverJustification: true
      steps:
        - approvalsNeeded: 1
          approverEmailRecipients:
            # Replace ${PROJECT_ID?} with your project ID.
            - pame-dep2-folder@${PROJECT_ID?}.iam.gserviceaccount.com
          approvers:
            - principals:
                # Replace ${GROUP_EMAIL?} with your group email.
                - "group:${GROUP_EMAIL?}"
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  annotations:
    # Replace ${PROJECT_ID?} with your project ID.
    cnrm.cloud.google.com/project-id: "${PROJECT_ID?}"
  name: pame-dep1-folder
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  annotations:
    # Replace ${PROJECT_ID?} with your project ID.
    cnrm.cloud.google.com/project-id: "${PROJECT_ID?}"
  name: pame-dep2-folder
```

### Org Level Entitlement
```yaml
# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: privilegedaccessmanager.cnrm.cloud.google.com/v1beta1
kind: PrivilegedAccessManagerEntitlement
metadata:
  name: privilegedaccessmanagerentitlement-sample-org
spec:
  organizationRef:
    # Replace ${ORG_ID?} with your organization ID.
    external: organizations/${ORG_ID?}
  location: global
  maxRequestDuration: 1800s
  privilegedAccess:
    gcpIAMAccess:
      roleBindings:
        - role: roles/pubsub.viewer
          conditionExpression: "request.time > timestamp(\"2019-12-31T12:00:00.000Z\")"
  requesterJustificationConfig:
    unstructured: {}
  eligibleUsers:
    - principals:
        # Replace ${PROJECT_ID?} with your project ID.
        - serviceAccount:pame-dep-org@${PROJECT_ID?}.iam.gserviceaccount.com
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  annotations:
    # Replace ${PROJECT_ID?} with your project ID.
    cnrm.cloud.google.com/project-id: "${PROJECT_ID?}"
  name: pame-dep-org
```

### Project Level Entitlement
```yaml
# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: privilegedaccessmanager.cnrm.cloud.google.com/v1beta1
kind: PrivilegedAccessManagerEntitlement
metadata:
  name: privilegedaccessmanagerentitlement-sample-project
spec:
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID
    external: "projects/${PROJECT_ID?}"
  location: global
  maxRequestDuration: 1800s
  privilegedAccess:
    gcpIAMAccess:
      roleBindings:
        - role: roles/pubsub.admin
  requesterJustificationConfig:
    notMandatory: {}
  eligibleUsers:
    - principals:
        # Replace ${PROJECT_ID?} with your project ID
        - serviceAccount:pame-dep-project@${PROJECT_ID?}.iam.gserviceaccount.com
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  annotations:
    # Replace ${PROJECT_ID?} with your project ID.
    cnrm.cloud.google.com/project-id: "${PROJECT_ID?}"
  name: pame-dep-project
```


Note: If you have any trouble with instantiating the resource, refer to <a href="/config-connector/docs/troubleshooting">Troubleshoot Config Connector</a>.

{% endblock %}
