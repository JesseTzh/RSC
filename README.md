<p align="center">
<img alt="RSC" src="https://user-images.githubusercontent.com/27763415/224210541-f0167570-33c2-42d7-b17a-33585fa82915.png">
<br>
Remote Shell Commander
<br>
一个极简的通过 Web 远程执行 Shell 命令的工具
<br>
A minimalist tool for executing Shell commands remotely via the Web.
</p>

## 💡 简介

[RSC](https://github.com/JesseTzh/RSC) Remote Shell Commander 的缩写，旨在通过 Web 接口直接执行 Shell 脚本或命令。

适用于不方便 SSH 访问的情景中远程执行脚本或 Shell 命令，甚至可以用来做一些简单运维工作。

>⚠️警告：由于程序本身不具有鉴权功能，请在使用时注意安全问题。

## ✨ 使用
在 [Releases](https://github.com/JesseTzh/RSC/releases) 中下载最新的编译好的程序，并上传至服务器中。

可以使用 nothup 命令运行程序，例如：
```
nohup ./rsc &
```
并可以通过以下启动参数更改默认配置
```
Usage of .\RSC.exe:
  -c string
        The default executed command. (default "echo hello world")
        默认要执行的 shell 命令（不配置则为 "echo hello world"）
  -f    Enable the free command mode? (Use with caution)
        是否开启自由命令模式（开启后任何人均可以直接向服务器执行任何命令，请谨慎使用）
  -p string
        Change the default(6021) server port. (default "6021")
        修改默认端口
```
