package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type OrderRepository struct {
	*GenericRepository[Order]
}

func NewOrderRepository(client *Client) *OrderRepository {
	return &OrderRepository{
		GenericRepository: NewGenericRepository[Order](client),
	}
}

func (t *OrderRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Order], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "order")
}

func (t *OrderRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Order], *http.Response, error) {
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

func (t *OrderRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "order")
}

func (t *OrderRepository) Upsert(ctx ApiContext, entity []Order) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "order")
}

func (t *OrderRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "order")
}

type Order struct {
	AmountTotal float64 `json:"amountTotal,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CreatedBy *User `json:"createdBy,omitempty"`

	CreatedById string `json:"createdById,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	CurrencyFactor float64 `json:"currencyFactor,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	DeepLinkCode string `json:"deepLinkCode,omitempty"`

	Id string `json:"id,omitempty"`

	ItemRounding interface{} `json:"itemRounding,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	LineItems []OrderLineItem `json:"lineItems,omitempty"`

	OrderCustomer *OrderCustomer `json:"orderCustomer,omitempty"`

	OrderDate time.Time `json:"orderDate,omitempty"`

	OrderDateTime time.Time `json:"orderDateTime,omitempty"`

	OrderNumber string `json:"orderNumber,omitempty"`

	PositionPrice float64 `json:"positionPrice,omitempty"`

	Price interface{} `json:"price,omitempty"`

	PrimaryOrderTransaction *OrderTransaction `json:"primaryOrderTransaction,omitempty"`

	PrimaryOrderTransactionId string `json:"primaryOrderTransactionId,omitempty"`

	PrimaryOrderTransactionVersionId string `json:"primaryOrderTransactionVersionId,omitempty"`

	RuleIds interface{} `json:"ruleIds,omitempty"`

	Source string `json:"source,omitempty"`

	StateId string `json:"stateId,omitempty"`

	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	TotalRounding interface{} `json:"totalRounding,omitempty"`

	Transactions []OrderTransaction `json:"transactions,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	UpdatedBy *User `json:"updatedBy,omitempty"`

	UpdatedById string `json:"updatedById,omitempty"`

	VersionId string `json:"versionId,omitempty"`
}
