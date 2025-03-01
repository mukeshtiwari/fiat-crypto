// Code generated by Fiat Cryptography. DO NOT EDIT.
//
// Autogenerated: 'src/ExtractionOCaml/unsaturated_solinas' --lang Go --no-wide-int --relax-primitive-carry-to-bitwidth 32,64 --cmovznz-by-mul --internal-static --package-case flatcase --public-function-case UpperCamelCase --private-function-case camelCase --public-type-case UpperCamelCase --private-type-case camelCase --no-prefix-fiat --doc-newline-in-typedef-bounds --doc-prepend-header 'Code generated by Fiat Cryptography. DO NOT EDIT.' --doc-text-before-function-name '' --doc-text-before-type-name '' --package-name poly1305 '' 64 3 '2^130 - 5' carry_mul carry_square carry add sub opp selectznz to_bytes from_bytes relax carry_add carry_sub carry_opp
//
// curve description (via package name): poly1305
//
// machine_wordsize = 64 (from "64")
//
// requested operations: carry_mul, carry_square, carry, add, sub, opp, selectznz, to_bytes, from_bytes, relax, carry_add, carry_sub, carry_opp
//
// n = 3 (from "3")
//
// s-c = 2^130 - [(1, 5)] (from "2^130 - 5")
//
// tight_bounds_multiplier = 1 (from "")
//
//
//
// Computed values:
//
//   carry_chain = [0, 1, 2, 0, 1]
//
//   eval z = z[0] + (z[1] << 44) + (z[2] << 87)
//
//   bytes_eval z = z[0] + (z[1] << 8) + (z[2] << 16) + (z[3] << 24) + (z[4] << 32) + (z[5] << 40) + (z[6] << 48) + (z[7] << 56) + (z[8] << 64) + (z[9] << 72) + (z[10] << 80) + (z[11] << 88) + (z[12] << 96) + (z[13] << 104) + (z[14] << 112) + (z[15] << 120) + (z[16] << 128)
//
//   balance = [0x1ffffffffff6, 0xffffffffffe, 0xffffffffffe]
package poly1305

import "math/bits"

type uint1 uint64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927
type int1 int64 // We use uint64 instead of a more narrow type for performance reasons; see https://github.com/mit-plv/fiat-crypto/pull/1006#issuecomment-892625927

// LooseFieldElement is a field element with loose bounds.
//
// Bounds:
//
//   [[0x0 ~> 0x300000000000], [0x0 ~> 0x180000000000], [0x0 ~> 0x180000000000]]
type LooseFieldElement [3]uint64

// TightFieldElement is a field element with tight bounds.
//
// Bounds:
//
//   [[0x0 ~> 0x100000000000], [0x0 ~> 0x80000000000], [0x0 ~> 0x80000000000]]
type TightFieldElement [3]uint64

// addcarryxU44 is an addition with carry.
//
// Postconditions:
//   out1 = (arg1 + arg2 + arg3) mod 2^44
//   out2 = ⌊(arg1 + arg2 + arg3) / 2^44⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0xfffffffffff]
//   arg3: [0x0 ~> 0xfffffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0xfffffffffff]
//   out2: [0x0 ~> 0x1]
func addcarryxU44(out1 *uint64, out2 *uint1, arg1 uint1, arg2 uint64, arg3 uint64) {
	x1 := ((uint64(arg1) + arg2) + arg3)
	x2 := (x1 & 0xfffffffffff)
	x3 := uint1((x1 >> 44))
	*out1 = x2
	*out2 = x3
}

// subborrowxU44 is a subtraction with borrow.
//
// Postconditions:
//   out1 = (-arg1 + arg2 + -arg3) mod 2^44
//   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^44⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0xfffffffffff]
//   arg3: [0x0 ~> 0xfffffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0xfffffffffff]
//   out2: [0x0 ~> 0x1]
func subborrowxU44(out1 *uint64, out2 *uint1, arg1 uint1, arg2 uint64, arg3 uint64) {
	x1 := ((int64(arg2) - int64(arg1)) - int64(arg3))
	x2 := int1((x1 >> 44))
	x3 := (uint64(x1) & 0xfffffffffff)
	*out1 = x3
	*out2 = (0x0 - uint1(x2))
}

