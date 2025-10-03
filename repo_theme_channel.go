package go_heyframe_admin_sdk

import (
	"net/http"
)

type ThemeChannelRepository struct {
	*GenericRepository[ThemeChannel]
}

func NewThemeChannelRepository(client *Client) *ThemeChannelRepository {
	return &ThemeChannelRepository{
		GenericRepository: NewGenericRepository[ThemeChannel](client),
	}
}

func (t *ThemeChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "theme-channel")
}

func (t *ThemeChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ThemeChannel], *http.Response, error) {
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

func (t *ThemeChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "theme-channel")
}

func (t *ThemeChannelRepository) Upsert(ctx ApiContext, entity []ThemeChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "theme_channel")
}

func (t *ThemeChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "theme_channel")
}

type ThemeChannel struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	Theme *Theme `json:"theme,omitempty"`

	ThemeId string `json:"themeId,omitempty"`
}
