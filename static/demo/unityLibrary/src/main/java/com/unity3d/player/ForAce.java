package com.unity3d.player;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class ForAce implements CreateSdk{

    @Override
    public void init(int gameid, String gamekey) {
//        TP2Sdk.initEx(gameid,gamekey);
        Class<?> cl = null;
        try {
            cl = Class.forName("com.tencent.tersafe2.TP2Sdk");
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        Method initex = null;
        try {
            initex = cl.getMethod("initEx", int.class, String.class);
        } catch (NoSuchMethodException e) {
            e.printStackTrace();
        }
        try {
            initex.invoke(cl.newInstance(), gameid, gamekey);
        } catch (IllegalAccessException e) {
            e.printStackTrace();
        } catch (InvocationTargetException e) {
            e.printStackTrace();
        } catch (InstantiationException e) {
            e.printStackTrace();
        }
    }


    @Override
    public void login(String openid) {
//        TP2Sdk.onUserLogin(1,88,openid,"test");
        Class<?> cl = null;
        try {
            cl = Class.forName("com.tencent.tersafe2.TP2Sdk");
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        Method initex = null;
        try {
            initex = cl.getMethod("onUserLogin", int.class, int.class, String.class, String.class);
        } catch (NoSuchMethodException e) {
            e.printStackTrace();
        }
        try {
            initex.invoke(cl.newInstance(), 1,88, openid, "test");
        } catch (IllegalAccessException e) {
            e.printStackTrace();
        } catch (InvocationTargetException e) {
            e.printStackTrace();
        } catch (InstantiationException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void ioctl(String str) {
//        TP2Sdk.ioctl(str);
        Class<?> cl = null;
        try {
            cl = Class.forName("com.tencent.tersafe2.TP2Sdk");
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        Method initex = null;
        try {
            initex = cl.getMethod("ioctl", String.class);
        } catch (NoSuchMethodException e) {
            e.printStackTrace();
        }
        try {
            initex.invoke(cl.newInstance(), str);
        } catch (IllegalAccessException e) {
            e.printStackTrace();
        } catch (InvocationTargetException e) {
            e.printStackTrace();
        } catch (InstantiationException e) {
            e.printStackTrace();
        }
    }
}
