package main

import (
    "fmt"
    "context"
    "os"
    "errors"

    "github.com/urfave/cli/v2"

    "github.com/wthys/project-euler-solutions/solutions"
    _ "github.com/wthys/project-euler-solutions/solutions/impl"
)

func main() {
    ctx := context.Background()

    app := cli.NewApp()
    app.Name = "euler"
    app.Description = "Run Project Euler problems"
    app.Authors = []*cli.Author{
        {
            Name: "Wim Thys",
            Email: "wim.thys@zardof.be",
        },
    }

    app.CommandNotFound = notFound(ctx)
    app.Commands = commands(ctx)
    app.After = onExit(ctx)

    if err := app.Run(os.Args); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}

func commands(ctx context.Context) []*cli.Command {
    return []*cli.Command{
        {
            Name: "run",
            Usage: `run PROBLEM`,
            Action: cmdRun(ctx),
            Flags: cmdRunFlags(),
            SkipFlagParsing: false,
        },
        {
            Name: "list",
            Usage: `list`,
            Action: cmdList(ctx),
            Flags: cmdListFlags(),
            SkipFlagParsing: false,
        },
    }
}

var flagDebug = cli.BoolFlag{
            Name: "debug",
            Aliases: []string{"d"},
            Usage: "turns on debug mode",
            EnvVars: []string{"DEBUG"},
            Required: false,
            HasBeenSet: false,
        }

func debugFromContext(c *cli.Context) bool {
    if c.Bool("debug") || c.Bool("d") {
        return true
    }
    return false
}

func cmdRun(ctx context.Context) cli.ActionFunc {
    return func (c *cli.Context) error {
        problem := c.Args().First()

        solution, err := solutions.Get(problem)
        if err != nil {
            return err
        }

        var options = solutions.Options{}
        options.Debug = debugFromContext(c)

        var result, solveError = solution.Solve(options)
        if solveError != nil {
            if errors.Is(solveError, solutions.ErrNotImplemented) {
                return fmt.Errorf("Problem %q is not yet implemented", problem)
            }
            return solveError
        }

        fmt.Printf("%v\t%v\n", problem, result)

        return nil
    }
}

func cmdRunFlags() []cli.Flag {
    return []cli.Flag{ &flagDebug }
}

func cmdList(ctx context.Context) cli.ActionFunc {
    return func (c *cli.Context) error {
        for _, problem := range solutions.Available() {
            fmt.Println(problem)
        }
        return nil
    }
}

func cmdListFlags() []cli.Flag {
    return []cli.Flag{}
}

func onExit(ctx context.Context) cli.AfterFunc {
    return func (c *cli.Context) error {
        return nil
    }
}

func notFound(ctx context.Context) cli.CommandNotFoundFunc {
    return func (c *cli.Context, command string) {
        fmt.Fprintf(os.Stderr, "command %q not supported. Try --help flag to see how to use the program\n", command)
    }
}
