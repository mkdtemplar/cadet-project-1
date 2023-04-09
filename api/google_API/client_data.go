package google_API

type Request struct {
	Origin      string
	Destination string
}

type Client struct {
	Apikey string
}

func New(apikey string) *Client {
	return &Client{Apikey: apikey}
}

// rename to client
func NewClientData(origin string, destination string, apikey string) *Request {
	return &Request{Origin: origin, Destination: destination}
}
