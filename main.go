package main

import (
	"os"

	cli "github.com/urfave/cli"
)

func main() {
	var host string
	var port int
	app := cli.NewApp()
	app.Name = "echo-udp"
	app.Version = "1.0.0"
	app.Usage = "UDP Echo Server"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
			Value:       "127.0.0.1",
			Destination: &host,
			Usage:       "Host to bind",
		},
		cli.IntFlag{
			Name:        "port",
			Value:       12345,
			Destination: &port,
			Usage:       "Port to listen to",
		},
	}

	app.Action = func(c *cli.Context) error {
		runServer(host, port)
		return nil
	}

	app.Run(os.Args)
}
