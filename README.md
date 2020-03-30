# tabcot

tabcot 是一个表格数据转换器，把 json, csv, excel, html table, sqlite 等转化为 json, csv, excel,html table, sqlite 等数据格式

平时有很多转换数据要处理，程序中大部分处理的都是 json 数据
如果要给其他非技术人员使用，他们不怎么看得懂的，所以，需要转换成 csv 或者 excel 的方式

本次，准备整理一个工具，把 json 数据，转换成其他的数据

## example

```
tabcot input.json output.csv
tabcot input.json output.csv -k data
tabcot http://example.com/data.json output.csv -k data
tabcot http://example.com/data.json -k data
```

## todo

1. 封装成一个 lib
2. 提供一个命令行工具
3. 数据过滤清洗

参考 python tablib
https://github.com/pubgo/gojsonq
