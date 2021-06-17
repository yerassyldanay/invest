package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"testing"
)

func TestOrganization(t *testing.T) {
	var organization = model.Organization {
		Bin: "190240004115",
		Lang: "rus",
	}

	// create organization
	msg := organization.Create_or_get_organization_from_db_by_bin(model.GetDB())

	// check
	require.Zero(t, msg.ErrMsg)

	// get org
	newOrg := model.Organization{
		Bin: organization.Bin,
		Lang: "rus",
	}
	err := newOrg.OnlyGetByBinAndLang(model.GetDB())

	// check
	require.NoError(t, err)
	require.Equal(t, organization.Bin, newOrg.Bin)
	require.Equal(t, organization.Name, newOrg.Name)
	require.Equal(t, organization.Fio, newOrg.Fio)
}
