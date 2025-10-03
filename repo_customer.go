package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type CustomerRepository struct {
	*GenericRepository[Customer]
}

func NewCustomerRepository(client *Client) *CustomerRepository {
	return &CustomerRepository{
		GenericRepository: NewGenericRepository[Customer](client),
	}
}

func (t *CustomerRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Customer], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "customer")
}

func (t *CustomerRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Customer], *http.Response, error) {
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

func (t *CustomerRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "customer")
}

func (t *CustomerRepository) Upsert(ctx ApiContext, entity []Customer) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "customer")
}

func (t *CustomerRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "customer")
}

type Customer struct {
	Active bool `json:"active,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	AvatarId string `json:"avatarId,omitempty"`

	AvatarMedia *Media `json:"avatarMedia,omitempty"`

	Birthday time.Time `json:"birthday,omitempty"`

	BoundChannel *Channel `json:"boundChannel,omitempty"`

	BoundChannelId string `json:"boundChannelId,omitempty"`

	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CreatedBy *User `json:"createdBy,omitempty"`

	CreatedById string `json:"createdById,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CustomerNumber string `json:"customerNumber,omitempty"`

	Email string `json:"email,omitempty"`

	ExtraFields interface{} `json:"extraFields,omitempty"`

	FirstLogin time.Time `json:"firstLogin,omitempty"`

	Group *CustomerGroup `json:"group,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	Hash string `json:"hash,omitempty"`

	Id string `json:"id,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	LastLogin time.Time `json:"lastLogin,omitempty"`

	LastOrderDate time.Time `json:"lastOrderDate,omitempty"`

	LastPaymentMethod *PaymentMethod `json:"lastPaymentMethod,omitempty"`

	LastPaymentMethodId string `json:"lastPaymentMethodId,omitempty"`

	LastUpdatedPasswordAt time.Time `json:"lastUpdatedPasswordAt,omitempty"`

	LegacyEncoder string `json:"legacyEncoder,omitempty"`

	LegacyPassword string `json:"legacyPassword,omitempty"`

	Name string `json:"name,omitempty"`

	Nickname string `json:"nickname,omitempty"`

	OrderCount float64 `json:"orderCount,omitempty"`

	OrderCustomers []OrderCustomer `json:"orderCustomers,omitempty"`

	OrderTotalAmount float64 `json:"orderTotalAmount,omitempty"`

	Password interface{} `json:"password,omitempty"`

	RemoteAddress interface{} `json:"remoteAddress,omitempty"`

	Roles []CustomerRole `json:"roles,omitempty"`

	TagIds interface{} `json:"tagIds,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UpdatedBy *User `json:"updatedBy,omitempty"`

	UpdatedById string `json:"updatedById,omitempty"`
}
