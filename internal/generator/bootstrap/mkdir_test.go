package bootstrap

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMkdir(t *testing.T) {
	require.NoError(t, Mkdir("test"))
	require.NoError(t, os.RemoveAll("test"))
}
