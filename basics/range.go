package main
import "fmt"

func main() {
	// Using range to iterate over a slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	// Using range to iterate over a map
	person := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for key, value := range person {
		fmt.Printf("Name: %s, Age: %d\n", key, value)
	}
	// Using range to iterate over a string
	str := "hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}