package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/kadel/devfile-utils/registry"
	"github.com/redhat-developer/alizer/go/pkg/apis/recognizer"
)

const DevfileRegistryUrl = "https://registry.devfile.io"

func printJson(data interface{}) {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

func main() {

	path := flag.String("path", "./", "Directory to analyze")
	outputFormat := flag.String("o", "", "output format")
	flag.Parse()

	if path == nil {
		panic("No path specified")
	}

	if outputFormat == nil {
		panic("No output format specified")
	}

	fmt.Printf("# Runing recognizer.Analyzer(%q)\n", *path)
	languages, err := recognizer.Analyze(*path)
	if err != nil {
		panic(err)
	}

	for _, lang := range languages {
		fmt.Printf("%+v\n", lang)
	}

	fmt.Println()
	fmt.Printf("# Runing recognizer.SelectDevFileFromTypes(%q)\n", *path)

	devfileRegistry := registry.NewIndex(DevfileRegistryUrl)

	types := []recognizer.DevFileType{}
	for _, d := range devfileRegistry.GetIndex() {
		types = append(types, recognizer.DevFileType{
			Name:        d.Name,
			ProjectType: d.ProjectType,
			Language:    d.Language,
			Tags:        d.Tags,
		})
	}

	detectedType, err := recognizer.SelectDevFileFromTypes(*path, types)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", detectedType)

	fmt.Println()
	fmt.Printf("# Runing recognizer.DetectComponents(%q)\n", *path)
	components, err := recognizer.DetectComponents(*path)
	if err != nil {
		panic(err)
	}

	for _, component := range components {
		fmt.Printf("%+v\n", component)
	}
}
