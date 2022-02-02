# Ebiten-cli Project Standard Thin 1.0

> _**Waiting Translate...**_

# Ebiten-cli 项目标准 精简 1.0

[![GitHub license](https://img.shields.io/github/license/EldersJavas/ebiten-cli?logo=apache&logoColor=red&style=flat-square)](https://github.com/EldersJavas/ebiten-cli/blob/master/LICENSE)

## 目的
为了方便项目管理和生成,便于项目的查看,并解决`Ebiten-cli 项目标准 完整 v1.0`臃肿的问题,特此制定项目标准.

## 目标
精简项目配置,适配更多Ebiten项目.

## 适用版本
同 `Standard 1.0`, 见 ./README.md

## 内容

```text
Ebiten-cli 项目标准 精简 
v1.0, 2022-02-02

本标准仅适用于Ebiten-cli项目的0.1版本以上.
本标准使用过程中遵循 Apache License 2.0 许可.


1. 不完整宽容
  "Ebiten-cli 项目标准 精简 v1.0"(以下简称本标准)不强调项目的完整性.
  任何使用本标准的Ebiten项目无需完整.
  
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

6. 缺失宽容
  本标准允许处于本标准内的文件缺失.
  
7. 禁止向上迭代
  本标准禁止直接更新或适配最新的同名标准.

8. 精简性
  本标准强调项目的精简.
  
  
制定者:EldersJavas
2022年2月2日
```

# 示例

## json
```json
{
  "game": {
    "name": "game2022",
    "repo": "github.com/game/game2022",
    "type": "classic-monofile"
  },
  "project": {
    "statustype": 1,
    "cli-version": "1.0",
    "standard": "ecst1"
  }
}
```

# 定义

## Go struct

```go
type ECSTProject struct {
    Game struct {
        Name string `json:"name"`
        Repo string `json:"repo"`
        Type string `json:"type"`
    } `json:"game"`
    Project struct {
        Statustype int    `json:"statustype"`
        CliVersion string `json:"cli-version"`
        Standard   string `json:"standard"`
    } `json:"project"`
}
```

By EldersJavas