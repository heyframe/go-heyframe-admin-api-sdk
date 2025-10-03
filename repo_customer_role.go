package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type CustomerRoleRepository struct {
	*GenericRepository[CustomerRole]
}

func NewCustomerRoleRepository(client *Client) *CustomerRoleRepository {
	return &CustomerRoleRepository{
		GenericRepository: NewGenericRepository[CustomerRole](client),
	}
}

func (t *CustomerRoleRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerRole], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-role")
}

func (t *CustomerRoleRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerRole], *http.Response, error) {
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

func (t *CustomerRoleRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-role")
}

func (t *CustomerRoleRepository) Upsert(ctx ApiContext, entity []CustomerRole) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_role")
}

func (t *CustomerRoleRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_role")
}

type CustomerRole struct {
	Active bool `json:"active,omitempty"`

	Config interface{} `json:"config,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	DeletedAt time.Time `json:"deletedAt,omitempty"`

	ExtraFields interface{} `json:"extraFields,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Privileges interface{} `json:"privileges,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
