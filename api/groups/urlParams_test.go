package groups

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/ElrondNetwork/elrond-proxy-go/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestParseBlockQueryOptions(t *testing.T) {
	options, err := parseBlockQueryOptions(createDummyGinContextWithQuery("withTxs=true&withLogs=true"))
	require.Nil(t, err)
	require.Equal(t, common.BlockQueryOptions{WithTransactions: true, WithLogs: true}, options)

	options, err = parseBlockQueryOptions(createDummyGinContextWithQuery("withTxs=true"))
	require.Nil(t, err)
	require.Equal(t, common.BlockQueryOptions{WithTransactions: true, WithLogs: false}, options)

	options, err = parseBlockQueryOptions(createDummyGinContextWithQuery("withTxs=foobar"))
	require.NotNil(t, err)
	require.Empty(t, options)
}

func TestParseHyperblockQueryOptions(t *testing.T) {
	options, err := parseHyperblockQueryOptions(createDummyGinContextWithQuery("withLogs=true"))
	require.Nil(t, err)
	require.Equal(t, common.HyperblockQueryOptions{WithLogs: true}, options)

	options, err = parseHyperblockQueryOptions(createDummyGinContextWithQuery(""))
	require.Nil(t, err)
	require.Empty(t, options)

	options, err = parseHyperblockQueryOptions(createDummyGinContextWithQuery("withLogs=foobar"))
	require.NotNil(t, err)
	require.Empty(t, options)
}

func TestParseAccountQueryOptions(t *testing.T) {
	options, err := parseAccountQueryOptions(createDummyGinContextWithQuery("onFinalBlock=true"))
	require.Nil(t, err)
	require.Equal(t, common.AccountQueryOptions{OnFinalBlock: true}, options)

	options, err = parseAccountQueryOptions(createDummyGinContextWithQuery(""))
	require.Nil(t, err)
	require.Empty(t, options)

	options, err = parseAccountQueryOptions(createDummyGinContextWithQuery("onFinalBlock=foobar"))
	require.NotNil(t, err)
	require.Empty(t, options)
}

func TestParseTransactionQueryOptions(t *testing.T) {
	options, err := parseTransactionQueryOptions(createDummyGinContextWithQuery("withResults=true"))
	require.Nil(t, err)
	require.Equal(t, common.TransactionQueryOptions{WithResults: true}, options)

	options, err = parseTransactionQueryOptions(createDummyGinContextWithQuery(""))
	require.Nil(t, err)
	require.Empty(t, options)

	options, err = parseTransactionQueryOptions(createDummyGinContextWithQuery("withResults=foobar"))
	require.NotNil(t, err)
	require.Empty(t, options)
}

func TestParseTransactionSimulationOptions(t *testing.T) {
	options, err := parseTransactionSimulationOptions(createDummyGinContextWithQuery("checkSignature=false"))
	require.Nil(t, err)
	require.Equal(t, common.TransactionSimulationOptions{CheckSignature: false}, options)

	options, err = parseTransactionSimulationOptions(createDummyGinContextWithQuery(""))
	require.Nil(t, err)
	require.Equal(t, options, common.TransactionSimulationOptions{CheckSignature: true})

	options, err = parseTransactionSimulationOptions(createDummyGinContextWithQuery("checkSignature=foobar"))
	require.NotNil(t, err)
	require.Empty(t, options)
}

func TestParseBoolUrlParam(t *testing.T) {
	c := createDummyGinContextWithQuery("a=true&b=false&c=foobar&d")

	value, err := parseBoolUrlParam(c, "a")
	require.Nil(t, err)
	require.True(t, value)

	value, err = parseBoolUrlParam(c, "b")
	require.Nil(t, err)
	require.False(t, value)

	value, err = parseBoolUrlParam(c, "c")
	require.NotNil(t, err)
	require.False(t, value)

	value, err = parseBoolUrlParam(c, "d")
	require.Nil(t, err)
	require.False(t, value)

	value, err = parseBoolUrlParam(c, "e")
	require.Nil(t, err)
	require.False(t, value)
}

