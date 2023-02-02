package utils

import "github.com/zalando/go-keyring"

const (
	serviceName = "go-gpt3-chat-cli"
)

type KeyMag struct {
	Service string
}

func (k KeyMag) GetKey(name string) string {
	re, err := keyring.Get(k.Service, name)
	if err != nil {
		return ""
	}
	return re
}

func (k KeyMag) SetKey(name string, value string) {
	err := keyring.Set(k.Service, name, value)
	if err != nil {
		panic(err)
	}
}

func (k KeyMag) DelKey(name string) {
	err := keyring.Delete(k.Service, name)
	if err != nil {
		panic(err)
	}
}

func NewKeyMag() *KeyMag {
	return &KeyMag{
		Service: serviceName,
	}
}

type KeyMagInterface interface {
	GetKey(name string) string
	SetKey(name string, value string)
	DelKey(name string)
}

var _ KeyMagInterface = &KeyMag{}
