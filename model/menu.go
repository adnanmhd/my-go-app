package model

import "time"

type Menu struct {
	Id          string    `json:"id" gorm:"primary_key"`
	MenuName    string    `json:"menu_name"`
	MenuCode    string    `json:"menu_code"`
	Price       uint64    `json:"price"`
	MenuTypeCd  string    `json:"menu_type_cd"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	MenuType    MenuType  `json:"menu_type" gorm:"ForeignKey:MenuTypeCd;AssociationForeignKey:MenuTypeCd"`
}
