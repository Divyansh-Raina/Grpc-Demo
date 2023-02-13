package model

import "gorm.io/gorm"

type typesofaccount int  

const (
	Free typesofaccount = iota+1
	Developer
	Intel
	Business_T1
	Business_T2
)
type User struct {
	gorm.Model
	Uid         uint64 `gorm:"uniqueIndex"`
	Name        string
	Age         uint32
	Type_of_acc string
}
