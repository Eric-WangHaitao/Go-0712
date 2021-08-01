# 学习笔记
## 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## 分析
DAO 不应该把异常处理掉，因为事务的边界不在 DAO 层上，而是在 Service 层上，假如一个业务逻辑调用了两个 DAO 的方法（A、B）组成一个事务体，若 A 正常执行了，而 B 执行时却发生了异常，这时 A 的执行就需要回滚，若将 B 中的异常捕获了，那根本就不知道 B 是正常执行而是异常执行了，也无法处理事务了。

因此Dao层往上抛，在Service层捕获。

## 代码示例
```cmd
git clone https://github.com/Eric-WangHaitao/Go-0712.git
cd Week02
go run main.go
```

