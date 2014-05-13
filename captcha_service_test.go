package captcha

import "testing"
import "reflect"

type MyRandom struct { }

func (m MyRandom) Intn(n int) int {
  return 2
}

func TestNewCaptchaService(t *testing.T) {
  var service *CaptchaService = NewCaptchaService()
  if service == nil {
    t.Errorf("expect captcha service is not nil")
  }
}

func TestGetCaptcha(t *testing.T) {
  var service *CaptchaService = NewCaptchaService()
  myRand := new(MyRandom)
  service.SetRandom(myRand)
  captcha := service.GetCaptcha()
  if "Captcha" != reflect.TypeOf(captcha).Name() {
    t.Errorf("return type is not Captcha %s", reflect.TypeOf(captcha))
  }
}



