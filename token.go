package amocrmlib

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/amoerrors"
	"github.com/zlietapki/amocrmlib/apimodel"
	"golang.org/x/sync/singleflight"
)

type Token struct {
	*apimodel.Token
	TokenFile string

	requestGroup singleflight.Group
}

func NewToken(apiToken *apimodel.Token, tokenFile string) *Token {
	return &Token{
		Token:     apiToken,
		TokenFile: tokenFile,
	}
}

func LoadTokenFile(file string) (*Token, error) {
	tokenData, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	apiToken := new(apimodel.Token)
	if err = apiToken.UnmarshalBinary(tokenData); err != nil {
		return nil, errors.WithStack(err)
	}

	if err = apiToken.Validate(nil); err != nil {
		return nil, err
	}

	token := NewToken(apiToken, file)
	return token, nil
}

func (t *Token) SaveTokenFile() error {
	data, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return errors.WithStack(err)
	}
	if err = ioutil.WriteFile(t.TokenFile, data, 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (t *Token) Expired() bool {
	return time.Now().After(time.Time(t.AccessTokenExpire))
}

func (t *Token) DoRequest(method string, url string, body []byte) ([]byte, error) {
	if t.Expired() {
		_, err, _ := t.requestGroup.Do("RefreshToken", func() (interface{}, error) {
			if err := t.RefreshAccessToken(); err != nil {
				return nil, err
			}
			if err := t.SaveTokenFile(); err != nil {
				return nil, err
			}
			return nil, nil
		})
		if err != nil {
			return nil, err
		}
	}

	realUrl := t.Domain + url
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(method, realUrl, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, realUrl, nil)
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Add("Authorization", "Bearer "+t.AccessToken)
	if method == http.MethodPost || method == http.MethodPatch {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := new(http.Client).Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// amo api response error
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if res.Header.Get("Content-Type") == "application/problem+json" {
			return nil, amoerrors.NewErrorApi(respBody)
		}
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}

func (t *Token) RefreshAccessToken() error {
	oAuth2Req := &apimodel.OAuth2Req{
		ClientSecret: t.ClientSecret,
		ClientID:     t.ClientID,
		RedirectURI:  t.RedirectURI,
		GrantType:    "refresh_token",
		RefreshToken: t.RefreshToken,
	}
	oAuth2Res, err := oAuth2Req.Do(t.Domain)
	if err != nil {
		return err
	}

	apiToken := apimodel.NewToken(
		t.Domain,
		t.RedirectURI,
		t.ClientSecret,
		t.ClientID,
		oAuth2Res.AccessToken,
		oAuth2Res.ExpiresIn,
		oAuth2Res.RefreshToken,
	)
	t.Token = apiToken
	return nil
}
