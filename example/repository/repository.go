package repository

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("item not found")

// Repository is a generic interface for basic CRUD operations.
// T is the entity type, ID is the type of the identifier.
type Repository[T any, ID comparable] interface {
	Create(id ID, entity T) error
	Get(id ID) (T, error)
	List() ([]T, error)
	Delete(id ID) error
}

// MemoryRepository is a thread-safe in-memory implementation of Repository.
type MemoryRepository[T any, ID comparable] struct {
	mu   sync.RWMutex
	data map[ID]T
}

func NewMemoryRepository[T any, ID comparable]() *MemoryRepository[T, ID] {
	return &MemoryRepository[T, ID]{
		data: make(map[ID]T),
	}
}

func (r *MemoryRepository[T, ID]) Create(id ID, entity T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; exists {
		return errors.New("already exists")
	}

	r.data[id] = entity
	return nil
}

func (r *MemoryRepository[T, ID]) Get(id ID) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entity, exists := r.data[id]
	if !exists {
		var zero T
		return zero, ErrNotFound
	}
	return entity, nil
}

func (r *MemoryRepository[T, ID]) List() ([]T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]T, 0, len(r.data))
	for _, item := range r.data {
		items = append(items, item)
	}
	return items, nil
}

func (r *MemoryRepository[T, ID]) Delete(id ID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return ErrNotFound
	}

	delete(r.data, id)
	return nil
}
