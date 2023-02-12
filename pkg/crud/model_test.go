package crud

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

type (
	A struct {
		gorm.Model
		Name string
	}
)

func F(m gorm.Model) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func TestModel(t *testing.T) {
	g, err := F(*new(gorm.Model))
	require.NoError(t, err)
	t.Log(g)

	m, err := F(*new(gorm.Model))
	require.NoError(t, err)
	t.Log(m)
}

func f(s string, b bool) bool {
	fmt.Printf("%s called\n", s)
	return b
}

func TestOrAnd(t *testing.T) {
	if f("a", true) && f("b", false) {

	}
}
