package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/cli"
	"./config"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "Location of configuration file",
	},
}

func main() {
	NewApp().Run(os.Args)
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "go-events-manager"
	app.Usage = ""
	app.Authors = []cli.Author{
		{"Marcin Malessa", "marcin@malessa.pl"},
	}
	app.Version = "0.0.1"
	app.Flags = flags
	app.Action = Action
	app.ExitErrHandler = ExitErrHandler
	return app
}

// Action is the function being run when the application gets executed.
func Action(c *cli.Context) error {
	cfg, err := LoadConfiguration(c);
	if err != nil {
		return err
	}
	fmt.Println("It working...")
	fmt.Println(cfg.RabbitMq.Host);
	return mainFunction()
}

// ExitErrHandler is a global error handler registered with the application.
func ExitErrHandler(_ *cli.Context, err error) {
	if err == nil {
		return
	}
	code := 1
	if exitErr, ok := err.(cli.ExitCoder); ok {
		code = exitErr.ExitCode()
	}

	os.Exit(code)
}

func LoadConfiguration(c *cli.Context) (*config.Config, error) {
	file := c.String("config")
	if file == "" {
		file = "dist.conf"
	}
	cfg, err := config.LoadAndParse(file);
	if err != nil {
		return nil, fmt.Errorf("failed parsing configuration: %s", err)
	}
	return cfg, nil
}

func mainFunction() error {
	done := make(chan error)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		fmt.Println(ctx)
		fmt.Println("go Func pending")
		// done <- doSomething()
		//done <- client.Consume(ctx)
	}()

	fmt.Println("Next line command")

	select {
		case <-sig:
			fmt.Println("Cancel consumption of messages.")
			cancel()
			return nil

		case <-done:
			fmt.Println("Done...")
			return nil
	}
}
