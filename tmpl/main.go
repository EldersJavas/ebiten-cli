// Created by EldersJavas(EldersJavas&gmail.com)

package tmpl

import (
	"encoding/json"
	"errors"
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

type TargetPT struct {
	Tmpl model.Tmpl
	path string
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
			// println(path + tool.DirFormat() + Tpath.Name() + tool.DirFormat() + "tmpl.json")
			PPPathJson := path + tool.DirFormat() + Tpath.Name() + tool.DirFormat() + "tmpl.json"
			if b, t, e := IsTmpl(PPPathJson); b {
				box = append(box, t)
			} else {
				log.Println(e)
			}
		}
	}
	return
}

// IsTmpl Check Tmpl
func IsTmpl(path string) (bool, model.Tmpl, error) {
	if tool.VaildFile(path) {
		var TTmpl model.Tmpl
		TTmplJ, err := os.ReadFile(path)
		if err != nil {
			return false, model.Tmpl{}, errors.New("unable to read file")
		}
		err = json.Unmarshal(TTmplJ, &TTmpl)
		if err != nil {
			return false, model.Tmpl{}, errors.New("unable to load JSON")
		}
		if TTmpl.CliVersion == model.Version {
			return true, TTmpl, nil
		} else if TTmpl.CliVersion != model.Version && TTmpl.CliVersion != "" {
			return false, model.Tmpl{}, errors.New("incompatible version")
		} else {
			return false, model.Tmpl{}, errors.New("not Ebiten-cli tmpl JSON file")
		}
	} else {
		return false, model.Tmpl{}, errors.New("not file")
	}
}

func (p TargetPT) CopyPjJSON() error {
	for _, s := range p.Tmpl.Target {
		if tool.GetFileBaseName(s)=="config"{
			f, err := os.Create("./output3.txt") //创建文件
			defer f.Close()
			n2, err := f.Write(d1) //写入文件(字节数组)
			n3, err := f.WriteString("writesn") //写入文件(字节数组)
			err = f.Sync()
			if err != nil {
				return err
			}
		}
		
	}
}

// StartTmplPrint StartTmpl
func (a App) StartTmplPrint(path string) error {
		
	
	return nil
}

func TsmplInit(app App) []model.Tmpl {
	return app.AllTmpl
	// parseFiles, err := template.ParseFiles()
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
