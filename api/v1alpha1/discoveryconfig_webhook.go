/*


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

package v1alpha1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var discoveryconfiglog = logf.Log.WithName("discoveryconfig-resource")

func (r *DiscoveryConfig) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-discovery-open-cluster-management-io-v1alpha1-discoveryconfig,mutating=false,failurePolicy=fail,sideEffects=None,groups=discovery.open-cluster-management.io,resources=discoveryconfigs,verbs=create;update,versions=v1alpha1,name=vdiscoveryconfig.kb.io

var _ webhook.Validator = &DiscoveryConfig{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DiscoveryConfig) ValidateCreate() error {
	discoveryconfiglog.Info("validate create", "name", r.Name)

	return r.validateDiscoveryConfig()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DiscoveryConfig) ValidateUpdate(old runtime.Object) error {
	discoveryconfiglog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *DiscoveryConfig) ValidateDelete() error {
	discoveryconfiglog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *DiscoveryConfig) validateDiscoveryConfig() error {
	var allErrs field.ErrorList
	if err := r.validateDiscoveryConfigName(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "batch.tutorial.kubebuilder.io", Kind: "CronJob"},
		r.Name, allErrs)
}

func (r *DiscoveryConfig) validateDiscoveryConfigName() *field.Error {
	if r.ObjectMeta.Name != "discovery" {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be 'discovery'")
	}
	return nil
}
