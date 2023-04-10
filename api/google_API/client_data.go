package google_API

type Request struct {
	Origin      string
	Destination string
}

func NewClientRequest(origin string, destination string) *Request {
	return &Request{Origin: origin, Destination: destination}
}
