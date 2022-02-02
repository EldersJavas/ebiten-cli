# Ebiten-cli Project Standard Full 1.0

> _**Waiting Translate...**_

# Ebiten-cli 项目标准 完整 1.0

[![GitHub license](https://img.shields.io/github/license/EldersJavas/ebiten-cli?logo=apache&logoColor=red&style=flat-square)](https://github.com/EldersJavas/ebiten-cli/blob/master/LICENSE)

## 目的
为了方便项目管理和生成,便于项目的查看,特此制定项目标准

## 目标
完整的功能性和可用性,考虑到一切因素

## 适用版本
同 `Standard 1.0`, 见 ./README.md

## 内容

```text
Ebiten-cli 项目标准 完整 
v1.0, 2022-02-02

本标准仅适用于Ebiten-cli项目的0.1版本以上.
本标准使用过程中遵循 Apache License 2.0 许可.


1. 完整性
  "Ebiten-cli 项目标准 完整 v1.0"(以下简称本标准)强调项目的完整性.
  任何使用本标准的Ebiten项目均应当确保项目的完整.
  使用任何工具生成本标准所规定的项目同样适用完整性.
  本标准不支持任何违背本标准的Ebiten项目.
  
2. 健壮性(请不要将robustness翻译成鲁棒性.)
  本标准强调项目整体和生成工具的健壮性.

3. 易读性
  本标准要求项目整体具有易读性.
  这也是本标准的设计目的.

4. 灵活性
  本标准应当可以适用于任何Ebiten框架.
  本标准可以与其他标准和工具并行.
  
5. 易变更性
  本标准的设计应便于更改项目配置和内容.

6. 禁止缺失
  本标准禁止任何处于本标准内的文件缺失.
  
7. 禁止向上迭代
  本标准禁止直接更新或适配最新的同名标准.


制定者:EldersJavas
2022年2月2日
```

# 示例

## json
```json
{
    "game": {
        "name": "game2022",
        "repo": "github.com/test/game2022",
        "go-version": "1.17",
        "ebiten-version": "2.2.4",
        "type": "classic-monofile",
        "version": "1.0",
        "language": [
            "en",
            "zh"
        ]
    },
    "dev": {
        "author": [
            "EldersJavas"
        ],
        "translator": {
          "zh":"EldersJavas"
        },
        "constructor": {
          "main":"EldersJavas"
        },
        "developer": {
          "Art":"EldersJavas"
        },
        "contributor": [
            "EldersJavas"
        ],
        "use": [
            "ebiten"
        ],
        "thanks": [
            "hajime hoshi"
        ]
    },
    "project": {
        "cli": {
            "autorewrite": false,
            "hotupdate": false,
            "autogomod": false,
            "autogofmt": false
        },
        "cgo": false,
        "makefile": "makefile",
        "statustype": 1,
        "cli-version": "1.0",
        "standard": "ecsf1"
    }
}
```

# 定义

## Go struct

```go
type Project struct {
	Game struct {
		Name          string   `json:"name"`
		Repo          string   `json:"repo"`
		GoVersion     string   `json:"go-version"`
		EbitenVersion string   `json:"ebiten-version"`
		Type          string   `json:"type"`
		Version       string   `json:"version"`
		Language      []string `json:"language"`
	} `json:"game"`
	Dev struct {
		Author      []string          `json:"author"`
		Translator  map[string]string `json:"translator"`
		Constructor map[string]string `json:"constructor"`
		Developer   map[string]string `json:"developer"`
		Contributor []string          `json:"contributor"`
		Use         []string          `json:"use"`
		Thanks      []string          `json:"thanks"`
	} `json:"dev"`
	Project struct {
		Cli struct {
			Autorewrite bool `json:"autorewrite"`
			Hotupdate   bool `json:"hotupdate"`
			Autogomod   bool `json:"autogomod"`
			Autogofmt   bool `json:"autogofmt"`
		} `json:"cli"`
		Cgo        bool   `json:"cgo"`
		Makefile   string `json:"makefile"`
		Statustype int    `json:"statustype"`
		CliVersion string `json:"cli-version"`
		Standard   string `json:"standard"`
	} `json:"project"`
}
```

By EldersJavas