package model

import (
	"errors"

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

func GetLeads() (leads []Lead, err error) {
	db.Find(&leads)

	return leads, nil
}

func GetLead(id int) (lead Lead, err error) {
	result := db.Find(&lead).Where("ID=?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Lead{}, nil // not found, return zero value
		}
		return Lead{}, result.Error // real error
	}

	return lead, nil
}

func NewLead(lead Lead) (newLead Lead, err error) {

	if lead.Name != "" {
		newLead.Name = lead.Name
	}

	if lead.Company != "" {
		newLead.Company = lead.Company
	}

	if lead.Email != "" {
		newLead.Email = lead.Email
	}

	if lead.Phone != "" {
		newLead.Phone = lead.Phone
	}

	result := db.Create(&newLead)

	if result.Error != nil {
		return Lead{}, result.Error
	}

	return newLead, nil
}

func DeleteLead(id int) (lead Lead, err error) {
	result := db.Where("ID=?", id).Delete(&lead)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Lead{}, errors.New("Lead not found") // not found, return zero value
		}
		return Lead{}, result.Error // real error
	}

	return lead, nil
}
