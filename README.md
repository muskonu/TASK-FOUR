# 图书管理

**接口地址**

> 42.192.127.60:8080

**说明**

- 返回格式: JSON

- 请求方式: GET

- 请求示例(随机获取一本书)

  42.192.127.60:8080/book

**功能说明**

*/book(无需验证登录)*

| 名称 | 必填 | 类型   | 说明         |
| ---- | ---- | ------ | ------------ |
| id   | 否   | int    |     编号     |
|class | 否   | string |     分类     |
| name | 否   | string |     名字     |
|author| 否   | string |     作者     |

- 一个都不填则随机返回

*/register*

| 名称 | 必填 | 类型     | 说明         |
| ------ | ---- | ------ | ------------ |
| account| 是   | string |     账号     |
|password| 是   | string |     密码     |

- 注册一个账号 

*/login*

| 名称 | 必填 | 类型     | 说明         |
| ------ | ---- | ------ | ------------ |
| account| 是   | string |     账号     |
|password| 是   | string |     密码     |

- 登录

*/borrow*

| 名称 | 必填 | 类型   | 说明         |
| ---- | ---- | ------ | ------------ |
| id   | 是   | int    |     编号     |

- 借书并返回所有借阅信息

*/return*

| 名称 | 必填 | 类型   | 说明         |
| ---- | ---- | ------ | ------------ |
| id   | 是   | int    |     编号     |

- 还书并返回所有还书信息

*/create*

| 名称 | 必填 | 类型   | 说明         |
| ---- | ---- | ------ | ------------ |
|class | 是   | string |     分类     |
| name | 是   | string |     名字     |
|author| 是   | string |     作者     |

- 捐书加入数据库,每加入一本书,账号级别加一,最高为4，可持有书本总数和持有天数增加

*/info*

- 返回账号信息

**过程总结**

- 用colly随便爬了一个网站填充数据库书籍信息
- 登录验证使用jwt(最后还是存在了cookie当中)
- controllers中调用的数据库的存储过程放在proceduer中(话说创造存储过程还要重新定义终止符好麻烦)
- 用了个ticker定时检查有无人过时不还书，未还则级别降为0，无法借书
- 为了访问方便所有都是get方法
- 调用数据库用的gorm

总之就是写的时候就感觉自己写了很多没用的东西，并且也有很多地方写的十分繁琐。权当锻炼和这段时间的总结吧。
