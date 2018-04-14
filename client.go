package marvel

import (
	"net/http"
	"fmt"
	"github.com/dghubble/sling"
)

const (
	BASEURL = "https://gateway.marvel.com/v1/public/"
)

// Marvel Client
type Client struct {
	auth Authenticator
	slg  *sling.Sling
	CharacterDetails *CharacterInfo
}

// NewClient returns an API Client with given authenticator
func NewClient(authenticator Authenticator, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	base := sling.New().Client(httpClient).Base(BASEURL).QueryStruct(authenticator.GetAuth())
	c := &Client{
		auth: authenticator,
		slg:  base,
		CharacterDetails: NewCharacterInfo(base.New()),
	}

	return c
}

// This function make Get request and unmarshal into wrapper
func getRequest(slg *sling.Sling, path string, wrapper interface{}) (*http.Response, error) {
	apiError := &APIError{}
	resp, err := slg.New().Get(path).Receive(wrapper, apiError)
	if err == nil && apiError.Code != nil {
		err = apiError
	}
	return resp, err
}

// struct for APIError
type APIError struct {
	Code    interface{}
	Message string
}

// Returns string with error code and error message
func (ae *APIError) Error() string {
	return fmt.Sprintf("error: %v %v", ae.Code, ae.Message)
}
