package main

type GeneralLedger struct {
  Data []AccountData `json:"data"`
}

type AccountData struct {
	AccountCategory   string  `json:"account_category"`
	AccountCode       int16   `json:"account_code"`
	AccountCurrency   string  `json:"account_currency"`
	AccountIdentifier string  `json:"account_identifier"`
	AccountStatus     string  `json:"account_status"`
	ValueType         string  `json:"credit"`
	AccountName       string  `json:"account_name"`
	AccountType       string  `json:"account_type"`
	AccountTypeBank   string  `json:"account_type_bank"`
	SystemAccount     string  `json:"system_account"`
	TotalValue        float32 `json:"total_value"`
}
