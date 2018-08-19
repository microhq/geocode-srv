// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/microhq/geocode-srv/proto/google/google.proto

/*
Package google is a generated protocol buffer package.

It is generated from these files:
	github.com/microhq/geocode-srv/proto/google/google.proto

It has these top-level messages:
	Point
	Bounds
	AddressComponent
	Geometry
	Result
	GeocodeRequest
	GeocodeResponse
	ReverseRequest
	ReverseResponse
*/
package google

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	api "github.com/micro/go-api"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Google service

type GoogleService interface {
	Geocode(ctx context.Context, in *GeocodeRequest, opts ...client.CallOption) (*GeocodeResponse, error)
	ReverseGeocode(ctx context.Context, in *ReverseRequest, opts ...client.CallOption) (*ReverseResponse, error)
}

type googleService struct {
	c    client.Client
	name string
}

func NewGoogleService(name string, c client.Client) GoogleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "google"
	}
	return &googleService{
		c:    c,
		name: name,
	}
}

func (c *googleService) Geocode(ctx context.Context, in *GeocodeRequest, opts ...client.CallOption) (*GeocodeResponse, error) {
	req := c.c.NewRequest(c.name, "Google.Geocode", in)
	out := new(GeocodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleService) ReverseGeocode(ctx context.Context, in *ReverseRequest, opts ...client.CallOption) (*ReverseResponse, error) {
	req := c.c.NewRequest(c.name, "Google.ReverseGeocode", in)
	out := new(ReverseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Google service

type GoogleHandler interface {
	Geocode(context.Context, *GeocodeRequest, *GeocodeResponse) error
	ReverseGeocode(context.Context, *ReverseRequest, *ReverseResponse) error
}

func RegisterGoogleHandler(s server.Server, hdlr GoogleHandler, opts ...server.HandlerOption) error {
	type google interface {
		Geocode(ctx context.Context, in *GeocodeRequest, out *GeocodeResponse) error
		ReverseGeocode(ctx context.Context, in *ReverseRequest, out *ReverseResponse) error
	}
	type Google struct {
		google
	}
	h := &googleHandler{hdlr}
	return s.Handle(s.NewHandler(&Google{h}, opts...))
	return s.Handle(s.NewHandler(&Google{h}, opts...))
}

type googleHandler struct {
	GoogleHandler
}

func (h *googleHandler) Geocode(ctx context.Context, in *GeocodeRequest, out *GeocodeResponse) error {
	return h.GoogleHandler.Geocode(ctx, in, out)
}

func (h *googleHandler) ReverseGeocode(ctx context.Context, in *ReverseRequest, out *ReverseResponse) error {
	return h.GoogleHandler.ReverseGeocode(ctx, in, out)
}
