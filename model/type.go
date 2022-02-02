package model

const (
	ECSF1Json string = `{
    "game": {
        "name": "value",
        "repo": "value",
        "go-version": "value",
        "ebiten-version": "value",
        "type": "value",
        "version": "value",
        "language": [
            "value"
        ]
    },
    "dev": {
        "author": [
            "value"
        ],
        "translator": {},
        "constructor": {},
        "developer": {},
        "contributor": [
            "value"
        ],
        "use": [
            "value"
        ],
        "thanks": [
            "value"
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
        "cli-version": "a",
        "standard": "ecsf1"
    }
}`
	ECST1Json string = `{
    "game": {
        "name": "value",
        "repo": "value",
        "type": "value"
    },
    "project": {
        "statustype": 1,
        "cli-version": "a",
        "standard": "ecsf1"
    }
}`
	TmplJson string = `{
		"name": "value",
		"target": [
			"value"
		],
		"author": [
			"value"
		],
		"version": 0,
		"cli-version": 0,
		"ebiten-version": 0
	}`
)

func init() {
	//log.Println(ECSF1Json,ECSF1Json, TmplJson)
}

type ECSFProject struct {
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

type Tmpl struct {
	Name          string   `json:"name"`
	Target        []string `json:"target"`
	Author        []string `json:"author"`
	Version       string   `json:"version"`
	CliVersion    string   `json:"cli-version"`
	EbitenVersion string   `json:"ebiten-version"`
}

const StatustypeAchieve = 0
const StatustypeActive = 1
const StatustypeDebug = 2
const StatustypeRelease = 3
const StatustypePreRelease = 4
const StatustypeHold = 128
const StatustypeDelete = 404  // Just 404
const StatustypeFixing = 502  // just 503
const StatustypeRewrite = 233 //
const StatustypePause = 1409  // just
const StatustypeCBT = 114233  // close beta test
const StatustypePBT = 114514  // public beta test
