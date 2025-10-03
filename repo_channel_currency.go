package go_heyframe_admin_sdk

import (
	"net/http"
)

type ChannelCurrencyRepository struct {
	*GenericRepository[ChannelCurrency]
}

func NewChannelCurrencyRepository(client *Client) *ChannelCurrencyRepository {
	return &ChannelCurrencyRepository{
		GenericRepository: NewGenericRepository[ChannelCurrency](client),
	}
}

func (t *ChannelCurrencyRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelCurrency], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-currency")
}

func (t *ChannelCurrencyRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelCurrency], *http.Response, error) {
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

func (t *ChannelCurrencyRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-currency")
}

func (t *ChannelCurrencyRepository) Upsert(ctx ApiContext, entity []ChannelCurrency) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_currency")
}

func (t *ChannelCurrencyRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_currency")
}

type ChannelCurrency struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`
}
