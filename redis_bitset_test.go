package bloom_test

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/bculberson/bloom"
	"github.com/garyburd/redigo/redis"
)

func TestRedisBitSet_New_Set_Test(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Error("Miniredis could not start")
	}
	defer s.Close()

	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", s.Addr()) },
	}
	bitSet := bloom.NewRedisBitSet("test_key", pool)
	isSetBefore, err := bitSet.Test([]uint{0})
	if err != nil {
		t.Error("Could not test bitset in redis")
	}
	if isSetBefore {
		t.Error("Bit should not be set")
	}
	err = bitSet.Set([]uint{512})
	if err != nil {
		t.Error("Could not set bitset in redis")
	}
	isSetAfter, err := bitSet.Test([]uint{512})
	if err != nil {
		t.Error("Could not test bitset in redis")
	}
	if !isSetAfter {
		t.Error("Bit should be set")
	}
}
