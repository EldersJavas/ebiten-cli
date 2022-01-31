// Created by EldersJavas(EldersJavas&gmail.com)

package model

import (
	"encoding/json"
)

func NewProjectFromJSON(jsonc []byte) (Project, error) {
	b := Project{}
	err := json.Unmarshal(jsonc, &b)
	return b, err
}

func (p Project) SaveToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p Project) ReadFromJSON(jsonc []byte) error {
	return json.Unmarshal(jsonc, &p)
}

func (p Project) init() {
	p.Project.Cli.Hotupdate = false
	p.Project.CliVersion= Version
}
