// Created by EldersJavas(EldersJavas&gmail.com)

package model

import (
	"encoding/json"
)

func NewECSFProjectFromJSON(jsonc []byte) (ECSFProject, error) {
	b := ECSFProject{}
	err := json.Unmarshal(jsonc, &b)
	return b, err
}

func (p ECSFProject) SaveToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p ECSFProject) ReadFromJSON(jsonc []byte) error {
	return json.Unmarshal(jsonc, &p)
}

func (p ECSFProject) init() {
	p.Project.CliVersion = Version
	p.Project.Cgo = false
	p.Project.Standard = "ecsf1"
	p.Project.Cli = struct {
		Autorewrite bool `json:"autorewrite"`
		Hotupdate   bool `json:"hotupdate"`
		Autogomod   bool `json:"autogomod"`
		Autogofmt   bool `json:"autogofmt"`
	}(struct {
		Autorewrite bool
		Hotupdate   bool
		Autogomod   bool
		Autogofmt   bool
	}{Autorewrite: false, Hotupdate: false, Autogomod: false, Autogofmt: false})

}
