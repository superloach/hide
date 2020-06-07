package state

var Filename string = "state.json"

type State struct {
	Keys  *Keys
	Frame int
	Name  string
	Scene string
}
