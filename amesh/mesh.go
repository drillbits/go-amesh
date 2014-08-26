package amesh

import (
	"fmt"
	"time"
)

type MeshService struct {
	client *Client
}

type MeshGetOptions struct {
	DateTime time.Time
}

func (s *MeshService) Get(opt *MeshGetOptions) (*Response, error) {
	var t time.Time
	if opt.DateTime.IsZero() {
		t = time.Now()
	} else {
		t = opt.DateTime
	}
	timefmt := "200601021504"
	d := t.Format(timefmt)
	path := fmt.Sprintf(
		"mesh/100/%s0.gif", d[:len(d)-1])
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}
