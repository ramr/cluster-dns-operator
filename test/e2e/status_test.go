package e2e

import (
	"testing"

	osv1 "github.com/openshift/cluster-version-operator/pkg/apis/operatorstatus.openshift.io/v1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	clusterOperatorName = "openshift-dns"
)

func TestClusterDNSStatus(t *testing.T) {
	co := &osv1.ClusterOperator{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterOperator",
			APIVersion: osv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: clusterDNSOperatorNamespace,
			Name:      clusterOperatorName,
		},
	}

	operatorAvailable := func() (bool, error) {
		if err := sdk.Get(co); err != nil {
			return false, nil
		}

		for _, cond := range co.Status.Conditions {
			if cond.Type == osv1.OperatorAvailable && cond.Status == osv1.ConditionTrue {
				return true, nil
			}
		}

		return false, nil
	}

	if err := waitForCondition(operatorAvailable); err != nil {
		t.Errorf("waiting for cluster operator %s/%s to become available",
			co.Namespace, co.Name)
	}
}
