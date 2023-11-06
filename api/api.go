package api

import (
	"github.com/mojocn/base64Captcha"

	"captcha/service"
	"captcha/util"
)

type Api struct {
	Svc     *service.Service
	Captcha *base64Captcha.Captcha
	jwt     *util.JWT
}

func Newapi(svc *service.Service, captcha *base64Captcha.Captcha, jwt *util.JWT) *Api {
	return &Api{
		Svc:     svc,
		Captcha: captcha,
		jwt:     jwt,
	}
}
