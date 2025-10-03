package go_heyframe_admin_sdk

import (
	"net/http"

	"time"
)

type NumberRangeChannelRepository struct {
	*GenericRepository[NumberRangeChannel]
}

func NewNumberRangeChannelRepository(client *Client) *NumberRangeChannelRepository {
	return &NumberRangeChannelRepository{
		GenericRepository: NewGenericRepository[NumberRangeChannel](client),
	}
}

func (t *NumberRangeChannelRepository) Search(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeChannel], *http.Response, error) {
	return t.GenericRepository.Search(ctx, criteria, "number-range-channel")
}

func (t *NumberRangeChannelRepository) SearchAll(ctx ApiContext, criteria Criteria) (*EntityCollection[NumberRangeChannel], *http.Response, error) {
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

func (t *NumberRangeChannelRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	return t.GenericRepository.SearchIds(ctx, criteria, "number-range-channel")
}

func (t *NumberRangeChannelRepository) Upsert(ctx ApiContext, entity []NumberRangeChannel) (*http.Response, error) {
	return t.GenericRepository.Upsert(ctx, entity, "number_range_channel")
}

func (t *NumberRangeChannelRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	return t.GenericRepository.Delete(ctx, ids, "number_range_channel")
}

type NumberRangeChannel struct {
	Channel *Channel `json:"channel,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Id string `json:"id,omitempty"`

	NumberRange *NumberRange `json:"numberRange,omitempty"`

	NumberRangeId string `json:"numberRangeId,omitempty"`

	NumberRangeType *NumberRangeType `json:"numberRangeType,omitempty"`

	NumberRangeTypeId string `json:"numberRangeTypeId,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
