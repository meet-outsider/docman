
    200 OK：请求成功，服务器已成功处理了请求。
    201 Created：请求成功，服务器已创建了新的资源。
    204 No Content：请求成功，但服务器没有返回任何内容。
    301 Moved Permanently：请求的资源已被分配了新的 URI，以后应使用资源现在所指的 URI。
    302 Found：请求的资源已被分配了新的 URI，希望用户能使用新的 URI 访问。
    304 Not Modified：请求的资源未发生变化，可以直接使用客户端缓存的内容。
    400 Bad Request：客户端发送的请求有错误，服务器无法处理该请求。
    401 Unauthorized：请求未经授权，需要用户进行身份验证。
    403 Forbidden：服务器拒绝了该请求，因为客户端没有足够的权限。
    404 Not Found：请求的资源不存在。
    405 Method Not Allowed：请求方法不被允许，例如使用了不支持的 HTTP 方法。
    406 Not Acceptable：请求的资源的内容特性无法满足请求头中的条件，因而无法生成响应实体。
    409 Conflict：请求冲突，例如在注册用户时，如果用户已经存在，则会返回该状态。
    500 Internal Server Error：服务器内部错误，无法完成请求。
    502 Bad Gateway：服务器作为网关或代理，从上游服务器收到了无效的响应。
    503 Service Unavailable：服务器当前无法处理请求，一段时间后可能恢复正常。
