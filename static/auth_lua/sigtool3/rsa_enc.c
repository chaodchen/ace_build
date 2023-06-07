#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>

#include "rsa.h"

// a[] -= mod
static void subM1(RSAPrivateKey key,
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
static int geM1(RSAPrivateKey key,
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
static void montMulAdd1(RSAPrivateKey key,
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
        subM1(key, c);
    }
}

// montgomery c[] = a[] * b[] / R % mod
static void montMul1(RSAPrivateKey key,
                    uint32_t* c,
                    const uint32_t* a,
                    const uint32_t* b) {
    int i;
    for (i = 0; i < key->len; ++i) {
        c[i] = 0;
    }
    for (i = 0; i < key->len; ++i) {
        montMulAdd1(key, c, a[i], b);
    }
}

// In-place public exponentiation.
// Input and output big-endian byte array in inout.
static void modpowF41(RSAPrivateKey key,
                 uint8_t* inout) {
    uint32_t a[TP_RSANUMWORDS];
    uint32_t aR[TP_RSANUMWORDS];
    uint32_t aaR[TP_RSANUMWORDS];
    uint32_t t[TP_RSANUMWORDS];
    uint32_t* aaa = aaR;  // Re-use location.
    int i = 0, j = 0;

    // Convert from big endian byte array to little endian word array.
    for (i = 0; i < key->len; ++i) {
        uint32_t tmp =
            (inout[((key->len - 1 - i) * 4) + 0] << 24) |
            (inout[((key->len - 1 - i) * 4) + 1] << 16) |
            (inout[((key->len - 1 - i) * 4) + 2] << 8) |
            (inout[((key->len - 1 - i) * 4) + 3] << 0);
            a[i] = tmp;
    }

    montMul1(key, aR, a, key->rr);
    montMul1(key, t, a, key->rr);
    int d_dword_len = key->d_bit_len / 32, d_bit_rem = key->d_bit_len % 32;
    for(i = d_bit_rem - 2; i >= 0; i--) {
        montMul1(key, aaR, t, t); //aaR = aR * aR /R mod M
        memcpy(t, aaR, TP_RSANUMBYTES);

        if (key->d[d_dword_len] & (1 << i)) {
            montMul1(key, aaR, aR, t);  // aaR = aR * aR / R mod M
            memcpy(t, aaR, TP_RSANUMBYTES);
        }
    }
    for(i = d_dword_len - 1; i >= 0; i--) {
        for(j = 31; j >= 0; --j) {
            montMul1(key, aaR, t, t); //aaR = aR * aR /R mod M
            memcpy(t, aaR, TP_RSANUMBYTES);

            if (key->d[i] & (1 << j)) {
                montMul1(key, aaR, aR, t);  // aR = aaR * aaR / R mod M
                memcpy(t, aaR, TP_RSANUMBYTES);
            }
        }
    }
    uint32_t one[TP_RSANUMWORDS] = {0};
    one[0] = 1;
    montMul1(key, aaa, t, one);

    // Make sure aaa < mod; aaa is at most 1x mod too large.
    if (geM1(key, aaa)) {
        subM1(key, aaa);
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

int RSA2048_Encrypt(RSAPrivateKey key,
               const uint8_t* pIn,
               const int len,
               uint8_t* pOut)
{
    uint8_t buf[TP_RSANUMBYTES];
    int i;

    if (key->len != TP_RSANUMWORDS) {
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
        buf[i] = pIn[i];
    }


    modpowF41(key, buf);  // In-place exponentiation.

    for (i = 0; i < len; ++i) {  // Copy input to local workspace.
        pOut[i] = buf[i];
    }
    return 1;  // All checked out OK.
}