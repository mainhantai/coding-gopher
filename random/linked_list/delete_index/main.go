package main

import "fmt"

type node struct {
	data     int
	nextNode *node
}

type linkedList struct {
	headnode *node
	length   int
}

func (l linkedList) printData() {
	toPrint := l.headnode
	fmt.Println("Print Process")
	for l.length != 0 {
		fmt.Printf("%d\n", toPrint.data)
		l.length--
		toPrint = toPrint.nextNode
	}
}

func (l *linkedList) DeleteIndex(index int) bool {
	if index >= l.length {
		return false
	}
	ptr := l.headnode
	for i := 0; i < l.length; i++ {
		if index == i+1 {
			ptr.nextNode = ptr.nextNode.nextNode
		} else {
			ptr = ptr.nextNode
		}
	}
	l.length--
	return true
}

func (l *linkedList) insert(value int) {
	n := node{value, nil}
	if l.length == 0 {
		l.headnode = &n
		l.length++
		return
	}
	ptr := l.headnode
	for i := 0; i < l.length; i++ {
		if ptr.nextNode == nil {
			ptr.nextNode = &n
			l.length++
			return
		}
		ptr = ptr.nextNode
	}
}

func main() {
	mylist := linkedList{}
	mylist.insert(10)
	mylist.printData()
	fmt.Println(mylist.DeleteIndex(2))
	mylist.insert(20)
	mylist.insert(450)
	mylist.insert(25)
	mylist.printData()
	fmt.Println(mylist.DeleteIndex(1))
	mylist.insert(96)
	mylist.printData()
}
