#ifndef __RAS_H__
#define __RAS_H__

#include <stdint.h>

#define TP_RSANUMBYTES 256  // 2048 bit key length
#define TP_RSANUMWORDS (TP_RSANUMBYTES / sizeof(uint32_t))

typedef struct RSAPrivateKeyInstance {
    int len;  // Length of n[] in number of uint32_t
    uint32_t n0inv;  // -1 / n[0] mod 2^32
    uint32_t n[TP_RSANUMWORDS];  // modulus as little endian array
    uint32_t rr[TP_RSANUMWORDS];  // R^2 as little endian array
    uint32_t d[TP_RSANUMWORDS];
    uint32_t d_bit_len;
} RSAPrivateKeyInstance;

typedef const RSAPrivateKeyInstance * const RSAPrivateKey;


#define RSANUMBYTES 256  // 2048 bit key length
#define RSANUMWORDS (RSANUMBYTES / sizeof(uint32_t))

typedef struct RSAPublicKeyInstance {
  int len;  // Length of n[] in number of uint32_t
  uint32_t n0inv;  // -1 / n[0] mod 2^32
  uint32_t n[RSANUMWORDS];  // modulus as little endian array
  uint32_t rr[RSANUMWORDS];  // R^2 as little endian array
} RSAPublicKeyInstance;

typedef const RSAPublicKeyInstance * const RSAPublicKey;


int RSA2048_Encrypt(RSAPrivateKey key,
               const uint8_t* pIn,
               const int len,
               uint8_t* pOut);

int RSA2048_Decrypt(RSAPublicKey key,
               const uint8_t* pEncryptedData,
               const int len,
               uint8_t* pOut);


#endif