package rpg

type createFormData struct {
	RPG string `json:"rpg"`
}

type responseError struct {
	Error string `json:"message"`
}