// addcarryxU43 is an addition with carry.
//
// Postconditions:
//   out1 = (arg1 + arg2 + arg3) mod 2^43
//   out2 = ⌊(arg1 + arg2 + arg3) / 2^43⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x7ffffffffff]
//   arg3: [0x0 ~> 0x7ffffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x7ffffffffff]
//   out2: [0x0 ~> 0x1]
func addcarryxU43(out1 *uint64, out2 *uint1, arg1 uint1, arg2 uint64, arg3 uint64) {
	x1 := ((uint64(arg1) + arg2) + arg3)
	x2 := (x1 & 0x7ffffffffff)
	x3 := uint1((x1 >> 43))
	*out1 = x2
	*out2 = x3
}

// subborrowxU43 is a subtraction with borrow.
//
// Postconditions:
//   out1 = (-arg1 + arg2 + -arg3) mod 2^43
//   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^43⌋
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0x7ffffffffff]
//   arg3: [0x0 ~> 0x7ffffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0x7ffffffffff]
//   out2: [0x0 ~> 0x1]
func subborrowxU43(out1 *uint64, out2 *uint1, arg1 uint1, arg2 uint64, arg3 uint64) {
	x1 := ((int64(arg2) - int64(arg1)) - int64(arg3))
	x2 := int1((x1 >> 43))
	x3 := (uint64(x1) & 0x7ffffffffff)
	*out1 = x3
	*out2 = (0x0 - uint1(x2))
}

// cmovznzU64 is a single-word conditional move.
//
// Postconditions:
//   out1 = (if arg1 = 0 then arg2 else arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [0x0 ~> 0xffffffffffffffff]
//   arg3: [0x0 ~> 0xffffffffffffffff]
// Output Bounds:
//   out1: [0x0 ~> 0xffffffffffffffff]
func cmovznzU64(out1 *uint64, arg1 uint1, arg2 uint64, arg3 uint64) {
	x1 := (uint64(arg1) * 0xffffffffffffffff)
	x2 := ((x1 & arg3) | ((^x1) & arg2))
	*out1 = x2
}

