host.primes:
	g++ -Wall -g primes.cpp -o primes

host.memory:
	g++ -Wall -g memory.cpp -o memory

emu.primes:
	riscv64-unknown-elf-g++ -g -O0 -Wl,--gc-sections \
	-mcmodel=medany \
    -nostartfiles -Wl,-T,../link/riscv64.ld \
    ../link/crt0.s primes_riscv.cpp -o primes_riscv
	riscv64-unknown-elf-objcopy -O binary primes_riscv primes-riscv64.bin

emu.memory:
	riscv64-unknown-elf-g++  -g -O0 --std=c++11 -Wl,--gc-sections \
	-mcmodel=medany \
    -nostartfiles -Wl,-T,../link/riscv64.ld \
    ../link/crt0.s /opt/riscv/lib/gcc/riscv64-unknown-elf/13.2.0/crti.o /opt/riscv/lib/gcc/riscv64-unknown-elf/13.2.0/crtbegin.o memory-riscv.cpp /opt/riscv/lib/gcc/riscv64-unknown-elf/13.2.0/crtend.o /opt/riscv/lib/gcc/riscv64-unknown-elf/13.2.0/crtn.o \
	-static-libstdc++ \
	-o memory_riscv
	riscv64-unknown-elf-objcopy -O binary memory_riscv memory-riscv64.bin

