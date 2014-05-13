package captcha

import "testing"



func TestNewCaptchaService(t *testing.T) {
  var service *CaptchaService = NewCaptchaService()
  if service == nil {
    t.Errorf("expect captcha service is not nil")
  }
}
