#!/usr/bin/python
# -*- coding:utf-8 -*-
import os,platform,sys
root = os.path.split(os.path.realpath(__file__))[0]
print(root + ";" + "%s/deps" %root)
os.system('cmd /C "set GOPATH=%s;%s&& cd %s && go build"' %(root, root + "/deps", root))
