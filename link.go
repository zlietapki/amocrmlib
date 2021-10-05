package amocrmlib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetLinksList(entityType string, entityId int64) ([]*apimodel.Link, error) {
	path := fmt.Sprintf("/api/v4/%s/%d/links", entityType, entityId)
	resp, err := t.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	linksListResp := new(apimodel.LinksListResponse)
	if err = linksListResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return linksListResp.Embedded.Links, nil
}

func (t *Token) AddLink(toEntityType string, toEntityId int64, fromEntityType string, fromEntityId int64) error {
	req, err := json.Marshal(apimodel.EntityLinksAdd{
		{
			ToEntityID:   toEntityId,
			ToEntityType: toEntityType,
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/%s/%d/link", fromEntityType, fromEntityId)
	resp, err := t.DoRequest(http.MethodPost, path, req)
	if err != nil {
		return err
	}

	linkAddRes := new(apimodel.LinksListResponse)
	if err := linkAddRes.UnmarshalBinary(resp); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (t *Token) Unlink(toEntityType string, toEntityId int64, fromEntityType string, fromEntityId int64) error {
	req, err := json.Marshal(apimodel.LinkUnlinks{
		{
			ToEntityID:   toEntityId,
			ToEntityType: toEntityType,
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/%s/%d/unlink", fromEntityType, fromEntityId)
	_, err = t.DoRequest(http.MethodPost, path, req)
	if err != nil {
		return err
	}

	return nil
}
