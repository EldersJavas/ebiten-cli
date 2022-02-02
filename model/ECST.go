// Created by EldersJavas(EldersJavas&gmail.com)

package model

import (
	"encoding/json"
)

func NewECSTProjectFromJSON(jsonc []byte) (ECSTProject, error) {
	b := ECSTProject{}
	err := json.Unmarshal(jsonc, &b)
	return b, err
}

func (p ECSTProject) SaveToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p ECSTProject) ReadFromJSON(jsonc []byte) error {
	return json.Unmarshal(jsonc, &p)
}

func (p ECSTProject) init() {
	p.Project.CliVersion = Version
	p.Project.Standard = "ecst1"
}
