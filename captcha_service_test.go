package captcha

import (
  "testing"
  "reflect"
  "github.com/stretchr/testify/assert"
  ."github.com/roofimon/captcha"
)

type MyRandom struct { }

func (m MyRandom) Intn(n int) int {
  return 0
}

func TestTestify(t *testing.T) {
  var service *CaptchaService = NewCaptchaService()
  myRandom := new(MyRandom)
  service.SetRandom(myRandom)
  captcha := service.GetCaptcha()
  assert.Equal(t, "One + 1", captcha.ToString())
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
