/*
Copyright 2019 Polyaxon, Inc.

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

package kf

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/polyaxon/polyaxon/operator/api/v1alpha1"

	pytorchjobv1 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1"
	kfcommonv1 "github.com/kubeflow/tf-operator/pkg/apis/common/v1"
)

// CopyPytorchJobFields copies the owned fields from one PytrochJob to another
// Returns true if the fields copied from don't match to.
func CopyPytorchJobFields(from, to *pytorchjobv1.PyTorchJob) bool {
	requireUpdate := false
	for k, v := range to.Labels {
		if from.Labels[k] != v {
			requireUpdate = true
		}
	}
	to.Labels = from.Labels

	if to.Spec.ActiveDeadlineSeconds != from.Spec.ActiveDeadlineSeconds {
		to.Spec.ActiveDeadlineSeconds = from.Spec.ActiveDeadlineSeconds
		requireUpdate = true
	}

	if to.Spec.BackoffLimit != from.Spec.BackoffLimit {
		to.Spec.BackoffLimit = from.Spec.BackoffLimit
		requireUpdate = true
	}

	if to.Spec.CleanPodPolicy != from.Spec.CleanPodPolicy {
		to.Spec.CleanPodPolicy = from.Spec.CleanPodPolicy
		requireUpdate = true
	}

	if to.Spec.TTLSecondsAfterFinished != from.Spec.TTLSecondsAfterFinished {
		to.Spec.TTLSecondsAfterFinished = from.Spec.TTLSecondsAfterFinished
		requireUpdate = true
	}

	if !reflect.DeepEqual(to.Spec, from.Spec) {
		requireUpdate = true
		to.Spec = from.Spec
	}

	return requireUpdate
}

// GeneratePytorchJob returns a a PytorchJob given a a KFSpec
func GeneratePytorchJob(
	name string,
	namespace string,
	labels map[string]string,
	spec corev1alpha1.KFSpec,
) *pytorchjobv1.PyTorchJob {
	tfReplicaSpecs := map[pytorchjobv1.PyTorchReplicaType]*kfcommonv1.ReplicaSpec{}
	for k, v := range spec.ReplicaSpecs {
		tfReplicaSpecs[pytorchjobv1.PyTorchReplicaType(k)] = generateKFReplica(v)
	}

	// copy all of the labels to the pod including pod default related labels
	for _, replicaSpec := range tfReplicaSpecs {
		l := &replicaSpec.Template.ObjectMeta.Labels
		for k, v := range labels {
			(*l)[k] = v
		}
	}

	tfJobSpec := pytorchjobv1.PyTorchJobSpec{
		ActiveDeadlineSeconds:   spec.ActiveDeadlineSeconds,
		BackoffLimit:            spec.BackoffLimit,
		CleanPodPolicy:          spec.CleanPodPolicy,
		TTLSecondsAfterFinished: spec.TTLSecondsAfterFinished,
		PyTorchReplicaSpecs:     tfReplicaSpecs,
	}

	tfJob := &pytorchjobv1.PyTorchJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: tfJobSpec,
	}

	return tfJob
}
