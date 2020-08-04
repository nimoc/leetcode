package main

import (
	"log"
)

func main () {
	q := Constructor(5);
	q.EnQueue(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	log.Print(q.Rear())
	log.Print(q.IsFull())
	q.DeQueue()
	q.EnQueue(4)
	log.Print(q.Rear())
	log.Printf("%+v", q)
}

type MyCircularQueue struct {
	cap int
	head int
	size int
	data []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap: k,
		data: make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {return false}
	this.size++
	this.data[this.Tail()] = value
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() { return  false}
	this.size--
	this.head = (this.head+1) % this.cap
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {return -1}
	return this.data[this.head]
}

func (this *MyCircularQueue) Tail() int {
	return (this.head + this.size -1) % this.cap
}
/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {return -1}
	return this.data[this.Tail()]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}
