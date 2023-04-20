package encryption

import (
	"github.com/goal-web/contracts"
	"sync"
)

type Manager struct {
	key        string
	drivers    map[string]contracts.EncryptDriver
	encryptors map[string]contracts.Encryptor
	mutex      sync.Mutex
}

func NewManager(key string, drivers map[string]contracts.EncryptDriver) contracts.EncryptManager {
	return &Manager{
		key:        key,
		drivers:    drivers,
		encryptors: make(map[string]contracts.Encryptor),
	}
}

func DefaultDrivers() map[string]contracts.EncryptDriver {
	return map[string]contracts.EncryptDriver{
		"AES": AES,
	}
}

func (factory *Manager) Extend(key string, provider contracts.EncryptDriver) {
	factory.mutex.Lock()
	defer factory.mutex.Unlock()
	factory.drivers[key] = provider
}

func (factory *Manager) Encryptor(key string) contracts.Encryptor {
	factory.mutex.Lock()
	defer factory.mutex.Unlock()
	if encryptor, ok := factory.encryptors[key]; ok {
		return encryptor
	}
	if driver, ok := factory.drivers[key]; ok {
		factory.encryptors[key] = driver(factory.key)
	}
	return factory.encryptors[key]
}

func (factory *Manager) Driver(driver string) contracts.EncryptDriver {
	return factory.drivers[driver]
}
