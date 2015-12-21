package main

import (
	"github.com/codegangsta/cli"
	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"

	"github.com/micro/geocode-srv/google"
	"github.com/micro/geocode-srv/handler"
	proto "github.com/micro/geocode-srv/proto/google"
)

func main() {
	app := cli.NewApp()
	app.HideVersion = true
	app.Flags = cmd.Flags
	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "google_api_key",
			EnvVar: "GOOGLE_API_KEY",
			Usage:  "Google maps API key",
		},
		cli.StringFlag{
			Name:   "google_client_id",
			EnvVar: "GOOGLE_CLIENT_ID",
			Usage:  "Google client id",
		},
		cli.StringFlag{
			Name:   "google_signature",
			EnvVar: "GOOGLE_SIGNATURE",
			Usage:  "Google signature",
		},
	)
	app.Before = cmd.Setup
	app.Action = func(c *cli.Context) {
		google.Key = c.String("google_api_key")
		google.ClientID = c.String("google_client_id")
		google.Signature = c.String("google_signature")
	}
	app.RunAndExitOnError()

	server.Init(
		server.Name("go.micro.srv.geocode"),
	)

	proto.RegisterGoogleHandler(server.DefaultServer, &handler.Google{})

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
