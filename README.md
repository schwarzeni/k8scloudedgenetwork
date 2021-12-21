项目结构

- crd: 监听自定义 service 创建的控制器，与 store 通信；与 k8s api server 通信
- proxyserver/cloud: 云端 ep 节点，与 store 通信，查询节点信息
- proxyserver/edge：各个节点上的 ep，主动与云端 ep 连接，跨网段传输数据
- simpledns：接收 pod 查询自定义 svc 服务 URL 的 dns 请求
- store：存储 service 信息，与 k8s api server 通信

