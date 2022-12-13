package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-users2/client-emailsettings-go/version1"
	"github.com/stretchr/testify/assert"
)

type EmailSettingsClientFixtureV1 struct {
	Client version1.IEmailSettingsClientV1
}

var SETTINGS = &version1.EmailSettingsV1{
	Id:       "1",
	Name:     "User 1",
	Email:    "user1@conceptual.vision",
	Language: "en",
	Verified: false,
}

func NewEmailSettingsClientFixtureV1(client version1.IEmailSettingsClientV1) *EmailSettingsClientFixtureV1 {
	return &EmailSettingsClientFixtureV1{
		Client: client,
	}
}

func (c *EmailSettingsClientFixtureV1) clear() {
	c.Client.DeleteSettingsById(context.Background(), "", SETTINGS.Id)
}

func (c *EmailSettingsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create email settings
	settings, err := c.Client.SetSettings(context.Background(), "", SETTINGS)
	assert.Nil(t, err)

	assert.NotNil(t, settings)
	assert.Equal(t, settings.Id, SETTINGS.Id)
	assert.Equal(t, settings.Email, SETTINGS.Email)
	assert.False(t, settings.Verified)

	settings1 := settings

	// Update verified email settings
	settings, err = c.Client.SetVerifiedSettings(context.Background(), "", settings1)
	assert.Nil(t, err)

	assert.NotNil(t, settings)
	assert.Equal(t, settings.Id, settings1.Id)
	assert.Equal(t, settings.Email, settings1.Email)
	assert.True(t, settings.Verified)

	//settings2 := settings

	// Update the email settings
	settings1.Subscriptions = map[string]any{"engagement": true}
	settings, err = c.Client.SetSettings(context.Background(), "", settings1)
	assert.Nil(t, err)

	assert.NotNil(t, settings)
	assert.NotNil(t, settings.Subscriptions)

	subscriptions := settings.Subscriptions.(map[string]any)
	assert.True(t, subscriptions["engagement"].(bool))

	settings1 = settings

	// Get email settings
	list, err1 := c.Client.GetSettingsByIds(context.Background(), "", []string{settings1.Id})
	assert.Nil(t, err1)

	assert.NotNil(t, list)
	assert.Len(t, list, 1)

	// Delete email settings
	err = c.Client.DeleteSettingsById(context.Background(), "", settings1.Id)
	assert.Nil(t, err)

	// Try to get deleted settings
	settings, err = c.Client.GetSettingsById(context.Background(), "", settings1.Id)
	assert.Nil(t, err)

	assert.Nil(t, settings)
}
