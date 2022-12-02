package problem61

import (
    "fmt"
)

type PolygonalFunc func (n int) int


func Triagonal(n int) int {
    return n*(n+1)/2
}

func Square(n int) int {
    return n*n
}

func Pentagonal(n int) int {
    return n*(3*n-1)/2
}

func Hexagonal(n int) int {
    return n*(2*n-1)
}

func Heptagonal(n int) int {
    return n*(5*n-3)/2
}

func Octagonal(n int) int {
    return n*(3*n-2)
}


var figuratives = map[int]PolygonalFunc{
    3: Triagonal,
    4: Square,
    5: Pentagonal,
    6: Hexagonal,
    7: Heptagonal,
    8: Octagonal,
}

func IsNPolygonal(n int, candidate int) (bool, error) {
    upper := 10000
    lower := 1000

    polys, err := NPolygonalsBetween(lower, upper, n)
    if err != nil {
        return false, err
    }

    for _, poly := range polys {
        if poly == candidate {
            return true, nil
        }
    }
    return false, nil

}

func NPolygonalsBetween(n int, lower int, upper int) ([]int, error) {
    poly, ok := figuratives[n]
    if !ok {
        return nil, fmt.Errorf("no %v-polygonal function", n)
    }

    i := 1
    numbers := make([]int, 0)

    for num := poly(i); num < upper; i, num = i+1, poly(i+1) {
        if num >= lower {
            numbers = append(numbers, num)
        }
    }

    return numbers, nil
}
