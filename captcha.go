package captcha

import "strconv"

const (
  STRING_NUMBER_PATTERN = 1
  NUMBER_STRING_PATTERN = 2
)

type Captcha struct {
  Pattern       int
  LeftOperand  int
  Operator      int
  RightOperand int
}

var integerString = map[int]string{
    1: "One",
    2: "Two",
    3: "Three",
    4: "Four",
    5: "Five",
    6: "Six",
    7: "Seven",
    8: "Eight",
    9: "Nine",
}

func (captcha *Captcha) FirstOperand() string {
  if captcha.Pattern == STRING_NUMBER_PATTERN {
    return integerString[captcha.LeftOperand]
  }else{
    return strconv.Itoa(captcha.LeftOperand)
  }
}

func (captcha *Captcha) GetOperator() string {
  if captcha.Operator == 1 {
    return "+"
  }else if captcha.Operator == 2 {
    return "-"
  }else{
    return "*"
  }
}

func (captcha *Captcha) SecondOperand() string {
  if captcha.Pattern == NUMBER_STRING_PATTERN {
      return integerString[captcha.RightOperand]
  }else{
    return strconv.Itoa(captcha.RightOperand)
  }
}

func (captcha *Captcha) ToString() string {
  return captcha.FirstOperand() + " "+captcha.GetOperator()+" "+captcha.SecondOperand()
}
