package types

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestDataSuite struct {
	suite.Suite
}

func TestData(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestDataSuite))
}

func (t *TestDataSuite) TestTWDataCreateWithGoBytes_ok() {
	goBytes := []byte("ten bytes.")
	twData := TWDataCreateWithGoBytes(goBytes)
	t.NotNil(twData)
}

func (t *TestDataSuite) TestTWDataGoBytes_ok() {
	goBytes := []byte("ten bytes.")
	twData := TWDataCreateWithGoBytes(goBytes)
	t.Require().NotNil(twData)

	goBytes2 := TWDataGoBytes(twData)
	t.Len(goBytes2, 10)
	t.Equal("ten bytes.", string(goBytes2))
}

func (t *TestDataSuite) TestTWDataHexString_ok() {
	goBytes := []byte("ten bytes.")
	twData := TWDataCreateWithGoBytes(goBytes)
	t.Require().NotNil(twData)

	goBytesHex := TWDataHexString(twData)
	t.Equal("74656e2062797465732e", goBytesHex)
}
