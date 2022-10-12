package db

type Model interface {
}

var models map[string]Model

func init() {
	models = make(map[string]Model)
}

// Register register module
func Register(name string, m Model) {
	models[name] = m
}

func GetModels() map[string]Model {
	return models
}
