# userApi
探探交友的一份golang必答题，用原生golang简单实现一下探探的用户核心模块
当时下载了探探app（**探探交友招聘有涉嫌的嫌疑用招聘拉用户量的嫌疑吆**），不同用户之间存在三种关系
##三种关系描述
1.a喜欢b 
2.b喜欢a
3.ab 相互喜欢
##逻辑实现
采用原生go语言进行路由，提供三种用户关系的api，并提供了简单的sql表设计
