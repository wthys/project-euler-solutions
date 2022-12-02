package problem61

import (
    "fmt"
    "strconv"
    "github.com/wthys/project-euler-solutions/solutions"
)

type solution struct{}

func init() {
    solutions.Register(solution{})
}

func IsCyclic(numbers []int) bool {
    l := len(numbers)

    for i, num := range numbers {
        if num % 100 != numbers[(i + 1) % l] {
            return false
        }
    }

     return true
}

func HasAllPolygonals(numbers []int) bool {
    var matched = make(map[int]int)
    var present = make(map[int]int)
    var polys = []int{3,4,5,6,7,8}

    for _, candidate := range numbers {
        for _, n := range polys {
            if ok, _ := IsNPolygonal(n, candidate); ok {
                _, exist := matched[candidate]
                if exist {
                    return false
                }
                matched[candidate] = n

                _, found := present[n]
                if found {
                    return false
                }
                present[n] = candidate
            }
        }
    }
    return true
}

func BuildCyclic(numbers []int) ([]int, error) {
    if len(numbers) <= 1 {
        return numbers, nil
    }

    var impossible = fmt.Errorf("cannot build cyclic set from %v", numbers)

    var cyclic = []int{numbers[0]}

    var candidates []int = numbers[1:]

    current := cyclic[0]

    for len(candidates) > 1 {
        var del = -1
        for i, num := range candidates {
            if num / 100 == current % 100 {
                cyclic = append(cyclic, num)
                del = i
                current = num
                break
            }
        }

        if del < 0 {
            return nil, impossible
        }

        last := len(candidates)-1
        candidates[del] = candidates[last]
        candidates[last] = 0
        candidates = candidates[:last]
    }

    num := candidates[0]
    if !(num / 100 == current % 100 && cyclic[0] / 100 == num % 100) {
        return nil, impossible
    }

    cyclic = append(cyclic, num)

    return cyclic, nil
}

func (s solution) Problem() string {
    return "61"
}

func (s solution) Solve(options solutions.Options) (string, error) {

    p3, _ := NPolygonalsBetween(3, 1000, 10000)
    p4, _ := NPolygonalsBetween(4, 1000, 10000)
    p5, _ := NPolygonalsBetween(5, 1000, 10000)
    p6, _ := NPolygonalsBetween(6, 1000, 10000)
    p7, _ := NPolygonalsBetween(7, 1000, 10000)
    p8, _ := NPolygonalsBetween(8, 1000, 10000)

    maxChecks := len(p3) * len(p4) * len(p5) * len(p6) * len(p7) * len(p8)
    fmt.Printf("Combinations to check: %v\n", maxChecks)

    checked := 0

    for _, n3 := range p3 {
        var present = map[int]bool{ n3: true }
        for _, n4 := range p4 {
            ok, found := present[n4]
            if found && ok { continue }

            present[n4] = true
            for _, n5 := range p5 {
                ok, found := present[n5]
                if found && ok { continue }

                present[n5] = true
                for _, n6 := range p6 {
                    ok, found := present[n6]
                    if found && ok { continue }

                    present[n6] = true
                    for _, n7 := range p7 {
                        ok, found := present[n7]
                        if found && ok { continue }

                        present[n7] = true
                        for _, n8 := range p8 {
                            ok, found := present[n8]
                            if found && ok { continue }

                            numbers := []int{n3, n4, n5, n6, n7, n8}

                            cyclic, err := BuildCyclic(numbers)

                            checked += 1

                            if err == nil {
                                fmt.Printf("Found cyclic set: %v\n", cyclic)
                                return strconv.Itoa(n3 + n4 + n5 + n6 + n7 + n8), nil
                            } else {
                                fmt.Printf("%12d/%d (%.6f%%) %v               \r", checked, maxChecks, 100.0 * float64(checked) / float64(maxChecks), err)
                            }

                        }
                        present[n7] = false
                    }
                    present[n6] = false
                }
                present[n5] = false
            }
            present[n4] = false
        }
    }

    return "", fmt.Errorf("no set found")
}
