
package main

import (
    "fmt"
    "math"
)

func main() {
    area := TriangleArea(10.0, 10.0, 10.0)

    if area < 0 {
        fmt.Println("Invalid lengths")
    } else {
        fmt.Println(area)
    }

}

// TriangleArea takes three integers representing sides of a triangle and returns the area 
// Returns <0 value if lengths cannot form a triangle or negative lengths
func TriangleArea(edge1, edge2, edge3 float64) float64 {
    if edge1 < 0 || edge2 < 0 || edge3 < 0 {
        return -1
    }

    // if any two edges add up to or less than the other edge, a triangle cannot be formed
    if ((edge1 + edge2) <= edge3) || ((edge3 + edge1) <= edge2) || ((edge2 + edge3) <= edge1) {
        return -1
    }

    // calculate and return are using Heron's formula
    perimeter := (edge1 + edge2 + edge3)/2
    return math.Sqrt(perimeter * (perimeter - edge1) * (perimeter - edge2) * (perimeter - edge3))
}