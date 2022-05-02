package schemas

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCrypt(t *testing.T) {
	pass := "asda25"
	hashStr, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	hash := string(hashStr)
	require.NoError(t, err)
	t.Logf("hash=%s\n", hash)
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	require.NoError(t, err)
}
