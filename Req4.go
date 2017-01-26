
package main

import (
    "fmt"
)

func main() {
    inputList := []int{5, 4, 3, 2, 4, 5, 1, 6, 1, 2, 5, 4}
    // inputList := []int{1, 2, 3, 4, 5, 1, 6, 7}
    // inputList := []int{1, 2, 3, 4, 5, 6, 7}
    // inputList := []int{}
    // inputList := []int{"a", "b", 3, 4}
    // nputList = nil

    fmt.Println(MostCommon(inputList))
}

// MostCommon takes in a list of integers and returns another list that contains those integers which are most common 
// on the input array
func MostCommon(input []int) []int {
    // use an int-to-int map (https://blog.golang.org/go-maps-in-action) to keep track of occurence frequencies of integers
    // for every integer in the input list, add as a key with value 1 if not existing in the map or add key with value 1
    freqTracker := make(map[int]int)
    repeatKeys := false // flag indicating if there are any integers with a frequency >1
    topFreq := 1    // highest frequency of an integer recorded

    for _, i := range input {
        if freqTracker[i] == 0 {
            // integer is not tracked for frequency, add
            freqTracker[i] = 1
        } else {
            // already tracked - increment count
            freqTracker[i]++
            repeatKeys = true
            if freqTracker[i] > topFreq {
                topFreq = freqTracker[i]
            }
        }
    }

    // no integers were repeated, then return the input as-is
    if !repeatKeys {
        return input
    }

    // compile a list of integers with the highest frequency of occurence and return it
    var topFreqList []int
    for k, v := range freqTracker {
        if v == topFreq {
            topFreqList = append(topFreqList, k)
        }
    }

    return topFreqList
}
