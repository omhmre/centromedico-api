package models

import "time"

type License struct {
	Id          int       `json:"id"`
	KeyHash     string    `json:"key_hash"`
	LicenseKey  string    `json:"license_key"`
	ClientName  string    `json:"client_name"`
	Rif         string    `json:"rif"`
	ValidUntil  time.Time `json:"valid_until"`
	IsPremium   bool      `json:"is_premium"`
	ActivatedAt time.Time `json:"activated_at"`
}

type LicenseStatus struct {
	IsActive      bool      `json:"is_active"`
	IsPremium     bool      `json:"is_premium"`
	ClientName    string    `json:"client_name"`
	Rif           string    `json:"rif"`
	ValidUntil    time.Time `json:"valid_until"`
	DaysRemaining int       `json:"days_remaining"`
	StatusMessage string    `json:"status_message"`
	LicenseKey    string    `json:"license_key"`
}

type LicenseActivationRequest struct {
	LicenseKey string `json:"license_key"`
}
