package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-emailsettings-go/version1"
)

type settingsMemoryClientV1Test struct {
	client  *version1.EmailSettingsMockClientV1
	fixture *EmailSettingsClientFixtureV1
}

func newsettingsMemoryClientV1Test() *settingsMemoryClientV1Test {
	return &settingsMemoryClientV1Test{}
}

func (c *settingsMemoryClientV1Test) setup(t *testing.T) {
	c.client = version1.NewEmailSettingsMockClientV1()

	c.fixture = NewEmailSettingsClientFixtureV1(c.client)
}

func TestMockCrudOperations(t *testing.T) {
	c := newsettingsMemoryClientV1Test()
	c.setup(t)

	c.fixture.TestCrudOperations(t)
}
