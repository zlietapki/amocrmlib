package amocrmlib

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func (t *Token) GetUsersList() (*apimodel.UsersListResponse, error) {
	resp, err := t.DoRequest(http.MethodGet, "/api/v4/users", nil)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}

	usersListResp := new(apimodel.UsersListResponse)
	if err = usersListResp.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return usersListResp, nil
}

func (t *Token) GetUserById(userId int64) (*apimodel.User, error) {
	path := fmt.Sprintf("/api/v4/users/%d", userId)
	resp, err := t.DoRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	//возвращает код 404, если не найдено, а не пустой ответ

	apiUser := new(apimodel.User)
	if err = apiUser.UnmarshalBinary(resp); err != nil {
		return nil, errors.WithStack(err)
	}

	return apiUser, nil
}

func (t *Token) GetUserByEmail(email string) (*apimodel.User, error) {
	usersList, err := t.GetUsersList()
	if err != nil {
		return nil, err
	}

	for _, user := range usersList.Embedded.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.Errorf("User not found by email: %s", email)
}
