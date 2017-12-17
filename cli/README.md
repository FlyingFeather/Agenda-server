### 命令设计
#### agenda [command]

可以使用的command:
1. help
2. list
3. register
4. login
5. logout
*** 
##### agenda help :列出帮助命令
***
##### agenda list: 列出所有的注册的用户
[flags]:
+ -h 帮助
###### 例子：agenda list
*** 
##### agenda register [flags]
[flags]:
+ -u 用户名
+ -p 密码
+ -m 邮箱地址
+ -t 手机号码
+ -h 帮助
###### 例子：agenda register -u young -p password -m 16368278@qq.com -t 13456726543
***
##### agenda login [flags]
[flags]:
+ -u 用户名
+ -p 密码
+ -h 帮助
###### 例子：agenda login -u young -p password 
***
##### agenda logout [flags]
[flags]:
+ -h 帮助
###### 例子：agenda logout
