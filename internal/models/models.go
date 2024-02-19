package models


type Omnibus struct {
	UPC string `json:"upc"`
	ISTUrl string `json:"isturl"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	PageCount int `json:"pagecount"`
	Publisher string `json:"publisher"`
	Sale bool `json:sale`
	Current float32 `json:"current"`
	Saving int `json:"saving"`
	DateCreated string `json:date`
}