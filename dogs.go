package dogs

import "github.com/Angy41/puppy"

func Bark() string {
	return "Whoa"
}

func Barks() string {
	return "Whoa Whoa"
}

func BigBark() string {
	return puppy.WhenGrownUp(Bark())
}

func BigBarks() string {
	return puppy.WhenGrownUp(Barks())
}
