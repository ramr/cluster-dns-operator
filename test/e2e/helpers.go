package e2e

import (
	"fmt"
	"time"

	"github.com/operator-framework/operator-sdk/pkg/sdk"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	// clusterDNSOperatorNamespace is cluster dns operator's namespace.
	clusterDNSOperatorNamespace = "openshift-cluster-dns-operator"

	// clusterDNSOperatorDeployment is the name of the cluster dns operator deployment.
	clusterDNSOperatorDeployment = "cluster-dns-operator"

	// pollInterval controls how often we poll.
	pollInterval = 5 * time.Second

	// waitTimeout controls how long we wait for a resource.
	waitTimeout = 15 * time.Minute
)

// waitForCondition waits for a match using a condition function.
func waitForCondition(conditionFn wait.ConditionFunc) error {
	return wait.PollImmediate(pollInterval, waitTimeout, conditionFn)
}

// waitForSDKResource waits for the sdk resource to be created.
func waitForSDKResource(obj runtime.Object) error {
	resourceCreated := func() (bool, error) {
		if err := sdk.Get(obj); err != nil {
			return false, nil
		}

		return true, nil
	}

	return waitForCondition(resourceCreated)
}

// waitForDNSOperatorDeployment waits for the cluster dns operator deployment
// to be created.
func waitForDNSOperatorDeployment() error {
	d := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: appsv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: clusterDNSOperatorNamespace,
			Name:      clusterDNSOperatorDeployment,
		},
	}

	if err := waitForSDKResource(d); err != nil {
		return fmt.Errorf("waiting for ClusterDNS operator deployment %s/%s to be created: %v", d.Namespace, d.Name, err)
	}

	return nil
}
