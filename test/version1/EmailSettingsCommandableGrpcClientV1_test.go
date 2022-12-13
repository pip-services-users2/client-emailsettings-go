package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-emailsettings-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type settingsCommandableGrpcClientV1Test struct {
	client  *version1.EmailSettingsCommandableGrpcClientV1
	fixture *EmailSettingsClientFixtureV1
}

func newSettingsCommandableGrpcClientV1Test() *settingsCommandableGrpcClientV1Test {
	return &settingsCommandableGrpcClientV1Test{}
}

func (c *settingsCommandableGrpcClientV1Test) setup(t *testing.T) {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewEmailSettingsCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEmailSettingsClientFixtureV1(c.client)
}

func (c *settingsCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcCrudOperations(t *testing.T) {
	c := newSettingsCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
