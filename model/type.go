package model

import "log"

const (
	ProjectJson string = `{
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
        "makefile": "value",
        "statustype": 1,
		"cli-version": ""
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
	log.Println(ProjectJson, TmplJson)
}

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
		Author     []string `json:"author"`
		Translator struct {
		} `json:"translator"`
		Constructor struct {
		} `json:"constructor"`
		Developer struct {
		} `json:"developer"`
		Contributor []string `json:"contributor"`
		Use         []string `json:"use"`
		Thanks      []string `json:"thanks"`
	} `json:"dev"`
	Project struct {
		Cli struct {
			Autorewrite bool `json:"autorewrite"`
			Hotupdate   bool `json:"hotupdate"`
			Autogomod   bool `json:"autogomod"`
			Autogofmt   bool `json:"autogofmt"`
		} `json:"cli"`
		Cgo          bool   `json:"cgo"`
		Makefile     string `json:"makefile"`
		Statustype   int    `json:"statustype"`
		CliVersion   string `json:"cli-version"`
		StandVersion string `json:"Standard-Version"`
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

const (
	StatustypeAchieve    = 0
	StatustypeActive     = 1
	StatustypeDebug      = 2
	StatustypeRelease    = 3
	StatustypePreRelease = 4
	StatustypeHold       = 128
	StatustypeDelete     = 404    //Just 404
	StatustypeFixing     = 502    //just 503
	StatustypeRewrite    = 233    //
	StatustypePause      = 1409   //just
	StatustypeCBT        = 114233 //close beta test
	StatustypePBT        = 114514 //public beta test

)
