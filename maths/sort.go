package maths

import "sort"

type Int32Slice []int32
type Int64Slice []int64
type Float32Slice []float32

func SortInt32s(s []int32) {
	sort.Sort(Int32Slice(s))
}

func SortInt64s(s []int64) {
	sort.Sort(Int64Slice(s))
}

func SortFloat32s(s []float32) {
	sort.Sort(Float32Slice(s))
}

func (this Int32Slice) Len() int {
	return len(this)
}

func (this Int32Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int32Slice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Int64Slice) Len() int {
	return len(this)
}

func (this Int64Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64Slice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Float32Slice) Len() int {
	return len(this)
}

func (this Float32Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Float32Slice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
