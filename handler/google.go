package handler

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/micro/geocode-srv/google"
	proto "github.com/micro/geocode-srv/proto/google"
	"github.com/micro/go-micro/errors"

	"golang.org/x/net/context"
)

type Google struct{}

func (g *Google) Geocode(ctx context.Context, req *proto.GeocodeRequest, rsp *proto.GeocodeResponse) error {
	u := url.Values{}

	if len(req.Address) > 0 {
		u.Set("address", req.Address)
	}
	if len(req.Language) > 0 {
		u.Set("language", req.Language)
	}
	if len(req.Region) > 0 {
		u.Set("region", req.Region)
	}
	if req.Components != nil {
		var components []string
		for component, value := range req.Components {
			components = append(components, component+":"+value)
		}
		u.Set("components", strings.Join(components, "|"))
	}
	if req.Bounds != nil {
		var bounds []string
		bounds = append(bounds, fmt.Sprintf("%.6f,%.6f", req.Bounds.Northeast.Lat, req.Bounds.Northeast.Lng))
		bounds = append(bounds, fmt.Sprintf("%.6f,%.6f", req.Bounds.Southwest.Lat, req.Bounds.Southwest.Lng))
		u.Set("bounds", strings.Join(bounds, "|"))
	}

	b, err := google.Do("geocode", u)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.geocode.Google.Geocode", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}

func (g *Google) ReverseGeocode(ctx context.Context, req *proto.ReverseRequest, rsp *proto.ReverseResponse) error {
	u := url.Values{}

	if req.Latlng != nil {
		u.Set("latlng", fmt.Sprintf("%.6f,%.6f", req.Latlng.Lat, req.Latlng.Lng))
	}
	if len(req.PlaceId) > 0 {
		u.Set("place_id", req.PlaceId)
	}
	if len(req.PlaceID) > 0 {
		u.Set("placeID", req.PlaceID)
	}
	if len(req.Language) > 0 {
		u.Set("language", req.Language)
	}
	if req.ResultType != nil {
		u.Set("result_type", strings.Join(req.ResultType, "|"))
	}
	if req.LocationType != nil {
		u.Set("location_type", strings.Join(req.LocationType, "|"))
	}

	b, err := google.Do("geocode", u)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.geocode.Google.ReverseGeocode", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.slack", err.Error())
	}
	return nil
}
