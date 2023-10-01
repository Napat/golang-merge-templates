package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTemplates(t *testing.T) {
	tests := []struct {
		name                string
		mainTemplateContent string
		template1Content    string
		template2Content    string
		expectedOutput      string
		expectedError       bool
	}{
		{
			name:                "Valid arguments",
			mainTemplateContent: "# Test Merge templates\n\n{{ template \"template_1.tpl\" }}\n{{ template \"template_2.tpl\" }}",
			template1Content:    "Template 1",
			template2Content:    "Template 2",
			expectedOutput:      "# Test Merge templates\n\nTemplate 1\nTemplate 2",
			expectedError:       false,
		},
		{
			name:                "Invalid main template",
			mainTemplateContent: "# Test Merge templates\n{{ template \"missing_template.tpl\" }}",
			template1Content:    "Template 1",
			template2Content:    "Template 2",
			expectedOutput:      "",
			expectedError:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testDataDir := t.TempDir()

			mainTemplateFile := createTempFile(t, testDataDir, "main.tpl", test.mainTemplateContent)
			template1File := createTempFile(t, testDataDir, "template_1.tpl", test.template1Content)
			template2File := createTempFile(t, testDataDir, "template_2.tpl", test.template2Content)

			tempOutputFile := filepath.Join(testDataDir, "test_output.html")

			err := mergeTemplates(tempOutputFile, mainTemplateFile, []string{template1File, template2File})
			if test.expectedError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			// Verify the output file
			expectedOutput := test.expectedOutput
			actualOutput, err := os.ReadFile(tempOutputFile)
			if err != nil {
				t.Fatalf("Error reading output file: %v", err)
			}

			assert.Equal(t, expectedOutput, string(actualOutput))
		})
	}
}

func createTempFile(t *testing.T, dir, fileName, content string) string {
	t.Helper()

	tempFile := filepath.Join(dir, fileName)
	err := os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}

	return tempFile
}
