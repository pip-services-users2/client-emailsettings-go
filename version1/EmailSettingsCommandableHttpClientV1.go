package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type EmailSettingsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewEmailSettingsCommandableHttpClientV1() *EmailSettingsCommandableHttpClientV1 {
	return NewEmailSettingsCommandableHttpClientV1ithConfig(nil)
}

func NewEmailSettingsCommandableHttpClientV1ithConfig(config *cconf.ConfigParams) *EmailSettingsCommandableHttpClientV1 {
	c := &EmailSettingsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/email_settings"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *EmailSettingsCommandableHttpClientV1) GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) (result []*EmailSettingsV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"recipient_ids", recipientIds,
	)

	res, err := c.CallCommand(ctx, "get_settings_by_ids", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableHttpClientV1) GetSettingsById(ctx context.Context, correlationId string, recipientId string) (result *EmailSettingsV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	)

	res, err := c.CallCommand(ctx, "get_settings_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableHttpClientV1) GetSettingsByEmail(ctx context.Context, correlationId string, email string) (result *EmailSettingsV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"email", email,
	)

	res, err := c.CallCommand(ctx, "get_settings_by_email", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableHttpClientV1) SetSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"settings", settings,
	)

	res, err := c.CallCommand(ctx, "set_settings", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableHttpClientV1) SetVerifiedSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"settings", settings,
	)

	res, err := c.CallCommand(ctx, "set_verified_settings", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableHttpClientV1) SetRecipient(ctx context.Context, correlationId string, recipientId string, name string, email string, language string) (result *EmailSettingsV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"name", name,
		"email", email,
		"language", language,
	)

	res, err := c.CallCommand(ctx, "set_recipient", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableHttpClientV1) SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (result *EmailSettingsV1, err error) {
	params := data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"subscriptions", subscriptions,
	)

	res, err := c.CallCommand(ctx, "set_subscriptions", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableHttpClientV1) DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error {
	params := data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	)

	_, err := c.CallCommand(ctx, "delete_settings_by_id", correlationId, params)
	return err
}

func (c *EmailSettingsCommandableHttpClientV1) ResendVerification(ctx context.Context, correlationId string, recipientId string) error {
	params := data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	)

	_, err := c.CallCommand(ctx, "resend_verification", correlationId, params)
	return err
}

func (c *EmailSettingsCommandableHttpClientV1) VerifyEmail(ctx context.Context, correlationId string, recipientId string, code string) error {
	params := data.NewAnyValueMapFromTuples(
		"verify_email", recipientId,
		"code", code,
	)

	_, err := c.CallCommand(ctx, "resend_verification", correlationId, params)
	return err
}
