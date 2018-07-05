package egret

type IComponent interface {
	Parse() error
}

// EgretComponent 组件类
type EgretComponent struct {
}

// Parse 解析
func (comp *EgretComponent) Parse() {
}
