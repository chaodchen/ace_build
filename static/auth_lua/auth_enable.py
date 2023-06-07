#coding:utf-8
import os
import hashlib
import sys

g_this_file = os.path.realpath(__file__)
g_this_folder = os.path.dirname(g_this_file)

g_log_id_ls = [
    # (0, "a9d7e95b000007b0000ce4b3"),
    # (0, "323c3b1b0000080160099ea7"),
    # (0, "ece66ee300000801604b2934"),
    # (0, "427d18840000b31a556fc031"),
    # (0, "650ae86500010304000e5319")
    (0, "650ae8650000fe004955ddcb"),
]

def enc_str(s):
    g_enc_table = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    str_len = len(s)
    i = 0
    ret_s = ""
    while i < str_len:
        ch_1 = s[i]
        ch_2 = g_enc_table[i % 36]
        num_1 = ord(ch_1)
        num_2 = ord(ch_2)
        num_1 ^= num_2
        ch_1 = chr(num_1)
        ret_s = ret_s + ch_1
        i += 1
    return ret_s

class lib:
    def __init__(self):
        self.m_buf = ""
    def write_u8(self, num):
        num = num & 0xff
        self.m_buf += chr(num)
    def write_u16(self, num):
        n1 = (num >> 8) & 0xff
        n2 = num & 0xff
        self.write_u8(n1)
        self.write_u8(n2)
    def write_u32(self, num):
        n1 = (num >> 16) & 0xffff
        n2 = num & 0xffff
        self.write_u16(n1)
        self.write_u16(n2)
    def write_str(self, str):
        self.write_u32(len(str))
        self.m_buf += str

class datastream:
    def __init__(self, buf):
        self.m_buf = buf
        self.m_offset = 0
    def read_u8(self):
        num = ord(self.m_buf[self.m_offset])
        self.m_offset += 1
        return num
    def read_u16(self):
        return (self.read_u8() << 8) | (self.read_u8())
    def read_u32(self):
        return (self.read_u16() << 16) | (self.read_u16())
    def read_u64(self):
        return (self.read_u32() << 32) | (self.read_u32())
    def read_str(self):
        size = self.read_u32()
        if size == 0:
            return ""
        str = self.m_buf[self.m_offset : self.m_offset + size]
        self.m_offset += size
        return str
    def read_enc_str(self):
        return enc_str(self.read_str())

def read_enable(path):
    f = open(path, "rb")
    data = f.read()
    f.close()
    ds = datastream(data)
    #
    magic_code = ds.read_u32()
    assert(magic_code == 0x20210928)
    #
    i = 0
    while i < 64:
        ds.read_u32()
        i += 1
    #
    cnt = ds.read_u16()
    i = 0
    while i < cnt:
        i += 1
        game_id = ds.read_u32()
        dev_id = ds.read_enc_str()
        print("%d:%s" % (game_id, dev_id))
    magic_code = ds.read_u32()
    assert(magic_code == 0x20210928)

def cal_md5(buf):
    m = hashlib.md5()
    m.update(buf)
    return m.hexdigest()

def run(ls):
    ds = lib()
    ds.write_u32(0x20210928)
    i = 0
    while i < 64:
        ds.write_u32(0)          #signature
        i += 1
    ds.write_u16(len(ls))
    for node in ls:
        game_id = node[0]
        dev_id = node[1]
        ds.write_u32(game_id)
        ds.write_str(enc_str(dev_id))
    ds.write_u32(0x20210928)

    #计算md5
    md5 = cal_md5(ds.m_buf)
    path = os.path.join(g_this_folder, "sig.dat")
    if os.path.exists(path): os.remove(path)
    print(md5)
    os.system("%s %s" % (os.path.join(g_this_folder, "sigtool3", "sigtool3"), md5))
    print(path)

    #
    f = open(path, "rb")
    sig = f.read()
    f.close()
    assert(len(sig) == 256)

    #
    f = open("enable.dat", "wb")
    f.write(ds.m_buf[:4])
    f.write(sig)
    f.write(ds.m_buf[4 + len(sig):])
    f.close()

    #
    read_enable("enable.dat")

    #
    # os.system("adb.exe push enable.dat /sdcard/sdk/")

def init_run():
    f = open('devices.txt')
    ret = 0
    devices = f.read()
    if len(devices) < 24:
        return -1
    else:
        devices = devices.split('\n')
        for devid in devices:
            if len(devid) < 24:
                continue
            if (0, devid) not in g_log_id_ls:
                g_log_id_ls.append((0, devid))
    
    f.close()

    # 获取命令行devid
    for index, arg in enumerate(sys.argv):
        if index == 0:
            continue
        if len(arg) != 24:
            continue
        if arg in g_log_id_ls:
            continue
        g_log_id_ls.append((0, arg))
    return ret


if __name__ == "__main__":
    if init_run() == 0:
        run(g_log_id_ls)
