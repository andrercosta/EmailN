package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email@email.com", "email2@email.com"}
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
}

func TestNewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func TestNewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func TestNewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func TestNewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func TestNewCampaign_MustValidatContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}
