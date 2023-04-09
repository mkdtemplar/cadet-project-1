package google_API

import (
	"cadet-project/pkg/config"
	"context"
	"fmt"

	"googlemaps.github.io/maps"
)

var totalDistance int

func (cl *Client) FindRoute(request Request) ([]Route, error) {

	c, err := maps.NewClient(maps.WithAPIKey(config.Config.MapsKey))
	if err != nil {
		return nil, err
	}

	r := &maps.DirectionsRequest{
		Origin:      request.Origin,
		Destination: request.Destination,
	}

	routes, _, err := c.Directions(context.Background(), r)
	if err != nil {
		return nil, err
	}

	//for i, j := range ToRoutes(routes) {
	//	fmt.Println("Total Distance: ", j.Legs[i].Distance.Value)
	//	legs := ToLegs(routes[i].Legs)
	//	fmt.Println("Printing legs: ", legs[i].Distance)
	//	fmt.Println("-----------------------")
	//	for k := range legs {
	//		steps := ToSteps(routes[i].Legs[k].Steps)
	//		for _, n := range steps {
	//			fmt.Println("Steps: ", n.Distance.Text)
	//			fmt.Println("Time: ", n.Duration.Seconds())
	//		}
	//	}
	//}

	totalDistance = ToRoutes(routes)[0].Legs[0].Distance.Value
	fmt.Println("Total distance: ", totalDistance)

	return ToRoutes(routes), nil
}

func GetTotalDistance() int {
	return totalDistance
}

func ToRoutes(routes []maps.Route) []Route {
	var output []Route
	for _, r := range routes {
		output = append(output, ToRoute(r))
	}
	return output
}

func ToRoute(route maps.Route) Route {
	return Route{
		Summary: route.Summary,
		Legs:    ToLegs(route.Legs),
	}
}

func ToLegs(legs []*maps.Leg) []Leg {
	var output []Leg
	for _, leg := range legs {
		output = append(output, ToLeg(leg))
	}
	return output
}

func ToLeg(leg *maps.Leg) Leg {
	return Leg{
		Steps: ToSteps(leg.Steps),
		Distance: Distance{
			Text:  leg.Distance.HumanReadable,
			Value: leg.Distance.Meters,
		},
		StartLocation: Location{
			Lat: leg.StartLocation.Lat,
			Lng: leg.StartLocation.Lng,
		},
		EndLocation: Location{
			Lat: leg.EndLocation.Lat,
			Lng: leg.EndLocation.Lng,
		},
		StartAddress: leg.StartAddress,
		EndAddress:   leg.EndAddress,
		Duration:     leg.Duration,
	}
}

func ToSteps(steps []*maps.Step) []Steps {
	var output []Steps
	for _, step := range steps {
		output = append(output, ToStep(step))
	}
	return output
}

func ToStep(step *maps.Step) Steps {
	return Steps{
		Distance: Distance{
			Text:  step.Distance.HumanReadable,
			Value: step.Distance.Meters,
		},
		StartLocation: Location{
			Lat: step.StartLocation.Lat,
			Lng: step.StartLocation.Lng,
		},
		EndLocation: Location{
			Lat: step.EndLocation.Lat,
			Lng: step.EndLocation.Lng,
		},
		Steps:    nil,
		Duration: step.Duration,
	}
}
