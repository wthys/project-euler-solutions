package solutions

import (
    "fmt"
    "errors"
)

var (
    ErrNotImplemented = errors.New("Not Implemented")
    
    registry = make(map[string]Solution)
)

type Options struct {
    Debug bool
}

type Solution interface {
    Solve(options Options) (string, error)
    Problem() string
}


func Register(solution Solution) error {
    problem := solution.Problem()
    if problem == "" {
        return fmt.Errorf("solution %T does not specify a problem", solution)
    }

    registry[problem] = solution

    return nil
}

func Get(problem string) (Solution, error) {
    if problem == "" {
        return nil, fmt.Errorf("problem not specified")
    }

    solution, exists := registry[problem]

    if !exists {
        return nil, fmt.Errorf("problem %q does not have a solution", problem)
    }

    return solution, nil
}

func Available() []string {
    problems := make([]string,0)
    for problem, _ := range registry {
        problems = append(problems, problem)
    }
    return problems
}
