package main

import "fmt"

func createDinner(e entree, s salad, d mainDish) fullDinner {
	result := fullDinner{
		name: fmt.Sprintf("%q with %q salad and a starter of %q.", d.name, s.name, e.name),
		/// THIS NEEDS TESTING
	}

	// This is tested
	if e.hasCheese || d.hasCheese {
		result.hasCheese = true
	}

	/// THIS NEEDS TESTING
	if e.hasFruit || s.hasFruit {
		result.hasFruit = true
	}

	/// THIS NEEDS TESTING
	if s.isVegetarian && d.isVegetarian {
		result.hasFruit = true
	}

	return result
}
