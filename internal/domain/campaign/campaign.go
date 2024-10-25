package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Status string

const (
	Status_Pending Status = "AWAITING"
	Status_Started Status = "STARTED"
	Status_Done    Status = "DONE"
)

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedAt time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    Status
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
		Status:    Status_Pending,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}

	return nil, err
}
