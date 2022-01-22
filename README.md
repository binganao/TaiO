# TaiO
TaiO 的定位是一款用于攻击方对靶标资产梳理，快速定位脆弱资产的工具

![](https://github.com/binganao/TaiO/blob/master/assets/Process.png?raw=true)

## 注意

在使用本项目之前，请确保已经得到授权。同时**本项目功能还未完善**，同时仅包含后端功能

## DONE:

Day1: 配置文件、Logger、数据库、端口扫描、服务扫描、指纹识别、单个 IP 探测、结果整理至数据库、通过 API 查询数据、自动探测主机、优化路由分组

Day2: 增加添加数据路由、Banner、解决删除 jobs 有时侯删除失败、异步 Nmap 扫描，实时出结果

## TODO:

### 重要

完成增加数据的操作（已完成，暂未测试），数据添加至 临时 表中，当验证过密钥之后再加入正式表

完成删除数据的操作

增加 jwt 认证

增加 jwt 中间件

增加 Cors 中间件

### 特点

增加更多信息：IP 定位、运营商、Header、Body、哈希

增加危险服务标记：RDP、MYSQL、MSSQL、REDIS、SSH

### 当前任务

休息了家人们

## 感谢

[https://github.com/EdgeSecurityTeam/EHole](https://github.com/EdgeSecurityTeam/EHole)

[https://github.com/Ullaakut/nmap](https://github.com/Ullaakut/nmap)

[https://github.com/zan8in/masscan](https://github.com/zan8in/masscan)
