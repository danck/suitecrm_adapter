package hawaicrm

type Address struct {
	ID          string `json:"Id"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Street      string `json:"Street"`
	HouseNumber string `json:"HouseNumber"`
	Zipcode     string `json:"Zipcode"`
	City        string `json:"City"`
	Country     string `json:"Country"`
}

type Customer struct {
	ID              string  `json:"Id"`
	FirstName       string  `json:"FirstName"`
	LastName        string  `json:"LastName"`
	Email           string  `json:"Email"`
	Password        string  `json:"Password"`
	Address         Address `json:"Address"`
	ShipmentAddress Address `json:"ShipmentAddress"`
	InvoiceAddress  Address `json:"InvoiceAddress"`
}

type Order struct {
	ID             string  `json:"Id"`
	Price          float64 `json:"Price"`
	ShippingCosts  float64 `json:"ShippingCosts"`
	ShipmentStatus string  `json:"ShipmentStatus"`
}
