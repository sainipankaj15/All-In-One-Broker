package tiqs

const ADMIN_TIQS string = "FB5650"

type ProductType string

var Product = struct {
	INTRADAY, MARGIN, CNC ProductType
}{
	INTRADAY: "I",
	MARGIN:   "M",
	CNC:      "C",
}
