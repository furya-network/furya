package simapp_test

import (
	"testing"

	simappUtil "github.com/furya-network/furya/testutil/simapp"
	"github.com/stretchr/testify/require"
)

func TestGetTestObjects(t *testing.T) {
	_, _, err := simappUtil.GetTestObjects()
	require.NoError(t, err)
}

func TestSetup(t *testing.T) {
	panicFunc := func() {
		simappUtil.Setup(true)
		simappUtil.Setup(false)
	}
	require.NotPanics(t, panicFunc)
}
