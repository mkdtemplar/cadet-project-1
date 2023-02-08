package google_API

import (
	"cadet-project/pkg/config"
	"context"
	"log"

	"googlemaps.github.io/maps"
)

func (cl *ClientData) FindRoute() []maps.Route {
	config.InitConfig("configurations")

	c, err := maps.NewClient(maps.WithAPIKey(config.Config.MapsKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil
	}

	r := &maps.DirectionsRequest{
		Origin:      cl.Origin,
		Destination: cl.Destination,
	}

	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil
	}

	return route
}
