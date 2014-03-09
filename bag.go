package utils

import (
	"fmt"
	"reflect"
)

type Bag struct {
	data []interface{}
	size int
}

func NewBag(capacity int) *Bag {
	return &Bag{
		data: make([]interface{}, capacity),
		size: 0,
	}
}

// Set index to value
func (this *Bag) Set(index int, value interface{}) {
	if index == this.size {
		this.Grow()
	}

	this.data[index] = value
	this.size += 1
}

// Add value to last avaible element
func (this *Bag) Add(value interface{}) {
	this.Set(this.size+1, value)
}

// Grow by 64
func (this *Bag) Grow() {
	oldData := this.data
	newData := make([]interface{}, this.size+64)
	copy(oldData, newData)
	this.data = newData
}

// Return array (data)
func (this *Bag) GetData() []interface{} {
	return this.data
}

// Get data from index
func (this *Bag) Get(index int) interface{} {
	return this.data[index]
}

// See if bag contains value
func (this *Bag) Contains(value interface{}) bool {
	checkval := reflect.TypeOf(value)
	for _, v := range this.data {
		if checkval == reflect.TypeOf(v) {
			return true
		}
	}

	return false
}

// Make sure bag has enough space, if not - grow
func (this *Bag) EnsureCapacity(index int) {
	if index > len(this.data) {
		this.Grow()
	}
}

// Print data
func (this *Bag) Debug() {
	for k, v := range this.data {
		fmt.Println(k, "[", reflect.TypeOf(v), "]")
	}
}
