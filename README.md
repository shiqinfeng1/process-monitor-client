# [process-monitor-client](https://github.com/shiqinfeng1/process-monitor-client)
## 功能
* 用于进程监控，管理

## 安装
> git clone https://github.com/shiqinfeng1/process-monitor-client.git


## 实现
* 被监控进程启动后，按每500ms执行一次状态检测（通过发signal0信号检测），每个被监控进程在一个独立的协程里被监测。
* monitor启动后接收管理命令（start|stop|restart|status|...）

## 编译
* 在项目目录下执行
```
./gradlew init
./gradlew build
```
*生成macos和linux平台下的2个可执行文件,路径在 ./build/bin .

## 使用方法
* 配置文件：[monitor.ini](https://github.com/shiqinfeng1/process-monitor-client/blob/master/monitor.ini) (ini格式，支持注释)

```
    [Demo]  ;监控模块名称
    ProcessName = Demo  ;监控进程名称
    command=php 监控程序目录/test/test.php  ;监控进程启动命令
    autostart=true       ;是否自启动
    autorestart=true     ;是否随进程启动而重启
    logfile=/var/log/monitor.log  ;日志文件存储，可以为空，默认为xlog目录下monitor.log
```
* 配置文件需要和可执行程序一起拷贝到同一目录
* 运行方法：./monitor [start|stop|restart|status|check]
* 进程管理：./monitor -[list|status|start|stop|restart]

## 注意
* 当monitor启动时，会根据monitor.ini配置启动所有被监控进程，当被监控进程已经启动过，并且符合配置要求时，monitor会自动将其加入监控列表
* monitor会定期检查进程运行状态，如果进程异常退出，monitor会反复重试拉起，并且记录日志
* 当被监控进程为多进程运行模式，monitor只监控管理父进程(子进程应实现检测父进程运行状态，并随父进程退出而退出）
* 被监控进程以nohup方式启动，所以你的程序就不要自己设定daemon运行了
* 每30秒通过ps方式检测一次进程状态，如果出现任何异常，比如有多份进程启动等，记日志

## demo
```
$ ./monitor start

Process:demo is already stop...
Process:demo start [success]
Process:test is already stop...
Process:test start [success]

```
