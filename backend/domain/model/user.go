package model

// User - user
type User struct {
	Age    int    `json:"age"`
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}

// Gender - gender
type Gender int

const (
	_            Gender = iota // ignore first value by assigning to blank identifier
	GenderMale                 // male
	GenderFemale               // female
	GenderCustom               // custom
)
