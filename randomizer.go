package captcha

type Randomizer interface {
  Intn(int) int
}
