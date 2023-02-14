package middleware

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/sirupsen/logrus"
)

type PermissionStatus struct {
	Status   bool
	TenantId string
	UserId   string
}

func HasPermissions(reqToken string) PermissionStatus {
	var ps PermissionStatus
	if len(reqToken) > 0 {
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(commons.GoDotEnvVariable("JWT_DECODE_KEY")), nil
		})
		if err != nil {
			if err.Error() == jwt.ErrSignatureInvalid.Error() {
				logs.ErrorLogger.Println(err.Error())
				return ps
			}
			logrus.Error(err.Error())
			logs.ErrorLogger.Println(err.Error())
			return ps
		}
		for key, val := range claims {
			if key == "userID" {
				ps.UserId = fmt.Sprintf("%v", val)
			}
			if key == "tenantID" {
				ps.TenantId = fmt.Sprintf("%v", val)
			}
			if key == "permissions" {
				v, ok := val.(map[string]interface{})["0"]
				if !ok {
					logs.ErrorLogger.Println("Permissions not found")
				}
				if v != nil {
					switch reflect.TypeOf(v).Kind() {
					case reflect.Slice:
						s := reflect.ValueOf(v)
						for i := 0; i < s.Len(); i++ {
							if s.Index(i).Interface().(string) == "97" {
								ps.Status = true
							}
						}
					}
				} else {
					logrus.Error("Permissions not found")
					logs.ErrorLogger.Println("Permissions not found")
					ps.Status = false
				}

			}
		}
	} else {
		logrus.Error("Bearer token not found")
		logs.ErrorLogger.Println("Bearer token not found")
		return ps
	}
	return ps
}
