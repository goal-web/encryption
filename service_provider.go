package encryption

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
}

func NewService() contracts.ServiceProvider {
	return &ServiceProvider{}
}

func (provider ServiceProvider) Stop() {

}

func (provider ServiceProvider) Start() error {
	return nil
}

func (provider ServiceProvider) Register(container contracts.Application) {
	container.Singleton("encryption", func(config contracts.Config, env contracts.Env) contracts.EncryptorFactory {
		factory := &Factory{encryptors: make(map[string]contracts.Encryptor)}

		factory.Extend("default", AES(env.GetString("app.key")))

		return factory
	})

	container.Singleton("encryption.default", func(factory contracts.EncryptorFactory) contracts.Encryptor {
		return factory.Driver("default")
	})
}
