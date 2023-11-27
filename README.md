<div style="text-align: center"><h1>Gin Example</h1></div>
<div style="text-align: center;">This is a sample project for Gin.</div>
<div style="text-align: center;">Gin is a web framework written in Golang.</div>
<div style="text-align: center;">It features a Martini-like API, but with performance up to 40 times faster than Martini.</div>
<div style="text-align: center;">If you need performance and productivity, you will love Gin.</div>
<div>
    conf：用于存储配置文件<br>
    middleware：应用中间件<br>
    models：应用数据库模型<br>
    pkg：第三方包<br>
    routers 路由逻辑处理<br>
    runtime：应用运行时数据<br>
</div>
``
version: '3.3'
services:
  database:
    image: amp:1.0.0
    container_name: amp-go
    restart: always
    ports:
      - '9000:9000'
    expose:
      - '9000'
networks:
  default:
    external:
      name: mysql_default
``
