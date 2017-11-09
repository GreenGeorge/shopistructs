// Package shopistructs :
// provides prewritten struct-types
// that conforms to Shopify's API response objects
// use to quickly unmarshal JSON response objects
// and explore the API
package shopistructs

type OrdersEnvelope struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	ID                    int            `json:"id"`
	Email                 string         `json:"email"`
	ClosedAt              string         `json:"closed_at"`
	CreatedAt             string         `json:"created_at"`
	Number                int            `json:"number"`
	Note                  string         `json:"note"`
	Token                 string         `json:"token"`
	Gateway               string         `json:"gateway"`
	Test                  bool           `json:"test"`
	TotalPrice            string         `json:"total_price"`
	SubtotalPrice         string         `json:"subtotal_price"`
	TotalWeight           int            `json:"total_weight"`
	TotalTax              string         `json:"total_tax"`
	TaxesIncluded         bool           `json:"taxes_included"`
	Currency              string         `json:"currency"`
	FinancialStatus       string         `json:"financial_status"`
	Confirmed             bool           `json:"confirmed"`
	TotalDiscounts        string         `json:"total_discounts"`
	TotalLineItemsPrice   string         `json:"total_line_items_price"`
	CartToken             string         `json:"cart_token"`
	BuyerAcceptsMarketing bool           `json:"buyer_accepts_marketing"`
	Name                  string         `json:"name"`
	ReferringSite         string         `json:"referring_site"`
	LandingSite           string         `json:"landing_site"`
	CancelledAt           string         `json:"cancelled_at"`
	CancelReason          string         `json:"cancel_reason"`
	TotalPriceUSD         string         `json:"total_price_usd"`
	CheckoutToken         string         `json:"checkout_token"`
	Reference             string         `json:"reference"`
	UserID                int            `json:"user_id"`
	LocationID            string         `json:"location_id"`
	SourceIdentifier      string         `json:"source_identifier"`
	SourceURL             string         `json:"source_url"`
	ProcessedAt           string         `json:"processed_at"`
	DeviceID              string         `json:"device_id"`
	Phone                 string         `json:"phone"`
	CustomerLocale        string         `json:"customer_locale"`
	AppID                 int            `json:"app_id"`
	BrowserIP             string         `json:"browser_ip"`
	LandingSiteRef        string         `json:"landing_site_ref"`
	OrderNumber           int            `json:"order_number"`
	DiscountCodes         []DiscountCode `json:"discount_codes"`
	NoteAttributes        []Property     `json:"note_attributes"`
	PaymentGatewayNames   []string       `json:"payment_gateway_names"`
	ProcessingMethod      string         `json:"processing_method"`
	CheckoutID            int            `json:"checkout_id"`
	SourceName            string         `json:"source_name"`
	FulfillmentStatus     string         `json:"fulfillment_status"`
	TaxLines              []TaxLine      `json:"tax_lines"`
	Tags                  string         `json:"tags"`
	ContactEmail          string         `json:"contact_email"`
	OrderStatusURL        string         `json:"order_status_url"`
	LineItems             []LineItem     `json:"line_items"`
	ShippingLines         []ShippingLine `json:"shipping_lines"`
	BillingAddress        Address        `json:"billing_address"`
	ShippingAddress       Address        `json:"shipping_address"`
	Fulfillments          []Fulfillment  `json:"fulfillments"`
	ClientDetails         struct {
		BrowserIP      string `json:"browser_ip"`
		AcceptLanguage string `json:"accept_language"`
		UserAgent      string `json:"user_agent"`
		SessionHash    string `json:"session_hash"`
		BrowserWidth   string `json:"browser_width"`
		BrowserHeight  string `json:"browser_height"`
	} `json:"client_details"`
	Refunds []Refund `json:"refunds"`
}

type DiscountCode struct {
	Code   string `json:"code"`
	Amount string `json:"amount"`
	Type   string `json:"percentage"`
}

type TaxLine struct {
	Title string  `json:"title"`
	Price string  `json:"price"`
	Rate  float64 `json:"rate"`
}

