package network

import (
	"crypto/sha512"
	"encoding/hex"
	"log"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"

	// "github.com/moondevgo/go-mods/basic"
	"github.com/moondevgo/go-mods/basic"
)

func UpbitGetPayload(data map[string]interface{}) url.Values {
	payload := url.Values{}
	for k, v := range data {
		payload.Add(k, v.(string))
	}
	return payload
}

func UpbitSimpleJwt(access_key, secret_key string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["access_key"] = access_key
	u := uuid.NewV4()
	claims["nonce"] = u.String()
	token.Claims = claims

	signedToken, err := token.SignedString([]byte(secret_key))
	if err != nil {
		log.Println("SignedString Error")
		log.Fatal(err)
		return ""
	}
	return signedToken
}

// func UpbitPayloadJwt(access_key, secret_key string, payload url.Values) string {
func UpbitPayloadJwt(access_key, secret_key string, data map[string]interface{}) string {
	payload := UpbitGetPayload(data)
	claims := make(jwt.MapClaims)
	claims["access_key"] = access_key
	claims["nonce"] = uuid.NewV4().String()
	qh := sha512.New()
	qh.Reset()
	qh.Write([]byte(payload.Encode()))
	claims["query_hash"] = hex.EncodeToString(qh.Sum(nil))
	claims["query_hash_alg"] = "SHA512"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret_key))
	//log.Println(signedToken)
	if err != nil {
		log.Println("SignedString Error")
		log.Fatal(err)
		return ""
	}
	return signedToken
}

func UpbitJwt(data map[string]interface{}) (signedToken string) {
	keys := basic.GetConfigYaml("coin_apis", "upbit")
	access_key := keys["access_key"].(string)
	secret_key := keys["secret_key"].(string)

	switch len(data) {
	case 0:
		signedToken = UpbitSimpleJwt(access_key, secret_key)
	default:
		signedToken = UpbitPayloadJwt(access_key, secret_key, data)
	}
	return signedToken
}
