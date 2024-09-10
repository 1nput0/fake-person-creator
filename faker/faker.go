package faker

import "github.com/Pallinder/go-randomdata"

// basic Person struct for generating fast fake data
type Person struct {
	firstName string
	lastName  string
	email     string
	country   string
}

// constructor for creating a new male person with random fake data
func NewMalePerson() Person {
	newMale := Person{
		firstName: randomdata.FirstName(randomdata.Male),
		lastName:  randomdata.LastName(),
		email:     randomdata.Email(),
		country:   randomdata.Country(randomdata.FullCountry),
	}

	return newMale
}
