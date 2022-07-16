package models

type Model interface {
	Insert()
	Update()
	Select()
}

type Default struct{}

func (d *Default) Insert() {}
func (d *Default) Update() {}
func (d *Default) Select() {}
