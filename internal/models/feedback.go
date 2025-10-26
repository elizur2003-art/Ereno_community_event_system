package models

type Feedback struct {
	ID      int
	EventID int
	Name    string
	Comment string
	Rating  int
}
