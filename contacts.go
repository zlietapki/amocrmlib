package amocrmlib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetContactsList(query string) (*apimodel.ContactsListResponse, error) {
	path := "/api/v4/contacts"
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

	contactsListResp := new(apimodel.ContactsListResponse)
	if err = contactsListResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return contactsListResp, nil
}

func (t *Token) GetContactByID(contactID int64) (*apimodel.Contact, error) {
	path := fmt.Sprintf("/api/v4/contacts/%d", contactID)
	resp, err := t.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, errors.Errorf("contact not found: %d", contactID)
	}

	contact := new(apimodel.Contact)
	if err = contact.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return contact, nil
}

func (t *Token) AddContact(name string) (*apimodel.ContactsAddResponse, error) {
	contacts := []apimodel.Contact{
		{
			Name: name,
		},
	}
	bodyReq, err := json.Marshal(contacts)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := t.DoRequest(http.MethodPost, "/api/v4/contacts", bodyReq)
	if err != nil {
		return nil, err
	}

	contactsAddedresp := new(apimodel.ContactsAddResponse)
	if err = contactsAddedresp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return contactsAddedresp, nil
}

func (t *Token) EditContact(contactID int64, apiData *apimodel.Contact) (*apimodel.EditResponse, error) {
	req, err := apiData.MarshalBinary()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path := fmt.Sprintf("/api/v4/contacts/%d", contactID)
	resp, err := t.DoRequest(http.MethodPatch, path, req)
	if err != nil {
		return nil, err
	}

	editContactResp := new(apimodel.EditResponse)
	if err := editContactResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := editContactResp.Validate(nil); err != nil {
		return nil, errors.WithStack(err)
	}

	return editContactResp, nil
}

func (t *Token) SetContactResponsibleUserID(contactID int64, userID int64) (*apimodel.EditResponse, error) {
	data := &apimodel.Contact{
		ResponsibleUserID: userID,
	}
	editContactResp, err := t.EditContact(contactID, data)
	if err != nil {
		return nil, err
	}

	return editContactResp, nil
}
