package hw04_lru_cache //nolint:golint,stylecheck
import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(2)

		c.Set("Not Found", 404)
		c.Set("OK", 200)
		c.Clear()
		c.Set("No Content", 204)

		val, ok := c.Get("Not Found")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("OK")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("No Content")
		require.True(t, ok)
		require.Equal(t, val, 204)
	})

	t.Run("pushing out rarely used items", func(t *testing.T) {
		capacity := 100
		c := NewCache(capacity)

		for i := 0; i < 1_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
			if i > capacity-1 {
				val, ok := c.Get(Key(strconv.Itoa(i - capacity)))
				require.False(t, ok)
				require.Nil(t, val)
			}
		}

		for i := 1_000 - capacity; i < 1_000; i++ {
			val, ok := c.Get(Key(strconv.Itoa(i)))
			require.True(t, ok)
			require.Equal(t, val, i)
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
