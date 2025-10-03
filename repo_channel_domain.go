package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type ChannelDomainRepository struct {
	*GenericRepository[ChannelDomain]
}

func NewChannelDomainRepository(client *Client) *ChannelDomainRepository {
	return &ChannelDomainRepository{
		GenericRepository: NewGenericRepository[ChannelDomain](client),
	}
}

func (t *ChannelDomainRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelDomain], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel-domain")
}

func (t *ChannelDomainRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[ChannelDomain], *http.Response, error) {
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

func (t *ChannelDomainRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel-domain")
}

func (t *ChannelDomainRepository) Upsert(ctx ApiContext, entity []ChannelDomain) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel_domain")
}

func (t *ChannelDomainRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel_domain")
}

type ChannelDomain struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Id string `json:"id,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	SnippetSet *SnippetSet `json:"snippetSet,omitempty"`

	SnippetSetId string `json:"snippetSetId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Url string `json:"url,omitempty"`
}
