package util

import "github.com/mojocn/base64Captcha"

type CaptchaService struct {
	Captcha *base64Captcha.Captcha
}

func Newcaptcha() *CaptchaService {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	store := base64Captcha.DefaultMemStore
	captcha := base64Captcha.NewCaptcha(driver, store)
	return &CaptchaService{Captcha: captcha}
}

func (s *CaptchaService) Genarate() (string, string, error) {
	return s.Captcha.Generate()
}
func (s *CaptchaService) Verify(id, answer string, clear bool) bool {
	return s.Captcha.Verify(id, answer, clear)
}
