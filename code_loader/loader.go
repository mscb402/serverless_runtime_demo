package code_loader

import (
	"errors"
	"github.com/spf13/viper"
)

type Loader interface {
	Load(funcName string) ([]byte, error)
}

func NewLoader(loaderName string, cfg *viper.Viper) (Loader, error) {
	switch loaderName {
	case "file":
		loader, err := NewFileLoader(cfg)
		if err != nil {
			return nil, err
		}
		return loader, nil
	}
	return nil, errors.New("loader not found")
}
