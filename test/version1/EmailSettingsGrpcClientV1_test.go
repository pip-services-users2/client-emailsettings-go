package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-emailsettings-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type settingsGrpcClientV1Test struct {
	client  *version1.EmailSettingsGrpcClientV1
	fixture *EmailSettingsClientFixtureV1
}

func newSettingsGrpcClientV1Test() *settingsGrpcClientV1Test {
	return &settingsGrpcClientV1Test{}
}

func (c *settingsGrpcClientV1Test) setup(t *testing.T) {
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

	c.client = version1.NewEmailSettingsGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEmailSettingsClientFixtureV1(c.client)
}

func (c *settingsGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcCrudOperations(t *testing.T) {
	c := newSettingsHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
