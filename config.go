package encryption

import "github.com/goal-web/contracts"

type Config struct {
	Default string

	Drivers map[string]contracts.EncryptDriver
}
