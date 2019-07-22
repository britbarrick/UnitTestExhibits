package main

import (
	"testing"
)

//This test does its job but when it fails we don't get any information about why or how.
func Test_Dinner_CanHasCheese(t *testing.T) {
	tests := map[string]struct {
		entreeHasCheese bool
		dishHasCheese   bool
		expected        bool
	}{
		"Neither has cheese": {false, false, false},
		"Both has cheese":    {true, true, true},
		"Entree has cheese":  {true, false, true},
		"Dish has cheese":    {false, true, true},
	}

	for caseName, test := range tests {
		t.Run(caseName, func(t *testing.T) {
			mockEntree := entree{hasCheese: test.entreeHasCheese}
			mockDish := mainDish{hasCheese: test.dishHasCheese}
			actual := createDinner(mockEntree, salad{}, mockDish)

			if actual.hasCheese != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual.hasCheese)
			}
		})
	}
}