type LineItem struct {
	ID                         int        `json:"id"`
	VariantID                  int        `json:"variant_id"`
	Title                      string     `json:"title"`
	Quantity                   int        `json:"quantity"`
	Price                      string     `json:"price"`
	Grams                      int        `json:"grams"`
	SKU                        string     `json:"sku"`
	VariantTitle               string     `json:"variant_title"`
	Vendor                     string     `json:"vendor"`
	FulfillmentService         string     `json:"fullfillment_service"`
	ProductID                  int        `json:"product_id"`
	RequiresShipping           bool       `json:"requires_shipping"`
	Taxable                    bool       `json:"taxable"`
	GiftCard                   bool       `json:"gift_card"`
	Name                       string     `json:"Name"`
	VariantInventoryManagement string     `json:"variant_inventory_management"`
	Properties                 []Property `json:"properties"`
	ProductExists              bool       `json:"product_exists"`
	FulfillableQuantity        int        `json:"fulfillable_quantity"`
	TotalDiscount              string     `json:"total_discount"`
	FulfillmentStatus          string     `json:"fulfillment_status"`
	TaxLines                   []TaxLine  `json:"tax_lines"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ShippingLine struct {
	ID                             int       `json:"id"`
	Title                          string    `json:"title"`
	Price                          string    `json:"price"`
	Code                           string    `json:"code"`
	Phone                          string    `json:"phone"`
	RequestedFullfillmentServiceID int       `json:"requested_fulfillment_service_id"`
	DeliveryCategory               string    `json:"delivery_category"`
	CarrierIdentifier              string    `json:"carrier_identifier"`
	DiscountedPrice                string    `json:"discounted_price"`
	TaxLines                       []TaxLine `json:"tax_lines"`
}

type Address struct {
	FirstName    string  `json:"first_name"`
	Address1     string  `json:"address1"`
	Phone        string  `json:"phone"`
	City         string  `json:"city"`
	Zip          string  `json:"zip"`
	Province     string  `json:"province"`
	Country      string  `json:"country"`
	LastName     string  `json:"last_name"`
	Address2     string  `json:"address2"`
	Company      string  `json:"company"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"country_code"`
	ProvinceCode string  `json:"province_code"`
}

type Fulfillment struct {
	ID              int      `json:"id"`
	OrderID         int      `json:"order_id"`
	Status          int      `json:"status"`
	CreatedAt       string   `json:"created_at"`
	Service         string   `json:"service"`
	UpdatedAt       string   `json:"updated_at"`
	TrackingCompany string   `json:"tracking_company"`
	ShipmentStatus  string   `json:"shipment_status"`
	TrackingNumber  string   `json:"tracking_number"`
	TrackingNumbers []string `json:"tracking_numbers"`
	TrackingURL     string   `json:"tracking_url"`
	TrackingURLS    []string `json:"tracking_urls"`
	Receipt         struct {
		TestCase      bool   `json:"testcase"`
		Authorization string `json:"authorization"`
	} `json:"receipt"`
	LineItems []LineItem `json:"line_items"`
}

type Refund struct {
	ID              int              `json:"id"`
	OrderID         int              `json:"order_id"`
	CreatedAt       string           `json:"created_at"`
	Note            string           `json:"note"`
	Restock         bool             `json:"restock"`
	UserID          int              `json:"user_id"`
	ProcessedAt     string           `json:"processed_at"`
	RefundLineItems []RefundLineItem `json:"refund_line_items"`
}

type RefundLineItem struct {
	ID           int           `json:"id"`
	Quantity     int           `json:"quantity"`
	LineItemID   int           `json:"line_item_id"`
	Subtotal     float64       `json:"subtotal"`
	TotalTax     float64       `json:"total_tax"`
	LineItem     LineItem      `json:"line_item"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID            int      `json:"id"`
	OrderID       int      `json:"id"`
	Amount        string   `json:"amount"`
	Kind          string   `json:"kind"`
	Gateway       string   `json:"gateway"`
	Status        string   `json:"status"`
	Message       string   `json:"message"`
	CreatedAt     string   `json:"created_at"`
	Test          bool     `json:"test"`
	Authorization string   `json:"authorization"`
	Currency      string   `json:"currency"`
	LocationID    int      `json:"location_id"`
	UserID        int      `json:"user_id"`
	ParentID      int      `json:"parent_id"`
	DeviceID      int      `json:"device_id"`
	Receipt       struct{} `json:"receipt"`
	ErrorCode     string   `json:"error_code"`
	SourceName    string   `json:"web"`
}
