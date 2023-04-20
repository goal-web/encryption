package encryption

import (
	"github.com/goal-web/application"
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
	container.Singleton("encryption", func(config contracts.Config, env contracts.Env) contracts.EncryptManager {
		appConfig := config.Get("app").(application.Config)
		manager := NewManager(appConfig.Key, DefaultDrivers())
		if encryptionConfig, ok := config.Get("encryption").(Config); ok {
			for key, driver := range encryptionConfig.Drivers {
				manager.Extend(key, driver)
			}
		}
		return manager
	})

	container.Singleton("encryption.default", func(config contracts.Config, factory contracts.EncryptManager) contracts.Encryptor {
		if encryptionConfig, ok := config.Get("encryption").(Config); ok {
			return factory.Encryptor(encryptionConfig.Default)
		}
		return factory.Encryptor("default")
	})
}
