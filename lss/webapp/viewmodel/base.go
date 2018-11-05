package viewmodel

// Base struct
type Base struct {
	Title string
}

// NewBase function
func NewBase() Base {
	return Base{
		Title: "Lemonade Stand Supply",
	}
}
