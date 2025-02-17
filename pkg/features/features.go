/*
Copyright 2024 The CodeFuture Authors.

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

package features

import (
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/component-base/featuregate"
	"k8s.io/klog/v2"
)

const (
	// PushReconciler ensures that managed resources in member clusters represent the state declared in federated resources.
	PushReconciler featuregate.Feature = "PushReconciler"

	// SchedulerPreferences Scheduler controllers which dynamically schedules workloads based on user preferences.
	SchedulerPreferences featuregate.Feature = "SchedulerPreferences"

	// RawResourceStatusCollection enables the collection of the status of target types when enabled
	RawResourceStatusCollection featuregate.Feature = "RawResourceStatusCollection"
)

func init() {
	if err := utilfeature.DefaultMutableFeatureGate.Add(DefaultKubeFedFeatureGates); err != nil {
		klog.Fatalf("Unexpected error: %v", err)
	}
}

// DefaultKubeFedFeatureGates consists of all known KubeFed-specific
// feature keys.  To add a new feature, define a key for it above and
// add it here.
var DefaultKubeFedFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	SchedulerPreferences:        {Default: true, PreRelease: featuregate.Alpha},
	PushReconciler:              {Default: true, PreRelease: featuregate.Beta},
	RawResourceStatusCollection: {Default: false, PreRelease: featuregate.Beta},
}
