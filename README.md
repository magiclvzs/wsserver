# wsserver  
有疑问可联系qq：441707528   
## 什么是wsserver
	wsserver是一个简单的实例性框架，实现了基于websocket+protobuf的聊天系统，目的是便于用户理解和使用antnet
## 编译
	本步可以跳过，直接解压wsserver.rar运行wsserver.exe也行
	本示例是基于windows的，首先安装好go python
	然后到wsserver目录运行python build.py 便可生产可执行文件
	python build.py会输出所需的依赖路径，这个方便粘贴复制到gogland等ide
## 测试
	双击test.html，打开界面，登陆输入框里面输入玩家id，会收到服务器回复
	再次双击test.html，打开界面，这时有两个web页面，登陆输入框里面输入玩家id，和上面不一样
	输入全局聊天，点击发送，两个页面应该都有输出
	输入玩家id，点击发送，对方应该能收到消息，同时自己也会收到确认回复
## 代码
	所有代码都在main.go，不超过60行