// CarryMul multiplies two field elements and reduces the result.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 * eval arg2) mod m
//
func CarryMul(out1 *TightFieldElement, arg1 *LooseFieldElement, arg2 *LooseFieldElement) {
	var x1 uint64
	var x2 uint64
	x2, x1 = bits.Mul64(arg1[2], (arg2[2] * 0x5))
	var x3 uint64
	var x4 uint64
	x4, x3 = bits.Mul64(arg1[2], (arg2[1] * 0xa))
	var x5 uint64
	var x6 uint64
	x6, x5 = bits.Mul64(arg1[1], (arg2[2] * 0xa))
	var x7 uint64
	var x8 uint64
	x8, x7 = bits.Mul64(arg1[2], arg2[0])
	var x9 uint64
	var x10 uint64
	x10, x9 = bits.Mul64(arg1[1], (arg2[1] * 0x2))
	var x11 uint64
	var x12 uint64
	x12, x11 = bits.Mul64(arg1[1], arg2[0])
	var x13 uint64
	var x14 uint64
	x14, x13 = bits.Mul64(arg1[0], arg2[2])
	var x15 uint64
	var x16 uint64
	x16, x15 = bits.Mul64(arg1[0], arg2[1])
	var x17 uint64
	var x18 uint64
	x18, x17 = bits.Mul64(arg1[0], arg2[0])
	var x19 uint64
	var x20 uint64
	x19, x20 = bits.Add64(x5, x3, uint64(0x0))
	var x21 uint64
	x21, _ = bits.Add64(x6, x4, uint64(uint1(x20)))
	var x23 uint64
	var x24 uint64
	x23, x24 = bits.Add64(x17, x19, uint64(0x0))
	var x25 uint64
	x25, _ = bits.Add64(x18, x21, uint64(uint1(x24)))
	x27 := ((x23 >> 44) | ((x25 << 20) & 0xffffffffffffffff))
	x28 := (x23 & 0xfffffffffff)
	var x29 uint64
	var x30 uint64
	x29, x30 = bits.Add64(x9, x7, uint64(0x0))
	var x31 uint64
	x31, _ = bits.Add64(x10, x8, uint64(uint1(x30)))
	var x33 uint64
	var x34 uint64
	x33, x34 = bits.Add64(x13, x29, uint64(0x0))
	var x35 uint64
	x35, _ = bits.Add64(x14, x31, uint64(uint1(x34)))
	var x37 uint64
	var x38 uint64
	x37, x38 = bits.Add64(x11, x1, uint64(0x0))
	var x39 uint64
	x39, _ = bits.Add64(x12, x2, uint64(uint1(x38)))
	var x41 uint64
	var x42 uint64
	x41, x42 = bits.Add64(x15, x37, uint64(0x0))
	var x43 uint64
	x43, _ = bits.Add64(x16, x39, uint64(uint1(x42)))
	var x45 uint64
	var x46 uint64
	x45, x46 = bits.Add64(x27, x41, uint64(0x0))
	x47 := (uint64(uint1(x46)) + x43)
	x48 := ((x45 >> 43) | ((x47 << 21) & 0xffffffffffffffff))
	x49 := (x45 & 0x7ffffffffff)
	var x50 uint64
	var x51 uint64
	x50, x51 = bits.Add64(x48, x33, uint64(0x0))
	x52 := (uint64(uint1(x51)) + x35)
	x53 := ((x50 >> 43) | ((x52 << 21) & 0xffffffffffffffff))
	x54 := (x50 & 0x7ffffffffff)
	x55 := (x53 * 0x5)
	x56 := (x28 + x55)
	x57 := (x56 >> 44)
	x58 := (x56 & 0xfffffffffff)
	x59 := (x57 + x49)
	x60 := uint1((x59 >> 43))
	x61 := (x59 & 0x7ffffffffff)
	x62 := (uint64(x60) + x54)
	out1[0] = x58
	out1[1] = x61
	out1[2] = x62
}

// CarrySquare squares a field element and reduces the result.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 * eval arg1) mod m
//
func CarrySquare(out1 *TightFieldElement, arg1 *LooseFieldElement) {
	x1 := (arg1[2] * 0x5)
	x2 := (x1 * 0x2)
	x3 := (arg1[2] * 0x2)
	x4 := (arg1[1] * 0x2)
	var x5 uint64
	var x6 uint64
	x6, x5 = bits.Mul64(arg1[2], x1)
	var x7 uint64
	var x8 uint64
	x8, x7 = bits.Mul64(arg1[1], (x2 * 0x2))
	var x9 uint64
	var x10 uint64
	x10, x9 = bits.Mul64(arg1[1], (arg1[1] * 0x2))
	var x11 uint64
	var x12 uint64
	x12, x11 = bits.Mul64(arg1[0], x3)
	var x13 uint64
	var x14 uint64
	x14, x13 = bits.Mul64(arg1[0], x4)
	var x15 uint64
	var x16 uint64
	x16, x15 = bits.Mul64(arg1[0], arg1[0])
	var x17 uint64
	var x18 uint64
	x17, x18 = bits.Add64(x15, x7, uint64(0x0))
	var x19 uint64
	x19, _ = bits.Add64(x16, x8, uint64(uint1(x18)))
	x21 := ((x17 >> 44) | ((x19 << 20) & 0xffffffffffffffff))
	x22 := (x17 & 0xfffffffffff)
	var x23 uint64
	var x24 uint64
	x23, x24 = bits.Add64(x11, x9, uint64(0x0))
	var x25 uint64
	x25, _ = bits.Add64(x12, x10, uint64(uint1(x24)))
	var x27 uint64
	var x28 uint64
	x27, x28 = bits.Add64(x13, x5, uint64(0x0))
	var x29 uint64
	x29, _ = bits.Add64(x14, x6, uint64(uint1(x28)))
	var x31 uint64
	var x32 uint64
	x31, x32 = bits.Add64(x21, x27, uint64(0x0))
	x33 := (uint64(uint1(x32)) + x29)
	x34 := ((x31 >> 43) | ((x33 << 21) & 0xffffffffffffffff))
	x35 := (x31 & 0x7ffffffffff)
	var x36 uint64
	var x37 uint64
	x36, x37 = bits.Add64(x34, x23, uint64(0x0))
	x38 := (uint64(uint1(x37)) + x25)
	x39 := ((x36 >> 43) | ((x38 << 21) & 0xffffffffffffffff))
	x40 := (x36 & 0x7ffffffffff)
	x41 := (x39 * 0x5)
	x42 := (x22 + x41)
	x43 := (x42 >> 44)
	x44 := (x42 & 0xfffffffffff)
	x45 := (x43 + x35)
	x46 := uint1((x45 >> 43))
	x47 := (x45 & 0x7ffffffffff)
	x48 := (uint64(x46) + x40)
	out1[0] = x44
	out1[1] = x47
	out1[2] = x48
}

