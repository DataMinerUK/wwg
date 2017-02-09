package animals

import (
	"reflect"
	"testing"
)

// go test ./src/animals
func TestSetHobbiesSetsTheHobbies(t *testing.T) {
	hobbies := []string{"Sniffing"}
	dog := Dog{}
	dog.SetHobbies(hobbies)

	if !reflect.DeepEqual(dog.GetHobbies(), hobbies) {
		t.Error("Expected hobby to be sniffing")
	}
}

func TestDogBarks(t *testing.T) {
	noise := "bark"
	dog := Dog{}

	if dog.MakeNoise() != noise {
		t.Error("Expected dog to bark")
	}
}
