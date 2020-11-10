package pkg

import (
	"errors"
	"sync"
)

func NewCache() (CacheStorage, error) {
	newCacheStorage := CacheStorage{}
	newCacheStorage.Storage = make(map[string]interface{})
	return newCacheStorage, nil
}

//Основное хранилище данных
type CacheStorage struct {
	Mutex   sync.Mutex
	Storage map[string]interface{}
}

//Получение значения по ключу
func (storage *CacheStorage) Get(key string) (interface{}, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	val, ok := storage.Storage[key]
	if !ok {
		return nil, errors.New("Элемент не найден ")
	}
	return val, nil
}

//Получение значения по ключу
func (storage *CacheStorage) Set(key string, value interface{}) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.Storage[key] = value
	return nil
}

//Получение значения по ключу
func (storage *CacheStorage) Delete(key string) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if _, ok := storage.Storage[key]; !ok {
		return errors.New("Ключ не найден ")
	} else {
		delete(storage.Storage, key)
		return nil
	}
}
