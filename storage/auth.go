package storage

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"time"
)
var jwtsecret []byte
func init(){
	godotenv.Load()
}
func Loadsecret(){
	secretstr:=os.Getenv("jwt_secret")
	if secretstr == "" {
        panic("JWT_SECRET not in .env!")
    }
    jwtsecret= []byte(secretstr)
}
func GenerateToken(username string)(string ,error){
	claims:=jwt.MapClaims{}
	claims["username"]=username
	claims["exp"] = time.Now().Add(6 * time.Hour).Unix()
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtsecret)
}
func ValidateToken(tokenstr string) (string,error){
	token,err:=jwt.Parse(tokenstr,func(t *jwt.Token) (interface{}, error) {
		return jwtsecret ,nil
	})
	if err!=nil{
		return "",err
	}
	 claims,ok:=token.Claims.(jwt.MapClaims)
	 if ok && token.Valid{
		return claims["username"].(string),nil
	 }
	 return "", fmt.Errorf("invalid token")
}



