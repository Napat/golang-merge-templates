package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	tests := []struct {
		name         string
		args         []string
		isError      bool
		expectedFile string
	}{
		{
			name:         "Happy case with correct number of arguments",
			args:         []string{"golang-merge-templates", "output.html", "sampledata/main.tpl", "sampledata/template_1.tpl", "sampledata/template_2.tpl"},
			isError:      false,
			expectedFile: "output.html",
		},
		{
			name:         "Incorrect number of arguments",
			args:         []string{"golang-merge-templates"},
			isError:      true,
			expectedFile: "",
		},
		{
			name:         "Incorrect template path",
			args:         []string{"golang-merge-templates", "output.html", "missing/main.tpl", "sampledata/template_1.tpl", "sampledata/template_2.tpl"},
			isError:      true,
			expectedFile: "",
		},
		{
			name:         "Incorrect output path",
			args:         []string{"golang-merge-templates", "missing/output.html", "sampledata/main.tpl", "sampledata/template_1.tpl", "sampledata/template_2.tpl"},
			isError:      true,
			expectedFile: "",
		},
	}

	oldArgs := os.Args
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if !tt.isError {
				// redirect output to automatically removed temp directory
				testDataDir := t.TempDir()
				outputFile := createTempFile(t, testDataDir, tt.args[1], "")
				tt.args[1] = outputFile
				tt.expectedFile = outputFile
			}

			os.Args = tt.args
			t.Log("os.Args: ", os.Args)

			// Create a new buffer to capture output
			ri, wi, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			_ = wi

			ro, wo, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}

			os.Stdin = ri
			os.Stdout = wo

			main()

			wo.Close()
			o, err := io.ReadAll(ro)
			if err != nil {
				log.Fatal(err)
			}
			got := string(o)
			expected := fmt.Sprintf("Output file %s generated successfully!\n", tt.expectedFile)

			if tt.isError {
				assert.NotEqual(t, expected, got)
			} else {
				assert.Equal(t, expected, got)
			}
		})
	}
}
