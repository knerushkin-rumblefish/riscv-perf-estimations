#include <cstdint>
#include <cstdlib>
#include <stdint.h>

int primes(int32_t upto_prime_num) {

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

  return 0;
}

int main(int argc, char *argv[]) {
  int upto_prime_num = atoi(argv[1]);
  primes(upto_prime_num);

  return 0;
}
