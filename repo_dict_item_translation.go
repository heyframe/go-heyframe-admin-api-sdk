package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type DictItemTranslationRepository struct {
	*GenericRepository[DictItemTranslation]
}

func NewDictItemTranslationRepository(client *Client) *DictItemTranslationRepository {
	return &DictItemTranslationRepository{
		GenericRepository: NewGenericRepository[DictItemTranslation](client),
	}
}

func (t *DictItemTranslationRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[DictItemTranslation], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "dict-item-translation")
}

func (t *DictItemTranslationRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[DictItemTranslation], *http.Response, error) {
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

func (t *DictItemTranslationRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "dict-item-translation")
}

func (t *DictItemTranslationRepository) Upsert(ctx ApiContext, entity []DictItemTranslation) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "dict_item_translation")
}

func (t *DictItemTranslationRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "dict_item_translation")
}

type DictItemTranslation struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	CustomFields interface{} `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	DictItem *DictItem `json:"dictItem,omitempty"`

	DictItemId string `json:"dictItemId,omitempty"`

	Label string `json:"label,omitempty"`

	Language *Language `json:"language,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	Position float64 `json:"position,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
