# TaiO - 三合
TaiO 的定位是一款用于攻击方对靶标资产梳理，快速定位脆弱资产的网络空间测绘安全工具

![](https://github.com/binganao/TaiO/blob/master/assets/Snipaste_2022-01-23_18-16-48.png?raw=true)

## 注意

在使用本项目之前，请确保已经得到授权。同时**本项目功能还未完善**，同时仅包含后端功能

### 当前进度：

#### 后端返回结果
```json
{
  "Code": 200,
  "Msg": "操作成功",
  "Hosts": "打码",
  "Ports": [
    "111",
    "3306",
    "80",
    "443",
    "1234",
    "22"
  ],
  "Services": [
    {
      "Port": "3306",
      "ServiceName": "mysql"
    },
    {
      "Port": "22",
      "ServiceName": "ssh"
    },
    {
      "Port": "111",
      "ServiceName": "rpcbind"
    },
    {
      "Port": "80",
      "ServiceName": "http"
    },
    {
      "Port": "443",
      "ServiceName": "http"
    },
    {
      "Port": "1234",
      "ServiceName": "hotline"
    }
  ],
  "Fingers": [
    {
      "Url": "https://打码:443",
      "App": "[]",
      "Server": "nginx"
    }
  ]
}
```

## DONE:

Day1: 配置文件、Logger、数据库、端口扫描、服务扫描、指纹识别、单个 IP 探测、结果整理至数据库、通过 API 查询数据、自动探测主机、优化路由分组

Day2: 增加添加数据路由、Banner、解决删除 jobs 有时侯删除失败、异步 Nmap 扫描，实时出结果、修复 异步 Nmap 导致服务器网络崩溃的问题

Day3: 完成增加数据的操作，数据添加至 临时 表中、完成临时数据的查找功能

Day4: 完成删除临时数据的操作，优化代码结构、完成 增加 真实数据的功能、完成 删除 真实数据的功能

Day5: 增加部分路由、修正 增加临时数据的提示、测试 增加 真实数据的功能、测试 删除 真实数据的功能

## TODO:

### 重要

加入节点判断，若是子节点则对应子节点功能，若是主节点，则不进行数据采集

增加 jwt 认证

增加 jwt 中间件

增加 Cors 中间件

### 特点

增加更多信息：IP 定位、运营商、Header、Body、哈希、Title

增加危险服务标记：RDP、MYSQL、MSSQL、REDIS、SSH

### 当前任务

增加 获取全部 host 的功能

增加 获取全部临时 host 的功能

## 感谢

[https://github.com/EdgeSecurityTeam/EHole](https://github.com/EdgeSecurityTeam/EHole)

[https://github.com/Ullaakut/nmap](https://github.com/Ullaakut/nmap)

[https://github.com/zan8in/masscan](https://github.com/zan8in/masscan)

