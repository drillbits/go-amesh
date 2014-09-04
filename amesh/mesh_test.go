package amesh

import (
	"testing"
	"time"
)

func TestFormatMeshDatetime(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	d := time.Date(2014, 5, 23, 12, 6, 0, 0, jst)
	actual := FormatMeshDatetime(d)
	expected := "201405231200"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
