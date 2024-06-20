package control

type Controller interface {
	IsPressed(id int) bool
}

type Default struct {
}

func (d *Default) IsPressed(id int) {

}
