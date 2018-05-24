# datastash

Go 语言编写的一个微服务，用于接收业务数据日志存入 MongoDB，便于以后数据分析

## Dep

已将 `vendor` 目录加入版本控制，检出项目后不用安装依赖

如想增加依赖必须使用代理

```bash
$ HTTP_PROXY=<your proxy> dep ensure -add <repository url>
```

## Api

### HTTP Request

这个接口为异步接口，收到请求正确解析 json 后直接返回 Http Status 200， 不返回是否成功插入 MongoDB 结果

`POST http://datastash/rpc/stash`

### Request Body (application/json)

```json
{
	"database": "foo",
	"collection": "person",
	"document": {
		"key": "value"
	}
}
```

Property | Require | Description
--------- | ------- | -----------
database | true | The mongodb database name.
collection | true | The mongodb collection name.
document | true | MongoDB document object, can't be array.