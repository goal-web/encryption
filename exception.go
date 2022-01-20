package encryption

import (
	"github.com/goal-web/contracts"
)

type EncryptException struct {
	error
	fields contracts.Fields
}

func (e EncryptException) Error() string {
	return e.error.Error()
}

func (e EncryptException) Fields() contracts.Fields {
	return e.fields
}
