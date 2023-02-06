package directions

import (
	"cadet-project/pkg/config"
	"context"
	"log"

	"googlemaps.github.io/maps"
)

func GetDirections(start string, end string) []maps.Route {
	config.InitConfig("configurations")

	c, err := maps.NewClient(maps.WithAPIKey(config.Config.MapsKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      start,
		Destination: end,
	}

	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return route
}
