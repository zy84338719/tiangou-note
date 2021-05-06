## 舔狗日记
<h1 align="center">
  <br>
  <a href="https://cloudreve.org/" alt="logo" ><img src="https://raw.githubusercontent.com/cloudreve/frontend/master/public/static/img/logo192.png" width="150"/></a>
  <br>
  upftp
  <br>
</h1>

<h4 align="center">舔狗日记.</h4>

<p align="center">
  <a href="https://github.com/zy84338719/tiangou-note">
  <a href="https://codecov.io/gh/zy84338719/tiangou-note">
    <img src="https://img.shields.io/codecov/c/github/zy84338719/tiangou-note?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/zy84338719/tiangou-note">
      <img src="https://goreportcard.com/badge/github.com/zy84338719/tiangou-note?style=flat-square">
  </a>
  <a href="https://github.com/zy84338719/tiangou-note/releases">
    <img src="https://img.shields.io/github/v/release/zy84338719/tiangou-note?include_prereleases&style=flat-square">
  </a>
</p>

<p align="center">
  <a href="#scroll-许可证">许可证</a>
</p>

![Screenshot](https://raw.githubusercontent.com/zy84338719/tiangou-note/master/img.png)

数据库信息
```sql
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tiangou
-- ----------------------------
DROP TABLE IF EXISTS `tiangou`;
CREATE TABLE `tiangou` (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `key` varchar(255) DEFAULT NULL,
                                 `value` varchar(1024) DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `uni_key` (`key`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;


INSERT INTO `tiangou`(`id`, `key`, `value`) VALUES (4, '19c0dfb6bd9c5b170e6fa03f099f6767', ' 她刚刚说去洗澡了，我等了她3个小时。我问她玩不玩，她说要睡觉了。可当我上线看了她的战绩，原来已经有别的璇璇了。没关系，明天我也会舔你。');
INSERT INTO `tiangou`(`id`, `key`, `value`) VALUES (5, 'a8c4d734fe9f20854e3a5f9432baeacf', ' 小时候抓周抓了个方向盘，爸妈都以为我长大了会当赛车手，最差也是个司机，没想到我长大了当了你的备胎。');
INSERT INTO `tiangou`(`id`, `key`, `value`) VALUES (6, 'ef45c7b39ba5c36c4bcdd77c42771fd5', ' 我可以再见你一面吗，我可以站远一点，不让你同事发现我。');
INSERT INTO `tiangou`(`id`, `key`, `value`) VALUES (7, '668ab3519e2f935c673b433a6d0e3808', ' 今天有点口腔溃疡，不太想舔了，和你的旧爱好好的啊，不要不开心了。');
```

## 编译
```bash
go build
```


### 如何使用
#### 程序运行在 59998 端口
#### 运行目录为 http://127.0.0.1:59998/web
