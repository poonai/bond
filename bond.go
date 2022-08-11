package bond

import (
	"io"

	"github.com/cockroachdb/pebble"
)

type Options struct {
	pebble.Options

	Serializer Serializer[any]
}

type DB struct {
	*pebble.DB

	serializer Serializer[any]
}

func Open(dirname string, opts *Options) (*DB, error) {
	if opts.Comparer == nil {
		opts.Comparer = pebble.DefaultComparer
		opts.Comparer.Split = _KeyPrefixSplitIndex
	}

	pdb, err := pebble.Open(dirname, &opts.Options)
	if err != nil {
		return nil, err
	}

	var serializer Serializer[any]
	if opts.Serializer != nil {
		serializer = opts.Serializer
	} else {
		serializer = &JsonSerializer{}
	}

	return &DB{DB: pdb, serializer: serializer}, nil
}

func (db *DB) Serializer() Serializer[any] {
	return db.serializer
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func (db *DB) getKV(key []byte, batch *pebble.Batch) (data []byte, closer io.Closer, err error) {
	if batch != nil {
		data, closer, err = batch.Get(key)
	} else {
		data, closer, err = db.Get(key)
	}
	return
}
