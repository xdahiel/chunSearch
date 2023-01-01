package model

import (
	"chunsearch/common"
	"testing"
)

func TestAdd(t *testing.T) {
	common.Init()
	InitUser()

	err := Add("chunsheng", "123456")
	if err != nil {
		t.Error(err)
	}
}
