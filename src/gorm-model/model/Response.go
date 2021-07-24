package model

// Response struct
type Response struct {
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`

	// User
	UserID   uint   `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`

	// Company
	CompanyID   int    `json:"companyId,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
}
