package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type DictItemRepository struct {
	*GenericRepository[DictItem]
}

func NewDictItemRepository(client *Client) *DictItemRepository {
	return &DictItemRepository{
		GenericRepository: NewGenericRepository[DictItem](client),
	}
}

func (t *DictItemRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DictItem], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "dict-item")
}

func (t *DictItemRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DictItem], *http.Response, error) {
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

func (t *DictItemRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "dict-item")
}

func (t *DictItemRepository) Upsert(ctx ApiContext, entity []DictItem) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "dict_item")
}

func (t *DictItemRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "dict_item")
}

type DictItem struct {
	Active bool `json:"active,omitempty"`

	ChildCount float64 `json:"childCount,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	Dict *Dict `json:"dict,omitempty"`

	DictId string `json:"dictId,omitempty"`

	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	Level float64 `json:"level,omitempty"`

	Path string `json:"path,omitempty"`

	Position float64 `json:"position,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []DictItemTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	Value string `json:"value,omitempty"`
}
