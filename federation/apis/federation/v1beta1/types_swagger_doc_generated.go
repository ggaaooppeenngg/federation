/*
Copyright 2016 The Kubernetes Authors.

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

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Cluster = map[string]string{
	"":         "Information about a registered cluster in a federated kubernetes setup. Clusters are not namespaced and have unique names in the federation.",
	"metadata": "Standard object's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata",
	"spec":     "Spec defines the behavior of the Cluster.",
	"status":   "Status describes the current status of a Cluster",
}

func (Cluster) SwaggerDoc() map[string]string {
	return map_Cluster
}

var map_ClusterCondition = map[string]string{
	"":                   "ClusterCondition describes current state of a cluster.",
	"type":               "Type of cluster condition, Complete or Failed.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastProbeTime":      "Last time the condition was checked.",
	"lastTransitionTime": "Last time the condition transit from one status to another.",
	"reason":             "(brief) reason for the condition's last transition.",
	"message":            "Human readable message indicating details about last transition.",
}

func (ClusterCondition) SwaggerDoc() map[string]string {
	return map_ClusterCondition
}

var map_ClusterList = map[string]string{
	"":         "A list of all the kubernetes clusters registered to the federation",
	"metadata": "Standard list metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#types-kinds",
	"items":    "List of Cluster objects.",
}

func (ClusterList) SwaggerDoc() map[string]string {
	return map_ClusterList
}

var map_ClusterSpec = map[string]string{
	"": "ClusterSpec describes the attributes of a kubernetes cluster.",
	"serverAddressByClientCIDRs": "A map of client CIDR to server address. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR.",
	"secretRef":                  "Name of the secret containing kubeconfig to access this cluster. The secret is read from the kubernetes cluster that is hosting federation control plane. Admin needs to ensure that the required secret exists. Secret should be in the same namespace where federation control plane is hosted and it should have kubeconfig in its data with key \"kubeconfig\". This will later be changed to a reference to secret in federation control plane when the federation control plane supports secrets. This can be left empty if the cluster allows insecure access.",
}

func (ClusterSpec) SwaggerDoc() map[string]string {
	return map_ClusterSpec
}

var map_ClusterStatus = map[string]string{
	"":           "ClusterStatus is information about the current status of a cluster updated by cluster controller periodically.",
	"conditions": "Conditions is an array of current cluster conditions.",
	"zones":      "Zones is the list of availability zones in which the nodes of the cluster exist, e.g. 'us-east1-a'. These will always be in the same region.",
	"region":     "Region is the name of the region in which all of the nodes in the cluster exist.  e.g. 'us-east1'.",
}

func (ClusterStatus) SwaggerDoc() map[string]string {
	return map_ClusterStatus
}

var map_ServerAddressByClientCIDR = map[string]string{
	"":              "ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.",
	"clientCIDR":    "The CIDR with which clients can match their IP to figure out the server address that they should use.",
	"serverAddress": "Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.",
}

func (ServerAddressByClientCIDR) SwaggerDoc() map[string]string {
	return map_ServerAddressByClientCIDR
}

// AUTO-GENERATED FUNCTIONS END HERE
