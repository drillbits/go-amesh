package amesh

import (
	"testing"
)

func TestGetArea(t *testing.T) {
	actual, _ := GetArea(0)
	expected := Area{0, "全体", wholeWidth, wholeHeight, 0, 0}
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestGetNGArea(t *testing.T) {
	_, actual := GetArea(666)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestGetAreaOrWholeGetArea(t *testing.T) {
	actual := GetAreaOrWhole(1)
	expected := Area{1, "台東", zoomWidth, zoomHeight, 1708, 555}
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestGetAreaOrWholeGetWhole(t *testing.T) {
	actual := GetAreaOrWhole(666)
	expected := Area{0, "全体", wholeWidth, wholeHeight, 0, 0}
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
