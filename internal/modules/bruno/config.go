package bruno

import (
	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/config"
)

type BrunoConfig struct {
	Version string
	Name    string
	Type    string
	Ignore  []string
}

func CreateBrunoConfigIfNotExists(cfg config.Config) error {
	data := BrunoConfig{
		Version: "1",
		Name:    cfg.ForwardAddr,
		Type:    "collection",
		Ignore:  []string{"node_modules", ".git"},
	}

	if err := CreateFileFromTemplate(
		cfg.ForwardAddr,
		"bruno.json",
		data,
		nil,
	); err != nil {
		log.Error("Error creating bruno.json file")
	}

	return nil
}
