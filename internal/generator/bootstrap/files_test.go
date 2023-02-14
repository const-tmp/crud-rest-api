package bootstrap

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestFiles(t *testing.T) {
	require.NoError(t, Mkdir("test"))
	require.NoError(t, Files("test"))
	require.NoError(t, os.RemoveAll("test"))
}
