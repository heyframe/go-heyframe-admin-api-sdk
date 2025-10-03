package go_heyframe_admin_sdk

import (
	"net/http"
)

type ChannelPaymentMethodRepository struct {
	*GenericRepository[ChannelPaymentMethod]
}

func NewChannelPaymentMethodRepository(client *Client) *ChannelPaymentMethodRepository {
	return &ChannelPaymentMethodRepository{
		GenericRepository: NewGenericRepository[ChannelPaymentMethod](client),
	}
}

func (t *ChannelPaymentMethodRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelPaymentMethod], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-payment-method")
}

func (t *ChannelPaymentMethodRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelPaymentMethod], *http.Response, error) {
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

func (t *ChannelPaymentMethodRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-payment-method")
}

func (t *ChannelPaymentMethodRepository) Upsert(ctx ApiContext, entity []ChannelPaymentMethod) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_payment_method")
}

func (t *ChannelPaymentMethodRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_payment_method")
}

type ChannelPaymentMethod struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`
}
