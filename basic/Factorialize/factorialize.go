package factorialize

// ******** Factorialize Number ********
// Take an int and return factorial of it
// Just for now there is no need to check for errors
//

func Factorialize(n int) int {
	fact := 1
	for n > 1 {
		fact *= n
		n--
	}
	return fact
}

// ******** Factorialize With Errors ********
//
// Now let's try to do the same function, but check for errors
// Take a look at errors package and make brand New error
func FactorializeWithError(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Can't factorialize negative number")
	}
	fact := 1
	for n > 1 {
		fact *= n
		n--
	}
	return fact, nil
}