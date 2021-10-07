package amocrmlib

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func TestToken_GetContactsList(t *testing.T) {
	testContact := getRandomContact()
	contListResp, err := token.GetContactsList("")
	require.NoError(t, err)

	found := false
	for _, cont := range contListResp.Embedded.Contacts {
		if cont.ID == testContact.ID {
			found = true
		}
	}
	require.True(t, found)
}

func TestToken_GetContactByID(t *testing.T) {
	testContact := getRandomContact()
	cont, err := token.GetContactByID(testContact.ID)
	require.NoError(t, err)
	require.Equal(t, testContact.ID, cont.ID)
}

func TestToken_AddContact(t *testing.T) {
	testName := "some contact"
	contAddresp, err := token.AddContact(testName)
	require.NoError(t, err)
	contID := contAddresp.Embedded.Contacts[0].ID

	cont2, err := token.GetContactByID(contID)
	require.NoError(t, err)
	require.Equal(t, testName, cont2.Name)
}

func TestToken_EditContact(t *testing.T) {
	testName := "new name"
	testContact := getRandomContact()
	contData := &apimodel.Contact{
		Name: testName,
	}
	editContactResp, err := token.EditContact(testContact.ID, contData)
	require.NoError(t, err)
	require.Equal(t, testContact.ID, editContactResp.ID)

	cont2, err := token.GetContactByID(testContact.ID)
	require.NoError(t, err)
	require.Equal(t, testName, cont2.Name)
}

func TestToken_SetContactResponsibleUserID(t *testing.T) {
	testContact := getRandomContact()
	oldUserID := testContact.ResponsibleUserID
	testUser := getRandomUser()

	editResp, err := token.SetContactResponsibleUserID(testContact.ID, testUser.ID)
	require.NoError(t, err)
	require.Equal(t, testContact.ID, editResp.ID)

	//check
	cont2, err := token.GetContactByID(testContact.ID)
	require.NoError(t, err)
	require.Equal(t, testUser.ID, cont2.ResponsibleUserID)

	//cleanup
	_, err = token.SetContactResponsibleUserID(testContact.ID, oldUserID)
	require.NoError(t, err)
}

func getRandomContact() *apimodel.Contact {
	contactsList, _ := token.GetContactsList("")
	if contactsList == nil {
		_, _ = token.AddContact("test contact")
		contactsList, _ = token.GetContactsList("")
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(contactsList.Embedded.Contacts))
	return contactsList.Embedded.Contacts[n]
}
