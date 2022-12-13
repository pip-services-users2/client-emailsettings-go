package version1

import (
	"time"
)

type EmailSettingsV1 struct {
	/* Recipient information */
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Language string `json:"language"`

	/* EmailSettings management */
	Subscriptions any       `json:"subscriptions"`
	Verified      bool      `json:"verified"`
	VerCode       string    `json:"ver_code"`
	VerExpireTime time.Time `json:"ver_expire_time"`

	/* Custom fields */
	CustomHdr any `json:"custom_hdr"`
	CustomDat any `json:"custom_dat"`
}

func NewEmptyEmailSettingsV1() *EmailSettingsV1 {
	return &EmailSettingsV1{}
}

func NewEmailSettingsV1(recipientId string, name string, email string, language string) *EmailSettingsV1 {
	return &EmailSettingsV1{
		Id:       recipientId,
		Name:     name,
		Email:    email,
		Language: language,
		Verified: false,
	}
}
