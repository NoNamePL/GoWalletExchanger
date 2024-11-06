package currencyExchanger

type Request struct {
	FromCurrency string `json:"from"`
	ToCurrency   string `json:"to"`
}

type Response struct {
	FromCurrency       string `json:"from"`
	ToCurrency         string `json:"to"`
	ConversionMultiple string `json:"conversion_multiple"`
}
