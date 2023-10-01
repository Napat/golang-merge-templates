package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func mergeTemplates(outputFile, mainTemplatePath string, templatePaths []string) error {
	// Parse all the templates at once
	allTemplatePaths := append([]string{mainTemplatePath}, templatePaths...)
	mergedTemplate, err := template.ParseFiles(allTemplatePaths...)
	if err != nil {
		return fmt.Errorf("error parsing templates: %v", err)
	}

	// Execute the merged template with data and store the output in a buffer
	var outputContent bytes.Buffer
	if err := mergedTemplate.Execute(&outputContent, nil); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	// Write the output buffer to the output file
	err = os.WriteFile(outputFile, outputContent.Bytes(), 0600)
	if err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}

	return nil
}
