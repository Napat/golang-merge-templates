package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: golang-merge-templates <output_file> <main_template> <template_path_1> [<template_path_2> ...]")
		return
	}

	outputFile := os.Args[1]
	mainTemplatePath := os.Args[2]
	templatePaths := os.Args[3:]

	err := mergeTemplates(outputFile, mainTemplatePath, templatePaths)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Output file %s generated successfully!\n", outputFile)
}
