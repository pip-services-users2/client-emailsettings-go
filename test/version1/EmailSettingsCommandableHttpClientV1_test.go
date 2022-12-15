package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-emailsettings-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type settingsCommandableHttpClientV1Test struct {
	client  *version1.EmailSettingsCommandableHttpClientV1
	fixture *EmailSettingsClientFixtureV1
}

func newSettingsCommandableHttpClientV1Test() *settingsCommandableHttpClientV1Test {
	return &settingsCommandableHttpClientV1Test{}
}

func (c *settingsCommandableHttpClientV1Test) setup(t *testing.T) {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewEmailSettingsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewEmailSettingsClientFixtureV1(c.client)
}

func (c *settingsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpCrudOperations(t *testing.T) {
	c := newSettingsCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
