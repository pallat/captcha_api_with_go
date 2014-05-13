package captcha 

import (
  "github.com/stretchr/testify/assert"
  "strconv"
  "testing"
)

const (
  STRING_NUMBER_PATTERN = 1
  NUMBER_STRING_PATTERN = 2
)

type Captcha struct {
  pattern       int
  firstOperand  int
  operator      int
  secondOperand int
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
  if captcha.pattern == STRING_NUMBER_PATTERN {
    return integerString[captcha.firstOperand]
  }else{
    return strconv.Itoa(captcha.firstOperand)
  }
}

func (captcha *Captcha) Operator() string {
  if captcha.operator == 1 {
    return "+"
  }else if captcha.operator == 2 {
    return "-"
  }else{
    return "*"
  }
}

func (captcha *Captcha) SecondOperand() string {
  if captcha.pattern == NUMBER_STRING_PATTERN {
    return integerString[captcha.secondOperand]
  }else{
    return strconv.Itoa(captcha.secondOperand)
  }
}

func (captcha *Captcha) ToString() string {
  return captcha.FirstOperand() + " "+captcha.Operator()+" "+captcha.SecondOperand()
}

func Test1111ShouldBeFirstPatternOnePlus1(test *testing.T) {
  captcha := Captcha{1, 1, 1, 1}
  if captcha.ToString() == "One + 1" {
    test.Log("Test Passed")
  } else {
    test.Error("Failed")
  }
}

func Test1211ShouldBeFirstPatternTwoPlus1(test *testing.T) {
  captcha := Captcha{1, 2, 1, 1}
  assert.Equal(test, "Two", captcha.FirstOperand(), "They should be equal")
}

func Test1311ShouldBeFirstPatternAndFirstOperandShouldBeThree(test *testing.T) {
  captcha := Captcha{1, 3, 1, 1}
  assert.Equal(test, "Three", captcha.FirstOperand(), "They should be equal")
  assert.Equal(test, "Three + 1", captcha.ToString(), "They should ne equal")
}

func Test1411ShouldBeFirstPatternAndFirstOperandShouldBeFour(test *testing.T) {
  captcha := Captcha{1, 4, 1, 1}
  assert.Equal(test, "Four", captcha.FirstOperand(), "They should be equal")
  assert.Equal(test, "Four + 1", captcha.ToString(), "They should ne equal")
}

func Test1512ShouldBeFirstPatternAndSecondOperandShouldBe2(test *testing.T) {
  captcha := Captcha{1, 5, 1, 2}
  assert.Equal(test, "2", captcha.SecondOperand(), "They should be equal")
  assert.Equal(test, "Five + 2", captcha.ToString(), "They should ne equal")
}

func Test1613ShouldBeFirstPatternAndSecondOperandShouldBe3(test *testing.T) {
  captcha := Captcha{1, 6, 1, 3}
  assert.Equal(test, "3", captcha.SecondOperand(), "They should be equal")
  assert.Equal(test, "Six + 3", captcha.ToString(), "They should ne equal")
}

func Test1724ShouldBeFirstPatternAndOperatorShouldBeMinus(test *testing.T) {
  captcha := Captcha{1, 7, 2, 4}
  assert.Equal(test, "-", captcha.Operator(), "They should be equal")
  assert.Equal(test, "Seven - 4", captcha.ToString(), "They should ne equal")
}

func Test1835ShouldBeFirstPatternAndOperatorShouldBeMultiply(test *testing.T) {
  captcha := Captcha{1, 8, 3, 5}
  assert.Equal(test, "*", captcha.Operator(), "They should be equal")
  assert.Equal(test, "Eight * 5", captcha.ToString(), "They should ne equal")
}

func Test2835ShouldBeSecondPatternAndOperatorShouldBeMultiply(test *testing.T) {
  captcha := Captcha{2, 8, 3, 5}
  assert.Equal(test, "*", captcha.Operator(), "They should be equal")
  assert.Equal(test, "8 * Five", captcha.ToString(), "They should ne equal")
}
