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
	"strconv"
)

type BitSlice struct {
	array []int
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
	fmt.Printf("%#v\n", a)
	return a
}

// Temporary function while I define idiomatic names for accessing the values.
func (bs BitSlice) Get(pos int) int {
	return bs.array[pos]
}

// Temporary function while I define idiomatic names for accessing the values.
func (bs BitSlice) Set(pos, elem int) {
	bs.array[pos] = elem
}

// Function to print the slice in an idiomatic way.
func (bs BitSlice) String() string {
	buffer := "["

	for i := 0; i < Len(bs); i++ {
		buffer += strconv.Itoa(bs.Get(i))
		if i+1 < Len(bs) {
			buffer += " "
		}
	}
	buffer += "]"

	return buffer
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
//	slice = Append(slice, elem1, elem2)
//	slice = Append(slice, anotherBitSlice...)
// As a special case, it is legal to append a int to a bit slice, like this:
//	slice = append(slice, 0x01101100101)
// And the Append function will automatically append each separate bit.
func Append(slice BitSlice, elems ...int) BitSlice {
	return slice
}

func Copy(dst, src BitSlice) int {
	return 0
}
