package version1

import "context"

type EmailSettingsNullClientV1 struct {
}

func (c *EmailSettingsNullClientV1) GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) (result []*EmailSettingsV1, err error) {
	return make([]*EmailSettingsV1, 0), nil
}

func (c *EmailSettingsNullClientV1) GetSettingsById(ctx context.Context, correlationId string, recipientId string) (result *EmailSettingsV1, err error) {
	return nil, nil
}

func (c *EmailSettingsNullClientV1) GetSettingsByEmail(ctx context.Context, correlationId string, email string) (result *EmailSettingsV1, err error) {
	return nil, nil
}

func (c *EmailSettingsNullClientV1) SetSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	return settings, nil
}

func (c *EmailSettingsNullClientV1) SetVerifiedSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	return settings, nil
}

func (c *EmailSettingsNullClientV1) SetRecipient(ctx context.Context, correlationId string, recipientId string, name string, email string, language string) (result *EmailSettingsV1, err error) {
	return NewEmailSettingsV1(recipientId, name, email, language), nil
}

func (c *EmailSettingsNullClientV1) SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (result *EmailSettingsV1, err error) {
	s := NewEmptyEmailSettingsV1()
	s.Subscriptions = subscriptions

	return s, nil
}

func (c *EmailSettingsNullClientV1) DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error {
	return nil
}

func (c *EmailSettingsNullClientV1) ResendVerification(ctx context.Context, correlationId string, recipientId string) error {
	return nil
}

func (c *EmailSettingsNullClientV1) VerifyEmail(ctx context.Context, correlationId string, recipientId string, code string) error {
	return nil
}
