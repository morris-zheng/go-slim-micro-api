# go-slim-micro-api
## 基于go-slim基础，加入etcd服务发现，实现 BFF 层。
#### · 微服务（如 usersvc 等）通过 NewServiceContext 注入至 domain.ServiceContext 容器中。
#### · delivery.Register() 实例化handler，注册路由（restful）并绑定相应handler。
