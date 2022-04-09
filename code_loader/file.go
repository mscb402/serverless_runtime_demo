package code_loader

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

type FileLoader struct {
	dir string
}

func NewFileLoader(cfg *viper.Viper) (Loader, error) {
	return &FileLoader{
		dir: cfg.GetString("file_loader.dir"),
	}, nil
}

func (this *FileLoader) Load(hash string) ([]byte, error) {
	p := path.Join(this.dir, hash)

	data, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}
