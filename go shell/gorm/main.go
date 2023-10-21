package main
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)


/** 官方文档
http://gorm.io/docs/models.html
**/
// gorm.Model是基本结构体
// 结构体成员：ID，CreatedAt，UpdatedAt，DeletedAt
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(10)"`
	Age int
	Address string `gorm:"type:varchar(255);default:'GZ'"`
}



/** 官方文档：
https://gorm.io/zh_CN/docs/conventions.html
**/
// 结构体User默认的数据表名为Users
// 如果自定义数据表名，可自定义TableName方法
func (User) TableName() string{
	return "my_user"
}
func main(){
	// 连接数据库
	// 连接方式1：使用database/sql和go-sql-driver/mysql连接数据库
	//dataSourceName := "root:1234@(127.0.0.1:3306)/test"
	//mydb, _ := sql.Open("mysql", dataSourceName)
	//db, _ := gorm.Open(mysql.New(mysql.Config{
	//   Conn: mydb,
	//}), &gorm.Config{})
	// 连接方式2：使用gorm.io/driver/mysql连接数据库


	dsn := `root:1234@tcp(127.0.0.1:3306)/test?
			  charset=utf8mb4&parseTime=True&loc=Local`
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()


	// 关闭数据库，释放资源
	defer sqlDB.Close()



	// 设置连接池
	// SetMaxIdleConns设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	


	// 执行数据迁移
	db.AutoMigrate(&User{})
	


	// 创建数据
	u1 := User{Name: "Tom", Age: 10}
	db.Create(&u1)
	// 创建数据
	u2 := User{Name: "Tim"}
	// 批量创建
	db.Create(&u2)
	u3 := []User{{Name: "Lily", Age: 18},
				 {Name: "Lucy", Age: 22},
				 {Name: "Mary", Age: 20}}
	db.Create(&u3)
	


	// 查询数据
	// 创建数组对象l，数组元素为结构体User
	var l []User
	// 查询my_User的字段name不等于Tom的数据，并将结果写入数组对象l
	db.Where("name <> ?","Tom").First(&l)
	// 输出查询结果
	fmt.Printf("查询结果：%v\n", l)
	// Scan将查询结果转移到数组对象ls
	var ls []User
	db.Model(&User{}).Where("id = ?","1").Scan(&ls)
	// 上述查询方式等价于db.Where("id = ?","1").Find(&ls)
	fmt.Printf("查询结果：%v\n", ls)
	


	// 更新数据
	// Update用于更新某个字段的数据
	db.Where("id = ?","1").Find(&l).Update("name", "TomTom")
	// Updates用于批量更新（更新多列数据或多行数据）
	u4 := User{Name: "Jim", Age: 30}
	db.Model(&User{}).Where("id IN ?", []int{2, 3}).Updates(u4)
	


	// 删除数据是设置结构体成员DeletedAt，并不会真正删除数据
	// 因此执行数据查询会自动筛选结构体成员DeletedAt为Null的数据
	db.Where("name = ?","Jim").Delete(&User{})
	// 通过主键删除数据
	db.Delete(&User{}, []int{1, 5})
	// 使用Unscoped()能永久删除数据表的数据
	db.Unscoped().Where("name = ?","Lucy").Delete(&User{})
	

	
	// 执行原生的SQL语句
	var name string
	// 查询数据使用Raw()方法
	// 如果查询单行数据，使用Row()即可，如果查询多行数据，则使用Rows()
	db.Raw("select name from my_User where id=5").Row().Scan(&name)
	fmt.Printf("查询结果：%v\n", name)
	// 删除、创建或更新数据使用Exec()方法
	db.Exec("delete from my_User where id=1")
}