package amesh

type TopographyMapService struct {
	client *Client
}

func (s *TopographyMapService) Get() (*Response, error) {
	req, err := s.client.NewRequest("map/map100.jpg")
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}

type PlaceNameMapService struct {
	client *Client
}

func (s *PlaceNameMapService) Get() (*Response, error) {
	req, err := s.client.NewRequest("map/msk100.png")
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}
