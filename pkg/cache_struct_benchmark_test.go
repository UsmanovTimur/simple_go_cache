package pkg

import (
	"testing"
)

// Комплексное тестирование создания, кэша, получения ключа и удаления.
func BenchmarkCacheStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newCache, err := NewCache()
		if err != nil {
			b.Fatalf("Ошибка создания: %s", err)
		}
		err = newCache.Set("key", 123)
		if err != nil {
			b.Fatalf("Ошибка добавления элемента: %s", err)
		}
		if _, err = newCache.Get("key"); err != nil {
			b.Fatalf("Ошибка получения элемента: %s", err)
		}
		if err = newCache.Delete("key"); err != nil {
			b.Fatalf("Ошибка удаления элемента: %s", err)
		}
	}
}
