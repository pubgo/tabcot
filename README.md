# tabcot

tabcot 是一个表格数据转换器，把 json, csv, excel, html table, sqlite 等转化为 json, csv, excel,html table, sqlite 等数据格式

平时有很多转换数据要处理，程序中大部分处理的都是 json 数据
如果要给其他非技术人员使用，他们不怎么看得懂的，所以，需要转换成 csv 或者 excel 的方式

本次，准备整理一个工具，把 json 数据，转换成其他的数据

```
-k jsonpath 表达式
-s 是否流式输出
```

## example

默认输出到 output.csv 这个文件当中
如果添加-s 参数，那么，就通过流的方式打印输出
tabcot 会对输入文件进行识别，如果是 url，那么，就获取内容，然后再进行处理
tabcot 会识别是否是管道输入，如果是，那么就从管道获取内容作为输入源，然后进行处理

```
tabcot input.json output.csv
tabcot input.json output.csv -k data
tabcot http://example.com/data.json output.csv -k data
tabcot http://example.com/data.json -k data
curl http://example.com/data.json | tabcot -k data
curl http://example.com/data.json | tabcot -k data -s
```

## todo

1. 封装成一个 lib
2. 提供一个命令行工具
3. 数据过滤清洗

### 输入

1. 文件
2. url
3. 管道流

## 参考

```
tablib
gojsonq
jq
```
