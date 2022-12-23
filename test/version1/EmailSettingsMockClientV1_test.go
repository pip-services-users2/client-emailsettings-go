package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-emailsettings-go/version1"
)

type settingsMockClientV1Test struct {
	client  *version1.EmailSettingsMockClientV1
	fixture *EmailSettingsClientFixtureV1
}

func newsettingsMockClientV1Test() *settingsMockClientV1Test {
	return &settingsMockClientV1Test{}
}

func (c *settingsMockClientV1Test) setup(t *testing.T) {
	c.client = version1.NewEmailSettingsMockClientV1()

	c.fixture = NewEmailSettingsClientFixtureV1(c.client)
}

func TestMockCrudOperations(t *testing.T) {
	c := newsettingsMockClientV1Test()
	c.setup(t)

	c.fixture.TestCrudOperations(t)
}
