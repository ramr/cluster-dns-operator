package e2e

import (
	"testing"

	dnsv1alpha1 "github.com/openshift/cluster-dns-operator/pkg/apis/dns/v1alpha1"
	"github.com/openshift/cluster-dns-operator/pkg/manifests"
	"github.com/openshift/cluster-dns-operator/pkg/util"
)

func TestClusterDNSDefaultCR(t *testing.T) {
	// InstallConfig is needed to get the default ClusterDNS, so create a dummy one.
	ic := &util.InstallConfig{
		Networking: util.NetworkingConfig{
			ServiceCIDR: "10.3.0.0/16",
		},
	}

	f := manifests.NewFactory()
	defaultCR, err := f.ClusterDNSDefaultCR(ic)
	if err != nil {
		t.Errorf("getting default ClusterDNS CR: %v", err)
	}

	// Rather than hardcode name/namespace/type/version, just use the
	// ones in the default CR.
	cr := &dnsv1alpha1.ClusterDNS{
		TypeMeta:   defaultCR.TypeMeta,
		ObjectMeta: defaultCR.ObjectMeta,
	}

	if err := waitForSDKResource(cr); err != nil {
		t.Errorf("waiting for default ClusterDNS CR %s/%s to be created: %v",
			cr.Namespace, cr.Name, err)
	}
}
