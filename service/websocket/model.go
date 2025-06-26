package websocket

// Message represents a trade event.
type (
	Message struct {
		Stream string `json:"stream"`
		Data   Data   `json:"data"`
	}

	Data struct {
		EventType    string `json:"e"`
		EventTime    int64  `json:"E"`
		Symbol       string `json:"s"`
		TradeID      int64  `json:"t"`
		Price        string `json:"p"`
		Quantity     string `json:"q"`
		TradeTime    int64  `json:"T"`
		IsBuyerMaker bool   `json:"m"`
		Ignore       bool   `json:"M"`
	}
)
