package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type ChannelRepository struct {
	*GenericRepository[Channel]
}

func NewChannelRepository(client *Client) *ChannelRepository {
	return &ChannelRepository{
		GenericRepository: NewGenericRepository[Channel](client),
	}
}

func (t *ChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Channel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "channel")
}

func (t *ChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Channel], *http.Response, error) {
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

func (t *ChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "channel")
}

func (t *ChannelRepository) Upsert(ctx ApiContext, entity []Channel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "channel")
}

func (t *ChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "channel")
}

type Channel struct {
	AccessKey string `json:"accessKey,omitempty"`

	Active bool `json:"active,omitempty"`

	BoundCustomers []Customer `json:"boundCustomers,omitempty"`

	Configuration interface{} `json:"configuration,omitempty"`

	Countries []Country `json:"countries,omitempty"`

	Country *Country `json:"country,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Currencies []Currency `json:"currencies,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	CurrencyId string `json:"currencyId,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CustomerGroup *CustomerGroup `json:"customerGroup,omitempty"`

	CustomerGroupId string `json:"customerGroupId,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	Domains []ChannelDomain `json:"domains,omitempty"`

	FooterNavigationId string `json:"footerNavigationId,omitempty"`

	FooterNavigationVersionId string `json:"footerNavigationVersionId,omitempty"`

	HomeEnabled bool `json:"homeEnabled,omitempty"`

	HomeKeywords string `json:"homeKeywords,omitempty"`

	HomeMetaDescription string `json:"homeMetaDescription,omitempty"`

	HomeMetaTitle string `json:"homeMetaTitle,omitempty"`

	HomeName string `json:"homeName,omitempty"`

	HomeSlotConfig interface{} `json:"homeSlotConfig,omitempty"`

	Id string `json:"id,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Languages []Language `json:"languages,omitempty"`

	Maintenance bool `json:"maintenance,omitempty"`

	MaintenanceIpWhitelist interface{} `json:"maintenanceIpWhitelist,omitempty"`

	Name string `json:"name,omitempty"`

	NavigationDepth float64 `json:"navigationDepth,omitempty"`

	NavigationId string `json:"navigationId,omitempty"`

	NavigationVersionId string `json:"navigationVersionId,omitempty"`

	NumberRangeChannels []NumberRangeChannel `json:"numberRangeChannels,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	PaymentMethodId string `json:"paymentMethodId,omitempty"`

	PaymentMethodIds interface{} `json:"paymentMethodIds,omitempty"`

	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`

	ProductVisibilities []ProductVisibility `json:"productVisibilities,omitempty"`

	ServiceNavigationId string `json:"serviceNavigationId,omitempty"`

	ServiceNavigationVersionId string `json:"serviceNavigationVersionId,omitempty"`

	ShortName string `json:"shortName,omitempty"`

	SystemConfigs []SystemConfig `json:"systemConfigs,omitempty"`

	Themes []Theme `json:"themes,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []ChannelTranslation `json:"translations,omitempty"`

	Type *ChannelType `json:"type,omitempty"`

	TypeId string `json:"typeId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
