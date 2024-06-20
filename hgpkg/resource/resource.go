package resource

import (
	"os"

	"github.com/google/uuid"
)

type ResourceState int

const (
	Initial ResourceState = iota
	Loading
	Loaded
	Released
)

type Resourcer interface {
	Load()
	Release()
}

type Resource struct {
	id    string
	path  string
	data  []byte
	state ResourceState
}

func NewResource(path string) *Resource {
	r := new(Resource)
	uuid, _ := uuid.NewRandom()
	r.id = uuid.String()
	r.path = path
	r.state = Initial
	return r
}

func (r *Resource) Load() {
	r.state = Loading
	bytes, err := os.ReadFile(r.path)
	r.data = bytes
	if err != nil {
		return
	}
	r.state = Loaded
}

func (r *Resource) Release() {
	r.data = nil
	r.state = Released
}

func (r *Resource) IsLoaded() bool {
	return r.state == Loaded
}
