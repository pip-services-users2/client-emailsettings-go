package version1

import (
	"encoding/json"

	"github.com/pip-services-users2/client-emailsettings-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := make(map[string]string, 0)

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]any {
	r := make(map[string]any, 0)

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromEmailSettings(settings *EmailSettingsV1) *protos.EmailSettings {
	if settings == nil {
		return nil
	}

	obj := &protos.EmailSettings{
		Id:            settings.Id,
		Name:          settings.Name,
		Email:         settings.Email,
		Language:      settings.Language,
		Subscriptions: toJson(settings.Subscriptions),
		Verified:      settings.Verified,
		VerCode:       settings.VerCode,
		VerExpireTime: convert.StringConverter.ToString(settings.VerExpireTime),
		CustomHdr:     toJson(settings.CustomHdr),
		CustomDat:     toJson(settings.CustomDat),
	}

	return obj
}

func toEmailSettings(obj *protos.EmailSettings) *EmailSettingsV1 {
	if obj == nil {
		return nil
	}

	settings := &EmailSettingsV1{
		Id:            obj.Id,
		Name:          obj.Name,
		Email:         obj.Email,
		Language:      obj.Language,
		Subscriptions: fromJson(obj.Subscriptions),
		Verified:      obj.Verified,
		VerCode:       obj.VerCode,
		VerExpireTime: convert.DateTimeConverter.ToDateTime(obj.VerExpireTime),
		CustomHdr:     fromJson(obj.CustomHdr),
		CustomDat:     fromJson(obj.CustomDat),
	}

	return settings
}

func fromEmailSettingsList(settings []*EmailSettingsV1) []*protos.EmailSettings {
	if settings == nil {
		return nil
	}

	data := make([]*protos.EmailSettings, len(settings))

	for i, v := range settings {
		data[i] = fromEmailSettings(v)
	}

	return data
}

func toEmailSettingsList(obj []*protos.EmailSettings) []*EmailSettingsV1 {
	if obj == nil {
		return nil
	}

	settings := make([]*EmailSettingsV1, len(obj))

	for i, v := range obj {
		settings[i] = toEmailSettings(v)
	}

	return settings
}
