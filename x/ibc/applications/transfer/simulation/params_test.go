package simulation_test

import (
	"math/rand"
	"testing"

	simulation2 "github.com/pundix/pundix/x/ibc/applications/transfer/simulation"

	"github.com/stretchr/testify/require"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"transfer/SendEnabled", "SendEnabled", "false", "transfer"},
		{"transfer/ReceiveEnabled", "ReceiveEnabled", "true", "transfer"},
	}

	paramChanges := simulation2.ParamChanges(r)

	require.Len(t, paramChanges, 2)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r), p.Key())
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
