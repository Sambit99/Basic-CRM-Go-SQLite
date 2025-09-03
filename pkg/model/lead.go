package model

import (
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name,omitempty"`
	Company string `json:"company,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
}
