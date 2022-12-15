package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type EmailSettingsCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewEmailSettingsCommandableGrpcClientV1() *EmailSettingsCommandableGrpcClientV1 {
	return NewEmailSettingsCommandableGrpcClientV1WithConfig(nil)
}

func NewEmailSettingsCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *EmailSettingsCommandableGrpcClientV1 {
	c := &EmailSettingsCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/email_settings"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *EmailSettingsCommandableGrpcClientV1) GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) (result []*EmailSettingsV1, err error) {
	res, err := c.CallCommand(ctx, "get_settings_by_ids", correlationId, data.NewAnyValueMapFromTuples(
		"recipient_ids", recipientIds,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableGrpcClientV1) GetSettingsById(ctx context.Context, correlationId string, recipientId string) (result *EmailSettingsV1, err error) {
	res, err := c.CallCommand(ctx, "get_settings_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableGrpcClientV1) GetSettingsByEmail(ctx context.Context, correlationId string, email string) (result *EmailSettingsV1, err error) {
	res, err := c.CallCommand(ctx, "get_settings_by_email", correlationId, data.NewAnyValueMapFromTuples(
		"email", email,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableGrpcClientV1) SetSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	res, err := c.CallCommand(ctx, "set_settings", correlationId, data.NewAnyValueMapFromTuples(
		"settings", settings,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableGrpcClientV1) SetVerifiedSettings(ctx context.Context, correlationId string, settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	res, err := c.CallCommand(ctx, "set_verified_settings", correlationId, data.NewAnyValueMapFromTuples(
		"settings", settings,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableGrpcClientV1) SetRecipient(ctx context.Context, correlationId string, recipientId string, name string, email string, language string) (result *EmailSettingsV1, err error) {
	res, err := c.CallCommand(ctx, "set_recipient", correlationId, data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"name", name,
		"email", email,
		"language", language,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableGrpcClientV1) SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (result *EmailSettingsV1, err error) {
	res, err := c.CallCommand(ctx, "set_subscriptions", correlationId, data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"subscriptions", subscriptions,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*EmailSettingsV1](res, correlationId)
}

func (c *EmailSettingsCommandableGrpcClientV1) DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) (err error) {
	_, err = c.CallCommand(ctx, "delete_settings_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	))

	return err
}

func (c *EmailSettingsCommandableGrpcClientV1) ResendVerification(ctx context.Context, correlationId string, recipientId string) (err error) {
	_, err = c.CallCommand(ctx, "resend_verification", correlationId, data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	))

	return err
}

func (c *EmailSettingsCommandableGrpcClientV1) VerifyEmail(ctx context.Context, correlationId string, recipientId string, code string) (err error) {
	_, err = c.CallCommand(ctx, "verify_email", correlationId, data.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"code", code,
	))

	return err
}
