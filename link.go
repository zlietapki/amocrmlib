package amocrmlib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetLinksList(entityType string, entityID int64) ([]*apimodel.Link, error) {
	path := fmt.Sprintf("/api/v4/%s/%d/links", entityType, entityID)
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

func (t *Token) AddLink(toEntityType string, toEntityID int64, fromEntityType string, fromEntityID int64) error {
	req, err := json.Marshal(apimodel.EntityLinksAdd{
		{
			ToEntityID:   toEntityID,
			ToEntityType: toEntityType,
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/%s/%d/link", fromEntityType, fromEntityID)
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

func (t *Token) Unlink(toEntityType string, toEntityID int64, fromEntityType string, fromEntityID int64) error {
	req, err := json.Marshal(apimodel.LinkUnlinks{
		{
			ToEntityID:   toEntityID,
			ToEntityType: toEntityType,
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/%s/%d/unlink", fromEntityType, fromEntityID)
	_, err = t.DoRequest(http.MethodPost, path, req)
	if err != nil {
		return err
	}

	return nil
}
