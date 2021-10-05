package amocrmlib

import (
	"math/rand"
	"time"

	"github.com/zlietapki/amocrmlib/apimodel"
)

func getRandomUser() *apimodel.User {
	usersList, _ := token.GetUsersList()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(usersList.Embedded.Users))
	return usersList.Embedded.Users[n]
}
