package e2e

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// When the e2e tests are run, we have no guarantees that the cluster
	// dns operator is started up. First baby step here is to ensure that
	// the cluster dns operator deployment exists.
	if err := waitForDNSOperatorDeployment(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(74)
	}

	os.Exit(m.Run())
}
