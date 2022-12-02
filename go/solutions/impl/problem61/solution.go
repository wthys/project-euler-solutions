package problem61

import (
    "fmt"
    "strconv"
    "golang.org/x/text/message"

    "github.com/wthys/project-euler-solutions/solutions"
)

var p = message.NewPrinter(message.MatchLanguage("en"))

func Printfln(format string, args ...any) {
    fmt.Println(p.Sprintf(format, args...))
}

func Printf(format string, args ...any) {
    fmt.Print(p.Sprintf(format, args...))
}

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


    l3, l4, l5, l6, l7, l8 := len(p3), len(p4), len(p5), len(p6), len(p7), len(p8)
    maxChecks := l3 * l4 * l5 * l6 * l7 * l8
    Printfln("Combinations to check: %d", maxChecks)

    checked := 0
    skipped := 0

    s4 := l3
    s5 := l4 * l3
    s6 := l5 * l4 * l3
    s7 := l6 * l5 * l4 * l3

    for i8, n8 := range p8 {
        var present = map[int]bool{ n8: true }
        for i7, n7 := range p7 {
            ok, found := present[n7]
            if found && ok { skipped += s7; continue }

            present[n7] = true
            for i6, n6 := range p6 {
                ok, found := present[n6]
                if found && ok { skipped += s6; continue }

                present[n6] = true
                for i5, n5 := range p5 {
                    ok, found := present[n5]
                    if found && ok { skipped += s5; continue }

                    present[n5] = true
                    for i4, n4 := range p4 {
                        ok, found := present[n4]
                        if found && ok { skipped += s4; continue }

                        present[n4] = true
                        for i3, n3 := range p3 {

                            ok, found := present[n3]
                            if found && ok { skipped += 1; continue }
                            checked += 1

                            numbers := []int{n3, n4, n5, n6, n7, n8}

                            cyclic, err := BuildCyclic(numbers)


                            if err == nil {
                                fmt.Printf("Found cyclic set: %v\n", cyclic)
                                return strconv.Itoa(n3 + n4 + n5 + n6 + n7 + n8), nil
                            } else {
                                Printf("%d/%d (%d + %d) %2d/%d %2d/%d %2d/%d %2d/%d %2d/%d %2d/%d               \r",
                                    (checked + skipped),
                                    (maxChecks),
                                    (checked),
                                    (skipped),
                                    i3+1, l3,
                                    i4+1, l4,
                                    i5+1, l5,
                                    i6+1, l6,
                                    i7+1, l7,
                                    i8+1, l8,
                                )
                            }

                        }
                        present[n4] = false
                    }
                    present[n5] = false
                }
                present[n6] = false
            }
            present[n7] = false
        }
    }

    return "", fmt.Errorf("no set found")
}
