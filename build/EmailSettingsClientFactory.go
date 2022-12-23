package build

import (
	clients1 "github.com/pip-services-users2/client-emailsettings-go/version1"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type EmailSettingsClientFactory struct {
	*cbuild.Factory
}

func NewEmailSettingsClientFactory() *EmailSettingsClientFactory {
	c := &EmailSettingsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	memoryClientDescriptor := cref.NewDescriptor("service-emailsettings", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-emailsettings", "client", "commandable-http", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-emailsettings", "client", "commandable-grpc", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-emailsettings", "client", "grpc", "*", "1.0")

	c.RegisterType(memoryClientDescriptor, clients1.NewEmailSettingsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewEmailSettingsCommandableHttpClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewEmailSettingsCommandableGrpcClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewEmailSettingsGrpcClientV1)

	return c
}
