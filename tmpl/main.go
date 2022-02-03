// Created by EldersJavas(EldersJavas&gmail.com)

package tmpl

import (
	"encoding/json"
	"github.com/EldersJavas/ebiten-cli/cmd"
	"github.com/EldersJavas/ebiten-cli/cmd/tool"
	"github.com/EldersJavas/ebiten-cli/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var AllTmpl []model.Tmpl
var (
	TmplDir []string
)

func init()  {

}

// GetTaskOfDir 从项目目录读取任务
func GetTmplFromDir(path string)  {
	Tpathlist,_:=ioutil.ReadDir(path)
	for _,Tpath:=range Tpathlist{
		if Tpath.IsDir(){
			if tool.VaildFile(Tpath.Name()+tool.DirFormat()+"config.json"){
				var TTmpl model.Tmpl
				TTmplJ,_:=os.ReadFile(Tpath.Name()+tool.DirFormat()+"config.json")
				json.Unmarshal(TTmplC,&TTmpl)
			}else {
				log.Println(Tpath.Name()+" is not tmpl")}
		}
	}
}

func isTmpl(path string) bool {

}

func TmplInit() {
	parseFiles, err := template.ParseFiles()
	if err != nil {
		return
	}
	// 创建模板对象, parse关联模板
	//tmpl, err := template.New("test").Parse("{{.Name}} ID is {{ .ID }}")

	if err != nil {
		return
	}
	log.Fatal()
	// 渲染stu为动态数据, 标准输出到终端
	//err = parseFiles.Execute(os.Stdout, stu)
	//if err != nil {
	panic(err)
}
