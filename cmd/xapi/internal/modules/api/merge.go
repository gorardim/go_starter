package api

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type openapi struct {
	Openapi    string                 `yaml:"openapi"`
	Info       interface{}            `yaml:"info"`
	Tags       interface{}            `yaml:"tags"`
	Paths      map[string]interface{} `yaml:"paths"`
	Components struct {
		Schemas map[string]interface{} `yaml:"schemas"`
	} `yaml:"components"`
}

// MergeOpenApiYamlContent merge openapi yaml content
// impl path merge just now
func MergeOpenApiYamlContent(src, patch []byte) ([]byte, error) {
	// parse yml
	yaml1 := &openapi{}

	fmt.Println(string(src))

	if err := yaml.Unmarshal(src, yaml1); err != nil {
		return nil, err
	}

	yaml2 := &openapi{}
	if err := yaml.Unmarshal(patch, yaml2); err != nil {
		return nil, err
	}

	// merge paths
	for k, v := range yaml2.Paths {
		yaml1.Paths[k] = v
	}
	// merge components schemas
	for k, v := range yaml2.Components.Schemas {
		yaml1.Components.Schemas[k] = v
	}
	out, err := yaml.Marshal(yaml1)
	if err != nil {
		return nil, err
	}
	return out, nil
}