// Carry reduces a field element.
//
// Postconditions:
//   eval out1 mod m = eval arg1 mod m
//
func Carry(out1 *TightFieldElement, arg1 *LooseFieldElement) {
	x1 := arg1[0]
	x2 := ((x1 >> 44) + arg1[1])
	x3 := ((x2 >> 43) + arg1[2])
	x4 := ((x1 & 0xfffffffffff) + ((x3 >> 43) * 0x5))
	x5 := (uint64(uint1((x4 >> 44))) + (x2 & 0x7ffffffffff))
	x6 := (x4 & 0xfffffffffff)
	x7 := (x5 & 0x7ffffffffff)
	x8 := (uint64(uint1((x5 >> 43))) + (x3 & 0x7ffffffffff))
	out1[0] = x6
	out1[1] = x7
	out1[2] = x8
}

// Add adds two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 + eval arg2) mod m
//
func Add(out1 *LooseFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := (arg1[0] + arg2[0])
	x2 := (arg1[1] + arg2[1])
	x3 := (arg1[2] + arg2[2])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
}

// Sub subtracts two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 - eval arg2) mod m
//
func Sub(out1 *LooseFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := ((0x1ffffffffff6 + arg1[0]) - arg2[0])
	x2 := ((0xffffffffffe + arg1[1]) - arg2[1])
	x3 := ((0xffffffffffe + arg1[2]) - arg2[2])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
}

// Opp negates a field element.
//
// Postconditions:
//   eval out1 mod m = -eval arg1 mod m
//
func Opp(out1 *LooseFieldElement, arg1 *TightFieldElement) {
	x1 := (0x1ffffffffff6 - arg1[0])
	x2 := (0xffffffffffe - arg1[1])
	x3 := (0xffffffffffe - arg1[2])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
}

// Selectznz is a multi-limb conditional select.
//
// Postconditions:
//   eval out1 = (if arg1 = 0 then eval arg2 else eval arg3)
//
// Input Bounds:
//   arg1: [0x0 ~> 0x1]
//   arg2: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
//   arg3: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
func Selectznz(out1 *[3]uint64, arg1 uint1, arg2 *[3]uint64, arg3 *[3]uint64) {
	var x1 uint64
	cmovznzU64(&x1, arg1, arg2[0], arg3[0])
	var x2 uint64
	cmovznzU64(&x2, arg1, arg2[1], arg3[1])
	var x3 uint64
	cmovznzU64(&x3, arg1, arg2[2], arg3[2])
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
}

