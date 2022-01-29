package model

const (
	json string = `{
    "Game": {
        "name": "value",
        "repo": "value",
        "website": "value",
        "go-version": "value",
        "ebiten-version": "value",
        "type": "value"
    },
    "dev": {
        "author": [],
        "translator": {},
        "constructor": {},
        "developer": {},
        "contributor": []
    }
}`
)

func init() {
	println(json)
}


type Project struct {
	Game struct {
		Name          string `json:"name"`
		Repo          string `json:"repo"`
		Website       string `json:"website"`
		GoVersion     string `json:"go-version"`
		EbitenVersion string `json:"ebiten-version"`
		Type          string `json:"type"`
	} `json:"Game"`
	Dev struct {
		Author     []string `json:"author"`
		Translator struct {
		} `json:"translator"`
		Constructor struct {
		} `json:"constructor"`
		Developer struct {
		} `json:"developer"`
		Contributor []string `json:"contributor"`
	} `json:"dev"`
}


