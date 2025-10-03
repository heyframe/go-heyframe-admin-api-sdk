package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type NavigationRepository struct {
	*GenericRepository[Navigation]
}

func NewNavigationRepository(client *Client) *NavigationRepository {
	return &NavigationRepository{
		GenericRepository: NewGenericRepository[Navigation](client),
	}
}

func (t *NavigationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Navigation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "navigation")
}

func (t *NavigationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Navigation], *http.Response, error) {
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

func (t *NavigationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "navigation")
}

func (t *NavigationRepository) Upsert(ctx ApiContext, entity []Navigation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "navigation")
}

func (t *NavigationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "navigation")
}

type Navigation struct {
	Active bool `json:"active,omitempty"`

	AfterNavigationId string `json:"afterNavigationId,omitempty"`

	AfterNavigationVersionId string `json:"afterNavigationVersionId,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	Children []Navigation `json:"children,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	ExternalLink string `json:"externalLink,omitempty"`

	FooterChannels []Channel `json:"footerChannels,omitempty"`

	Id string `json:"id,omitempty"`

	InternalLink string `json:"internalLink,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	Level float64 `json:"level,omitempty"`

	LinkNewTab bool `json:"linkNewTab,omitempty"`

	LinkType string `json:"linkType,omitempty"`

	MetaDescription string `json:"metaDescription,omitempty"`

	MetaTitle string `json:"metaTitle,omitempty"`

	Name string `json:"name,omitempty"`

	NavigationChannels []Channel `json:"navigationChannels,omitempty"`

	Parent *Navigation `json:"parent,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Path string `json:"path,omitempty"`

	ServiceChannels []Channel `json:"serviceChannels,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []NavigationTranslation `json:"translations,omitempty"`

	Type string `json:"type,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Visible bool `json:"visible,omitempty"`
}