// ToBytes serializes a field element to bytes in little-endian order.
//
// Postconditions:
//   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..16]
//
// Output Bounds:
//   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x3]]
func ToBytes(out1 *[17]uint8, arg1 *TightFieldElement) {
	var x1 uint64
	var x2 uint1
	subborrowxU44(&x1, &x2, 0x0, arg1[0], 0xffffffffffb)
	var x3 uint64
	var x4 uint1
	subborrowxU43(&x3, &x4, x2, arg1[1], 0x7ffffffffff)
	var x5 uint64
	var x6 uint1
	subborrowxU43(&x5, &x6, x4, arg1[2], 0x7ffffffffff)
	var x7 uint64
	cmovznzU64(&x7, x6, uint64(0x0), 0xffffffffffffffff)
	var x8 uint64
	var x9 uint1
	addcarryxU44(&x8, &x9, 0x0, x1, (x7 & 0xffffffffffb))
	var x10 uint64
	var x11 uint1
	addcarryxU43(&x10, &x11, x9, x3, (x7 & 0x7ffffffffff))
	var x12 uint64
	var x13 uint1
	addcarryxU43(&x12, &x13, x11, x5, (x7 & 0x7ffffffffff))
	x14 := (x12 << 7)
	x15 := (x10 << 4)
	x16 := (uint8(x8) & 0xff)
	x17 := (x8 >> 8)
	x18 := (uint8(x17) & 0xff)
	x19 := (x17 >> 8)
	x20 := (uint8(x19) & 0xff)
	x21 := (x19 >> 8)
	x22 := (uint8(x21) & 0xff)
	x23 := (x21 >> 8)
	x24 := (uint8(x23) & 0xff)
	x25 := uint8((x23 >> 8))
	x26 := (x15 + uint64(x25))
	x27 := (uint8(x26) & 0xff)
	x28 := (x26 >> 8)
	x29 := (uint8(x28) & 0xff)
	x30 := (x28 >> 8)
	x31 := (uint8(x30) & 0xff)
	x32 := (x30 >> 8)
	x33 := (uint8(x32) & 0xff)
	x34 := (x32 >> 8)
	x35 := (uint8(x34) & 0xff)
	x36 := uint8((x34 >> 8))
	x37 := (x14 + uint64(x36))
	x38 := (uint8(x37) & 0xff)
	x39 := (x37 >> 8)
	x40 := (uint8(x39) & 0xff)
	x41 := (x39 >> 8)
	x42 := (uint8(x41) & 0xff)
	x43 := (x41 >> 8)
	x44 := (uint8(x43) & 0xff)
	x45 := (x43 >> 8)
	x46 := (uint8(x45) & 0xff)
	x47 := (x45 >> 8)
	x48 := (uint8(x47) & 0xff)
	x49 := uint8((x47 >> 8))
	out1[0] = x16
	out1[1] = x18
	out1[2] = x20
	out1[3] = x22
	out1[4] = x24
	out1[5] = x27
	out1[6] = x29
	out1[7] = x31
	out1[8] = x33
	out1[9] = x35
	out1[10] = x38
	out1[11] = x40
	out1[12] = x42
	out1[13] = x44
	out1[14] = x46
	out1[15] = x48
	out1[16] = x49
}

// FromBytes deserializes a field element from bytes in little-endian order.
//
// Postconditions:
//   eval out1 mod m = bytes_eval arg1 mod m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x3]]
func FromBytes(out1 *TightFieldElement, arg1 *[17]uint8) {
	x1 := (uint64(arg1[16]) << 41)
	x2 := (uint64(arg1[15]) << 33)
	x3 := (uint64(arg1[14]) << 25)
	x4 := (uint64(arg1[13]) << 17)
	x5 := (uint64(arg1[12]) << 9)
	x6 := (uint64(arg1[11]) * uint64(0x2))
	x7 := (uint64(arg1[10]) << 36)
	x8 := (uint64(arg1[9]) << 28)
	x9 := (uint64(arg1[8]) << 20)
	x10 := (uint64(arg1[7]) << 12)
	x11 := (uint64(arg1[6]) << 4)
	x12 := (uint64(arg1[5]) << 40)
	x13 := (uint64(arg1[4]) << 32)
	x14 := (uint64(arg1[3]) << 24)
	x15 := (uint64(arg1[2]) << 16)
	x16 := (uint64(arg1[1]) << 8)
	x17 := arg1[0]
	x18 := (x16 + uint64(x17))
	x19 := (x15 + x18)
	x20 := (x14 + x19)
	x21 := (x13 + x20)
	x22 := (x12 + x21)
	x23 := (x22 & 0xfffffffffff)
	x24 := uint8((x22 >> 44))
	x25 := (x11 + uint64(x24))
	x26 := (x10 + x25)
	x27 := (x9 + x26)
	x28 := (x8 + x27)
	x29 := (x7 + x28)
	x30 := (x29 & 0x7ffffffffff)
	x31 := uint1((x29 >> 43))
	x32 := (x6 + uint64(x31))
	x33 := (x5 + x32)
	x34 := (x4 + x33)
	x35 := (x3 + x34)
	x36 := (x2 + x35)
	x37 := (x1 + x36)
	out1[0] = x23
	out1[1] = x30
	out1[2] = x37
}

