package com.unity3d.player;

import java.lang.reflect.InvocationTargetException;

public interface CreateSdk {
    public void init(int gameid, String gamekey);
    public void login(String openid);
    public void ioctl(String str);
}
