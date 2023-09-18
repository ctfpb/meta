# Meta

题目信息 Meta

## Yaml

### Single 单文件

```yaml
author:
  name: 陌竹
  contact: mozhu233@outlook.com
task:
  name: skill_web_advanced_jwt_infoleakage
  type: Web
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
  type: Web
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
  type: Web
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
