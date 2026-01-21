package bruno

import (
	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/config"
)

type CollectionConfig struct {
	BaseURL     string
	BearerToken *string
}

func CreateCollectionBruIfNotExists(cfg config.Config) error {
	data := CollectionConfig{
		BaseURL: cfg.ForwardAddr,
	}

	if err := CreateFileFromTemplate(
		cfg.ForwardAddr,
		"collection.bru",
		data,
		nil,
	); err != nil {
		log.Error("Error creating collection.bru file")
	}
	return nil
}
