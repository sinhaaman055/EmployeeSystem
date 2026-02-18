package storage

import (
	"fmt"
	"os"

	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)
var jwtsecret []byte
var encsecret []byte
func init(){
	godotenv.Load()
}
func Loadsecret(){
	secretstr:=os.Getenv("jwt_secret")
	if secretstr == "" {
        panic("JWT_SECRET not in .env!")
    }
    jwtsecret= []byte(secretstr)
	encstr:=os.Getenv("jwt_encsecret")
	if encstr==""{
		panic("Enc_secret not in .env")
	}
	encsecret=[]byte(encstr)
}
func GenerateToken(username string)(string ,error){
	claims:=jwt.MapClaims{}
	claims["username"]=username
	claims["exp"] = time.Now().Add(6 * time.Hour).Unix()
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenstr,_:= token.SignedString(jwtsecret)
    encryptor,err:=jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{Algorithm: jose.DIRECT,Key: encsecret},
		nil,
	)
	if err!=nil{
		return "",err
	}
    object,err:=encryptor.Encrypt([]byte(tokenstr))
	if err!=nil{
		return "",err
	}
	return object.CompactSerialize()
	
}
func ValidateToken(tokenstr string) (string,error){
	jwe,err:=jose.ParseEncrypted(tokenstr)
	if err!=nil{
		return "",err
	}
	jwebyte,err:=jwe.Decrypt(encsecret)
	if err!=nil{
		return "Not abe to decrypt ",err
	}

	token,err:=jwt.Parse(string(jwebyte),func(t *jwt.Token) (interface{}, error) {
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



