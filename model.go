package main

// GeneralLedger represents the top level structure of the JSON file
// Data is the only attribute included as it is the only one necessary
// for the calculations
type GeneralLedger struct {
	Data []AccountData `json:"data"`
}

// AccountData is the representation of the main account data
// which is used in the calcations
type AccountData struct {
	AccountCategory   string  `json:"account_category"`
	AccountCode       int16   `json:"account_code"`
	AccountCurrency   string  `json:"account_currency"`
	AccountIdentifier string  `json:"account_identifier"`
	AccountStatus     string  `json:"account_status"`
	ValueType         string  `json:"value_type"`
	AccountName       string  `json:"account_name"`
	AccountType       string  `json:"account_type"`
	AccountTypeBank   string  `json:"account_type_bank"`
	SystemAccount     string  `json:"system_account"`
	TotalValue        float32 `json:"total_value"`
}

// LedgerTotals is the struct that is returned from the calculation function
// It contains all the required values formatted back as floats, ready for output
type LedgerTotals struct {
	Revenue             float32
	Expenses            float32
	GrossProfitMargin   float32
	NetProfitMargin     float32
	WorkingCaptialRatio float32
}
