package main

import (
	"context"
	"os"

	"captcha/util"
)

func main() {
	ctx := context.Background()
	db := util.Newdb(ctx)
	privateKety, _ := os.ReadFile("private.rsa")
	jwt := util.NewJWTFromKeyBytes(privateKety)
	router := Newrouter(db, jwt)
	router.Run(":8089")
}
