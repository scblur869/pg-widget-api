package structs

// NeoHandler allows for passing the neo driver to receiver functions

type Widget struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Color    string `json:"color"`
}
