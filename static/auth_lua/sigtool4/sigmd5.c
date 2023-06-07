#include <stdio.h>
#include <string.h>
#include <stdlib.h>


#include "rsa.h"
#include "private.h"
#include "public.h"


void usage()
{
    printf("usage: sigmd5 md5 [work directory]\n");
}

unsigned char Hex2Ch(const char *hex)
{
    unsigned char ch = 0;

    unsigned char n1 = 0;
    unsigned char n2 = 0;

    if (hex[0] >= '0' && hex[0] <= '9') n1 = hex[0] - '0';
    else if (hex[0] >= 'A' && hex[0] <= 'F') n1 = hex[0] - 'A' + 10;
    else if (hex[0] >= 'a' && hex[0] <= 'f') n1 = hex[0] - 'a' + 10;

    if (hex[1] >= '0' && hex[1] <= '9') n2 = hex[1] - '0';
    else if (hex[1] >= 'A' && hex[1] <= 'F') n2 = hex[1] - 'A' + 10;
    else if (hex[1] >= 'a' && hex[1] <= 'f') n2 = hex[1] - 'a' + 10;


    return n1 * 16 + n2;
}

void HexBinary(const unsigned char *binary, size_t binary_len, char *hex)
{
    size_t i;
    for (i = 0; i < binary_len; i++)
    {
        snprintf(hex + i * 2, 3, "%02X", binary[i]);
    }
}

void Hex2Binary(const char *md5, unsigned char *binary)
{
    size_t i = 0;
    char tmp[33] = {0};
    strcpy(tmp, md5);

    for (i = 0; i < 16; i++)
    {
        char *p = tmp + i * 2;
        char bk = *(p + 2);
        *(p + 2) = 0;

        binary[i] = Hex2Ch(p);

        *(p + 2) = bk;
    }
}

void PrintHex(const unsigned char *binary, size_t len)
{
    char *hex = (char *)malloc(len * 2 + 1);
    for (size_t i = 0; i < len; i++)
    {
        sprintf(hex + 2 * i, "%02X", binary[i]);
    }
    printf("HEX:%s\n", hex);
    free(hex);
}


int main(int argc, char **argv)
{
    if (argc < 2 || argc > 4)
    {
        usage();
        return -1;
    }
    const char *md5 = argv[1];
    const char *work_dir = NULL;
    if (argc == 3)
    {
        work_dir = argv[2];
    }
    char sig_path[128] = {0};
    unsigned char md5_binary[16] = {0};
    char enc_md5[33] = {0};
    unsigned char raw_data[TP_RSANUMBYTES] = {0};
    unsigned char sig_data[TP_RSANUMBYTES] = {0};
    unsigned char dec_data[TP_RSANUMBYTES] = {0};


    if (strlen(md5) != 32)
    {
        printf("strlen(md5) should be 32, now:%zd\n", strlen(md5));
        return -1;
    }
    Hex2Binary(md5, md5_binary);
    memcpy(raw_data + 1, md5_binary, 16);


    int r1 = RSA2048_Encrypt(&privateKeys[1], raw_data, TP_RSANUMBYTES, sig_data);
    int r2 = RSA2048_Decrypt(&publicKeys[1], sig_data, TP_RSANUMBYTES, dec_data);

    if (r1 == 1 && r2 == 1 && memcmp(raw_data, dec_data, TP_RSANUMBYTES) == 0)
    {
        printf("done.\n");
        if (work_dir != NULL){
            strncpy(sig_path, work_dir, sizeof(sig_path));
            strcat(sig_path, "/sig.dat");         
        } else {
            strncpy(sig_path, "sig.dat", sizeof(sig_path));
        }
        FILE *f = fopen(sig_path, "wb");
        fwrite(sig_data, 1, TP_RSANUMBYTES, f);
        fclose(f);

        return 0;
    }
    else
    {
        printf("r1:%d, r2:%d, r3:%d\n", r1, r2, memcmp(raw_data, dec_data, TP_RSANUMBYTES));
        PrintHex(raw_data, TP_RSANUMBYTES);
        PrintHex(dec_data, TP_RSANUMBYTES);
    }

    return -1;
}

// f24e92fd326738dc052829545a200058 ok 
// f24e92fd326738dc052829545a200059 wrong