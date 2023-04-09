package google_API

import "time"

type GoogleRouteResponse struct {
	Legs []Leg `json:"legs"`
}

type Route struct {
	Summary string
	Legs    []Leg
}

type Leg struct {
	Steps         []Steps       `json:"steps"`
	Distance      Distance      `json:"distance"`
	StartLocation Location      `json:"start_location"`
	EndLocation   Location      `json:"end_location"`
	StartAddress  string        `json:"start_address"`
	EndAddress    string        `json:"end_address"`
	Duration      time.Duration `json:"duration"`
}
type Steps struct {
	Distance      Distance      `json:"distance"`
	StartLocation Location      `json:"start_location"`
	EndLocation   Location      `json:"end_location"`
	Steps         interface{}   `json:"steps"`
	Duration      time.Duration `json:"duration"`
}
type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

// GasStationsObject used in method to find all gas stations in certain radius
type GasStationsObject struct {
	HtmlAttributions []interface{} `json:"html_attributions"`
	NextPageToken    string        `json:"next_page_token"`
	Results          []Results     `json:"results"`
	Status           string        `json:"status"`
}

type Results struct {
	BusinessStatus      string       `json:"business_status"`
	Geometry            Geometry     `json:"geometry"`
	Icon                string       `json:"icon"`
	IconBackgroundColor string       `json:"icon_background_color"`
	IconMaskBaseUri     string       `json:"icon_mask_base_uri"`
	Name                string       `json:"name"`
	OpeningHours        OpeningHours `json:"opening_hours"`
	Photos              []Photos     `json:"photos"`
	PlaceId             string       `json:"place_id"`
	PlusCode            PlusCode     `json:"plus_code"`
	Rating              float64      `json:"rating"`
	Reference           string       `json:"reference"`
	Scope               string       `json:"scope"`
	Types               []string     `json:"types"`
	UserRatingsTotal    int          `json:"user_ratings_total"`
	Vicinity            string       `json:"vicinity"`
}

type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

type Viewport struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}

type Northeast struct {
	Location Location `json:"location"`
}

type Southwest struct {
	Location Location `json:"location"`
}

type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

type Photos struct {
	Height           int      `json:"height"`
	HtmlAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

// DistanceMatrixObject  used for finding distance of the gas stations
type DistanceMatrixObject struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []Rows   `json:"rows"`
	Status               string   `json:"status"`
}

type Rows struct {
	Elements []Elements `json:"elements"`
}

type Elements struct {
	Distance Distance `json:"distance"`
	Duration Duration `json:"duration"`
	Status   string   `json:"status"`
}
