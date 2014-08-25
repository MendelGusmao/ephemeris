package models

var Models = make([]interface{}, 0)

func register(model interface{}) {
	Models = append(Models, model)
}
