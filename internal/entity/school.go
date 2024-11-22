package entity

import (
	"github.com/google/uuid"
	pb "github.com/grzegab/GO_Module_Example/internal/pb"
	"time"
)

type School struct {
	UUID         string    `json:"id"`
	Owner        string    `json:"admin_id"`
	Name         string    `json:"school_name"`
	IsActive     bool      `json:"is_active"`
	Street       string    `json:"street"`
	Town         string    `json:"town"`
	Postcode     string    `json:"postcode"`
	Config       Config    `json:"config"` //when config change old is still available for historical data
	RegisterCode string    `json:"register_code"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	DeletedAt    time.Time `json:"-"`
}

func CreateNewFromRequest(rr *pb.RegisterRequest) *School {
	uid, _ := uuid.NewUUID()
	c := NewBasicConfig()

	s := &School{
		UUID:         uid.String(),
		Name:         "School",
		IsActive:     false,
		RegisterCode: rr.GetCode(),
		Owner:        rr.GetAdminId(),
		Config:       *c,
	}

	return s
}
