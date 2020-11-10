package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Проверяем создание кэша
func TestNewCache(t *testing.T) {
	newCache, err := NewCache()
	assert.Nil(t, err)
	assert.Equal(t, newCache.Storage, map[string]interface{}{})
}

//Проверка получения значения по ключу
func TestGetCache(t *testing.T) {
	newCache, err := NewCache()
	assert.Nil(t, err)
	newCache.Storage["key"] = 123
	val, err := newCache.Get("key")
	assert.Nil(t, err)
	assert.Equal(t, val, 123)
	newCache.Storage["key"] = "123"
	val, err = newCache.Get("key")
	assert.Nil(t, err)
	assert.Equal(t, val, "123")
}

//Проверка получения значения по ключу
func TestSetCache(t *testing.T) {
	newCache, err := NewCache()
	assert.Nil(t, err)
	err = newCache.Set("key", 123)
	assert.Nil(t, err)
	assert.Equal(t, newCache.Storage["key"], 123)
	err = newCache.Set("key", "123")
	assert.Nil(t, err)
	assert.Equal(t, newCache.Storage["key"], "123")
}

//Проверка удаления значения по ключу
func TestDeleteCache(t *testing.T) {
	newCache, err := NewCache()
	assert.Nil(t, err)
	newCache.Storage["key"] = 123
	err = newCache.Delete("key")
	assert.Nil(t, err)
	_, ok := newCache.Storage["key"]
	assert.False(t, ok)
}
