package sqlserver

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Demo() {
	// 安装和配置mssql
	// 1. 下载安装包，一键安装
	// 2. 虚拟机连接主机数据库时，使用虚拟机上`ip a`查到的`VMnet8`地址：192.168.232.1
	// 3. 在主机上，使用sqlserver配置管理器开启tcp/ip配置，并配置端口，重启生效
	// 4. 安装SSMS管理器，新增数据库和用户，并配置用户权限等

	// 打开连接
	// * `sqlserver://username:password@host/instance?param1=value&param2=value`
	// * `sqlserver://username:password@host:port?param1=value&param2=value`
	// * `sqlserver://sa@localhost/SQLExpress?database=master&connection+timeout=30` // `SQLExpress instance.
	// * `sqlserver://sa:mypass@localhost?database=master&connection+timeout=30`     // username=sa, password=mypass.
	// * `sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30` // port 1234 on localhost.
	// * `sqlserver://sa:my%7Bpass@somehost?connection+timeout=30` // password is "my{pass"
	//   A string of this format can be constructed using the `URL` type in the `net/url` package.
	connector, err := mssql.NewConnector(dsn())
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)

	// 关闭
	defer db.Close()

	// 读
	r, err := db.Exec("select * from dbo.[user]")
	if err != nil {
		panic(err)
	}
	n, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("n: %d\n", n)

	now := time.Now()

	// 写
	r, err = db.Exec(`insert into dbo.[user] (id, name, created_at) values (@id, @name, @created_at)`, sql.Named("id", 1), sql.Named("name", "jd"), sql.Named("created_at", now))
	if err != nil {
		panic(err)
	}
	n, err = r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("n: %d\n", n)

	// 读
	rows, err := db.Query("select id, name, created_at from dbo.[user]")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var id uint
	var name string
	var createdAt time.Time
	for rows.Next() {
		rows.Scan(&id, &name, &createdAt) // 变量顺序与Query的sql语句里的select字段保持一致
	}
	fmt.Printf("id: '%v', name: '%v', createdAt: '%v'\n", id, name, createdAt)

	// rows, err = db.Query(`select id, name, created_at from dbo.[user] where id = @id`, sql.Named("id", 1)) // 可以
	// rows, err = db.Query(`select id, name, created_at from dbo.[user] where created_at = @created_at`, sql.Named("created_at", now)) // 不可以
	rows, err = db.Query(`select id, name, created_at from dbo.[user] where created_at = @created_at`, sql.Named("created_at", now.Format("2006-01-02 15:04:05.000"))) // 可以，时间格式的最后部分不单'.000'可以，'.999'也可以
	// rows, err = db.Query(`select id, name, created_at from dbo.[user] where created_at = @created_at`, sql.Named("created_at", now.Format("2006-01-02 15:04:05"))) // 不可以，时间格式不对
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]User, 0, 8)
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.Name, &user.CreatedAt); err != nil { // 变量顺序与Query的sql语句里的select字段保持一致
			panic(err)
		}
		users = append(users, user)
	}
	fmt.Printf("now: %v, users: %d, '%+v'\n", now, len(users), users)
}

type User struct {
	Id        uint      `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func (User) TableName() string {
	return `user`
}

func DemoGorm() {
	gormDB, err := gorm.Open(sqlserver.Open(dsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.Debug()

	var user User
	if err = gormDB.Table("user").First(&user).Error; err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v\n", user)

	// now := time.Now()
	// user2 := User{
	// 	Id:        2,
	// 	Name:      "jj",
	// 	CreatedAt: now,
	// }
	// if err := gormDB.Create(&user2).Error; err != nil {
	// 	panic(err)
	// }

	var user3 User
	if err = gormDB.Table("user").Where("created_at = ?", user.CreatedAt).Scan(&user3).Error; err != nil {
		panic(err)
	}
	fmt.Printf("user3: %+v\n", user3)
}

func dsn() string {
	query := url.Values{}
	query.Add("database", "my_test")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("jd", "jd123JD"),
		Host:   fmt.Sprintf("%s:%d", "192.168.232.1", 1433),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	fmt.Printf("dsn: %v\n", u.String())

	return u.String()
}
