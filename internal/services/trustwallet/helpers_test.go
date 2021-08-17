package trustwallet

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestHelpersSuite struct {
	suite.Suite
}

func TestHelpers(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestHelpersSuite))
}

func (t *TestHelpersSuite) TestTWStringCreateWithGoString_ok() {
	twStr := TWStringCreateWithGoString("test string")
	t.NotNil(twStr)
}

func (t *TestHelpersSuite) TestTWStringGoString_ok() {
	twStr := TWStringCreateWithGoString("test string")
	t.Require().NotNil(twStr)
	goStr := TWStringGoString(twStr)
	t.Equal("test string", goStr)
}

func (t *TestHelpersSuite) TestTWStringGoString_fail() {
	twStr := TWStringCreateWithGoString("test string")
	t.Require().NotNil(twStr)
	goStr := TWStringGoString(twStr)
	t.NotEqual("another string", goStr)
}

func (t *TestHelpersSuite) TestTWDataCreateWithGoBytes_ok() {
	goBytes := []byte("ten bytes.")
	twData := TWDataCreateWithGoBytes(goBytes)
	t.NotNil(twData)
}

func (t *TestHelpersSuite) TestTWDataGoBytes_ok() {
	goBytes := []byte("ten bytes.")
	twData := TWDataCreateWithGoBytes(goBytes)
	t.Require().NotNil(twData)
	goBytes2 := TWDataGoBytes(twData)
	t.Len(goBytes2, 10)
	t.Equal("ten bytes.", string(goBytes2))
}

func (t *TestHelpersSuite) TestTWDataHexString_ok() {
	goBytes := []byte("ten bytes.")
	twData := TWDataCreateWithGoBytes(goBytes)
	t.Require().NotNil(twData)
	goBytesHex := TWDataHexString(twData)
	t.Equal("74656e2062797465732e", goBytesHex)
}
