package google_API

type ClientData struct {
	Origin      string
	Destination string
}

func NewClientData(origin string, destination string) *ClientData {
	return &ClientData{Origin: origin, Destination: destination}
}
