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

func FormatMeshDatetime(t time.Time) string {
	const timefmt = "200601021504"
	s := t.Format(timefmt)
	return fmt.Sprintf("%s0", s[:len(s)-1])
}

func (s *MeshService) Get(opt *MeshGetOptions) (*Response, error) {
	var t time.Time
	if opt.DateTime.IsZero() {
		t = time.Now()
	} else {
		t = opt.DateTime
	}
	d := FormatMeshDatetime(t)
	path := fmt.Sprintf("mesh/100/%s.gif", d)
	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}
