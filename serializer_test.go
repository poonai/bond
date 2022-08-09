package bond

import (
	"bytes"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmihailenco/msgpack/v5"
)

func TestMsgpackSerializer_SerializerWithClosable(t *testing.T) {
	s := MsgpackSerializer{
		EncoderFunc: msgpack.GetEncoder,
		DecoderFunc: msgpack.GetDecoder,
		BufferPool: &SyncPoolWrapper[bytes.Buffer]{
			Pool: sync.Pool{New: func() interface{} { return bytes.Buffer{} }},
		},
	}

	tb := &TokenBalance{
		ID:              5,
		AccountID:       3,
		ContractAddress: "abc",
		AccountAddress:  "xyz",
		TokenID:         12,
		Balance:         7,
	}

	buff, closeBuff, err := s.SerializerWithCloseable(tb)
	require.NoError(t, err)
	require.NotNil(t, buff)
	require.NotNil(t, closeBuff)

	var tb2 *TokenBalance
	err = s.Deserialize(buff, &tb2)
	require.NoError(t, err)

	closeBuff()

	assert.Equal(t, tb, tb2)
}

func TestMsgpackGenSerializer_SerializerWithClosable(t *testing.T) {
	s := MsgpackGenSerializer{
		BufferPool: &SyncPoolWrapper[bytes.Buffer]{
			Pool: sync.Pool{New: func() interface{} { return bytes.Buffer{} }},
		},
	}

	tb := &TokenBalance{
		ID:              5,
		AccountID:       3,
		ContractAddress: "abc",
		AccountAddress:  "xyz",
		TokenID:         12,
		Balance:         7,
	}

	buff, closeBuff, err := s.SerializerWithCloseable(tb)
	require.NoError(t, err)
	require.NotNil(t, buff)
	require.NotNil(t, closeBuff)

	var tb2 *TokenBalance
	err = s.Deserialize(buff, &tb2)
	require.NoError(t, err)

	closeBuff()

	assert.Equal(t, tb, tb2)
}
