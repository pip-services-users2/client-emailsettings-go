package version1

import "context"

type IEmailSettingsClientV1 interface {
	GetSettingsByIds(ctx context.Context, correlationId string,
		recipientIds []string) (result []*EmailSettingsV1, err error)

	GetSettingsById(ctx context.Context, correlationId string,
		recipientId string) (result *EmailSettingsV1, err error)

	GetSettingsByEmail(ctx context.Context, correlationId string,
		email string) (result *EmailSettingsV1, err error)

	SetSettings(ctx context.Context, correlationId string,
		settings *EmailSettingsV1) (result *EmailSettingsV1, err error)

	SetVerifiedSettings(ctx context.Context, correlationId string,
		settings *EmailSettingsV1) (result *EmailSettingsV1, err error)

	SetRecipient(ctx context.Context, correlationId string,
		recipientId string, name string, email string,
		language string) (result *EmailSettingsV1, err error)

	SetSubscriptions(ctx context.Context, correlationId string, recipientId string,
		subscriptions any) (result *EmailSettingsV1, err error)

	DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error

	ResendVerification(ctx context.Context, correlationId string, recipientId string) error

	VerifyEmail(ctx context.Context, correlationId string, recipientId string, code string) error
}
