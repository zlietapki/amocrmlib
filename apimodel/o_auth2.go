package apimodel

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/amoerrors"
)

func (o *OAuth2Req) Do(domain string) (*OAuth2Res, error) {
	body, err := o.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	oauth2Url := domain + "/oauth2/access_token"
	req, err := http.NewRequest(http.MethodPost, oauth2Url, bytes.NewReader(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// amo api response error
	if resp.StatusCode != http.StatusOK {
		if resp.Header.Get("Content-Type") == "application/problem+json" {
			return nil, amoerrors.NewErrorApi(respBody)
		}
		return nil, errors.Errorf("%s", respBody)
	}

	auth2resp := new(OAuth2Res)
	if err = json.Unmarshal(respBody, auth2resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return auth2resp, nil
}
