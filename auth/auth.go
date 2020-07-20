package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"iguard_global/iredis"
	"iguard_global/utils"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var nameOfFunc = "JWT_TOKEN"
		var requestPath = request.URL.Path

		fmt.Println("Received: ", requestPath)

		for _, url := range utils.NoNeedToAuth {
			if requestPath == url {
				next.ServeHTTP(writer, request)
				return
			} else if strings.HasPrefix(requestPath, "/" + utils.FolderInWhichNotificationsAreStored) || strings.HasPrefix(requestPath, "/debug") {
				next.ServeHTTP(writer, request)
				return
			} else if strings.HasPrefix(requestPath, "/" + utils.FolderInWhichLogosAreStored) {
				next.ServeHTTP(writer, request)
				return
			}
		}

		var tokenHeader = request.Header.Get("session_id")

		if tokenHeader == "" {
			utils.RespondExtended(writer, request, utils.Message(http.StatusBadRequest, "Missing Auth Token"),
				&utils.LogMessage{
					Ok:       	false,
					FuncName: 	nameOfFunc,
					Message:  	"Missing Auth Token",
				})
			return
		}

		var splitted = strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			utils.RespondExtended(writer, request, utils.Message(http.StatusBadRequest, "Token does not contain 2 strings"),
				&utils.LogMessage{
					Ok:       	false,
					FuncName: 	nameOfFunc,
					Message:  	"Token does not contain 2 strings",
				})
			return
		}

		var tokenNeeded = splitted[1]
		var tokenStruct = &models.Token{}

		var token, err = jwt.ParseWithClaims(tokenNeeded, tokenStruct, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			utils.RespondExtended(writer, request, utils.Message(http.StatusBadRequest, "Invalid Auth Token"),
				&utils.LogMessage{
					Ok:       	false,
					FuncName: 	nameOfFunc,
					Message:  	"ParseWithClaims. Invalid Auth Token",
				})
			return
		}

		if !token.Valid {
			utils.RespondExtended(writer, request, utils.Message(http.StatusBadRequest, "Invalid Auth Token"),
				&utils.LogMessage{
					Ok:       	false,
					FuncName: 	nameOfFunc,
					Message:  	"Valid. Invalid Auth Token",
				})
			return
		}

		var cntx = context.WithValue(request.Context(), "id", tokenStruct.UserID)
		request = request.WithContext(cntx)

		var redis_key = fmt.Sprintf("%v_%v", tokenStruct.Role, tokenStruct.UserID)
		redis_result, err := iredis.GetRedis().Get(redis_key).Result()
		if err == redis.Nil {
			utils.RespondExtended(writer, request, utils.Message(http.StatusBadRequest, "You have not authenticated"),
				&utils.LogMessage{
					Ok:       	false,
					FuncName: 	nameOfFunc,
					Message:  	"Session Token is Old",
				})
			return
		}

		if tokenHeader[len(tokenHeader) - utils.RedisSliceLength:] != redis_result {
			utils.RespondExtended(writer, request, utils.Message(http.StatusBadRequest, "You have authenticated with another device. Please, authenticate again"),
				&utils.LogMessage{
					Ok:       	false,
					FuncName: 	nameOfFunc,
					Message:  	"Session Token is Old",
				})
			return
		}

		request.Header.Add("role", tokenStruct.Role)
		next.ServeHTTP(writer, request)
	})
}

