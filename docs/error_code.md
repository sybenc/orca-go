# 错误码

## 功能说明

如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：

```json
{
  "code": 100101,
  "message": "Database error"
}
```

上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

## 错误码列表

系统支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| Success | 100001 | 200 | 请求成功 |
| ErrInternalServer | 100002 | 500 | 服务器内部错误 |
| ErrBadRequest | 100003 | 400 | 请求存在错误 |
| ErrNotFound | 100004 | 404 | 资源未找到 |
| ErrValidation | 100005 | 400 | 字段验证错误 |
| ErrBind | 100006 | 400 | 绑定参数错误 |

