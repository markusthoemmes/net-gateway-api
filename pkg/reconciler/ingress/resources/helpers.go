/*
Copyright 2021 The Knative Authors

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

package resources

import (
	"sort"

	gwv1alpha1 "sigs.k8s.io/gateway-api/apis/v1alpha1"
)

func gatewayAllowTypePtr(val gwv1alpha1.GatewayAllowType) *gwv1alpha1.GatewayAllowType {
	return &val
}

func headerMatchTypePtr(val gwv1alpha1.HeaderMatchType) *gwv1alpha1.HeaderMatchType {
	return &val
}

func portNumPtr(port int) *gwv1alpha1.PortNumber {
	pn := gwv1alpha1.PortNumber(port)
	return &pn
}

func pathMatchTypePtr(val gwv1alpha1.PathMatchType) *gwv1alpha1.PathMatchType {
	return &val
}

// LongestHost returns the most specific host.
// The length is:
// 1. the length of the hostnames.
// 2. the first alphabetical order.
//
// For example, "hello-example.default.svc.cluster.local" will be
// returned from the following hosts in KIngress.
//
//  - hosts:
//    - hello.default
//    - hello.default.svc
//    - hello.default.svc.cluster.local
//
func LongestHost(hosts []string) string {
	sort.Strings(hosts)
	return hosts[len(hosts)-1]
}
