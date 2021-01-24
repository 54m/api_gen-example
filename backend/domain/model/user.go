package model

// User - user
type User struct {
	ID     string `json:"id"`
	Age    int    `json:"age"`
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}

// Gender - gender
type Gender int

const (
	GenderMale   Gender = iota + 1 // male
	GenderFemale                   // female
	GenderCustom                   // custom
)