func TestParseUint32UrlParam(t *testing.T) {
	c := createDummyGinContextWithQuery("a=7&b=0&c=foobar&d=-1&e=12345678987654321")

	value, err := parseUint32UrlParam(c, "a")
	require.Nil(t, err)
	require.True(t, value.HasValue)
	require.Equal(t, uint32(7), value.Value)

	value, err = parseUint32UrlParam(c, "b")
	require.Nil(t, err)
	require.True(t, value.HasValue)
	require.Equal(t, uint32(0), value.Value)

	value, err = parseUint32UrlParam(c, "c")
	require.NotNil(t, err)
	require.False(t, value.HasValue)
	require.Equal(t, uint32(0), value.Value)

	value, err = parseUint32UrlParam(c, "d")
	require.NotNil(t, err)
	require.False(t, value.HasValue)
	require.Equal(t, uint32(0), value.Value)

	value, err = parseUint32UrlParam(c, "e")
	require.NotNil(t, err)
	require.False(t, value.HasValue)
	require.Equal(t, uint32(0), value.Value)
}

func TestParseUint64UrlParam(t *testing.T) {
	c := createDummyGinContextWithQuery("a=7&b=0&c=foobar&d=-1&e=12345678987654321")

	value, err := parseUint64UrlParam(c, "a")
	require.Nil(t, err)
	require.True(t, value.HasValue)
	require.Equal(t, uint64(7), value.Value)

	value, err = parseUint64UrlParam(c, "b")
	require.Nil(t, err)
	require.True(t, value.HasValue)
	require.Equal(t, uint64(0), value.Value)

	value, err = parseUint64UrlParam(c, "c")
	require.NotNil(t, err)
	require.False(t, value.HasValue)
	require.Equal(t, uint64(0), value.Value)

	value, err = parseUint64UrlParam(c, "d")
	require.NotNil(t, err)
	require.False(t, value.HasValue)
	require.Equal(t, uint64(0), value.Value)

	value, err = parseUint64UrlParam(c, "e")
	require.Nil(t, err)
	require.True(t, value.HasValue)
	require.Equal(t, uint64(12345678987654321), value.Value)
}

func TestParseHexBytesUrlParam(t *testing.T) {
	c := createDummyGinContextWithQuery("a=aaaa&b=test&c")

	value, err := parseHexBytesUrlParam(c, "a")
	require.Nil(t, err)
	require.Equal(t, []byte{0xaa, 0xaa}, value)

	value, err = parseHexBytesUrlParam(c, "b")
	require.NotNil(t, err)
	require.Nil(t, value)

	value, err = parseHexBytesUrlParam(c, "c")
	require.Nil(t, err)
	require.Equal(t, []byte(nil), value)
}

func TestParseTransactionsPoolQueryOptions(t *testing.T) {
	c := createDummyGinContextWithQuery("")
	expectedValue := common.TransactionsPoolOptions{}
	value, err := parseTransactionsPoolQueryOptions(c)
	require.Nil(t, err)
	require.Equal(t, expectedValue, value)

	c = createDummyGinContextWithQuery("by-sender=some_sender&fields=sender,receiver&last-nonce=true&nonce-gaps=true&shard-id=333")
	expectedValue = common.TransactionsPoolOptions{
		ShardID:   "333",
		Sender:    "some_sender",
		Fields:    "sender,receiver",
		LastNonce: true,
		NonceGaps: true,
	}
	value, err = parseTransactionsPoolQueryOptions(c)
	require.Nil(t, err)
	require.Equal(t, expectedValue, value)
}

func TestParseStringUrlParam(t *testing.T) {
	c := createDummyGinContextWithQuery("a=dummy")

	require.Equal(t, "dummy", parseStringUrlParam(c, "a"))
}

func createDummyGinContextWithQuery(rawQuery string) *gin.Context {
	return &gin.Context{Request: &http.Request{URL: &url.URL{RawQuery: rawQuery}}}
}
