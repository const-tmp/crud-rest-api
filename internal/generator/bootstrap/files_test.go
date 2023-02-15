package bootstrap

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestFiles(t *testing.T) {
	require.NoError(t, Mkdir("test"))
	require.NoError(t, OpenAPIFiles("test"))
	require.NoError(t, os.RemoveAll("test"))
}

func TestEmbed(t *testing.T) {
	data, err := f.ReadFile("api/src/openapi.yaml")
	require.NoError(t, err)
	t.Log(string(data))
}
