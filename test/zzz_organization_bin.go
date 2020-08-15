package test

import (
	"fmt"
	"invest/model"
	"testing"
)

func TestOrganizationByBin(t *testing.T) {
	var o = model.Organization{}
	fmt.Println(o.Get_and_assign_info_on_organization_by_bin())

}
