package amesh

type area struct {
	Id     int
	Name   string
	Width  int
	Height int
	Left   int
	Top    int
}

const (
	wholeWidth  = 3080
	wholeHeight = 1920
	zoomWidth   = 770
	zoomHeight  = 480
)

var areas = map[int]area{
	0:  area{0, "全体", wholeWidth, wholeHeight, 0, 0},
	1:  area{1, "台東", zoomWidth, zoomHeight, 1708, 555},
	2:  area{2, "江東", zoomWidth, zoomHeight, 1739, 710},
	3:  area{3, "板橋", zoomWidth, zoomHeight, 1500, 570},
	4:  area{4, "新宿", zoomWidth, zoomHeight, 1516, 679},
	5:  area{5, "世田谷", zoomWidth, zoomHeight, 1516, 819},
	6:  area{6, "東村山", zoomWidth, zoomHeight, 1197, 570},
	7:  area{7, "府中", zoomWidth, zoomHeight, 1197, 679},
	8:  area{8, "町田", zoomWidth, zoomHeight, 1117, 850},
	9:  area{9, "青梅", zoomWidth, zoomHeight, 910, 523},
	10: area{10, "あきるの", zoomWidth, zoomHeight, 878, 601},
	11: area{11, "八王子", zoomWidth, zoomHeight, 958, 741},
	12: area{12, "檜原", zoomWidth, zoomHeight, 686, 617},
	13: area{13, "奥多摩", zoomWidth, zoomHeight, 638, 461},
}

func GetArea(id int) (area, bool) {
	a, ok := areas[id]
	return a, ok
}

func GetAreaOrWhole(id int) area {
	a, ok := GetArea(id)
	if ok {
		return a
	} else {
		a, _ := GetArea(0)
		return a
	}
}
