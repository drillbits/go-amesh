package amesh

type Area struct {
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

var Areas = map[int]Area{
	0:  Area{0, "全体", wholeWidth, wholeHeight, 0, 0},
	1:  Area{1, "台東", zoomWidth, zoomHeight, 1708, 555},
	2:  Area{2, "江東", zoomWidth, zoomHeight, 1739, 710},
	3:  Area{3, "板橋", zoomWidth, zoomHeight, 1500, 570},
	4:  Area{4, "新宿", zoomWidth, zoomHeight, 1516, 679},
	5:  Area{5, "世田谷", zoomWidth, zoomHeight, 1516, 819},
	6:  Area{6, "東村山", zoomWidth, zoomHeight, 1197, 570},
	7:  Area{7, "府中", zoomWidth, zoomHeight, 1197, 679},
	8:  Area{8, "町田", zoomWidth, zoomHeight, 1117, 850},
	9:  Area{9, "青梅", zoomWidth, zoomHeight, 910, 523},
	10: Area{10, "あきるの", zoomWidth, zoomHeight, 878, 601},
	11: Area{11, "八王子", zoomWidth, zoomHeight, 958, 741},
	12: Area{12, "檜原", zoomWidth, zoomHeight, 686, 617},
	13: Area{13, "奥多摩", zoomWidth, zoomHeight, 638, 461},
}

func GetArea(id int) (Area, bool) {
	a, ok := Areas[id]
	return a, ok
}

func GetAreaOrWhole(id int) Area {
	a, ok := GetArea(id)
	if ok {
		return a
	} else {
		a, _ := GetArea(0)
		return a
	}
}
