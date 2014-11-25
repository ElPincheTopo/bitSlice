/*
   bitSlice
   Copyright (C) 2014 Roberto Lapuente topo@asustin.net

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public License
   along with this program. If not, see <http://www.gnu.org/licenses/>.

   For more information visit http://github.com/ElPincheTopo/bitSlice
   or send an e-mail to topo@asustin.net.
*/

// append, copy, String
package bitSlice

import (
	"fmt"
)

type BitSlice struct {
	array []uint8
	len   int
	cap   int
}

// New returns an empty bit slice with length and capacity value of 0.
func New() *BitSlice {
	return new(BitSlice)
}

// Make returns a bit slice with the specified length and capacity.
func Make(length, capacity int) BitSlice {
	a := BitSlice{make([]int, length, capacity), length, capacity}
	return a
}

// Temporary function while I define idiomatic names for accessing the values.
func (bs BitSlice) Get(pos int) int {
	if pos >= len {
		panic("runtime error: index out of range")
	}
	
	subArray := pos / 8    // Get position of the corresponding unit8 in array
	pos = 7 - (pos - 8*subArray) // Transform the array position to the subArray 
								 // position. Then get the "math" position.
	return ((underlyingArray[subArray] & (uint8(1) << uint8(pos))) >> uint8(pos))
}

// Temporary function while I define idiomatic names for accessing the values.
func (bs *BitSlice) Set(pos int) {
	if pos >= bs.len {
		panic("runtime error: index out of range")
	}

	subArray := pos/8 // Get position of the corresponding unit8
	pos = pos - 8*subArray // Transform the array position to the subArray position
	bs.array[subArray] |= 1 << uint8(7-pos)
}

func (bs *BitSlice) Unset(pos int) {
	if pos >= bs.len {
		panic("runtime error: index out of range")
	}

	subArray := pos/8 // Get position of the corresponding unit8
	pos = pos - 8*subArray // Transform the array position to the subArray position
	bs.array[subArray] &^= 1 << uint8(7-pos)
}

// Function to print the slice in an idiomatic way.
func (bs BitSlice) String() string {
	return fmt.Sprintf("[%08b]", array)
}

// The Len functions returns the number of elements on the bit slice.
func Len(slice BitSlice) int {
	return slice.len
}

// The Cap function returns the maximum length the slice can reach when resliced;
func Cap(slice BitSlice) int {
	return slice.cap
}

// The Append function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//  slice = Append(slice, elem1, elem2)
//  slice = Append(slice, anotherBitSlice...)
// As a special case, it is legal to append a int to a bit slice, like this:
//  slice = append(slice, 0x01101100101)
// And the Append function will automatically append each separate bit.
func Append(slice BitSlice, elems ...int) BitSlice {
	return slice
}

func Copy(dst, src BitSlice) int {
	return 0
}
