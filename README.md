# ssbcV2

## 启动
* 构建

``` go build ```

* 单链：在根目录开启4个终端，分别运行4个节点和1个客户端
```
./ssbcV2 N0
./ssbcV2 N1
./ssbcV2 N2
./ssbcV2 N3
./ssbcV2 client1
```
或者
```
./ssbcV2 N4
./ssbcV2 N5
./ssbcV2 N6
./ssbcV2 N7
./ssbcV2 client2
```
  
* 跨链：在根目录开启10个终端，分别运行8个节点和2个客户端
```
./ssbcV2 N0
./ssbcV2 N1
./ssbcV2 N2
./ssbcV2 N3
./ssbcV2 N4
./ssbcV2 N5
./ssbcV2 N6
./ssbcV2 N7
./ssbcV2 client1
./ssbcV2 client2
```

* 在根目录启动前端项目visual-bctt
```
npm run dev
```

* 通过 http://localhost:9528/，即可访问后端


* 其他
```
./ssbcV2 clear  # 删除数据
```
