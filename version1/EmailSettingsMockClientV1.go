package version1

import "context"

type EmailSettingsMockClientV1 struct {
	settings []*EmailSettingsV1
}

func NewEmailSettingsMockClientV1() *EmailSettingsMockClientV1 {
	return &EmailSettingsMockClientV1{
		settings: make([]*EmailSettingsV1, 0),
	}
}

func (c *EmailSettingsMockClientV1) GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) (result []*EmailSettingsV1, err error) {
	settings := make([]*EmailSettingsV1, 0)
	for _, s := range c.settings {
		for _, recepientId := range recipientIds {
			if s.Id == recepientId {
				settings = append(settings, s)
				break
			}
		}
	}

	return settings, nil
}

func (c *EmailSettingsMockClientV1) GetSettingsById(ctx context.Context, correlationId string, recipientId string) (result *EmailSettingsV1, err error) {
	for _, s := range c.settings {
		if s.Id == recipientId {
			return s, nil
		}
	}

	return nil, nil
}

func (c *EmailSettingsMockClientV1) GetSettingsByEmail(ctx context.Context, correlationId string, email string) (result *EmailSettingsV1, err error) {
	for _, s := range c.settings {
		if s.Email == email {
			return s, nil
		}
	}

	return nil, nil
}

func (c *EmailSettingsMockClientV1) SetSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	settings.Verified = false

	for index, s := range c.settings {
		if s.Id == settings.Id {
			c.settings = append(c.settings[:index], c.settings[index+1:]...)
		}
	}

	c.settings = append(c.settings, settings)

	return settings, nil
}

func (c *EmailSettingsMockClientV1) SetVerifiedSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	settings.Verified = true

	for index, s := range c.settings {
		if s.Id == settings.Id {
			c.settings = append(c.settings[:index], c.settings[index+1:]...)
		}
	}

	c.settings = append(c.settings, settings)

	return settings, nil
}

func (c *EmailSettingsMockClientV1) SetRecipient(ctx context.Context, correlationId string, recipientId string, name string, email string, language string) (result *EmailSettingsV1, err error) {
	var settings *EmailSettingsV1

	for _, s := range c.settings {
		if s.Id == recipientId {
			settings = s
			break
		}
	}

	if settings != nil {
		settings.Name = name
		settings.Email = email
		settings.Language = language
	} else {
		settings = NewEmailSettingsV1(recipientId, name, email, language)
		c.settings = append(c.settings, settings)
	}

	return settings, nil
}

func (c *EmailSettingsMockClientV1) SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (result *EmailSettingsV1, err error) {
	var settings *EmailSettingsV1

	for _, s := range c.settings {
		if s.Id == recipientId {
			settings = s
			break
		}
	}

	if settings != nil {
		settings.Subscriptions = subscriptions
	} else {
		settings = NewEmailSettingsV1(recipientId, "", "", "")
		settings.Subscriptions = subscriptions

		c.settings = append(c.settings, settings)
	}

	return settings, nil
}

func (c *EmailSettingsMockClientV1) DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error {
	for index, s := range c.settings {
		if s.Id == recipientId {
			c.settings = append(c.settings[:index], c.settings[index+1:]...)
		}
	}

	return nil
}

func (c *EmailSettingsMockClientV1) ResendVerification(ctx context.Context, correlationId string, recipientId string) error {
	return nil
}

func (c *EmailSettingsMockClientV1) VerifyEmail(ctx context.Context, correlationId string, recipientId string, code string) error {
	var settings *EmailSettingsV1

	for _, s := range c.settings {
		if s.Id == recipientId {
			settings = s
			break
		}
	}

	if settings != nil && settings.VerCode == code {
		settings.Verified = true
		settings.VerCode = ""
	}

	return nil
}
