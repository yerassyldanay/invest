package intest

import (
	"invest/model"
	"testing"
)

type organizationTestStruct struct {
	Organization 		model.Organization
	ErrMsg				string
}

func TestOrganization(t *testing.T) {
	var cases = []organizationTestStruct{
		{
			Organization: model.Organization {
				Bin: "190240004115",
			},
			ErrMsg: "",
		},
	}

	for _, c := range cases {
		msg := c.Organization.Create_or_get_organization_from_db_by_bin(model.GetDB())
		if msg.ErrMsg != c.ErrMsg {
			t.Error("expected: ", c.ErrMsg, " but got: ", msg.ErrMsg)
		}
	}
}
