package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixEvaluate(t *testing.T) {
	res, err := Prefix_Eval("+ 3 4")
	if assert.Nil(t, err) {
		assert.Equal(t, 7.0, res)
	}

	res, err = Prefix_Eval("- 4 2")
	if assert.Nil(t, err) {
		assert.Equal(t, 2.0, res)
	}

	res, err = Prefix_Eval("* 8 9")
	if assert.Nil(t, err) {
		assert.Equal(t, 72.0, res)
	}

	res, err = Prefix_Eval("/ 64 4")
	if assert.Nil(t, err) {
		assert.Equal(t, 16.0, res)
	}

	res, err = Prefix_Eval("^ 2 5")
	if assert.Nil(t, err) {
		assert.Equal(t, 32.0, res)
	}

	res, err = Prefix_Eval("^ + 2 5 2")
	if assert.Nil(t, err) {
		assert.Equal(t, 49.0, res)
	}

	res, err = Prefix_Eval("- * 2 17 2")
	if assert.Nil(t, err) {
		assert.Equal(t, 32.0, res)
	}

	res, err = Prefix_Eval("- * / 2 16 2 - 8 2")
	if assert.Nil(t, err) {
		assert.Equal(t, -5.75, res)
	}

	res, err = Prefix_Eval("f")
	assert.Equal(t, 0.0, res)
	assert.Equal(t, "incorrect input", err.Error())

	res, err = Prefix_Eval("++ 4 2")
	assert.Equal(t, 0.0, res)
	assert.Equal(t, "incorrect input", err.Error())

	res, err = Prefix_Eval("+ 4 2 - 2 4 9")
	assert.Equal(t, 0.0, res)
	assert.Equal(t, "incorrect input", err.Error())

	res, err = Prefix_Eval("+ - 4 2 3 ^ 2")
	assert.Equal(t, 0.0, res)
	assert.Equal(t, "Something is wrong, you need to check input", err.Error())
}
func ExamplePrefixEvaluate() {
	res, _ := Prefix_Eval("+ 10 3")
	fmt.Println(res)

	// Output:
	// 13
}
