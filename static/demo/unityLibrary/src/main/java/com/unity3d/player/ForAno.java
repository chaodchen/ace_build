package com.unity3d.player;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class ForAno implements CreateSdk{
    @Override
    public void init(int gameid, String gamekey) {
        Class<?> cl = null;
        try {
            cl = Class.forName("com.gamesafe.ano.AnoSdk");
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
//        AnoSdk.setUserInfo(1, openid);
        Class<?> cl = null;
        try {
            cl = Class.forName("com.gamesafe.ano.AnoSdk");
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }
        Method initex = null;
        try {
            initex = cl.getMethod("setUserInfo", int.class, String.class);
        } catch (NoSuchMethodException e) {
            e.printStackTrace();
        }
        try {
            initex.invoke(cl.newInstance(), 1, openid);
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
//        AnoSdk.ioctl(str);
        Class<?> cl = null;
        try {
            cl = Class.forName("com.gamesafe.ano.AnoSdk");
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
