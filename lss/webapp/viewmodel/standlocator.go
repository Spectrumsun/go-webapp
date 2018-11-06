package viewmodel

type StandLocator struct {
	Title  string
	Active string
}

func NewStandLocator() StandLocator {
	result := StandLocator{
		Active: "stand_locator",
		Title:  "Lemonade Stand Supply - Stand Locator",
	}
	return result
}
