package opendsd

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const (
	APIRoot                      = "http://opendsd.sandiego.gov/api"
	HeaderExtractTimestampFormat = "1/2/2006 3:04:05 PM"
	TimestampFormat              = "2006-01-02T15:04:05"
)

type APIError struct {
	Message string `json:"ErrorMessage"`
}

func (a APIError) Error() string {
	return a.Message
}

type Client struct {
	APIRoot    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		APIRoot: APIRoot,
		HTTPClient: &http.Client{
			//	20 second timeout
			Timeout: time.Second * 20,
			//	the API returns a 302 (Temporary Redirect) status code but does
			//	not provide a Location header for the redirect. Because of this
			//	we need to use a DefaultTransport to avoid Go's automatic handling
			//	of 3XX status codes
			Transport: http.DefaultTransport,
		},
	}
}

//	get makes a GET request to the API using the supplied URI
//	and decodes the JSON response into the provided interface
func (c *Client) get(uri string, v interface{}) error {
	//	build our request
	req, err := http.NewRequest("GET", c.APIRoot+uri, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")

	//	the API returns a 302 (Temporary Redirect) status code but does
	//	not provide a Location header for the redirect. Because of this
	//	we need to use a DefaultTransport rather than client.Do()
	//	make our request
	resp, err := c.HTTPClient.Transport.RoundTrip(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//	decode our response body into the provided interface
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		//	check if we have an error message. This API returns a 302
		//	for both successes and fails
		var apiError APIError
		if err = json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return err
		}

		return apiError
	}

	return nil
}

type HeaderExtractTimestamp time.Time

func (aet HeaderExtractTimestamp) String() string {
	return time.Time(aet).Format(time.RFC3339)
}

func (aet *HeaderExtractTimestamp) UnmarshalJSON(data []byte) error {
	tstr := strings.Trim(string(data), "\"")
	if tstr == "null" || tstr == "" {
		return nil
	}

	v, err := time.Parse(HeaderExtractTimestampFormat, tstr)
	if err != nil {
		return err
	}
	*aet = HeaderExtractTimestamp(v)

	return nil
}

type Timestamp time.Time

func (at Timestamp) String() string {
	return time.Time(at).Format(time.RFC3339)
}

func (at *Timestamp) UnmarshalJSON(data []byte) error {
	tstr := strings.Trim(string(data), "\"")
	if tstr == "null" || tstr == "" {
		return nil
	}

	v, err := time.Parse(TimestampFormat, tstr)
	if err != nil {
		return err
	}
	*at = Timestamp(v)

	return nil
}
