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

var Index = struct {
	BANKNIFTY, NIFTY, MIDCPNIFTY, FINNIFTY, NIFTYNXT string
}{
	BANKNIFTY:  "BANKNIFTY",
	NIFTY:      "NIFTY",
	MIDCPNIFTY: "MIDCPNIFTY",
	FINNIFTY:   "FINNIFTY",
	NIFTYNXT:   "NIFTYNXT50",
}