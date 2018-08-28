package operator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/openshift/machine-config-operator/pkg/operator/assets"
)

type renderConfig struct {
	TargetNamespace string
}

func renderAsset(config renderConfig, path string) ([]byte, error) {
	objBytes, err := assets.Asset(path)
	if err != nil {
		return nil, fmt.Errorf("error getting asset %s: %v", path, err)
	}

	tmpl, err := template.New(path).Parse(string(objBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to parse asset %s: %v", path, err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, config); err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.Bytes(), nil
}