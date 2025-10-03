package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type NavigationTranslationRepository struct {
	*GenericRepository[NavigationTranslation]
}

func NewNavigationTranslationRepository(client *Client) *NavigationTranslationRepository {
	return &NavigationTranslationRepository{
		GenericRepository: NewGenericRepository[NavigationTranslation](client),
	}
}

func (t *NavigationTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NavigationTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "navigation-translation")
}

func (t *NavigationTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NavigationTranslation], *http.Response, error) {
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

func (t *NavigationTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "navigation-translation")
}

func (t *NavigationTranslationRepository) Upsert(ctx ApiContext, entity []NavigationTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "navigation_translation")
}

func (t *NavigationTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "navigation_translation")
}

type NavigationTranslation struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	ExternalLink string `json:"externalLink,omitempty"`

	InternalLink string `json:"internalLink,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	LinkNewTab bool `json:"linkNewTab,omitempty"`

	LinkType string `json:"linkType,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	Name string `json:"name,omitempty"`

	Navigation *Navigation `json:"navigation,omitempty"`

	NavigationId string `json:"navigationId,omitempty"`

	NavigationVersionId string `json:"navigationVersionId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
