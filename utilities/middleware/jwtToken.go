package middleware

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
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
			return []byte("bsof2sJXPp0T5G38L6RKq21mqayXyr4u"), nil //todo move this to env file
		})
		if err != nil {
			if err.Error() == jwt.ErrSignatureInvalid.Error() {
				logs.ErrorLogger.Println(err.Error())
				return ps
			}
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
				v, ok := val.(map[string]interface{})["0"]//Todo create permission in admin
				if !ok {
					logs.ErrorLogger.Println("Permissions not found")
				}
				if v != nil{
					switch reflect.TypeOf(v).Kind() {
					case reflect.Slice:
						s := reflect.ValueOf(v)
						for i := 0; i < s.Len(); i++ {
							if s.Index(i).Interface().(string) == "94"{//Todo create permission in admin
								ps.Status = true
							}
						}
					}
				}else {
					logs.ErrorLogger.Println("Permissions not found")
					ps.Status = false
				}

			}
		}
	}else{
		logs.ErrorLogger.Println("Bearer token not found")
		return ps
	}
	return ps
}