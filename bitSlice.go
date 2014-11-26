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

// Package bitSlice implements a slice of bits implemented with an array of
// uint8. This way, bitSlice, at most only wastes 7 bits more than the
// space the user asked to allocate.
package bitSlice

import (
	"fmt"
)

type BitSlice struct {
	array []uint8
	len   int
	cap   int
	first int
}

// Make returns a bit slice with the specified length and capacity. The slice in
// the parameters is there to simulate an "optional" capacity parameter.
func Make(length int, capacity ...int) BitSlice {
	var cap int

	if len(capacity) > 1 {
		panic("To many arguments in Make function")
	} else if len(capacity) < 1 {
		cap = length
	} else {
		cap = capacity[0]
		if cap < length {
			panic("Slice capacity can't be less than it's length")
		}
	}
	arrayLen := length / 8
	if length%8 != 0 {
		arrayLen++
	}
	//arrayCap := cap/8.0
	return BitSlice{make([]uint8, arrayLen), length, cap, 0}
}

// Get returns the value of the pos position of the array.
func (bs BitSlice) Get(pos int) int {
	if pos >= bs.len {
		panic("runtime error: index out of range")
	}

	// Convert slice position into the position in the underlying array
	pos = pos + bs.first

	subArray := pos / 8          // Get position of the corresponding unit8 in array
	pos = 7 - (pos - 8*subArray) // Transform the array position to the subArray
	// position. Then get the "math" position.
	return int((bs.array[subArray] & (uint8(1) << uint8(pos))) >> uint8(pos))
}

// Set sets the pos position in the array to 1.
func (bs BitSlice) Set(pos int) {
	if pos >= bs.len {
		panic("runtime error: index out of range")
	}

	// Convert slice position into the position in the underlying array
	pos = pos + bs.first

	subArray := pos / 8    // Get position of the corresponding unit8
	pos = pos - 8*subArray // Transform the array position to the subArray position
	bs.array[subArray] |= 1 << uint8(7-pos)
}

// Unset sets the pos position in the array to 0.
func (bs BitSlice) Unset(pos int) {
	if pos >= bs.len {
		panic("runtime error: index out of range")
	}

	// Convert slice position into the position in the underlying array
	pos = pos + bs.first

	subArray := pos / 8    // Get position of the corresponding unit8
	pos = pos - 8*subArray // Transform the array position to the subArray position
	bs.array[subArray] &^= 1 << uint8(7-pos)
}

// The slice function returns a new slice with the same underlying array starting
// in begin and ending in end, the new slice is like a subset of the original but
// with different pointers for the start and end position. The slice in the
// parameters is there to simulate optional parameters.
func (bs BitSlice) Slice(positions ...int) BitSlice {
	var begin, end int
	if len(positions) > 2 {
		panic("To many arguments in Slice function")
	}
	switch len(positions) {
	case 2:
		begin = positions[0]
		end = positions[1]
	case 1:
		begin = positions[0]
		end = bs.len
	case 0:
		begin = 0
		end = bs.len
	}

	if begin >= bs.len || end > bs.len {
		panic("runtime error: slice bounds out of range")
	}
	if begin > end {
		panic(fmt.Sprintf("invalid slice index: %d > %d", begin, end))
	}

	begin = begin + bs.first
	end = end + bs.first
	return BitSlice{bs.array, end - begin, bs.cap - begin, begin}
}

// Function to print the slice in an idiomatic way.
func (bs BitSlice) String() string {
	var str string
	for _, byte := range bs.array {
		str += fmt.Sprintf("%08b", byte)
	}
	str = str[bs.first : bs.first+bs.len]

	var fmtStr = "["
	for i := 1; i <= len(str); i++ {
		fmtStr += string(str[i-1])
		if i%4 == 0 && i != len(str) {
			fmtStr += " "
		}
	}

	fmtStr += "]"
	return fmtStr
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

// The copy function copies bits from a source slice into a destination slice.
// (As a special case, it also will copy the corresponding bits from integers..)
// The source and destination may overlap. Copy returns the number of elements
// copied, which will be the minimum of len(src) and len(dst).
func Copy(dst, src BitSlice) int {
	return 0
}
