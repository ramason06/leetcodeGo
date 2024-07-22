package common

type Queue[T any] struct {
	Items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.Items) == 0 {
		return zero, false
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if len(q.Items) == 0 {
		return zero, false
	}
	return q.Items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Items) == 0
}
