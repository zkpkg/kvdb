package bigkv

import (
	"io"
	"testing"

	"github.com/dfuse-io/kvdb/store"
	"github.com/dfuse-io/kvdb/store/storetest"
	"github.com/stretchr/testify/require"
)

func init() {
	//	if os.Getenv("DEBUG") != "" {
	//logging.Override(logging.MustCreateLoggerWithLevel("test", zap.NewAtomicLevelAt(zap.DebugLevel)))
	//	}
}

func TestAll(t *testing.T) {
	storetest.TestAll(t, "bigkv", newTestFactory(t))
}

func newTestFactory(t *testing.T) storetest.DriverFactory {
	return func() (store.KVStore, storetest.DriverCleanupFunc) {
		kvStore, err := NewStore("bigkv://dev.dev/dev?createTable=true&compression=none")
		if err != nil {
			t.Skip("bigtable unreachable, cannot run tests")
			return nil, nil
		}
		require.NoError(t, err)
		return kvStore, func() {
			kvStore.(io.Closer).Close()
		}
	}
}