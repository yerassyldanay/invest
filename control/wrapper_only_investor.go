package control

import (
	"invest/service"
	"invest/utils"
	"net/http"
)

var OnlyInvestorCanAccess = func(next http.Handler) (http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {

		// only investor can create a project
		roleName := service.Get_header_parameter(r, utils.KeyRoleName, "").(string)
		if roleName != utils.RoleInvestor {
			utils.Respond(w, r, utils.Msg{utils.ErrorMethodNotAllowed, 405, "", "role must be investor. role is " + roleName})
			return
		}

		next.ServeHTTP(w, r)

	}
}
