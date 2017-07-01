package rtda

import (
	"go/scanner"
	"math"
)

type OperandStack struct {
	size uint
	slots []Slot
}

func newOperandStack(maxStackSize uint) *OperandStack {
	if maxStackSize > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStackSize),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size ++
}

func (self *OperandStack) PopInt() int32 {
	self.size --
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val>>32)
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := self.slots[self.size].num
	high := self.slots[self.size+1].num
	return int64(high)<<32 | int64(low)
}

