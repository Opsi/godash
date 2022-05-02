package godash

import "sync"

// A generic atomic type.
type Atomic[T any] struct {
	value T
	mu    sync.RWMutex
}

// Get the value of the atomic in a thread-safe manner.
func (a *Atomic[T]) Get() T {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.value
}

// Set the value of the atomic in a thread-safe manner.
func (a *Atomic[T]) Set(newValue T) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.value = newValue
}

// Calculate a new value based on the old value in a thread-safe manner. Returns the new value.
func (a *Atomic[T]) CalcAndSet(calc func(T) T) T {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.value = calc(a.value)
	return a.value
}

// Creates a read-only reference of the Atomic
func (a *Atomic[T]) ReadOnly() RAtomic[T] {
	return RAtomic[T]{a}
}

// A wrapper for an Atomic[T] that can be used as a thread-safe read-only reference.
type RAtomic[T any] struct {
	a *Atomic[T]
}

// Get the value of the atomic in a thread-safe manner.
func (roa RAtomic[T]) Get() T {
	return roa.a.Get()
}
