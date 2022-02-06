// Created by EldersJavas(EldersJavas&gmail.com)

package tmpl

import (
	"encoding/json"
	"github.com/EldersJavas/ebiten-cli/cmd/tool"
	"github.com/EldersJavas/ebiten-cli/model"
	"io/ioutil"
	"log"
	"os"
)

var Gobal App

type App struct {
	AllTmpl []model.Tmpl
	TmplDir []string
}

func init() {
	Gobal.AllTmpl = append(Gobal.AllTmpl, GetTmplFromDir(tool.GetAppRootDir()+tool.DirFormat()+"tmpl")...)
}

// GetTmplFromDir GetTmplOfDir
func GetTmplFromDir(path string) (box []model.Tmpl) {
	Tpathlist, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	log.Println(Tpathlist)
	for _, Tpath := range Tpathlist {
		if Tpath.IsDir() {
			//println(path + tool.DirFormat() + Tpath.Name() + tool.DirFormat() + "tmpl.json")
			if PPPathJson := path + tool.DirFormat() + Tpath.Name() + tool.DirFormat() + "tmpl.json"; tool.VaildFile(PPPathJson) {
				var TTmpl model.Tmpl
				TTmplJ, err := os.ReadFile(PPPathJson)
				//log.Println(TTmplJ)
				if err != nil {
					log.Println(err)
				}
				err = json.Unmarshal(TTmplJ, &TTmpl)
				if err != nil {
					log.Println(err)
				}
				if TTmpl.CliVersion != "" {
					box = append(box, TTmpl)
				}
			} else {
				log.Println(Tpath.Name() + " is not tmpl")
			}
		}
	}
	return
}

func isTmpl(path string) bool {
	return true
}

func TmplInit(app App) []model.Tmpl {
	return app.AllTmpl
	//parseFiles, err := template.ParseFiles()
	//if err != nil {
	//	return
	//}
	//// 创建模板对象, parse关联模板
	////tmpl, err := template.New("test").Parse("{{.Name}} ID is {{ .ID }}")
	//
	//if err != nil {
	//	return
	//}
	//log.Fatal()
	//// 渲染stu为动态数据, 标准输出到终端
	////err = parseFiles.Execute(os.Stdout, stu)
	////if err != nil {
	//panic(err)
}
