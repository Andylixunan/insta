package jwt

import (
	"testing"
	"time"

	"github.com/Andylixunan/insta/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestManager_Generate(t *testing.T) {
	conf := &config.Config{
		JWT: config.JWT{
			Secret: "secret",
			Expire: 5 * time.Second,
		},
	}
	jwtManager := NewManager(conf)
	id := uint32(6)
	tokenStr, err := jwtManager.Generate(id)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
	t.Logf("token string: %v", tokenStr)
}

func TestManager_Validate(t *testing.T) {
	conf := &config.Config{
		JWT: config.JWT{
			Secret: "secret",
			Expire: 5 * time.Second,
		},
	}
	jwtManager := NewManager(conf)
	id := uint32(6)
	tokenStr, err := jwtManager.Generate(id)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
	t.Logf("token string: %v", tokenStr)
	claims, err := jwtManager.Validate(tokenStr)
	require.NoError(t, err)
	require.Equal(t, id, claims.ID)
	t.Logf("decoded: %+v", claims)
}
