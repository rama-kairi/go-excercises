package exce

//  1. Write a function that takes a slice of integers and returns a new slice with only the even numbers.
//     Input: [1, 2, 3, 4, 5]
//     Output: [2, 4]
func FilterEvenNumbers(numbers []int) []int {
	var newNumbers []int

	for _, number := range numbers {
		if number%2 == 0 {
			newNumbers = append(newNumbers, number)
		}
	}
	return newNumbers
}
