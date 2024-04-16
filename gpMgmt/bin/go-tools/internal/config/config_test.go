package config_test

import (
	"fmt"
	"testing"

	"github.com/greenplum-db/gpdb/gp/internal/config"
	"github.com/greenplum-db/gpdb/gp/internal/enums"
	"github.com/greenplum-db/gpdb/gp/testutils"
	"github.com/greenplum-db/gpdb/gp/testutils/exectest"
)

func init() {
	exectest.RegisterMains()
}

func TestConfig(t *testing.T) {

	t.Run("config not present at location", func(t *testing.T) {

		testConfig := config.New()
		err := testConfig.Load(".")
		if err == nil {
			t.Fatalf("Expected %s, got %d", "error", err)
		}

	})

	t.Run("config loaded succesfully with defaults", func(t *testing.T) {

		testConfig := config.New()
		testConfig.SetName("gpdeploy.testconf")
		err := testConfig.Load("../../test/data")

		testutils.Assert(t, nil, err, "Failed to load configs")
		testutils.Assert(t, enums.DeploymentTypeMirrorless, testConfig.GetDatabaseConfig().GetDeploymentType(), "")

		testutils.Assert(t, 4506, testConfig.GetInfraConfig().GetRequestPort(), "")
		testutils.Assert(t, 4505, testConfig.GetInfraConfig().GetPublishPort(), "")
		testutils.Assert(t, "custom-name", testConfig.GetInfraConfig().GetStandby().GetHostname(), "")
		testutils.Assert(t, "10.202.89.77", testConfig.GetInfraConfig().GetCoordinator().GetIp(), "")
		testutils.Assert(t, "192.168.100.1/24", testConfig.GetInfraConfig().GetSegmentHost().GetNetwork().GetInternalCidr(), "")
	})

	t.Run("config loaded succesfully overriding defaults", func(t *testing.T) {

		testConfig := config.New()
		testConfig.SetName("gpdeploy_override.testconf")
		err := testConfig.Load("../../test/data")

		testutils.Assert(t, nil, err, "Failed to load configs")
		testutils.Assert(t, enums.DeploymentTypeMirrorless, testConfig.GetDatabaseConfig().GetDeploymentType(), "")
		testutils.Assert(t, 5001, testConfig.GetInfraConfig().GetRequestPort(), "")
		testutils.Assert(t, 5002, testConfig.GetInfraConfig().GetPublishPort(), "")
		testutils.Assert(t, "cdw", testConfig.GetInfraConfig().GetCoordinator().GetHostname(), "")
		testutils.Assert(t, "&{password dssdfwef}", fmt.Sprintf("%v", testConfig.GetInfraConfig().GetStandby().GetAuth()), "")

	})
}
