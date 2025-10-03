package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type CategoryRepository struct {
	*GenericRepository[Category]
}

func NewCategoryRepository(client *Client) *CategoryRepository {
	return &CategoryRepository{
		GenericRepository: NewGenericRepository[Category](client),
	}
}

func (t *CategoryRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Category], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "category")
}

func (t *CategoryRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Category], *http.Response, error) {
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

func (t *CategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "category")
}

func (t *CategoryRepository) Upsert(ctx ApiContext, entity []Category) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "category")
}

func (t *CategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "category")
}

type Category struct {
	Active bool `json:"active,omitempty"`

	AfterCategoryId string `json:"afterCategoryId,omitempty"`

	AfterCategoryVersionId string `json:"afterCategoryVersionId,omitempty"`

	AutoIncrement float64 `json:"autoIncrement,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	Children []Category `json:"children,omitempty"`

	CmsPage *CmsPage `json:"cmsPage,omitempty"`

	CmsPageId string `json:"cmsPageId,omitempty"`

	CmsPageIdSwitched bool `json:"cmsPageIdSwitched,omitempty"`

	CmsPageVersionId string `json:"cmsPageVersionId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	DisplayNestedProducts bool `json:"displayNestedProducts,omitempty"`

	Id string `json:"id,omitempty"`

	Level float64 `json:"level,omitempty"`

	Media *Media `json:"media,omitempty"`

	MediaId string `json:"mediaId,omitempty"`

	Name string `json:"name,omitempty"`

	NestedProducts []Product `json:"nestedProducts,omitempty"`

	Parent *Category `json:"parent,omitempty"`

	ParentId string `json:"parentId,omitempty"`

	ParentVersionId string `json:"parentVersionId,omitempty"`

	Path string `json:"path,omitempty"`

	Products []Product `json:"products,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	TechnicalName string `json:"technicalName,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []CategoryTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	VersionId string `json:"versionId,omitempty"`

	Visible bool `json:"visible,omitempty"`

	VisibleChildCount float64 `json:"visibleChildCount,omitempty"`
}
