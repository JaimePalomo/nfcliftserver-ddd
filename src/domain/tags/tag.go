package tags

// Tag struct para representar las etiquetas NFC
type Tag struct {
	Id     string `json:"id"`
	Rae    int    `json:"rae"`
	Planta string `json:"planta"`
}

func (tag Tag) IsFloor() bool {
	if tag.Planta != "" {
		return true
	}
	return false
}
