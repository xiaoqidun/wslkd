# WslKD

Windows Subsystem for Linux Keep Daemon

# 简要说明

WSL在没有使用的情况下守护进程会被关闭，本工具在指定的WSL发行版打开一个SHELL为WSL守护进程保活。

# 参数说明

-d，需要保活守护进程的WSL发行版，默认值：Debian，同时为多个WSL发行版保活守护进程用英文逗号分割。
