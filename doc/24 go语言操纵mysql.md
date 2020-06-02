###  sql的分类

```html
1 DDL（Data Definition Language）：数据定义语言，用来定义数据库对象：库、表、列等，例如CREATE、ALTER、DROP
2 DML（Data Manipulation Language）：数据操作语言，用来定义数据库记录（数据），例如INSERT、UPDATE、DELETE
3 DQL（Data Query Language）：数据查询语言，用来查询记录（数据），例如SELECT
4 DCL（Data Control Language）：数据控制语言，用来定义访问权限和安全级别
```

### 事务

```html
1 事务：一组逻辑操作单元，使数据从一种状态变换到另一种状态
2 当在一个事务中执行多个操作时，要么所有的事务都被提交(commit)，那么这些修改就永久地保存下来；要么数据库管理系统将放弃所作的所有修改，整个事务回滚(rollback)到最初状态
3 为确保数据库中数据的一致性，数据的操纵应当是离散的成组的逻辑单元，当它全部完成时，数据的一致性可以保持，而当这个单元中的一部分操作失败，整个事务应全部视为错误，所有从起始点以后的操作应全部回退到开始状态
```

### 事务的ACID

```html
	原子性（Atomicity），要么都成，要么都不成
	一致性（Consistency）
	隔离性（Isolation），多个事务不能互相干扰
	持久性（Durability），事务提交，改变是永久性的
```

### 事务的隔离级别

```html
1 对于同时运行的多个事务, 当这些事务访问数据库中相同的数据时, 如果没有采取必要的隔离机制, 就会导致各种并发问题
		脏读，T1读取到了T2未提交的修改数据
		不可重复读：T1读取到了T2已提交的修改数据
		幻读：T1读取到了T2新添加的数据
2 数据库事务的隔离性：数据库系统必须具有隔离并发运行各个事务的能力，使它们不会相互影响，避免各种并发问题
3 一个事务与其他事务隔离的程度称为隔离级别，数据库规定了多种事务隔离级别，不同隔离级别对应不同的干扰程度，隔离级别越高，数据一致性就越好，但并发性越弱
4 数据库事务4个隔离级别：
		读未提交
		读已提交，oracle是默认这个级别的	
		可重复读，是mysql的默认隔离级别
		串行化		// 效率非常低下
```

### mysql常见的数据类型

```go
	int：整型
	double：浮点型，例如double(4,2)表示最多4位，其中必须有2位小数
	char：固定长度字符串类型，例如char(1)
	varchar：可变长度字符串类型，例如varchar(10)
	text：大文本类型
	blob：字节类型
	boolean：使用0或1表示真或假
	date：日期类型，格式为：yyyy-MM-dd
	time：时间类型，格式为：hh:mm:ss
	timestamp：时间戳类型，格式为：yyyy-MM-dd hh:mm:ss 会自动赋值
	datetime：日期时间类型，格式为：yyyy-MM-dd hh:mm:ss
```

### mysql 表操作

```mysql
	查看所有的表：SHOW tables;
	创建表：CREATE table 表名(字段 数据类型,……); 
CREATE table my_table1(
	id int,
	name varchar(10)
);
	查看表的字段信息：DESC 表名;
	创建带注释的表：COMMENT ‘注释内容’;
	为上表增加一个列：ALTER TABLE 表名 ADD 列名 数据类型;
	修改刚增加列的数据类型：ALTER TABLE 表名 MODIFY 列名 数据类型;
	修改列名：ALTER TABLE 表名 CHANGE 原列名 新列名 数据类型;
	将刚增加的列删除：ALTER TABLE 表名 DROP 列名;
	修改表名：RENAME TABLE 原表名 TO 新表名;
	查看表创建信息：SHOW CREATE TABLE 表名;
	修改表的编码集：ALTER TABLE 表名 character set 编码集;
	删除表：DROP TABLE 表名;
```

## go操作mysql

```go
// 安装相应的依赖
go get -u github.com/go-sql-driver/mysql
go get github.com/jmoiron/sqlx
_"github.com/go-sql-driver/mysql"    //只是加载驱动
"github.com/jmoiron/sqlx"						 //里面是操作mysql的具体	
```

`1 连接数据库and数据的增删改`

```go
import _ "github.com/go-sql-driver/mysql"
import "github.com/jmoiron/sqlx"

func main()  {
	// 1 连接数据库
	db, _ := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/firstTest")
	defer db.Close()
	// 2 添加 使用? 进行占位避免sql注入，返回值是受影响的行数
	result, _ := db.Exec("insert into person(name,age,money,birthday) values(?,?,?,?)",
		"张三", 20, 100.5, 20190101)
	// 查看受影响的行数
	row, _ := result.RowsAffected()
	// 返回受影响的最后一个id
	lastId,_ := result.LastInsertId()
	fmt.Printf("首影响的函数%d，受影响的最后一个id:%d\n", row, lastId)
  // 3 修改数据
  db.Exec("update person set name=? where name=?","李四", "张三")
  // 4 删除数据
  db.Exec("delete from person where id=?", 2)
  // 5 事务操作
  tx := db.MustBegin()
  tx.MustExec("要执行的sql")
  err := tx.Commit()
  if err != nil{
    fmt.Println(",,,,,")
    tx.RollBack()
  }
}
//首影响的函数1，受影响的最后一个id是1
```

`2 数据库的查询`

```go
// 查询出来的结果要映射到具体的对象上
import _ "github.com/go-sql-driver/mysql"
import "github.com/jmoiron/sqlx"

type Person struct {
	Id       int       `db:"id"`
	Name     string    `db:"name"`
	Age      int       `db:"age"`
	Money    float64   `db:"money"`
	Birthday time.Time `db:"birthday"`
}
func main() {
	// 1 建立数据库连接
	// go默认不支持时间解析，需要在建立链接的时候加上时间解析的参数 parseTime=true
	db, _ := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/firstTest?parseTime=true")
	defer db.Close()
	// 2 查询
	var ps []Person
	err := db.Select(&ps, "select name,age from person where id=?", 1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%v\n", ps)
}
```





