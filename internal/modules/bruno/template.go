package bruno

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/utils"
)

func getTemplateFile(fileName string) (*template.Template, error) {
	tmplFile := filepath.Join(
		"templates",
		"bruno",
		fmt.Sprintf("%s.tmpl", fileName),
	)

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Error("Error finding tmplFile",
			"file", tmplFile,
			"err", err,
		)
		return nil, err
	}

	return tmpl, nil
}

func CreateFileFromTemplate[T any](forwardURL string, templateName string, data T, outputName *string) error {
	slug := utils.Slugify(forwardURL)

	name := templateName
	if outputName != nil {
		name = *outputName
	}
	filePath := GetFilePath(slug, name)

	fileExists, err := CheckIfFileExists(filePath)
	if err != nil {
		log.Error("Error checking if file exists")
		return err
	}
	if fileExists {
		return nil
	}

	if err := CreateDirIfFileNotExists(filePath); err != nil {
		log.Error("Error creating directory", "path", filePath)
	}

	tmplFile, err := getTemplateFile(templateName)
	if err != nil {
		return nil
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		cerr := f.Close()
		if cerr != nil && err == nil {
			err = cerr
		}
	}()

	if err := tmplFile.Execute(f, data); err != nil {
		log.Error(err)
		return err
	}

	log.Printf("Created %s at %s", name, filePath)
	return nil
}