// Relax is the identity function converting from tight field elements to loose field elements.
//
// Postconditions:
//   out1 = arg1
//
func Relax(out1 *LooseFieldElement, arg1 *TightFieldElement) {
	x1 := arg1[0]
	x2 := arg1[1]
	x3 := arg1[2]
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
}

// CarryAdd adds two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 + eval arg2) mod m
//
func CarryAdd(out1 *TightFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := (arg1[0] + arg2[0])
	x2 := ((x1 >> 44) + (arg1[1] + arg2[1]))
	x3 := ((x2 >> 43) + (arg1[2] + arg2[2]))
	x4 := ((x1 & 0xfffffffffff) + ((x3 >> 43) * 0x5))
	x5 := (uint64(uint1((x4 >> 44))) + (x2 & 0x7ffffffffff))
	x6 := (x4 & 0xfffffffffff)
	x7 := (x5 & 0x7ffffffffff)
	x8 := (uint64(uint1((x5 >> 43))) + (x3 & 0x7ffffffffff))
	out1[0] = x6
	out1[1] = x7
	out1[2] = x8
}

// CarrySub subtracts two field elements.
//
// Postconditions:
//   eval out1 mod m = (eval arg1 - eval arg2) mod m
//
func CarrySub(out1 *TightFieldElement, arg1 *TightFieldElement, arg2 *TightFieldElement) {
	x1 := ((0x1ffffffffff6 + arg1[0]) - arg2[0])
	x2 := ((x1 >> 44) + ((0xffffffffffe + arg1[1]) - arg2[1]))
	x3 := ((x2 >> 43) + ((0xffffffffffe + arg1[2]) - arg2[2]))
	x4 := ((x1 & 0xfffffffffff) + ((x3 >> 43) * 0x5))
	x5 := (uint64(uint1((x4 >> 44))) + (x2 & 0x7ffffffffff))
	x6 := (x4 & 0xfffffffffff)
	x7 := (x5 & 0x7ffffffffff)
	x8 := (uint64(uint1((x5 >> 43))) + (x3 & 0x7ffffffffff))
	out1[0] = x6
	out1[1] = x7
	out1[2] = x8
}

// CarryOpp negates a field element.
//
// Postconditions:
//   eval out1 mod m = -eval arg1 mod m
//
func CarryOpp(out1 *TightFieldElement, arg1 *TightFieldElement) {
	x1 := (0x1ffffffffff6 - arg1[0])
	x2 := (uint64(uint1((x1 >> 44))) + (0xffffffffffe - arg1[1]))
	x3 := (uint64(uint1((x2 >> 43))) + (0xffffffffffe - arg1[2]))
	x4 := ((x1 & 0xfffffffffff) + (uint64(uint1((x3 >> 43))) * 0x5))
	x5 := (uint64(uint1((x4 >> 44))) + (x2 & 0x7ffffffffff))
	x6 := (x4 & 0xfffffffffff)
	x7 := (x5 & 0x7ffffffffff)
	x8 := (uint64(uint1((x5 >> 43))) + (x3 & 0x7ffffffffff))
	out1[0] = x6
	out1[1] = x7
	out1[2] = x8
}
