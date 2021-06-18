package app

import (
	"github.com/yerassyldanay/invest/service"
	"github.com/yerassyldanay/invest/utils/constants"
	"github.com/yerassyldanay/invest/utils/errormsg"
	"github.com/yerassyldanay/invest/utils/message"
	"net/http"
)

var OnlyInvestorCanAccess = func(next http.Handler) (http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {

		// only investor can create a project
		roleName := service.Get_header_parameter(r, constants.KeyRoleName, "").(string)
		if roleName != constants.RoleInvestor {
			message.Respond(w, r, message.Msg{errormsg.ErrorMethodNotAllowed, 405, "", "role must be investor. role is " + roleName})
			return
		}

		next.ServeHTTP(w, r)

	}
}