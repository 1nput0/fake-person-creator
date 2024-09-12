package faker

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

// basic Person struct for generating fast fake data
type Person struct {
	firstName string
	lastName  string
	email     string
	country   string
	state     string
	address   string
}

// constructor for creating a new male person with random fake data
func NewMalePerson() Person {
	newMale := Person{
		firstName: randomdata.FirstName(randomdata.Male),
		lastName:  randomdata.LastName(),
		email:     randomdata.Email(),
		country:   "USA",
		state:     randomdata.State(randomdata.Large),
		address:   randomdata.Address(),
	}

	return newMale
}

func NewFemalePerson() Person {
	newMale := Person{
		firstName: randomdata.FirstName(randomdata.Female),
		lastName:  randomdata.LastName(),
		email:     randomdata.Email(),
		country:   "USA",
		state:     randomdata.State(randomdata.Large),
		address:   randomdata.Address(),
	}

	return newMale
}

// prints the attributes of the person to the console
func (p Person) DisplayPerson() {
	displayInfo := fmt.Sprintf("\nFirst name: %s\nLast name: %s\nE-Mail: %s\nCountry: %s\nState: %s\nAddress: %s", p.firstName, p.lastName, p.email, p.country, p.state, p.address)
	fmt.Println(displayInfo)
}
