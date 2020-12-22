package epic

//Variables export
type Variables struct {
	Category       string `json:"category"`
	Count          int    `json:"count"`
	Country        string `json:"country"`
	Keywords       string `json:"keywords"`
	Locale         string `json:"locale"`
	SortDir        string `json:"sortDir"`
	AllowCountries string `json:"allowCountries"`
	Start          int    `json:"start"`
	Tag            string `json:"tag"`
	WithPrice      bool   `json:"withPrice"`
}

//SearchQuery export
type SearchQuery struct {
	Query    string    `json:"query"`
	Variable Variables `json:"variables"`
}
