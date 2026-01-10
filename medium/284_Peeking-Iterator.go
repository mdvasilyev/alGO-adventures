package main

//   Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	panic("does not matter")
	// Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	panic("does not matter")
	// Returns the next element in the iteration.
}

type PeekingIterator struct {
	pos int
	buf []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	peekingIterator := &PeekingIterator{
		pos: 0,
		buf: make([]int, 0),
	}
	for iter.hasNext() {
		peekingIterator.buf = append(peekingIterator.buf, iter.next())
	}
	return peekingIterator
}

func (this *PeekingIterator) hasNext() bool {
	return this.pos < len(this.buf)
}

func (this *PeekingIterator) next() int {
	if this.hasNext() {
		val := this.buf[this.pos]
		this.pos++
		return val
	}
	return 0
}

func (this *PeekingIterator) peek() int {
	if this.hasNext() {
		return this.buf[this.pos]
	}
	return 0
}
