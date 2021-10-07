amocrmlib
=========

amoCRM module for Golang

Install
-------

```bash
go get github.com/zlietapki/amocrmlib
```

Basic usage
-----------

Obtain `token.json` with <https://github.com/zlietapki/amocrmtoken>

```go
package main

import (
	"github.com/zlietapki/amocrmlib"
)

func main() {
	token, err := amocrmlib.LoadTokenFile("token.json")
	if err != nil {
		panic(err)
    }

	usersListResp, err := token.GetUsersList()
	println(usersListResp.Embedded.Users[0].Name)
}
```

Functions
---------

```text
token.GetCompaniesList(query string) (*apimodel.CompaniesListResponse, error)
token.GetCompanyByID(companyID int64) (*apimodel.Company, error)
token.AddCompany(name string) (*apimodel.CompaniesAddResponse, error)
token.EditCompany(companyID int64, companyData *apimodel.Company) (*apimodel.EditResponse, error)
token.SetCompanyResponsibleUserID(companyID int64, userID int64) (*apimodel.EditResponse, error)

token.GetContactsList(query string) (*apimodel.ContactsListResponse, error)
token.GetContactByID(contactID int64) (*apimodel.Contact, error)
token.AddContact(name string) (*apimodel.ContactsAddResponse, error)
token.EditContact(contactID int64, apiData *apimodel.Contact) (*apimodel.EditResponse, error)
token.SetContactResponsibleUserID(contactID int64, userID int64) (*apimodel.EditResponse, error)

token.GetCustomFieldsList(entityType string) (*apimodel.CustomFieldsListResponse, error)
token.GetCustomFieldID(entityType string, fieldName string) (int64, error)

token.GetLeadsList(query string) ([]*apimodel.Lead, error)
token.GetLeadByID(leadID int64) (*apimodel.Lead, error)
token.AddLead(name string) (*apimodel.LeadsAddResponse, error)
token.EditLead(leadID int64, apiData *apimodel.Lead) error

```