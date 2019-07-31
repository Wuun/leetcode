package protocol

import "fmt"

//Element is the inner element of a Queue,which could be every type.
type Element interface{}

//Queue is a queue implement
type Queue interface {
	Pop() Element
	Push(e Element)
	IsEmpty() bool
	Clear()
	Size() int
}

//ElementEntry is an implement of QUeue
type ElementEntry struct {
	elementSlice []Element
}

//NewQueue is a constructor of Queue
func NewQueue() *ElementEntry {
	return new(ElementEntry)
}

//Push can send an Element to Queue
func (entry *ElementEntry) Push(e Element) {
	entry.elementSlice = append(entry.elementSlice, e)
}

//Pop can get an Element from Queue
func (entry *ElementEntry) Pop() Element {
	if entry.IsEmpty() {
		return nil
	}

	firstElement := entry.elementSlice[0]
	entry.elementSlice = entry.elementSlice[1:len(entry.elementSlice)]
	return firstElement
}

//IsEmpty can judge if the Queue is empty or not
func (entry *ElementEntry) IsEmpty() bool {
	if len(entry.elementSlice) == 0 {
		return true
	}

	return false
}

//Clear is used to empty the Queue
func (entry *ElementEntry) Clear() {
	if entry.IsEmpty() {
		fmt.Print("this queue is empty,which don/'t need to clear")
		return
	}

	entry.elementSlice = nil
}

//Size is used  to count the Elements in Queue
func (entry *ElementEntry) Size() int {
	return len(entry.elementSlice)
}
