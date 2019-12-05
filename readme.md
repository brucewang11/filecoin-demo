程序启动流程<br>
1.程序分为两部分：主节点DMineMain和子节点DMineSub<br>
2.包管理工具使用go mod<br>
3.先启动DMineMain,启动方式为： ./DMineMain 5001(启动grpc服务的端口号) 5002(可选,http测试端口，测试使用)
4.后启动DMineSub,启动方式为：./DMineSub 127.0.0.1:10001（子节点grpc服务的地址）127.0.0.1:5001(主节点注册服务的地址)<br>
5.测试：http get方式请求方法 http://127.0.0.1:5002/seal_pre_commit<br>
                           http://127.0.0.1:5002/seal_commit  <br>
                           http://127.0.0.1:5002/gen_candidates  <br>
                           http://127.0.0.1:5002/gen_post  <br>

原理：主节点暴露grpc接口，子节点启动grpc服务，并将grpc地址通过http请求发送到主节点，<br>
主节点通过grpc连接子节点，并保存conn，可以注册多个子节点<br>

调用add_piece ,主节点通过一致性哈希随机调用子节点的add_piece方法,add_piece完成后，通知主节点sealResult方法<br>
调用gen_post,主节点通过轮询所有conn，调用子节点gen_post方法,gen_post完成后，通知主节点postResult方法 <br>

目前还有的问题：<br>
1.主节点单点问题<br>


                           
                             
