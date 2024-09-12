package entities

type Person struct {
	name string
	age  int // private
}

func (person Person) GetAge() int {
	return person.age
}

func (person Person) GetName() string {
	return person.name
}
