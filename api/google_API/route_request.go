package google_API

import (
	"cadet-project/pkg/config"
	"context"
	"errors"

	"googlemaps.github.io/maps"
)

var totalDistance int

var stop = 0

func (rq *Request) FindRoute(request Request, mileage float64) ([]maps.Route, []GasStationsObject, error) {
	var gasStations []GasStationsObject
	config.InitConfig("pkg/config")
	c, err := maps.NewClient(maps.WithAPIKey(config.Config.MapsKey))
	if err != nil {
		return nil, []GasStationsObject{}, err
	}

	r := &maps.DirectionsRequest{
		Origin:      request.Origin,
		Destination: request.Destination,
	}

	routes, _, err := c.Directions(context.Background(), r)
	if err != nil {
		return nil, []GasStationsObject{}, err
	}

	polyline, err := decodePolyline(routes)
	if err != nil {
		return nil, []GasStationsObject{}, err
	}

	var sum = -1

	for i := 0; i < len(polyline)-1; i++ {
		distance, err := DistanceMatrix(polyline[i].Lat, polyline[i].Lng, polyline[i+1].Lat, polyline[i+1].Lng)
		if err != nil {
			return nil, []GasStationsObject{}, err
		}
		for _, row := range distance.Rows {
			for _, element := range row.Elements {
				sum += element.Distance.Value
				if sum/1000 >= (int(mileage) - (element.Distance.Value / 1000)) {
					lat := polyline[i].Lat
					lng := polyline[i].Lng
					gasStation, err := GasStations(lat, lng, element.Distance.Value)
					if gasStation.Results == nil || err != nil {
						return nil, []GasStationsObject{}, errors.New("route not possible no gas stations available")
					}
					stop++
					sum = 0
					for _, station := range gasStation.Results {
						if station.BusinessStatus == "OPERATIONAL" && station.OpeningHours.OpenNow == true {
							gasStations = append(gasStations, gasStation)
						}
					}
				}
			}
		}
	}
	if stop == -1 {
		stop = 0
	}

	totalDistance = ToRoutes(routes)[0].Legs[0].Distance.Value

	return routes, gasStations, nil
}
func decodePolyline(r []maps.Route) ([]maps.LatLng, error) {
	var decode []maps.LatLng
	var err error
	for _, route := range r {
		decode, err = route.OverviewPolyline.Decode()
		if err != nil {
			return nil, err
		}
	}
	return decode, nil
}

func GetStops() int {
	return stop
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
