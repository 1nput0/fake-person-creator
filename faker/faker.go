package faker

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

// basic Person struct for generating fast fake data
type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	State     string `json:"state"`
	Address   string `json:"address"`
}

// constructor for creating a new male person with random fake data
func NewMalePerson() Person {
	newMale := Person{
		FirstName: randomdata.FirstName(randomdata.Male),
		LastName:  randomdata.LastName(),
		Email:     randomdata.Email(),
		Country:   "USA",
		State:     randomdata.State(randomdata.Large),
		Address:   randomdata.Address(),
	}

	return newMale
}

func NewFemalePerson() Person {
	newMale := Person{
		FirstName: randomdata.FirstName(randomdata.Female),
		LastName:  randomdata.LastName(),
		Email:     randomdata.Email(),
		Country:   "USA",
		State:     randomdata.State(randomdata.Large),
		Address:   randomdata.Address(),
	}

	return newMale
}

// prints the attributes of the person to the console
func (p Person) DisplayPerson() {
	displayInfo := fmt.Sprintf("\nFirst name: %s\nLast name: %s\nE-Mail: %s\nCountry: %s\nState: %s\nAddress: %s", p.FirstName, p.LastName, p.Email, p.Country, p.State, p.Address)
	fmt.Println(displayInfo)
}
