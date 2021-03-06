package main

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gmbh-micro/gmbh"
)

func main() {
	runtime := gmbh.SetRuntime(gmbh.RuntimeOptions{Blocking: true, Verbose: true})
	client, err := gmbh.NewClient("./demo.toml", runtime)
	if err != nil {
		panic(err)
	}

	client.Route("test", handleTest)
	client.Route("two", handleTest2)
	client.Route("tkn", handleTkn)
	client.Start()
}

func handleTest(req gmbh.Request, resp *gmbh.Responder) {
	resp.Result = "Hello from _cabal-generic, we received: " + req.Data1
}
func handleTest2(req gmbh.Request, resp *gmbh.Responder) {
	resp.Result = "Hello from _cabal-generic TEST2, we received: " + req.Data1
}

func handleTkn(req gmbh.Request, resp *gmbh.Responder) {

	tkn, err := generateToken(req.Data1, "guest")
	if err != nil {
		panic(err)
	}

	resp.Result = tkn

}

func generateToken(username string, admin string) (string, error) {
	tokenSecret := "gmbH"
	key := []byte(tokenSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	header := token.Header
	claims := token.Claims.(jwt.MapClaims)

	header["kid"] = os.Getenv("USERKID")
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	claims["id"] = username
	claims["usergroup"] = admin
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "_cabal-authentication-module"

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("Could not generate valid token")
	}
	return tokenString, nil
}
