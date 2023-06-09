#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>

#include "rsa.h"


// a[] -= mod
static void subM2(RSAPublicKey key,
                 uint32_t* a) {
  int64_t A = 0;
  int i;
  for (i = 0; i < key->len; ++i) {
    A += (uint64_t)a[i] - key->n[i];
    a[i] = (uint32_t)A;
    A >>= 32;
  }
}

// return a[] >= mod
static int geM2(RSAPublicKey key,
               const uint32_t* a) {
  int i;
  for (i = key->len; i;) {
    --i;
    if (a[i] < key->n[i]) return 0;
    if (a[i] > key->n[i]) return 1;
  }
  return 1;  // equal
}

// montgomery c[] += a * b[] / R % mod
static void montMulAdd2(RSAPublicKey key,
                       uint32_t* c,
                       const uint32_t a,
                       const uint32_t* b) {
  uint64_t A = (uint64_t)a * b[0] + c[0];
  uint32_t d0 = (uint32_t)A * key->n0inv;
  uint64_t B = (uint64_t)d0 * key->n[0] + (uint32_t)A;
  int i;

  for (i = 1; i < key->len; ++i) {
    A = (A >> 32) + (uint64_t)a * b[i] + c[i];
    B = (B >> 32) + (uint64_t)d0 * key->n[i] + (uint32_t)A;
    c[i - 1] = (uint32_t)B;
  }

  A = (A >> 32) + (B >> 32);

  c[i - 1] = (uint32_t)A;

  if (A >> 32) {
    subM2(key, c);
  }
}

// montgomery c[] = a[] * b[] / R % mod
static void montMul2(RSAPublicKey key,
                    uint32_t* c,
                    const uint32_t* a,
                    const uint32_t* b) {
  int i;
  for (i = 0; i < key->len; ++i) {
    c[i] = 0;
  }
  for (i = 0; i < key->len; ++i) {
    montMulAdd2(key, c, a[i], b);
  }
}

// In-place public exponentiation.
// Input and output big-endian byte array in inout.
static void modpowF42(RSAPublicKey key,
                 uint8_t* inout) {
  uint32_t a[RSANUMWORDS];
  uint32_t aR[RSANUMWORDS];
  uint32_t aaR[RSANUMWORDS];
  uint32_t* aaa = aaR;  // Re-use location.
  int i;

  // Convert from big endian byte array to little endian word array.
  for (i = 0; i < key->len; ++i) {
    uint32_t tmp =
      (inout[((key->len - 1 - i) * 4) + 0] << 24) |
      (inout[((key->len - 1 - i) * 4) + 1] << 16) |
      (inout[((key->len - 1 - i) * 4) + 2] << 8) |
      (inout[((key->len - 1 - i) * 4) + 3] << 0);
    a[i] = tmp;
  }

  montMul2(key, aR, a, key->rr);  // aR = a * RR / R mod M
  for (i = 0; i < 16; i += 2) {
    montMul2(key, aaR, aR, aR);  // aaR = aR * aR / R mod M
    montMul2(key, aR, aaR, aaR);  // aR = aaR * aaR / R mod M
  }
  montMul2(key, aaa, aR, a);  // aaa = aR * a / R mod M

  // Make sure aaa < mod; aaa is at most 1x mod too large.
  if (geM2(key, aaa)) {
    subM2(key, aaa);
  }

  // Convert to bigendian byte array
  for (i = key->len - 1; i >= 0; --i) {
    uint32_t tmp = aaa[i];
    *inout++ = tmp >> 24;
    *inout++ = tmp >> 16;
    *inout++ = tmp >> 8;
    *inout++ = tmp >> 0;
  }
}

// Decode a 2048 bit RSA PKCS1.5 DATA.
// Returns 0 on failure, 1 on success.
int RSA2048_Decrypt(RSAPublicKey key,
               const uint8_t* pEncryptedData,
               const int len,
               uint8_t* pOut)
{
  uint8_t buf[RSANUMBYTES];
  int i;

  if (key->len != RSANUMWORDS) {
    return 0;  // Wrong key passed in.
  }

  if (len != sizeof(buf)) {
    return 0;  // Wrong input length.
  }

  if (pOut == 0)
  {
	  return 0;
  }

  for (i = 0; i < len; ++i) {  // Copy input to local workspace.
    buf[i] = pEncryptedData[i];
  }


  modpowF42(key, buf);  // In-place exponentiation.

  for (i = 0; i < len; ++i) {  // Copy input to local workspace.
    pOut[i] = buf[i];
  }
  return 1;  // All checked out OK.
}