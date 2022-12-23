# go-Community
go——Community
  目录：
     功能实现
     接口说明
        注册
        登录
        当前登录用户提出问题GET
        当前登录用户回答问题GET

        查看当前登录用户问题GET
        查看当前登录用户回答GET

        删除当前登录用户问题GET
        删除当前登录用户回答GET

        修改当前登录用户问题GET
        修改当前登录用户回答GET

一、功能实现
  利用Gin框架和MYSQL，在本地创建了名叫Community的database,在其中创建了users、qas两table，分别存放用户数据和问答数据
  在vscode利用go代码与本地Community数据库连接
  利用原生SQL语句对数据库中两个表进行操作
  即单体
二、接口说明
  注册GET
  http://127.0.0.1:8080/user/register?username=(必填)&password=(必填)


  登录GET
  http://127.0.0.1:8080/user/login?username=(必填)&password=(必填)


  当前登录用户提出问题GET
  http://127.0.0.1:8080/user/question?username=(必填)&question=(必填)


  当前登录用户回答问题GET 
  http://127.0.0.1:8080/user/answer?username=(必填)&answer=(必填)&questionID=(必填)


  查看当前登录用户问题GET
  http://127.0.0.1:8080/user/getquestion?username=(必填)


  查看当前登录用户回答GET
  http://127.0.0.1:8080/user/getanswer?username=(必填)


  删除当前登录用户问题GET
  http://127.0.0.1:8080/user/questiondelete?username=(必填)&questionID=(必填)


  删除当前登录用户回答GET
  http://127.0.0.1:8080/user/answerdelete?username=(必填)&questionID=(必填)


  修改当前登录用户问题GET
  http://127.0.0.1:8080/user/questionalter?username=(必填)&questionID=(必填)&newquestion=(必填)


  修改当前登录用户回答GET
  http://127.0.0.1:8080/user/answeralter?username=(必填)&questionID=(必填)
  &newanswer=(必填)
