
// write a function that takes a single positive integer, and returns a collection of all positive divisors of the input integer

package main

import (
    "fmt"
)

func main(){
    input_test := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20, 1000000000}
    
    for _, i := range input_test {
        fmt.Printf("%d -> %v\n", i, GetDivisors(i))
    }
}

// GetDivisors takes a single positive integer and returns a slice of integers containing the positice divisors of the input. 
func GetDivisors(input int) []int {
    var divisors []int

    if input <= 0 {
        return divisors
    }

    for i := 1; i <= input; i++ {
        if input % i == 0 {
            divisors = append(divisors, i)
        } 
    }

    return divisors
} 