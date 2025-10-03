package go_heyframe_admin_sdk

import (
	"net/http"
)

type ChannelCountryRepository struct {
	*GenericRepository[ChannelCountry]
}

func NewChannelCountryRepository(client *Client) *ChannelCountryRepository {
	return &ChannelCountryRepository{
		GenericRepository: NewGenericRepository[ChannelCountry](client),
	}
}

func (t *ChannelCountryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelCountry], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-country")
}

func (t *ChannelCountryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelCountry], *http.Response, error) {
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

func (t *ChannelCountryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-country")
}

func (t *ChannelCountryRepository) Upsert(ctx ApiContext, entity []ChannelCountry) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_country")
}

func (t *ChannelCountryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_country")
}

type ChannelCountry struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	Country *Country `json:"country,omitempty"`

	CountryId string `json:"countryId,omitempty"`
}
