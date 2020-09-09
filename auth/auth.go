package auth

import (
	"github.com/dgrijalva/jwt-go"
	"invest/model"
	"invest/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var fname = "JWT_TOKEN"

		/*
			there are some urls that do not require authentication.
				E.g. getting static files or sign in/up urls
		 */
		for _, url := range utils.NoNeedToAuth {
			//fmt.Println(url, r.URL.Path)
			if url == r.URL.Path {
				//fmt.Println(url, r.URL.Path)
				next.ServeHTTP(w, r)
				return
			}
		}

		/*
			Authentication header must contain JWT token
		 */
		var tokenHeader = r.Header.Get("Authorization")

		var splits = strings.Split(tokenHeader, " ")
		if len(splits) != 2 {
			utils.Respond(w, r, &utils.Msg{
				Message: map[string]interface{}{
					"eng": "invalid token",
				},
				Status:  http.StatusMisdirectedRequest,
				Fname:   fname + " 1",
				ErrMsg:  "could not be split correctly",
			})
			return
		}

		var tokenNeeded = splits[1]
		var tokenStruct = &model.Token{}

		if check := os.Getenv("TOKEN_PASSWORD"); check == "" {
			_ = model.Load_env_values()
		}

		var token, err = jwt.ParseWithClaims(tokenNeeded, tokenStruct, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		if err != nil {
			utils.Respond(w, r, &utils.Msg{
				Message: utils.ErrorTokenInvalidOrExpired,
				Status:  http.StatusMisdirectedRequest,
				Fname:   fname + " 2",
				ErrMsg:  err.Error(),
			})
			return
		}

		if !token.Valid {
			utils.Respond(w, r, &utils.Msg{
				Message: utils.ErrorTokenInvalidOrExpired,
				Status:  http.StatusMisdirectedRequest,
				Fname:   fname + " 3",
				ErrMsg:  "token has been expired",
			})
			return
		}

		/*
			pass parameters using context
		 */
		r = utils.SetHeader(r, utils.KeyId, strconv.FormatUint(tokenStruct.UserId, 10))
		r = utils.SetHeader(r, utils.KeyRoleId, strconv.FormatUint(tokenStruct.RoleId, 10))

		//var redis_key = fmt.Sprintf("%v_%v", tokenStruct.Role, tokenStruct.UserID)
		//redis_result, err := iredis.GetRedis().Get(redis_key).Result()
		//if err == redis.Nil {
		//	utils.RespondExtended(writer, request, utils.Message(http.StatusBadRequest, "You have not authenticated"),
		//		&utils.LogMessage{
		//			Ok:       	false,
		//			FuncName: 	fname + " 4,
		//			Message:  	"Session Token is Old",
		//		})
		//	return
		//}
		//
		//if tokenHeader[len(tokenHeader) - utils.RedisSliceLength:] != redis_result {
		//	utils.Respond(w, r, &utils.Msg{
		//		Message: map[string]interface{}{
		//			"eng": "token has been expired",
		//		},
		//		Status:  http.StatusBadRequest,
		//		Fname:   fname + " 5,
		//		ErrMsg:  "token has been expired",
		//	})
		//	return
		//}
		EmailVerifiedWrapper(next, w, r)
	})
}

