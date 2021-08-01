package bitter

import (
	"fmt"
)

type (
	QuartetPosition int
	OctetIndex      int
	QuartetIndex    int
)

const (
	Lo    QuartetPosition = 0
	Hi    QuartetPosition = 1
	QBit0 QuartetIndex    = 0
	QBit1 QuartetIndex    = 1
	QBit2 QuartetIndex    = 2
	QBit3 QuartetIndex    = 3
	Bit0  OctetIndex      = 0
	Bit1  OctetIndex      = 1
	Bit2  OctetIndex      = 2
	Bit3  OctetIndex      = 3
	Bit4  OctetIndex      = 4
	Bit5  OctetIndex      = 5
	Bit6  OctetIndex      = 6
	Bit7  OctetIndex      = 7
)

type (
	Octet struct {
		bits byte
	}
	Quartet struct {
		bits byte
	}
	QuartetMap struct {
		B0 bool
		B1 bool
		B2 bool
		B3 bool
	}
	OctetMap struct {
		B0 bool
		B1 bool
		B2 bool
		B3 bool
		B4 bool
		B5 bool
		B6 bool
		B7 bool
	}
)

func OctetFromQuartets(hi, lo *Quartet) *Octet {
	newbyte := (hi.bits << 4) + lo.bits
	return &Octet{
		bits: newbyte,
	}
}

func OctetFromByte(b byte) *Octet {
	return &Octet{
		bits: b,
	}
}

func QuartetsFromByte(b byte) (hi, lo *Quartet) {
	hi = &Quartet{
		bits: b >> 4,
	}
	lo = &Quartet{
		bits: (b << 4) >> 4,
	}
	return hi, lo
}

///////////////////////////////////////
// Bit funcs

// Set the bit to 1 at a given position in the byte
func (o *Octet) Set(index OctetIndex) *Octet {
	o.bits |= (1 << index)
	return o
}

// Unset the bit at a given position in the byte
func (o *Octet) Unset(index OctetIndex) *Octet {
	o.bits &^= (1 << index)
	return o
}

// Toggle the value of the bit at the given position in the byte
func (o *Octet) Toggle(index OctetIndex) *Octet {
	o.bits ^= 1 << index
	return o
}

// Set the bit to 1 at a given position in the byte
func (q *Quartet) Set(index QuartetIndex) *Quartet {
	q.bits |= (1 << index)
	return q
}

// Unset the bit at a given position in the byte
func (q *Quartet) Unset(index QuartetIndex) *Quartet {
	q.bits &^= (1 << index)
	return q
}

// Toggle the value of the bit at the given position in the byte
func (q *Quartet) Toggle(index QuartetIndex) *Quartet {
	q.bits ^= 1 << index
	return q
}

///////////////////////////////////////
// Overwriter funcs

func (q *Quartet) toHi() byte {
	return q.bits << 4
}

func clearHi(b byte) byte {
	ob := b
	return (ob << 4) >> 4
}

// Overwrite all bits in the octet with the values from a byte
func (o *Octet) Overwrite(val byte) *Octet {
	o.bits = val
	return o
}

// OverwriteQuartet all bits in the octet with the values from a byte
func (o *Octet) OverwriteQuartet(val *Quartet, pos QuartetPosition) *Octet {
	// clear the part of the quartet that is supposed to be zeroed just in case
	val.bits = clearHi(val.bits)
	if pos == Hi {
		// clear the hi quartet
		o.bits = (o.bits << 4) >> 4
		// add the quartet to the Octet
		o.bits = o.bits + val.toHi()
	} else if pos == Lo {
		// clear the lo quartet
		o.bits = (o.bits >> 4) << 4
		o.bits = o.bits + val.bits
	}
	return o
}

// Nullify zeroes out (ie Unset) all the bits in the Octet
func (o *Octet) Nullify() *Octet {
	o.bits = 0b00000000
	return o
}

///////////////////////////////////////
// Reader funcs

// IsSet checks if the bit at a given position is set
func (o *Octet) IsSet(index OctetIndex) bool {
	return ((o.bits>>index)&1 == 1)
}

// IsSet checks if the bit at a given position is set
func (q *Quartet) IsSet(index QuartetIndex) bool {
	return ((q.bits>>index)&1 == 1)
}

// QuartetMaps gets the hi and lo nibbles in an Octet as boolean maps
func (o *Octet) QuartetMaps() (hi, lo QuartetMap) {
	lo = QuartetMap{
		B0: ((o.bits>>0)&1 == 1),
		B1: ((o.bits>>1)&1 == 1),
		B2: ((o.bits>>2)&1 == 1),
		B3: ((o.bits>>3)&1 == 1),
	}
	hi = QuartetMap{
		B0: ((o.bits>>4)&1 == 1),
		B1: ((o.bits>>5)&1 == 1),
		B2: ((o.bits>>6)&1 == 1),
		B3: ((o.bits>>7)&1 == 1),
	}
	return hi, lo
}

func (o *Octet) Quartets() (hi, lo *Quartet) {
	hi = &Quartet{
		bits: o.bits >> 4,
	}
	lo = &Quartet{
		bits: (o.bits << 4) >> 4,
	}
	return hi, lo
}

// OctetMap
func (o *Octet) OctetMap() OctetMap {
	return OctetMap{
		B0: ((o.bits>>0)&1 == 1),
		B1: ((o.bits>>1)&1 == 1),
		B2: ((o.bits>>2)&1 == 1),
		B3: ((o.bits>>3)&1 == 1),
		B4: ((o.bits>>4)&1 == 1),
		B5: ((o.bits>>5)&1 == 1),
		B6: ((o.bits>>6)&1 == 1),
		B7: ((o.bits>>7)&1 == 1),
	}
}

// BinaryString returns the binary representation of the octet
func (o *Octet) BinaryString() string {
	return fmt.Sprintf("%08b", o.bits)
}

func (o *Octet) HexString() string {
	return fmt.Sprintf("%08b", o.bits)
}

// Byte returns the protected byte that holds the Octet bit values
func (o *Octet) Byte() byte {
	return o.bits
}
