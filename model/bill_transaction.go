package model

import "time"

type BillTransaction struct {
	Id                 string               `json:"id" gorm:"primary_key"`
	BillCode           string               `json:"bill_code"`
	CustomerName       string               `json:"customer_name"`
	IsDineIn           string               `json:"is_dine_in"`
	CreatedDate        time.Time            `json:"created_date,omitempty"`
	BillTransactionDtl []BillTransactionDtl `json:"bill_transaction_dtl" gorm:"ForeignKey:BillCode;AssociationForeignKey:BillCode"`
}
