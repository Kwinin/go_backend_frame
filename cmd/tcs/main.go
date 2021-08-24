package main

import (
	"backend_ft_tcs/server"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	setLogLevel()
	setDefaultLocation()

	app := cli.NewApp()
	app.Version = server.AppVersion
	app.Name = server.AppName
	configPathFlag := cli.StringFlag{
		Name:   "configPath",
		Usage:  "config file path",
		EnvVar: "configPath",
	}
	portFlag := cli.StringFlag{
		Name:   "port",
		Usage:  "port",
		EnvVar: "port",
		Value:  "8081",
	}
	app.Flags = []cli.Flag{
		configPathFlag,
		portFlag,
	}
	app.Action = Start
	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func Start(ctx *cli.Context) error {
	sigCh := make(chan os.Signal)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	configPath := ctx.GlobalString("configPath")

	server := server.NewServer(server.AppName)

	go func() {
		select {
		case <-sigCh:
			if err := server.Close(); err != nil {
				logrus.Println(err.Error())
			}
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()
	err := server.Run(ctx.GlobalString("port"), configPath)
	if err != nil {
		return err
	}
	return nil
}

func setLogLevel() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "level",
		},
		TimestampFormat: time.RFC3339Nano,
	})
	logrus.SetLevel(logrus.InfoLevel)
}
func setDefaultLocation() {
	time.Local = time.FixedZone("UTC", 8*3600)
}
