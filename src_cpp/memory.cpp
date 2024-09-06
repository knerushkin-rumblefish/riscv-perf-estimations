#include <cstdint>
#include <cstdio>
#include <cstdlib>
#include <malloc.h>
#include <vector>

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

int main(int argc, char *argv[]) {
  long int size = 10L * 1024 * 1024 * 1024;
  std::vector<int32_t> large_vector;

  long int i;
  for (i = 0; i < size; i += 4) {
    large_vector.push_back(0);
  }
}
