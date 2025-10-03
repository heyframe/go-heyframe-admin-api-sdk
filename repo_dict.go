package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type DictRepository struct {
	*GenericRepository[Dict]
}

func NewDictRepository(client *Client) *DictRepository {
	return &DictRepository{
		GenericRepository: NewGenericRepository[Dict](client),
	}
}

func (t *DictRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[Dict], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "dict")
}

func (t *DictRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[Dict], *http.Response, error) {
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

func (t *DictRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "dict")
}

func (t *DictRepository) Upsert(ctx ApiContext, entity []Dict) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "dict")
}

func (t *DictRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "dict")
}

type Dict struct {
	Active bool `json:"active,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	Items []DictItem `json:"items,omitempty"`

	Key string `json:"key,omitempty"`

	Position float64 `json:"position,omitempty"`

	Translated interface{} `json:"translated,omitempty"`

	Translations []DictTranslation `json:"translations,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
