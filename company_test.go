package amocrmlib

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zlietapki/amocrmlib/apimodel"
)

func TestToken_GetCompaniesList(t *testing.T) {
	t.Run("no query", func(t *testing.T) {
		_, err := token.GetCompaniesList("")
		require.NoError(t, err)
	})

	t.Run("some query", func(t *testing.T) {
		_, err := token.GetCompaniesList("")
		require.NoError(t, err)
	})

	t.Run("all companies", func(t *testing.T) {
		company := getRandomCompany()
		compsList, err := token.GetCompaniesList("")
		require.NoError(t, err)
		require.NotEmpty(t, compsList)

		found := false
		for _, c := range compsList.Embedded.Companies {
			if c.ID == company.ID {
				found = true
				break
			}
		}
		require.True(t, found)
	})

}

func TestToken_GetCompanyByID(t *testing.T) {
	testCompany := getRandomCompany()

	cases := []struct {
		name        string
		companyID   int64
		wantIsFound bool
	}{
		{
			name:        "existing company",
			companyID:   testCompany.ID,
			wantIsFound: true,
		},
		{
			name:        "non existing testCompany",
			companyID:   9999999999,
			wantIsFound: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			comp, err := token.GetCompanyByID(c.companyID)
			if c.wantIsFound {
				require.NotNil(t, comp)
				require.NoError(t, err)
				require.Equal(t, testCompany.Name, comp.Name)
			} else {
				require.Nil(t, comp)
				require.Error(t, err)
				require.Equal(t, fmt.Sprintf("company not found: %d", c.companyID), err.Error())
			}
		})
	}
}

func TestToken_AddCompany(t *testing.T) {
	testName := "some company"
	companyAddRes, err := token.AddCompany(testName)
	require.NoError(t, err)
	require.NotNil(t, companyAddRes)

	compID := companyAddRes.Embedded.Companies[0].ID
	company, err := token.GetCompanyByID(compID)
	require.NoError(t, err)
	require.Equal(t, testName, company.Name)
}

func TestToken_EditCompany(t *testing.T) {
	t.Run("edit name", func(t *testing.T) {
		company := getRandomCompany()
		oldName := company.Name
		compData := &apimodel.Company{
			Name: "new name",
		}
		editCompanyResp, err := token.EditCompany(company.ID, compData)
		require.NoError(t, err)
		require.Equal(t, company.ID, editCompanyResp.ID)

		compData2 := &apimodel.Company{
			Name: oldName,
		}
		_, err = token.EditCompany(company.ID, compData2)
		require.NoError(t, err)
	})
}

func TestToken_SetCompanyResponsibleUserID(t *testing.T) {
	testUser := getRandomUser()
	company := getRandomCompany()
	oldRespUserID := company.ResponsibleUserID
	editResp, err := token.SetCompanyResponsibleUserID(company.ID, testUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, editResp)

	//check
	comp2, err := token.GetCompanyByID(company.ID)
	require.NoError(t, err)
	require.Equal(t, testUser.ID, comp2.ResponsibleUserID)

	//cleanup
	editResp, err = token.SetCompanyResponsibleUserID(company.ID, oldRespUserID)
	require.NoError(t, err)
	require.NotEmpty(t, editResp)
}

func getRandomCompany() *apimodel.Company {
	companiesList, _ := token.GetCompaniesList("")
	if companiesList == nil {
		_, _ = token.AddCompany("test company")
		companiesList, _ = token.GetCompaniesList("")
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(companiesList.Embedded.Companies))
	return companiesList.Embedded.Companies[n]
}
