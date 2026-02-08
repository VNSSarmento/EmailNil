package campaign

import (
	internaerrors "emailNil/internal/internaErrors"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	ID       string    `validate:"required"`
	Name     string    `validate:"min=5,max=24"`
	Content  string    `validate:"min=5,max=1024"`
	CreatAt  time.Time `validate:"required"`
	Contacts []Contact `validate:"min=1,dive"`
}

type Contact struct {
	Email string `validate:"email"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contact := make([]Contact, len(emails))

	for index, email := range emails {
		contact[index].Email = email
	}

	campaign := &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		CreatAt:  time.Now(),
		Contacts: contact,
	}

	err := internaerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err

}
