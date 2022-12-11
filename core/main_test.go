package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameUpdate(t *testing.T) {
	game := Game{}
	game.Update()
	require.Equal(t, 1, game.count)
}
