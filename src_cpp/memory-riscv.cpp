#include <cstddef>
#include <cstdint>
#include <vector>

#define INPUT_ADDRESS 0x8040000
#define OUTPUT_ADDRESS 0x8048000

int32_t primes(int32_t upto_prime_num) {

  int64_t num, i;

  int32_t prime_count = 1;

  for (num = 2; prime_count <= upto_prime_num; num++) {

    for (i = 2; i <= (num / 2); i++) {

      if (num % i == 0) {
        i = num;
        break;
      }
    }

    if (i != num) {
      prime_count++;
    }
  }

  return prime_count;
}

int _fstat(int fd, struct stat *st) {
  (void)fd, (void)st;
  return -1;
}

void *_sbrk(int incr) {
  extern char _end;
  static unsigned char *heap = NULL;
  unsigned char *prev_heap;
  if (heap == NULL)
    heap = (unsigned char *)&_end;
  prev_heap = heap;
  heap += incr;
  return prev_heap;
}

int _close(int fd) {
  (void)fd;
  return -1;
}

int _isatty(int fd) {
  (void)fd;
  return 1;
}

int _read(int fd, char *ptr, int len) {
  (void)fd, (void)ptr, (void)len;
  return -1;
}

int _lseek(int fd, int ptr, int dir) {
  (void)fd, (void)ptr, (void)dir;
  return 0;
}

int main() {
  int32_t *input = (int32_t *)INPUT_ADDRESS;

  long int size = 1l * 32 * 1024 * 1024;
  std::vector<long int> large_vector;

  large_vector.push_back(*input);

  long int i;
  for (i = 0; i < size; i += 4) {
    large_vector.push_back(i);
  }

  int64_t *output = (int64_t *)OUTPUT_ADDRESS;
  *output = (int64_t)large_vector.size();

  int64_t *output_1 = (int64_t *)(OUTPUT_ADDRESS + 8);
  *output_1 = (int64_t)large_vector[0];

  int64_t *output_2 = (int64_t *)(OUTPUT_ADDRESS + 16);
  *output_2 = (int64_t)large_vector[large_vector.size() - 1];
}
