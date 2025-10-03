package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type ChannelTypeTranslationRepository struct {
	*GenericRepository[ChannelTypeTranslation]
}

func NewChannelTypeTranslationRepository(client *Client) *ChannelTypeTranslationRepository {
	return &ChannelTypeTranslationRepository{
		GenericRepository: NewGenericRepository[ChannelTypeTranslation](client),
	}
}

func (t *ChannelTypeTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelTypeTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-type-translation")
}

func (t *ChannelTypeTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelTypeTranslation], *http.Response, error) {
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

func (t *ChannelTypeTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-type-translation")
}

func (t *ChannelTypeTranslationRepository) Upsert(ctx ApiContext, entity []ChannelTypeTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_type_translation")
}

func (t *ChannelTypeTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_type_translation")
}

type ChannelTypeTranslation struct {
	ChannelType *ChannelType `json:"channelType,omitempty"`

	ChannelTypeId string `json:"channelTypeId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	DescriptionLong string `json:"descriptionLong,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Manufacturer string `json:"manufacturer,omitempty"`

	Name string `json:"name,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
