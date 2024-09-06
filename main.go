package main

import (
	"bytes"
	"encoding/binary"
	"math"
	"runtime/debug"
	"strconv"

	"errors"
	"fmt"
	"os"

	// "github.com/deadsy/rvda"
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

func LoadFile(mu uc.Unicorn, fn string, base uint64) uint64 {
	dat, err := os.ReadFile(fn)
	if err != nil {
		fmt.Println("Not loaded program", err)
	}
	mu.MemWrite(base, dat)
	return uint64(len(dat))
}

func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int64
	binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return x
}

func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return x
}
func Int64ToBytes(n uint64) []byte {
	x := uint64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})

	binary.Write(bytesBuffer, binary.LittleEndian, x)

	return bytesBuffer.Bytes()
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)

	return bytesBuffer.Bytes()
}

func print_X_registers(mu uc.Unicorn) {
	x1, _ := mu.RegRead(uc.RISCV_REG_X1)
	x2, _ := mu.RegRead(uc.RISCV_REG_X2)
	x3, _ := mu.RegRead(uc.RISCV_REG_X3)
	x4, _ := mu.RegRead(uc.RISCV_REG_X4)
	x5, _ := mu.RegRead(uc.RISCV_REG_X5)
	x6, _ := mu.RegRead(uc.RISCV_REG_X6)
	x7, _ := mu.RegRead(uc.RISCV_REG_X7)
	x8, _ := mu.RegRead(uc.RISCV_REG_X8)
	x9, _ := mu.RegRead(uc.RISCV_REG_X9)

	fmt.Printf("-----------Registers X------------\n")
	fmt.Printf("reg x1: 0x%x\n", x1)
	fmt.Printf("reg x2: 0x%x\n", x2)
	fmt.Printf("reg x3: 0x%x\n", x3)
	fmt.Printf("reg x4: 0x%x\n", x4)
	fmt.Printf("reg x5: 0x%x\n", x5)
	fmt.Printf("reg x6: 0x%x\n", x6)
	fmt.Printf("reg x7: 0x%x\n", x7)
	fmt.Printf("reg x8: 0x%x\n", x8)
	fmt.Printf("reg x9: 0x%x\n", x9)
	fmt.Printf("-----------Registers X------------\n")

}

func print_A_registers(mu uc.Unicorn) {
	reg_a0_raw, _ := mu.RegRead(uc.RISCV_REG_A0)
	reg_a1_raw, _ := mu.RegRead(uc.RISCV_REG_A1)
	reg_a2_raw, _ := mu.RegRead(uc.RISCV_REG_A2)
	reg_a3_raw, _ := mu.RegRead(uc.RISCV_REG_A3)
	reg_a4_raw, _ := mu.RegRead(uc.RISCV_REG_A4)
	reg_a5_raw, _ := mu.RegRead(uc.RISCV_REG_A5)
	reg_a6_raw, _ := mu.RegRead(uc.RISCV_REG_A6)
	reg_a7_raw, _ := mu.RegRead(uc.RISCV_REG_A7)

	fmt.Printf("-----------Registers A------------\n")
	fmt.Printf("reg A0: 0x%x \n", reg_a0_raw)
	fmt.Printf("reg A1: 0x%x \n", reg_a1_raw)
	fmt.Printf("reg A2: 0x%x \n", reg_a2_raw)
	fmt.Printf("reg A3: 0x%x \n", reg_a3_raw)
	fmt.Printf("reg A4: 0x%x \n", reg_a4_raw)
	fmt.Printf("reg A5: 0x%x \n", reg_a5_raw)
	fmt.Printf("reg A6: 0x%x \n", reg_a6_raw)
	fmt.Printf("reg A7: 0x%x \n", reg_a7_raw)
	fmt.Printf("-----------Registers A------------\n")
}

func print_S_registers(mu uc.Unicorn) {
	sp_raw_value, _ := mu.RegRead(uc.RISCV_REG_SP)
	sp_mem_value, _ := mu.MemRead(uint64(sp_raw_value), 4)

	ra_raw_value, _ := mu.RegRead(uc.RISCV_REG_RA)
	s0_raw_value, _ := mu.RegRead(uc.RISCV_REG_S0)
	s1_raw_value, _ := mu.RegRead(uc.RISCV_REG_S1)
	s2_raw_value, _ := mu.RegRead(uc.RISCV_REG_S2)
	s3_raw_value, _ := mu.RegRead(uc.RISCV_REG_S3)
	s4_raw_value, _ := mu.RegRead(uc.RISCV_REG_S4)
	s5_raw_value, _ := mu.RegRead(uc.RISCV_REG_S5)
	s6_raw_value, _ := mu.RegRead(uc.RISCV_REG_S6)
	s7_raw_value, _ := mu.RegRead(uc.RISCV_REG_S7)
	s8_raw_value, _ := mu.RegRead(uc.RISCV_REG_S8)
	s9_raw_value, _ := mu.RegRead(uc.RISCV_REG_S9)
	s10_raw_value, _ := mu.RegRead(uc.RISCV_REG_S10)
	s11_raw_value, _ := mu.RegRead(uc.RISCV_REG_S11)

	fmt.Printf("-----------Registers S------------\n")
	fmt.Printf("reg SP: 0x%x (%d)\n", sp_raw_value, sp_raw_value)
	fmt.Printf("reg SP mem: 0x%x\n", sp_mem_value)
	fmt.Printf("reg ra: 0x%x\n", ra_raw_value)
	fmt.Printf("reg s0: 0x%x\n", s0_raw_value)
	fmt.Printf("reg s1: 0x%x\n", s1_raw_value)
	fmt.Printf("reg s2: 0x%x\n", s2_raw_value)
	fmt.Printf("reg s3: 0x%x\n", s3_raw_value)
	fmt.Printf("reg s4: 0x%x\n", s4_raw_value)
	fmt.Printf("reg s5: 0x%x\n", s5_raw_value)
	fmt.Printf("reg s6: 0x%x\n", s6_raw_value)
	fmt.Printf("reg s7: 0x%x\n", s7_raw_value)
	fmt.Printf("reg s8: 0x%x\n", s8_raw_value)
	fmt.Printf("reg s9: 0x%x\n", s9_raw_value)
	fmt.Printf("reg s10: 0x%x\n", s10_raw_value)
	fmt.Printf("reg s11: 0x%x\n", s11_raw_value)
	fmt.Printf("-----------Registers S------------\n")

}

func pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, errors.New("invalid blocksize")
	}
	if b == nil || len(b) == 0 {
		return nil, errors.New("invalid blocksize")
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(0)}, n))
	return pb, nil
}

func run_unicorn(num int) {

	PROGRAM_PATH := "./src_cpp/memory-riscv64.bin"
	// PROGRAM_PATH := "./src_cpp/primes-riscv64.bin"

	MEMORY_SLOT := uint64(32 * 1024 * 1024 * 1024)
	ADDRESS_START := uint64(0x8000000)

	INPUT_ADDRESS := uint64(0x8040000)
	OUTPUT_ADDRESS := uint64(0x8048000)

	mu, _ := uc.NewUnicorn(uc.ARCH_RISCV, uc.MODE_RISCV64)

	end_bytes, _ := mu.MemRead(MEMORY_SLOT-4, 4)
	end_value := BytesToInt32(end_bytes)
	fmt.Println("End value: ", end_value)

	// isa, _ := rvda.New(32, rvda.RV64gc)
	// mu.HookAdd(uc.HOOK_CODE, func(mu uc.Unicorn, addr uint64, size uint32) {
	// 	// fmt.Printf("code: addr 0x%x size %d\n", addr, size)
	//
	// 	ins_bytes, _ := mu.MemRead(addr, uint64(size))
	// 	normalized_ins_bytes, _ := pad(ins_bytes, 4)
	// 	normalized_ins := BytesToInt32(normalized_ins_bytes)
	// 	// fmt.Printf("ins: 0x%v\n", hex.EncodeToString(normalized_ins_bytes))
	// 	normalized_da := isa.Disassemble(uint(addr), uint(normalized_ins))
	// 	fmt.Printf("normalized decode: %#v\n", normalized_da)
	//
	// 	// reg_a4_raw, _ := mu.RegRead(uc.RISCV_REG_A4)
	// 	// reg_a5_raw, _ := mu.RegRead(uc.RISCV_REG_A5)
	// 	//
	// 	// fmt.Println("reg A4: ", reg_a4_raw)
	// 	// fmt.Println("reg A5: ", reg_a5_raw)
	// }, ADDRESS_START, 0)

	mu.HookAdd(uc.HOOK_INTR, func(mu uc.Unicorn, intno uint32) {
		// fmt.Println("intr no:", intno)
	}, ADDRESS_START, 0)

	// mu.HookAdd(uc.HOOK_MEM_WRITE, func(mu uc.Unicorn, access int, addr uint64, size int, value int64) {
	// 	fmt.Printf("mem write: @0x%x, 0x%x = 0x%x\n", addr, size, value)
	// }, ADDRESS_START, 0)

	fmt.Println(mu)

	mu.MemMap(0, MEMORY_SLOT)

	mu.MemWrite(INPUT_ADDRESS, IntToBytes(num))
	mu.MemWrite(OUTPUT_ADDRESS, Int64ToBytes(0))

	LoadFile(mu, PROGRAM_PATH, ADDRESS_START)

	if err := mu.Start(ADDRESS_START, 0); err != nil {
		print_S_registers(mu)
		print_A_registers(mu)
		print_X_registers(mu)

		input_value_bytes, _ := mu.MemRead(INPUT_ADDRESS, 4)
		input_value := BytesToInt32(input_value_bytes)
		fmt.Println("Input value: ", input_value)

		output_value_bytes, _ := mu.MemRead(OUTPUT_ADDRESS, 8)
		output_value := BytesToInt64(output_value_bytes)
		fmt.Println("Output value: ", output_value)

		panic(err)
	}

	input_value_bytes, _ := mu.MemRead(INPUT_ADDRESS, 4)
	input_value := BytesToInt32(input_value_bytes)
	fmt.Println("Input value: ", input_value)

	output_value_bytes, _ := mu.MemRead(OUTPUT_ADDRESS, 8)
	output_value := BytesToInt64(output_value_bytes)
	fmt.Println("Output value: ", output_value)

	output_value_1_bytes, _ := mu.MemRead(OUTPUT_ADDRESS+8, 8)
	output_value_1 := BytesToInt64(output_value_1_bytes)
	fmt.Println("Output value (8): ", output_value_1)
	output_value_2_bytes, _ := mu.MemRead(OUTPUT_ADDRESS+16, 8)
	output_value_2 := BytesToInt64(output_value_2_bytes)
	fmt.Println("Output value (16): ", output_value_2)
}

func main() {
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(math.MaxInt64)

	num, _ := strconv.Atoi(os.Args[1])
	run_unicorn(num)
}
