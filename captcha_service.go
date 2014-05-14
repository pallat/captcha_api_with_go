package captcha

type CaptchaService struct {
  rand Randomizer
}

func NewCaptchaService() *CaptchaService {
  return new(CaptchaService)
}

func (s *CaptchaService) SetRandom(rand Randomizer) {
  s.rand = rand
}

func (s *CaptchaService) GetCaptcha() Captcha {
  captcha := Captcha{s.rand.Intn(1) + 1, s.rand.Intn(8)+1, s.rand.Intn(2) + 1, s.rand.Intn(8)+1}
  return captcha
}
