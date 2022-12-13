package version1

import (
	"context"

	"github.com/pip-services-users2/client-emailsettings-go/protos"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type EmailSettingsGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewEmailSettingsGrpcClientV1() *EmailSettingsGrpcClientV1 {
	return &EmailSettingsGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("emailsettings_v1.EmailSettingsx"),
	}
}

func (c *EmailSettingsGrpcClientV1) GetSettingsByIds(ctx context.Context, correlationId string,
	recipientIds []string) (result []*EmailSettingsV1, err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.get_settings_by_ids")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsIdsRequest{
		CorrelationId: correlationId,
		RecipientIds:  recipientIds,
	}

	reply := new(protos.EmailSettingsListReply)
	err = c.CallWithContext(ctx, "get_settings_by_ids", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toEmailSettingsList(reply.Settings)

	return result, nil
}

func (c *EmailSettingsGrpcClientV1) GetSettingsById(ctx context.Context, correlationId string,
	recipientId string) (result *EmailSettingsV1, err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.get_settings_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsIdRequest{
		CorrelationId: correlationId,
		RecipientId:   recipientId,
	}

	reply := new(protos.EmailSettingsObjectReply)
	err = c.CallWithContext(ctx, "get_settings_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toEmailSettings(reply.Settings)

	return result, nil
}

func (c *EmailSettingsGrpcClientV1) GetSettingsByEmail(ctx context.Context, correlationId string,
	email string) (result *EmailSettingsV1, err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.get_settings_by_email")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsEmailRequest{
		CorrelationId: correlationId,
		Email:         email,
	}

	reply := new(protos.EmailSettingsObjectReply)
	err = c.CallWithContext(ctx, "get_settings_by_email", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toEmailSettings(reply.Settings)

	return result, nil
}

func (c *EmailSettingsGrpcClientV1) SetSettings(ctx context.Context, correlationId string,
	settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.set_settings")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsObjectRequest{
		CorrelationId: correlationId,
		Settings:      fromEmailSettings(settings),
	}

	reply := new(protos.EmailSettingsObjectReply)
	err = c.CallWithContext(ctx, "set_settings", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toEmailSettings(reply.Settings)

	return result, nil
}

func (c *EmailSettingsGrpcClientV1) SetVerifiedSettings(ctx context.Context, correlationId string,
	settings *EmailSettingsV1) (result *EmailSettingsV1, err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.set_verified_settings")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsObjectRequest{
		CorrelationId: correlationId,
		Settings:      fromEmailSettings(settings),
	}

	reply := new(protos.EmailSettingsObjectReply)
	err = c.CallWithContext(ctx, "set_verified_settings", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toEmailSettings(reply.Settings)

	return result, nil
}

func (c *EmailSettingsGrpcClientV1) SetRecipient(ctx context.Context, correlationId string,
	recipientId string, name string, email string,
	language string) (result *EmailSettingsV1, err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.set_recipient")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsRecipientRequest{
		CorrelationId: correlationId,
		RecipientId:   recipientId,
		Name:          name,
		Email:         email,
		Language:      language,
	}

	reply := new(protos.EmailSettingsObjectReply)
	err = c.CallWithContext(ctx, "set_recipient", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toEmailSettings(reply.Settings)

	return result, nil
}

func (c *EmailSettingsGrpcClientV1) SetSubscriptions(ctx context.Context, correlationId string,
	recipientId string, subscriptions any) (result *EmailSettingsV1, err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.set_subscriptions")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsSubscriptionsRequest{
		CorrelationId: correlationId,
		RecipientId:   recipientId,
		Subscriptions: toJson(subscriptions),
	}

	reply := new(protos.EmailSettingsObjectReply)
	err = c.CallWithContext(ctx, "set_subscriptions", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toEmailSettings(reply.Settings)

	return result, nil
}

func (c *EmailSettingsGrpcClientV1) DeleteSettingsById(ctx context.Context, correlationId string,
	recipientId string) (err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.delete_settings_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsIdRequest{
		CorrelationId: correlationId,
		RecipientId:   recipientId,
	}

	reply := new(protos.EmailSettingsEmptyReply)
	err = c.CallWithContext(ctx, "delete_settings_by_id", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *EmailSettingsGrpcClientV1) ResendVerification(ctx context.Context, correlationId string,
	recipientId string) (err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.resend_verification")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsIdRequest{
		CorrelationId: correlationId,
		RecipientId:   recipientId,
	}

	reply := new(protos.EmailSettingsEmptyReply)
	err = c.CallWithContext(ctx, "resend_verification", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *EmailSettingsGrpcClientV1) VerifyEmail(ctx context.Context, correlationId string,
	recipientId string, code string) (err error) {
	timing := c.Instrument(ctx, correlationId, "emailsettings_v1.verify_email")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailSettingsCodeRequest{
		CorrelationId: correlationId,
		RecipientId:   recipientId,
		Code:          code,
	}

	reply := new(protos.EmailSettingsEmptyReply)
	err = c.CallWithContext(ctx, "verify_email", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}
