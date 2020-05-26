package listing

// Beer defines the properties of a beer to be added
type Beer struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	Abv       float32 `json:"abv"`
	ShortDesc string  `json:"short_description"`
	Created   string  `json:"created"`
}
