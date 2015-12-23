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
	cmd.Flags = append(cmd.Flags,
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

	cmd.Actions = append(cmd.Actions, func(c *cli.Context) {
		google.Key = c.String("google_api_key")
		google.ClientID = c.String("google_client_id")
		google.Signature = c.String("google_signature")
	})

	cmd.Init()

	server.Init(
		server.Name("go.micro.srv.geocode"),
	)

	proto.RegisterGoogleHandler(server.DefaultServer, &handler.Google{})

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
