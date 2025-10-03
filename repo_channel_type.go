package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type ChannelTypeRepository struct {
	*GenericRepository[ChannelType]
}

func NewChannelTypeRepository(client *Client) *ChannelTypeRepository {
	return &ChannelTypeRepository{
		GenericRepository: NewGenericRepository[ChannelType](client),
	}
}

func (t *ChannelTypeRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelType], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-type")
}

func (t *ChannelTypeRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelType], *http.Response, error) {
	if criteria.Limit == 0 {
		criteria.Limit = 50
	}

	if criteria.Page == 0 {
		criteria.Page = 1
	}

	c, resp, err := t.Search(ctx, criteria)

	if err != nil {
		return c, resp, err
	}

	for {
		criteria.Page++

		nextC, nextResp, nextErr := t.Search(ctx, criteria)

		if nextErr != nil {
			return c, nextResp, nextErr
		}

		if len(nextC.Data) == 0 {
			break
		}

		c.Data = append(c.Data, nextC.Data...)
	}

	c.Total = int64(len(c.Data))

	return c, resp, err
}

func (t *ChannelTypeRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-type")
}

func (t *ChannelTypeRepository) Upsert(ctx ApiContext, entity []ChannelType) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_type")
}

func (t *ChannelTypeRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_type")
}

type ChannelType struct {
	Channels []Channel `json:"channels,omitempty"`

	CoverUrl string `json:"coverUrl,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	DescriptionLong string `json:"descriptionLong,omitempty"`

	IconName string `json:"iconName,omitempty"`

	Id string `json:"id,omitempty"`

	Manufacturer string `json:"manufacturer,omitempty"`

	Name string `json:"name,omitempty"`

	ScreenshotUrls interface{} `json:"screenshotUrls,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []ChannelTypeTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
