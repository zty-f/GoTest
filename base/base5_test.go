package base

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestNilStrcut(t *testing.T) {
	type User struct {
		Name string
	}

	var u *User
	if u == nil {
		t.Log("u is nil")
	} else {
		t.Error("u should be nil")
	}

	u = &User{}

	fmt.Println(u.Name)
	fmt.Println(u.Name)
	fmt.Println(u.Name)
	fmt.Println(cast.ToString(0))

}
