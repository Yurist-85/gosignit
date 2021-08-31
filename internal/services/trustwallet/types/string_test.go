package types

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestStringSuite struct {
	suite.Suite
}

func TestString(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestStringSuite))
}

func (t *TestStringSuite) TestTWStringCreateWithGoString_ok() {
	twStr := TWStringCreateWithGoString("test string")
	t.NotNil(twStr)
}

func (t *TestStringSuite) TestTWStringGoString_ok() {
	twStr := TWStringCreateWithGoString("test string")
	t.Require().NotNil(twStr)

	goStr := TWStringGoString(twStr)
	t.Equal("test string", goStr)
}

func (t *TestStringSuite) TestTWStringGoString_fail() {
	twStr := TWStringCreateWithGoString("test string")
	t.Require().NotNil(twStr)

	goStr := TWStringGoString(twStr)
	t.NotEqual("another string", goStr)
}
