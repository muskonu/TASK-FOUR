## 图书管理
# version 0.1

**接口地址**

> http://42.192.127.60:8080/book

**说明**

- 返回格式: JSON

- 请求方式: GET

- 请求示例(随机返回一本图书)

  http://42.192.127.60:8080/book

- 请求示例(根据名字查找图书)

  http://42.192.127.60:8080/book/:name
  
- 请求示例(根据类别查找图书)(后可跟返回图书的数目，不输入即默认为一)

  http://42.192.127.60:8080/books/:class(/:number)

**参数说明**

| 名称 | 说明           |
| ---- | -------------- |
| class | v0.1版本只有自然科学，外国文学，电脑网络，诗歌散文四个分类 |

**返回数据**

```json
{
    "id":152,
    "author":"伊恩·麦克尤恩",
    "name":"甜牙",
    "classification":"外国文学"
}
```

**v0.1版本我使用colly进行爬取来填充数据库，后续版本会持续增多，目前数据库的数据模型只有图书，后续版本同样可能会增加。**

**经验**

- 使用colly爬取数据填充数据库时，我发现即使将数据库中的域设为not null，其中有一些仍然为空。
经翻找资料发现，mysql中的nut null不能限制其元组的值为空值，只能使其不能为空。

- 由于Linux环境下不同版本MySQL的兼容问题导致我数据从电脑转移到服务器几经波折，最后改变sql文件的编码得已解决。



