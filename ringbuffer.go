package ringbuffer

import (
	"errors"
)

type RingBuffer[T any] struct {
	buffer   []T
	size     int
	capacity int
	head     int
	tail     int
}

// Creates new Ring Buffer with default capacity of 32 elements.
func New[T any]() *RingBuffer[T] {
	return NewWithCapacity[T](32)
}

// Creates new Ring Buffer with given capacity.
func NewWithCapacity[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer:   make([]T, capacity),
		size:     0,
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// Pushes element to the end of the Ring Buffer.
// Extends capacity if needed.
func (rb *RingBuffer[T]) PushBack(element T) {
	if rb.full() {
		rb.extendCapacity()
	}
	if rb.empty() {
		rb.head = 0
		rb.tail = 0
	} else {
		rb.tail = (rb.tail + 1) % rb.capacity
	}
	rb.size++
	rb.buffer[rb.tail] = element
}

// Pushes element to the front of the Ring Buffer.
// Extends capacity if needed.
func (rb *RingBuffer[T]) PushFront(element T) {
	if rb.full() {
		rb.extendCapacity()
	}
	if rb.empty() {
		rb.head = 0
		rb.tail = 0
	} else {
		rb.head = ((rb.head-1)%rb.capacity + rb.capacity) % rb.capacity
	}
	rb.size++
	rb.buffer[rb.head] = element
}

// Removes and returns element from the end of the Ring Buffer.
// Returns error on empty Ring Buffer.
func (rb *RingBuffer[T]) PopBack() (T, error) {
	if rb.empty() {
		var zero T
		return zero, errors.New("PopBack on an empty RingBuffer.")
	}
	value := rb.buffer[rb.tail]
	rb.tail = ((rb.tail-1)%rb.capacity + rb.capacity) % rb.capacity
	rb.size--
	return value, nil
}

// Removes and returns element from the front of the Ring Buffer.
// Returns error on empty Ring Buffer.
func (rb *RingBuffer[T]) PopFront() (T, error) {
	if rb.empty() {
		var zero T
		return zero, errors.New("PopFront on an empty RingBuffer.")
	}
	value := rb.buffer[rb.head]
	rb.head = (rb.head + 1) % rb.capacity
	rb.size--
	return value, nil
}

// Returns element from the back of the Ring Buffer without removing it.
// Returns error on empty Ring Buffer.
func (rb *RingBuffer[T]) PeekFront() (T, error) {
	if rb.empty() {
		var zero T
		return zero, errors.New("PeekFront on an empty RingBuffer.")
	}
	return rb.buffer[rb.head], nil
}

// Returns element from the front of the Ring Buffer without removing it.
// Returns error on empty Ring Buffer.
func (rb *RingBuffer[T]) PeekBack() (T, error) {
	if rb.empty() {
		var zero T
		return zero, errors.New("PeekBack on an empty RingBuffer.")
	}
	return rb.buffer[rb.tail], nil
}

func (rb *RingBuffer[T]) empty() bool {
	return rb.size == 0
}

func (rb *RingBuffer[T]) full() bool {
	return rb.size == rb.capacity
}

func (rb *RingBuffer[T]) extendCapacity() {
	new_capacity := rb.capacity * 2
	new_buffer := make([]T, new_capacity)
	for offset := range rb.size {
		new_buffer[offset] = rb.buffer[(rb.head+offset)%rb.capacity]
	}
	rb.buffer = new_buffer
	rb.capacity = new_capacity
	rb.head = 0
	rb.tail = rb.size - 1
}
