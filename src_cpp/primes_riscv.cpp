#include <stdint.h>

#define INPUT_ADDRESS 0x8040000
#define OUTPUT_ADDRESS 0x8048000

int main() {
  int32_t *input = (int32_t *)INPUT_ADDRESS;

  int64_t num, i;

  int32_t prime_count = 1;
  int32_t upto_prime_num = *input;

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
  return 0;
}
