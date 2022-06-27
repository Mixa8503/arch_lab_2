package lab2

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var buf []byte

func TestHandlerCompute(t *testing.T) {
	reader := strings.NewReader("- * 2 17 2")
	writer := &Writer{}
	handler := ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	assert.Equal(t, nil, err)

	res, err := strconv.ParseFloat(string(buf), 64)
	if assert.Nil(t, err) {
		assert.Equal(t, 32.0, res)
	}

	buf = []byte{}
	reader = strings.NewReader("+ + 3 5")
	handler = ComputeHandler{
		Input:  reader,
		Output: writer,
	}
	err = handler.Compute()
	assert.Equal(t, "incorrect input", err.Error())
}

type Writer struct{}

func (w *Writer) Write(data []byte) (n int, err error) {
	buf = append(buf, data...)
	return len(buf), nil
}
