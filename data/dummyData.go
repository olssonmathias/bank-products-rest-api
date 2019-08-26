package data

type Product struct {
	ID   string `json:"id"`
	Bank string `json:"bank"`
	Name string `json:"name"`
	AER  string `json:"aer"`
}

type Products []Product

var ProductData Products

func PopulateDummyData() {
	ProductData = Products{
		Product{
			ID:   "1",
			Bank: "Nationwide",
			Name: "FlexDirect",
			AER:  "5%",
		},
		Product{
			ID:   "2",
			Bank: "TSB",
			Name: "TSB Classic Plus",
			AER:  "3%",
		},
		Product{
			ID:   "3",
			Bank: "Lloyds",
			Name: "Club Lloyds",
			AER:  "1.5%",
		},
	}
}
