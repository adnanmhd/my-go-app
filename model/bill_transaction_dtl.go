package model

type BillTransactionDtl struct {
	Id       string  `json:"id" gorm:"primary_key"`
	MenuName string  `json:"menu_name"`
	Price    float32 `json:"price" sql:"type:decimal(19,2);"`
	Qty      int     `json:"qty"`
	Total    float32 `json:"total" sql:"type:decimal(19,2);"`
	BillCode string  `json:"bill_code"`
}
