package go_heyframe_admin_sdk

import (
	"net/http"
)

type ChannelLanguageRepository struct {
	*GenericRepository[ChannelLanguage]
}

func NewChannelLanguageRepository(client *Client) *ChannelLanguageRepository {
	return &ChannelLanguageRepository{
		GenericRepository: NewGenericRepository[ChannelLanguage](client),
	}
}

func (t *ChannelLanguageRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelLanguage], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-language")
}

func (t *ChannelLanguageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelLanguage], *http.Response, error) {
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

func (t *ChannelLanguageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-language")
}

func (t *ChannelLanguageRepository) Upsert(ctx ApiContext, entity []ChannelLanguage) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_language")
}

func (t *ChannelLanguageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_language")
}

type ChannelLanguage struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`
}
