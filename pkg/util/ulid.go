package util

import (
	"crypto/rand"
	"io"
	"sync"
)

import (
	"github.com/oklog/ulid/v2"
)

var (
	entropy     io.Reader
	entropyOnce sync.Once
)

func defaultEntropy() io.Reader {
	entropyOnce.Do(func() {
		entropy = &ulid.LockedMonotonicReader{
			MonotonicReader: ulid.Monotonic(rand.Reader, 1),
		}
	})
	return entropy
}

func MakeULID() (id string, err error) {
	ulid, err := ulid.New(ulid.Now(), defaultEntropy())
	return ulid.String(), err
}
