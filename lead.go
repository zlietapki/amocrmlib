package amocrmlib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetLeadsList(query string) ([]*apimodel.Lead, error) {
	path := "/api/v4/leads"
	if query != "" {
		path += fmt.Sprintf("?query=%s", url.QueryEscape(query))
	}
	resp, err := t.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}

	leadsList := new(apimodel.LeadsListResponse)
	if err = json.Unmarshal(resp, leadsList); err != nil {
		return nil, errors.WithStack(err)
	}

	return leadsList.Embedded.Leads, nil
}

func (t *Token) GetLeadByID(leadID int64) (*apimodel.Lead, error) {
	path := fmt.Sprintf("/api/v4/leads/%d", leadID)
	resp, err := t.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, errors.Errorf("lead not found: %d", leadID)
	}

	lead := new(apimodel.Lead)
	if err = lead.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return lead, nil
}

func (t *Token) AddLead(name string) (*apimodel.LeadsAddResponse, error) {
	data := &[]apimodel.Lead{
		{
			Name: name,
		},
	}

	bodyReq, err := json.Marshal(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := t.DoRequest(http.MethodPost, "/api/v4/leads", bodyReq)
	if err != nil {
		return nil, err
	}

	addRes := new(apimodel.LeadsAddResponse)
	if err = json.Unmarshal(resp, addRes); err != nil {
		return nil, errors.WithStack(err)
	}

	return addRes, nil
}

func (t *Token) EditLead(leadID int64, apiData *apimodel.Lead) error {
	req, err := apiData.MarshalBinary()
	if err != nil {
		return errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/leads/%d", leadID)
	resp, err := t.DoRequest(http.MethodPatch, path, req)
	if err != nil {
		return err
	}

	apiEditResp := new(apimodel.EditResponse)
	if err := apiEditResp.UnmarshalBinary(resp); err != nil {
		return errors.WithStack(err)
	}
	if err := apiEditResp.Validate(nil); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
