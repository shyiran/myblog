package common

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"github.com/russross/blackfriday"
	"github.com/PuerkitoBio/goquery"

)

var db *sql.DB

func InitMysql() {
	mysqlDriver,_ :=web.AppConfig.String("driverName")
	user ,_:= web.AppConfig.String("mysqluser")
	pwd,_ := web.AppConfig.String("mysqlpass")
	host,_ := web.AppConfig.String("mysqlurls")
	port,_ := web.AppConfig.String("port")
	dbname,_:= web.AppConfig.String("mysqldb")
	dbCoon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db1, err := sql.Open(mysqlDriver, dbCoon)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		//创建用户表
		CreateTableWithUser()
		//创建文章表
		CreateTableWithArticle()
		//创建相册数据表
		CreateTableWithAlbum()
	}
}
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS tb_user(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}
//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists tb_articles(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}
//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists tb_album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}
//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
/**
 * 将文章详情的内容，转换成HTMl语句
 */
func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))

	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		//light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		light:= selection.Text()
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}