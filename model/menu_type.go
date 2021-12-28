package model

type MenuType struct {
	MenuTypeCd   string `json:"menu_type_cd" gorm:"primary_key"`
	MenuTypeName string `json:"menu_type_name"`
}
