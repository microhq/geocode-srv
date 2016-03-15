package main

import (
	"log"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"

	"github.com/micro/geocode-srv/google"
	"github.com/micro/geocode-srv/handler"
	proto "github.com/micro/geocode-srv/proto/google"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.geocode"),
		micro.Flags(
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
		),
		micro.Action(func(c *cli.Context) {
			google.Key = c.String("google_api_key")
			google.ClientID = c.String("google_client_id")
			google.Signature = c.String("google_signature")
		}),
	)

	service.Init()

	proto.RegisterGoogleHandler(service.Server(), new(handler.Google))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
