// Created by EldersJavas(EldersJavas&gmail.com)

package tmpl

import (
	"log"
	"os"
	"text/template"
)

func main() {
	parseFiles, err := template.ParseFiles()
	if err != nil {
		return
	}
	// 创建模板对象, parse关联模板
	//tmpl, err := template.New("test").Parse("{{.Name}} ID is {{ .ID }}")
	//parseFiles.Execute(,)
	log.Fatal()
	// 渲染stu为动态数据, 标准输出到终端
	//err = parseFiles.Execute(os.Stdout, stu)
	//if err != nil {
	panic(err)
}
