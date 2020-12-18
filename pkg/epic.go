package struktura

import "time"

//KeyImage export
type KeyImage struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

//Seller export
type Seller struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Item export2
type Item struct {
	ID        string `json:"id"`
	Namespace string `json:"namespace"`
}

//CustomAttribute export
type CustomAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Category export
type Category struct {
	Path string `json:"path"`
}

//CurrencyInfo export
type CurrencyInfo struct {
	Decimals int `json:"decimals"`
}

//FmtPrice export
type FmtPrice struct {
	OriginalPrice     string `json:"originalPrice"`
	DiscountPrice     string `json:"discountPrice"`
	IntermediatePrice string `json:"intermediatePrice"`
}

//TotalPrice export
type TotalPrice struct {
	DiscountPrice   int          `json:"discountPrice"`
	OriginalPrice   int          `json:"originalPrice"`
	VoucherDiscount int          `json:"voucherPrice"`
	Discount        int          `json:"discount"`
	CurrencyCode    string       `json:"currencyCode"`
	CurrencyInfo    CurrencyInfo `json:"currencyInfo"`
	FmtPrice        FmtPrice     `json:"fmtPrice"`
}

//LineOffer export
type LineOffer struct {
	AppliedRules []interface{} `json:"appliedRules"`
}

//Price export
type Price struct {
	TotalPrice TotalPrice  `json:"totalPrice"`
	LineOffers []LineOffer `json:"lineOffers"`
}

//DiscountSetting export
type DiscountSetting struct {
	DiscountType       string `json:"discountType"`
	DiscountPercentage int    `json:"discountPercentage"`
}

//PromotionalOffer export
type PromotionalOffer struct {
	StartDate        time.Time       `json:"startDate"`
	EndDate          time.Time       `json:"endDate"`
	DiscountSettings DiscountSetting `json:"discountSetting"`
}

//UpcomingPromotionalOffer export
type UpcomingPromotionalOffer struct {
	PromotionalOffers []PromotionalOffer
}

//Promotion export
type Promotion struct {
	PromotionalOffers         []interface{}              `json:"promotionalOffers"`
	UpcomingPromotionalOffers []UpcomingPromotionalOffer `json:"upcomingPromotionalOffers"`
}

//Element export
type Element struct {
	Title            string            `json:"title"`
	ID               string            `json:"id"`
	Namespace        string            `json:"namespace"`
	Description      string            `json:"description"`
	EffectiveDate    time.Time         `json:"effectiveDate"`
	KeyImages        []KeyImage        `json:"keyImages"`
	Seller           Seller            `json:"seller"`
	ProductSlug      string            `json:"productSlug"`
	URLSlug          string            `json:"urlSlug"`
	URL              interface{}       `json:"url"`
	Items            []Item            `json:"items"`
	CustomAttributes []CustomAttribute `json:"customAttributes"`
	Categories       []Category        `json:"categories"`
	Tags             []interface{}     `json:"tags"`
	Price            Price             `json:"price"`
	Promotions       Promotion         `json:"promotions"`
}

//SearchStore export
type SearchStore struct {
	Elements []Element `json:"elements"`
}

//Catalog export
type Catalog struct {
	SearchStore SearchStore `json:"searchStore"`
}

//Data export
type Data struct {
	Catalog Catalog `json:"Catalog"`
}

//Darmowe export
type Darmowe struct {
	Data Data `json:"data"`
}
