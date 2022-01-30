package model

const (
	ProjectJson string = `{
    "game": {
        "name": "value",
        "repo": "value",
        "website": "value",
        "go-version": "value",
        "ebiten-version": "value",
        "type": "value",
        "field": "value",
        "version": "value",
        "setting": {
            "GameWidth": 0,
            "GameHeight": 0,
            "WindowWidth": 0,
            "WindowHeight": 0,
            "GameTitle": "value",
            "field": "value"
        },
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
        "statustype": 1
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
	println(ProjectJson, TmplJson)
}

type Project struct {
	Game struct {
		Name          string `json:"name"`
		Repo          string `json:"repo"`
		Website       string `json:"website"`
		GoVersion     string `json:"go-version"`
		EbitenVersion string `json:"ebiten-version"`
		Type          string `json:"type"`
		Field         string `json:"field"`
		Version       string `json:"version"`
		Setting       struct {
			GameWidth    int    `json:"GameWidth"`
			GameHeight   int    `json:"GameHeight"`
			WindowWidth  int    `json:"WindowWidth"`
			WindowHeight int    `json:"WindowHeight"`
			GameTitle    string `json:"GameTitle"`
			Field        string `json:"field"`
		} `json:"setting"`
		Language []string `json:"language"`
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
		Cgo        bool   `json:"cgo"`
		Makefile   string `json:"makefile"`
		Statustype int    `json:"statustype"`
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
