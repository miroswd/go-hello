// Convert dog age to human age
package dog

// One year old for a dog equals seven years old for a human.
const equivalentAge = 7;

// Function that multiplies the age of the dog by the equivalent age of a human being
func DogAgeToHuman(age int) int {
	return age * equivalentAge;
}
