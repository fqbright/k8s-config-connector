{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ManagedKafkaCluster{% endblock %}
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
<td>Managed Kafka</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/managed-service-for-apache-kafka/docs/">/managed-service-for-apache-kafka/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1beta1.projects.locations.clusters</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/managed-service-for-apache-kafka/docs/reference/rest/v1/projects.locations.clusters">/managed-service-for-apache-kafka/docs/reference/rest/v1/projects.locations.clusters</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpmanagedkafkacluster<br>gcpmanagedkafkaclusters<br>managedkafkacluster</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>managedkafka.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>managedkafkaclusters.managedkafka.cnrm.cloud.google.com</td>
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
capacityConfig:
  memoryBytes: integer
  vcpuCount: integer
gcpConfig:
  accessConfig:
    networkConfigs:
    - subnetworkRef:
        external: string
        name: string
        namespace: string
  kmsKeyRef:
    external: string
    name: string
    namespace: string
labels:
  string: string
location: string
projectRef:
  external: string
  kind: string
  name: string
  namespace: string
rebalanceConfig:
  mode: string
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
            <p><code>capacityConfig</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Required. Capacity configuration for the Kafka cluster.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>capacityConfig.memoryBytes</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Required. The memory to provision for the cluster in bytes. The CPU:memory ratio (vCPU:GiB) must be between 1:1 and 1:8. Minimum: 3221225472 (3 GiB).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>capacityConfig.vcpuCount</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Required. The number of vCPUs to provision for the cluster. Minimum: 3.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Required. Configuration properties for a Kafka cluster deployed to Google Cloud Platform.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.accessConfig</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Required. Access configuration for the Kafka cluster.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.accessConfig.networkConfigs</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Required. Virtual Private Cloud (VPC) networks that must be granted direct access to the Kafka cluster. Minimum of 1 network is required. Maximum 10 networks can be specified.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.accessConfig.networkConfigs[]</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.accessConfig.networkConfigs[].subnetworkRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Required. Reference to the VPC subnet in which to create Private Service Connect
 (PSC) endpoints for the Kafka brokers and bootstrap address.

 The subnet must be located in the same region as the Kafka cluster. The
 project may differ. Multiple subnets from the same parent network must not
 be specified.

 The CIDR range of the subnet must be within the IPv4 address ranges for
 private networks, as specified in RFC 1918.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.accessConfig.networkConfigs[].subnetworkRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The ComputeSubnetwork selflink of form "projects/{{project}}/regions/{{region}}/subnetworks/{{name}}", when not managed by Config Connector.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.accessConfig.networkConfigs[].subnetworkRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `name` field of a `ComputeSubnetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.accessConfig.networkConfigs[].subnetworkRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `namespace` field of a `ComputeSubnetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.kmsKeyRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Immutable. The Cloud KMS Key name to use for encryption. The key must be located in the same region as the cluster and cannot be changed.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.kmsKeyRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A reference to an externally managed KMSCryptoKey. Should be in the format `projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id]/cryptoKeys/[key]`.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.kmsKeyRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `name` of a `KMSCryptoKey` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>gcpConfig.kmsKeyRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The `namespace` of a `KMSCryptoKey` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>labels</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">map (key: string, value: string)</code></p>
            <p>{% verbatim %}Optional. Labels as key value pairs.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>location</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The Project that this resource belongs to.{% endverbatim %}</p>
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
            <p><code>rebalanceConfig</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Rebalance configuration for the Kafka cluster.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rebalanceConfig.mode</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. The rebalance behavior for the cluster. When not specified, defaults to `NO_REBALANCE`.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The GCP resource identifier. If not given, the metadata.name will be used.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



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
            <p>{% verbatim %}A unique specifier for the ManagedKafkaCluster resource in GCP.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
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
            <p>{% verbatim %}Output only. The time when the cluster was created.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedState.state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The current state of the cluster.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedState.updateTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The time when the cluster was last updated.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### ManagedKafkaCluster Basic
```yaml
# Copyright 2025 Google LLC
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

apiVersion: managedkafka.cnrm.cloud.google.com/v1beta1
kind: ManagedKafkaCluster
metadata:
  name: managedkafkacluster-sample
spec:
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID.
    external: ${PROJECT_ID?}
  location: us-central1
  capacityConfig:
    vcpuCount: 3
    memoryBytes: 3221225472 # 3GB
  gcpConfig:
    accessConfig:
      networkConfigs:
      - subnetworkRef:
          name: managedkafkacluster-dep
  rebalanceConfig:
    mode: NO_REBALANCE
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: managedkafkacluster-dep
spec:
  autoCreateSubnetworks: false
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeSubnetwork
metadata:
  name: managedkafkacluster-dep
spec:
  region: us-central1
  networkRef:
    name: managedkafkacluster-dep
  ipCidrRange: 10.0.0.0/24
```

### ManagedKafkaCluster Cmek
```yaml
# Copyright 2025 Google LLC
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

apiVersion: managedkafka.cnrm.cloud.google.com/v1beta1
kind: ManagedKafkaCluster
metadata:
  name: managedkafkacluster-sample-cmek
spec:
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID
    external: ${PROJECT_ID?}
  location: us-central1
  capacityConfig:
    vcpuCount: 3
    memoryBytes: 3221225472 # 3GB
  gcpConfig:
    accessConfig:
      networkConfigs:
      - subnetworkRef:
          name: managedkafkacluster-dep-cmek
    kmsKeyRef:
      name: managedkafkacluster-dep-cmek
  rebalanceConfig:
    mode: NO_REBALANCE
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: managedkafkacluster-dep-cmek
spec:
  autoCreateSubnetworks: false
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeSubnetwork
metadata:
  name: managedkafkacluster-dep-cmek
spec:
  region: us-central1
  networkRef:
    name: managedkafkacluster-dep-cmek
  ipCidrRange: 10.0.0.0/24
---
# Replace ${PROJECT_ID?} and ${PROJECT_NUMBER?} below with your desired project
# ID and project number.
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: managedkafkacluster-dep-cmek
spec:
  resourceRef:
    apiVersion: kms.cnrm.cloud.google.com/v1beta1
    kind: KMSCryptoKey
    name: managedkafkacluster-dep-cmek
  memberFrom:
    serviceIdentityRef:
      name: managedkafkacluster-dep-cmek
  role: roles/cloudkms.cryptoKeyEncrypterDecrypter # required by managedkafka service agent to access KMS keys
---
apiVersion: kms.cnrm.cloud.google.com/v1beta1
kind: KMSCryptoKey
metadata:
  annotations:
    # Replace ${PROJECT_ID?} with your project ID.
    cnrm.cloud.google.com/project-id: ${PROJECT_ID?}
  name: managedkafkacluster-dep-cmek
spec:
  keyRingRef:
    name: managedkafkacluster-dep-cmek
  purpose: ENCRYPT_DECRYPT
---
apiVersion: kms.cnrm.cloud.google.com/v1beta1
kind: KMSKeyRing
metadata:
  annotations:
    # Replace ${PROJECT_ID?} with your project ID.
    cnrm.cloud.google.com/project-id: ${PROJECT_ID?}
  name: managedkafkacluster-dep-cmek
spec:
  location: us-central1
---
apiVersion: serviceusage.cnrm.cloud.google.com/v1beta1
kind: ServiceIdentity
metadata:
  name: managedkafkacluster-dep-cmek
  annotations:
    cnrm.cloud.google.com/deletion-policy: "abandon"
spec:
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID.
    external: ${PROJECT_ID?}
  resourceID: managedkafka.googleapis.com
```


Note: If you have any trouble with instantiating the resource, refer to <a href="/config-connector/docs/troubleshooting">Troubleshoot Config Connector</a>.

{% endblock %}
