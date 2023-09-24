package models

import "gorm.io/gorm"

type ComplaintIssue struct {
	gorm.Model
	User   User
	UserID uint

	ComplaintCategory   ComplaintCategory `json:"complaint_category"`
	ComplaintCategoryID uint

	Meteran   Meteran `json:"meteran"`
	MeteranID uint

	ComplaintName    string            `json:"complaint_name"`
	ShortDescription string            `json:"short_description"`
	PriorityLevel    uint              `json:"priority_level"`
	ComplaintImages  []*ComplaintImage `json:"complaint_images"`
}

type ComplaintCategory struct {
	gorm.Model
	Name  string `json:"name"`
	Class string `json:"class"`
}

type ComplaintImage struct {
	gorm.Model
	ComplaintIssue   ComplaintIssue
	ComplaintIssueID uint

	FileName  string `json:"FileName"`
	IsPrimary bool   `json:"is_primary"`
}
