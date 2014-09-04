package amesh

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "http://tokyo-ame.jwa.or.jp/"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
}

type Client struct {
	// HTTP client used to communicate with the amesh
	client  *http.Client
	BaseURL *url.URL

	TopographyMap *TopographyMapService
	PlaceNameMap  *PlaceNameMapService
	Mesh          *MeshService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.TopographyMap = &TopographyMapService{client: c}
	c.PlaceNameMap = &PlaceNameMapService{client: c}
	c.Mesh = &MeshService{client: c}
	return c
}

func (c *Client) NewRequest(urlStr string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type Response struct {
	*http.Response
	Image image.Image
}

func newResponse(r *http.Response) (*Response, error) {
	resp := &Response{Response: r}
	err := resp.populateImage()
	return resp, err
}

func (resp *Response) populateImage() error {
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	resp.Image = img
	return err
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return newResponse(resp)
}
