package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type LanguageRepository struct {
	*GenericRepository[Language]
}

func NewLanguageRepository(client *Client) *LanguageRepository {
	return &LanguageRepository{
		GenericRepository: NewGenericRepository[Language](client),
	}
}

func (t *LanguageRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Language], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "language")
}

func (t *LanguageRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Language], *http.Response, error) {
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

func (t *LanguageRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "language")
}

func (t *LanguageRepository) Upsert(ctx ApiContext, entity []Language) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "language")
}

func (t *LanguageRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "language")
}

type Language struct {
	Active bool `json:"active,omitempty"`

	ChannelDefaultAssignments []Channel `json:"channelDefaultAssignments,omitempty"`

	ChannelDomains []ChannelDomain `json:"channelDomains,omitempty"`

	ChannelTranslations []ChannelTranslation `json:"channelTranslations,omitempty"`

	ChannelTypeTranslations []ChannelTypeTranslation `json:"channelTypeTranslations,omitempty"`

	Channels []Channel `json:"channels,omitempty"`

	Children []Language `json:"children,omitempty"`

	CountryStateTranslations []CountryStateTranslation `json:"countryStateTranslations,omitempty"`

	CountryTranslations []CountryTranslation `json:"countryTranslations,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CurrencyTranslations []CurrencyTranslation `json:"currencyTranslations,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	CustomerGroupTranslations []CustomerGroupTranslation `json:"customerGroupTranslations,omitempty"`

	Customers []Customer `json:"customers,omitempty"`

	Id string `json:"id,omitempty"`

	Locale *Locale `json:"locale,omitempty"`

	LocaleId string `json:"localeId,omitempty"`

	LocaleTranslations []LocaleTranslation `json:"localeTranslations,omitempty"`

	MediaTranslations []MediaTranslation `json:"mediaTranslations,omitempty"`

	Name string `json:"name,omitempty"`

	NumberRangeTranslations []NumberRangeTranslation `json:"numberRangeTranslations,omitempty"`

	NumberRangeTypeTranslations []NumberRangeTypeTranslation `json:"numberRangeTypeTranslations,omitempty"`

	Orders []Order `json:"orders,omitempty"`

	Parent *Language `json:"parent,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	PaymentMethodTranslations []PaymentMethodTranslation `json:"paymentMethodTranslations,omitempty"`

	PluginTranslations []PluginTranslation `json:"pluginTranslations,omitempty"`

	ProductTranslations []ProductTranslation `json:"productTranslations,omitempty"`

	PropertyGroupOptionTranslations []PropertyGroupOptionTranslation `json:"propertyGroupOptionTranslations,omitempty"`

	PropertyGroupTranslations []PropertyGroupTranslation `json:"propertyGroupTranslations,omitempty"`

	StateMachineStateTranslations []StateMachineStateTranslation `json:"stateMachineStateTranslations,omitempty"`

	StateMachineTranslations []StateMachineTranslation `json:"stateMachineTranslations,omitempty"`

	ThemeTranslations []ThemeTranslation `json:"themeTranslations,omitempty"`

	TranslationCode *Locale `json:"translationCode,omitempty"`

	TranslationCodeId string `json:"translationCodeId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
