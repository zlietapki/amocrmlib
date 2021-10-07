package amocrmlib

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetCustomFieldsList(entityType string) (*apimodel.CustomFieldsListResponse, error) {
	path := fmt.Sprintf("/api/v4/%s/custom_fields", entityType)
	resp, err := t.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	apiCustomFieldsList := new(apimodel.CustomFieldsListResponse)
	if err := apiCustomFieldsList.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return apiCustomFieldsList, nil
}

func (t *Token) GetCustomFieldID(entityType string, fieldName string) (int64, error) {
	fieldsListResp, err := t.GetCustomFieldsList(entityType)
	if err != nil {
		return 0, err
	}

	for _, f := range fieldsListResp.Embedded.CustomFields {
		if f.Name == fieldName {
			return f.ID, nil
		}
	}

	return 0, errors.Errorf("field not found: %s entityType: %s", fieldName, entityType)
}
