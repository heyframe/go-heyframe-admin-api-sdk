package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type ChannelTranslationRepository struct {
	*GenericRepository[ChannelTranslation]
}

func NewChannelTranslationRepository(client *Client) *ChannelTranslationRepository {
	return &ChannelTranslationRepository{
		GenericRepository: NewGenericRepository[ChannelTranslation](client),
	}
}

func (t *ChannelTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-translation")
}

func (t *ChannelTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelTranslation], *http.Response, error) {
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

func (t *ChannelTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-translation")
}

func (t *ChannelTranslationRepository) Upsert(ctx ApiContext, entity []ChannelTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_translation")
}

func (t *ChannelTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_translation")
}

type ChannelTranslation struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	HomeEnabled bool `json:"homeEnabled,omitempty"`

	HomeKeywords string `json:"homeKeywords,omitempty"`

	HomeMetaDescription string `json:"homeMetaDescription,omitempty"`

	HomeMetaTitle string `json:"homeMetaTitle,omitempty"`

	HomeName string `json:"homeName,omitempty"`

	HomeSlotConfig interface{} `json:"homeSlotConfig,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Name string `json:"name,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
