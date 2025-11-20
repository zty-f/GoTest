package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorsIsDuplicate(t *testing.T) {

	assert.False(t, ErrorsIsDuplicate(nil))

	assert.True(t, ErrorsIsDuplicate(errors.New("Duplicate")))
	assert.True(t, ErrorsIsDuplicate(errors.New(" Duplicate")))
	assert.True(t, ErrorsIsDuplicate(errors.New(" Duplicate ")))

	assert.False(t, ErrorsIsDuplicate(errors.New("uplicate. ")))
	assert.False(t, ErrorsIsDuplicate(errors.New("Duplicat")))
}

func TestErrorsIsLevelNoMatch(t *testing.T) {

	assert.False(t, ErrorsIsLevelNoMatch(nil))
	assert.False(t, ErrorsIsLevelNoMatch(errors.New("等级不匹配")))
	assert.True(t, ErrorsIsLevelNoMatch(errors.New("等级不匹配,Lv5等级及以上可使用")))
	assert.True(t, ErrorsIsLevelNoMatch(errors.New("等级不匹配，Lv5等级及以上可使用")))
}
