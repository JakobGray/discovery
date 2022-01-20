// Copyright Contributors to the Open Cluster Management project

package v1alpha1

import (
	v1 "github.com/stolostron/discovery/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this DiscoveryConfig to the Hub version (v1).
func (src *DiscoveryConfig) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.DiscoveryConfig)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Credential = src.Spec.Credential

	openshiftVersions := make([]v1.Semver, len(src.Spec.Filters.OpenShiftVersions))
	for i := range src.Spec.Filters.OpenShiftVersions {
		openshiftVersions[i] = v1.Semver(src.Spec.Filters.OpenShiftVersions[i])
	}
	dst.Spec.Filters = v1.Filter{
		LastActive:        src.Spec.Filters.LastActive,
		OpenShiftVersions: openshiftVersions,
	}

	// Status

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *DiscoveryConfig) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.DiscoveryConfig)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Credential = src.Spec.Credential

	openshiftVersions := make([]Semver, len(src.Spec.Filters.OpenShiftVersions))
	for i := range src.Spec.Filters.OpenShiftVersions {
		openshiftVersions[i] = Semver(src.Spec.Filters.OpenShiftVersions[i])
	}
	dst.Spec.Filters = Filter{
		LastActive:        src.Spec.Filters.LastActive,
		OpenShiftVersions: openshiftVersions,
	}

	// Status

	return nil
}
