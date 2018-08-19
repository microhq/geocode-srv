# Geocode Server

Geocode server implements the Google Maps Geocoding API as a go-micro RPC service

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Get your an API key from Google - [doc](https://developers.google.com/maps/documentation/geocoding/get-api-key)

4. Download and start the service

	```shell
	go get github.com/microhq/geocode-srv
	geocode-srv --google_api_key=YOUR_API_TOKEN
	```

	OR as a docker container

	```shell
	docker run microhq/geocode-srv --google_api_key=YOUR_API_TOKEN --registry_address=YOUR_REGISTRY_ADDRESS
	```

## The API
Geocode server implements the [Google Geocoding API](https://developers.google.com/maps/documentation/geocoding) as RPC.

```shell
$ micro query go.micro.srv.geocode Google.Geocode '{"address": "1600 Amphitheatre Parkway"}'
{
	"results": [
		{
			"address_components": [
				{
					"long_name": "1600",
					"short_name": "1600",
					"types": [
						"street_number"
					]
				},
				{
					"long_name": "Amphitheatre Parkway",
					"short_name": "Amphitheatre Pkwy",
					"types": [
						"route"
					]
				},
				{
					"long_name": "Mountain View",
					"short_name": "Mountain View",
					"types": [
						"locality",
						"political"
					]
				},
				{
					"long_name": "Santa Clara County",
					"short_name": "Santa Clara County",
					"types": [
						"administrative_area_level_2",
						"political"
					]
				},
				{
					"long_name": "California",
					"short_name": "CA",
					"types": [
						"administrative_area_level_1",
						"political"
					]
				},
				{
					"long_name": "United States",
					"short_name": "US",
					"types": [
						"country",
						"political"
					]
				},
				{
					"long_name": "94043",
					"short_name": "94043",
					"types": [
						"postal_code"
					]
				}
			],
			"formatted_address": "1600 Amphitheatre Pkwy, Mountain View, CA 94043, USA",
			"geometry": {
				"location": {
					"lat": 37.4223607,
					"lng": -122.0841964
				},
				"location_type": "ROOFTOP",
				"viewport": {
					"northeast": {
						"lat": 37.42370968029149,
						"lng": -122.0828474197085
					},
					"southwest": {
						"lat": 37.4210117197085,
						"lng": -122.0855453802915
					}
				}
			}
		}
	],
	"status": "OK"
}
```
