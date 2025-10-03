package go_heyframe_admin_sdk

import (
	"net/http"
)

type LandingPageChannelRepository struct {
	*GenericRepository[LandingPageChannel]
}

func NewLandingPageChannelRepository(client *Client) *LandingPageChannelRepository {
	return &LandingPageChannelRepository{
		GenericRepository: NewGenericRepository[LandingPageChannel](client),
	}
}

func (t *LandingPageChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "landing-page-channel")
}

func (t *LandingPageChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[LandingPageChannel], *http.Response, error) {
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

func (t *LandingPageChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "landing-page-channel")
}

func (t *LandingPageChannelRepository) Upsert(ctx ApiContext, entity []LandingPageChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "landing_page_channel")
}

func (t *LandingPageChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "landing_page_channel")
}

type LandingPageChannel struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	LandingPage *LandingPage `json:"landingPage,omitempty"`

	LandingPageId string `json:"landingPageId,omitempty"`

	LandingPageVersionId string `json:"landingPageVersionId,omitempty"`
}
