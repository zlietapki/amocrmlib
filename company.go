package amocrmlib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetCompaniesList(query string) (*apimodel.CompaniesListResponse, error) {
	path := "/api/v4/companies"
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

	compsListResp := new(apimodel.CompaniesListResponse)
	if err := compsListResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return compsListResp, nil
}

func (t *Token) GetCompanyByID(companyID int64) (*apimodel.Company, error) {
	path := fmt.Sprintf("/api/v4/companies/%d", companyID)
	resp, err := t.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, errors.Errorf("company not found: %d", companyID)
	}

	company := new(apimodel.Company)
	if err := company.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return company, nil
}

func (t *Token) AddCompany(name string) (*apimodel.CompaniesAddResponse, error) {
	comps := []apimodel.Company{
		{
			Name: name,
		},
	}
	bodyReq, err := json.Marshal(comps)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := t.DoRequest(http.MethodPost, "/api/v4/companies", bodyReq)
	if err != nil {
		return nil, err
	}

	compAddResponse := new(apimodel.CompaniesAddResponse)
	if err = compAddResponse.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return compAddResponse, nil
}

func (t *Token) EditCompany(companyID int64, companyData *apimodel.Company) (*apimodel.EditResponse, error) {
	req, err := companyData.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/companies/%d", companyID)
	resp, err := t.DoRequest(http.MethodPatch, path, req)
	if err != nil {
		return nil, err
	}

	editResp := new(apimodel.EditResponse)
	if err := editResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return editResp, nil
}

func (t *Token) SetCompanyResponsibleUserID(companyID int64, userID int64) (*apimodel.EditResponse, error) {
	companyData := &apimodel.Company{
		ResponsibleUserID: userID,
	}

	editResp, err := t.EditCompany(companyID, companyData)
	if err != nil {
		return nil, err
	}

	return editResp, nil
}
