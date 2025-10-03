package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type CustomerGroupRepository struct {
	*GenericRepository[CustomerGroup]
}

func NewCustomerGroupRepository(client *Client) *CustomerGroupRepository {
	return &CustomerGroupRepository{
		GenericRepository: NewGenericRepository[CustomerGroup](client),
	}
}

func (t *CustomerGroupRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerGroup], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-group")
}

func (t *CustomerGroupRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerGroup], *http.Response, error) {
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

func (t *CustomerGroupRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-group")
}

func (t *CustomerGroupRepository) Upsert(ctx ApiContext, entity []CustomerGroup) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_group")
}

func (t *CustomerGroupRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_group")
}

type CustomerGroup struct {
	Channels []Channel `json:"channels,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []CustomerGroupTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
