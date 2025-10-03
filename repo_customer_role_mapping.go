package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type CustomerRoleMappingRepository struct {
	*GenericRepository[CustomerRoleMapping]
}

func NewCustomerRoleMappingRepository(client *Client) *CustomerRoleMappingRepository {
	return &CustomerRoleMappingRepository{
		GenericRepository: NewGenericRepository[CustomerRoleMapping](client),
	}
}

func (t *CustomerRoleMappingRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerRoleMapping], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer-role-mapping")
}

func (t *CustomerRoleMappingRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[CustomerRoleMapping], *http.Response, error) {
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

func (t *CustomerRoleMappingRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer-role-mapping")
}

func (t *CustomerRoleMappingRepository) Upsert(ctx ApiContext, entity []CustomerRoleMapping) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer_role_mapping")
}

func (t *CustomerRoleMappingRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer_role_mapping")
}

type CustomerRoleMapping struct {
	AclRoleId string `json:"aclRoleId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Customer *Customer `json:"customer,omitempty"`

	CustomerId string `json:"customerId,omitempty"`

	CustomerRole *CustomerRole `json:"customerRole,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
