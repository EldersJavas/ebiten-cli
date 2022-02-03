// Created by EldersJavas(EldersJavas&gmail.com)

package model

import (
	"encoding/json"
	"os"
)

func (t Tmpl) init() {

}
func (t Tmpl) ReadFromJSON(path string) {
	n, _ := os.ReadFile(path)
	json.Unmarshal(n, &t)
}
