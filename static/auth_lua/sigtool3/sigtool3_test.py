#coding:utf-8
# 由于出现了某些md5加解密前后结果不一致的问题，这里对sigtool3的结果进行二分测试
import os

g_this_file = os.path.realpath(__file__)
g_this_folder = os.path.dirname(g_this_file)

g_min_md5 = "00000000000000000000000000000000"
g_max_md5 = "ffffffffffffffffffffffffffffffff"

def mid_md5(a, b):
    int_a = int(a, 16)
    int_b = int(b, 16)
    int_c = (int_a + int_b) / 2
    return hex(int_c)[2:-1]

def main():
    a = g_min_md5
    b = g_max_md5
    c = mid_md5(a, b)
    while a != c:
        ret1 = os.system("%s/sigtool3 %s" % (g_this_folder, a))
        ret2 = os.system("%s/sigtool3 %s" % (g_this_folder, b))
        ret3 = os.system("%s/sigtool3 %s" % (g_this_folder, c))
        if ret3 == ret1 and ret3 != ret2:
            a, b, c = c, b, mid_md5(c, b)
        if ret3 != ret1 and ret3 == ret2:
            a, b, c = a, c, mid_md5(a, c)
    print a

if __name__ == "__main__":
    main()