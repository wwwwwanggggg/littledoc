## LittleDoc

一个简单的生成API文档的工具，将根据文件中的注释进行文档生成，生成的某个接口的文档示例如下：

#### 示例

为了写示例使用了六级标题，实际生产出来是四级标题

###### 请求方式 GET 

###### 请求URL {baseurl}/api/example

###### 传入格式 JSON

###### 参数格式

|名称|类型|是否必填|样例|说明|
|--|--|--|--|--|
|`example`|`string`|`true`|`喜欢就点个Star`|`Follow Please`|

###### Response

情况1
```json
{
    "success":true
} 
```

情况2
```json
{
    "success":false
}
```

#### 使用方式

###### 参数传递

###### 文件注释规范

