# Meta

题目信息 Meta

由[题目环境生成器 - cg](https://github.com/ctfhub-team/challenge_generate)自动生成

## 说明

```yaml
author:
  # 制作者ID，由cg自动生成
  name: l1n3
  # 制作者邮箱，由cg自动生成
  contact: yw9381@163.com
task:
  # 题目镜像名称，由cg自动生成，题目自身名称中应当不包含_
  name: challenge_2022_hitcon_web_rce-me
  # 访问类型 1 = 浏览器访问 | 2 = nc访问 | 3 = 附件
  type: 1
  # 题目分类，首字母需要大写，每个题目应当只有一个分类
  # 例如 Web, Pwn, Reverse, Misc, Crypto, Forensics, Blockchain, Android, Mobile, ICS, IoT
  category: Web
  # 题目描述
  description: aasasdsadas
  # 题目难度，由cg自动生成
  level: 签到
  # 题目flag
  # 如是静态flag在此处填写具体的flag值
  # 如是动态flag则此处为空字符串
  flag: 
  # 题目提示，如无提示则此处为空数组
  hints:
    - asdasdaz
    - asdas
    - asdad
challenge:
  # 题目显示名称
  name: rce me
  # 题目来源，来源格式: 年份-比赛名称简写-题目类型-题目显示名称
  # 例如 2021年强网杯的Web类的babysqli，则为2021-QWB-Web-baysqli
  # 例如 2019年SCTF的Pwn类的babyheap，则为2019-SCTF-Pwn-bayheap
  refer: 2022-HitCon-Web-rce
  # 题目标签，标签应当尽可能体现题目考点，如无标签则此处为空数组
  tags:
    - web
    - 2022
    - hitcon
```

## Yaml

### Single 单文件

```yaml
author:
  name: 陌竹
  contact: mozhu233@outlook.com
task:
  name: skill_web_advanced_jwt_infoleakage
  type: 1
  category: Web
  description: JWT的头部和有效载荷这两部分的数据是以明文形式传输的
  level: easy
  flag: ""
challenge:
  name: 敏感信息泄露
  refer: https://www.ctfhub.com
  tags:
    - web
    - jwt
```

### Multiple 多文件合并

```yaml
---
author:
  name: 陌竹
  contact: mozhu233@outlook.com
task:
  name: skill_web_advanced_jwt_infoleakage
  type: 1
  category: Web
  description: JWT的头部和有效载荷这两部分的数据是以明文形式传输的
  level: easy
  flag: ""
challenge:
  name: 敏感信息泄露
  refer: https://www.ctfhub.com
  tags:
    - web
    - jwt
---
author:
  name: 陌竹
  contact: mozhu233@outlook.com
task:
  name: skill_web_ssrf_file_read
  type: 1
  category: Web
  description: 尝试去读取一下Web目录下的flag.php吧
  level: easy
  flag: ""
challenge:
  name: 伪协议读取文件
  refer: https://www.ctfhub.com
  tags:
    - web
    - ssrf
```
