package psdui

// UIAst UI语法树
type UIAst struct {
	PSDID  int `json:"psdid"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`

	Children []interface{} `json:"children"`
}
