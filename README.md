# drpc
Cool RPC project

- client: 支持客户端发包模块
- server: 支持服务端收包模块
- transport: 提供底层通信能力模块
- codec: 自定义协议的解析、序列化和反序列化模块
- pool: 池技术、支持连接池、对象池等实现；提供客户端连接的复用，对象复用的能力
- log:  提供日志能力模块
- selector: 提供寻址能力、服务发现、负载均衡能力的模块
- stream: 提供客户端和服务端上下文数据透传能力的模块
- protocol: 提供自定义私有协议能力的模块
- plugin: 提供第三方插件化支持能力
- interceptor: 提供框架拦截器能力
- metadata: 提供客户端和服务端参数传递能力

## TODO

### 客户端
1. 客户端链接超时
2. 客户端重连机制
3. 客户端重复握手保护
4. 消息缓存转发
5. 心跳机制
6. Client automatically manages connections and automatically reconnects to the server on connection errors.
7. Client supports response timeouts.
8. Client supports RPC batching.
9. Client supports async requests' canceling.
10. Client prioritizes new requests over old pending requests if server fails to handle the given load.
11. Client detects stuck servers and immediately returns error to the caller.
12. Client supports fast message passing to the Server, i.e. requests without responses.
13. Both Client and Server provide network stats and RPC stats out of the box.


### 服务端
1. 服务端超时
2. 服务端调用失败
3. 失败自动切换
4. 失败通知
5. 失败缓存
6. 快速失败
7. Server provides graceful shutdown out of the box.
8. Server supports RPC handlers' councurrency throttling out of the box.
9. Server may pass client address to RPC handlers.
10. Server gracefully handles panic in RPC handlers.
11. Dispatcher accepts functions as RPC handlers.
12. Dispatcher supports registering multiple receiver objects of the same type under distinct names.
13. Dispatcher supports RPC handlers with zero, one (request) or two (client address and request) arguments and zero, one (either response or error) or two (response, error) return values.

### 负载均衡
1. 随机
2. 轮询
3. 服务调用时延
4. 一致性哈希
5. 粘滞连接 
