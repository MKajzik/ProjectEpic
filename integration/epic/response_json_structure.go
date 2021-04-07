package epic

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

//Promotion export
type Promotion struct {
	PromotionalOffers         []PromotionalOffer `json:"promotionalOffers"`
	UpcomingPromotionalOffers []PromotionalOffer `json:"upcomingPromotionalOffers"`
}

//Element export
type Element struct {
	Title                string            `json:"title"`
	ID                   string            `json:"id"`
	Namespace            string            `json:"namespace"`
	Description          string            `json:"description"`
	EffectiveDate        time.Time         `json:"effectiveDate"`
	OfferType            string            `json:"offerType"`
	ExpiryDate           interface{}       `json:"expiryDate"`
	Status               string            `json:"status"`
	IsCodeRedemptionOnly bool              `json:"isCodeRedemptionOnly"`
	KeyImages            []KeyImage        `json:"keyImages"`
	Seller               Seller            `json:"seller"`
	ProductSlug          string            `json:"productSlug"`
	URLSlug              string            `json:"urlSlug"`
	URL                  interface{}       `json:"url"`
	Items                []Item            `json:"items"`
	CustomAttributes     []CustomAttribute `json:"customAttributes"`
	Categories           []Category        `json:"categories"`
	Tags                 []interface{}     `json:"tags"`
	Price                Price             `json:"price"`
	Promotions           Promotion         `json:"promotions"`
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

//FreeGame export
type FreeGame struct {
	Data Data `json:"data"`
}

//GetElementTitle export
func (f FreeGame) GetElementTitle(i int) string {
	return f.Data.Catalog.SearchStore.Elements[i].Title
}

//GetElementStatus export
func (f FreeGame) GetElementStatus(i int) string {
	return f.Data.Catalog.SearchStore.Elements[i].Status
}

//GetProductSlug export
func (f FreeGame) GetProductSlug(i int) string {
	return f.Data.Catalog.SearchStore.Elements[i].ProductSlug
}

//GetAllElements export
func (f FreeGame) GetAllElements() []Element {
	return f.Data.Catalog.SearchStore.Elements
}

//GetAllKeyImages export
func (f FreeGame) GetAllKeyImages(i int) []KeyImage {
	return f.Data.Catalog.SearchStore.Elements[i].KeyImages
}

//GetKeyImageType export
func (f FreeGame) GetKeyImageType(i int, j int) string {
	return f.Data.Catalog.SearchStore.Elements[i].KeyImages[j].Type
}

//GetKeyImageURL export
func (f FreeGame) GetKeyImageURL(i int, j int) string {
	return f.Data.Catalog.SearchStore.Elements[i].KeyImages[j].URL
}

//GetPromotionalOffers export
func (f FreeGame) GetPromotionalOffers(i int) []PromotionalOffer {
	return f.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers
}

//GetPromotianlOfferStartDate export
func (f FreeGame) GetPromotianlOfferStartDate(i int) time.Time {
	return f.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].StartDate
}

//GetPromotianlOfferEndDate export
func (f FreeGame) GetPromotianlOfferEndDate(i int) time.Time {
	return f.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].EndDate
}

//GetPrice export
func (f FreeGame) GetPrice(i int) int {
	return f.Data.Catalog.SearchStore.Elements[i].Price.TotalPrice.DiscountPrice
}
