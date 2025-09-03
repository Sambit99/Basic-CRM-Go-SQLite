package model

import (
	"github.com/Sambit99/Basic-CRM-Go-SQLite/pkg/config"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name,omitempty"`
	Company string `json:"company,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

var db *gorm.DB

func init() {
	config.ConnectDB()
	db = config.GetDB()
}
