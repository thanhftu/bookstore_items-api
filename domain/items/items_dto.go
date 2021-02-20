package items

// Item contains item info
type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Picturea          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float64     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

// Description contain description of item
type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

// Picture contain picture info of item
type Picture struct {
	ID  int64  `json:"id"`
	URL string `json:""url`
}
