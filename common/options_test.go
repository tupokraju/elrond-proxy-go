package common

import (
	"net/url"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/stretchr/testify/require"
)

func TestBuildUrlWithBlockQueryOptions_ShouldWork(t *testing.T) {
	builtUrl := BuildUrlWithBlockQueryOptions("/block/by-nonce/15", BlockQueryOptions{})
	require.Equal(t, "/block/by-nonce/15", builtUrl)

	builtUrl = BuildUrlWithBlockQueryOptions("/block/by-nonce/15", BlockQueryOptions{
		WithTransactions: true,
	})
	require.Equal(t, "/block/by-nonce/15?withTxs=true", builtUrl)

	builtUrl = BuildUrlWithBlockQueryOptions("/block/by-nonce/15", BlockQueryOptions{
		WithTransactions: true,
		WithLogs:         true,
	})
	parsed, err := url.Parse(builtUrl)
	require.Nil(t, err)
	require.Equal(t, "/block/by-nonce/15", parsed.Path)
	require.Equal(t, "true", parsed.Query().Get("withTxs"))
	require.Equal(t, "true", parsed.Query().Get("withLogs"))
}

func TestBuildUrlWithAccountQueryOptions_ShouldWork(t *testing.T) {
	builtUrl := BuildUrlWithAccountQueryOptions("/address/erd1alice", AccountQueryOptions{})
	require.Equal(t, "/address/erd1alice", builtUrl)

	builtUrl = BuildUrlWithAccountQueryOptions("/address/erd1alice", AccountQueryOptions{
		BlockNonce: core.OptionalUint64{HasValue: true, Value: 42},
	})
	require.Equal(t, "/address/erd1alice?blockNonce=42", builtUrl)

	builtUrl = BuildUrlWithAccountQueryOptions("/address/erd1alice", AccountQueryOptions{
		BlockHash: []byte{0xab, 0xba},
	})
	require.Equal(t, "/address/erd1alice?blockHash=abba", builtUrl)

	// The following isn't a valid scenario in the real world, according to the validation defined in:
	// https://github.com/ElrondNetwork/elrond-go/blob/master/api/groups/addressGroupOptions.go
	// However, here, we are testing each code path.
	builtUrl = BuildUrlWithAccountQueryOptions("/address/erd1alice", AccountQueryOptions{
		OnFinalBlock:   true,
		OnStartOfEpoch: core.OptionalUint32{HasValue: true, Value: 1},
		BlockNonce:     core.OptionalUint64{HasValue: true, Value: 2},
		BlockHash:      []byte{0xaa, 0xbb},
		BlockRootHash:  []byte{0xbb, 0xaa},
		HintEpoch:      core.OptionalUint32{HasValue: true, Value: 3},
	})
	parsed, err := url.Parse(builtUrl)
	require.Nil(t, err)
	require.Equal(t, "/address/erd1alice", parsed.Path)
	require.Equal(t, "true", parsed.Query().Get("onFinalBlock"))
	require.Equal(t, "1", parsed.Query().Get("onStartOfEpoch"))
	require.Equal(t, "2", parsed.Query().Get("blockNonce"))
	require.Equal(t, "aabb", parsed.Query().Get("blockHash"))
	require.Equal(t, "bbaa", parsed.Query().Get("blockRootHash"))
	require.Equal(t, "3", parsed.Query().Get("hintEpoch"))
}
