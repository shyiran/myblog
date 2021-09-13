package models

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"log"
	"myblog/common"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}
//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	//从配置文件中获取每页的文章数量
	num, err:= web.AppConfig.Int("articleListPageNum")
	if err!=nil{
		err.Error()
	}else{
		err.Error()
	}

	page--
	// 2页面， 每页显示2条 {3,4}
	// 3页面，每页显示5条  {11,12,13,14,15}
	//fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)
}
/**
分页查询数据库
limit分页查询语句，
    语法：limit m，n

    m代表从多少位开始获取，与id值无关
    n代表获取多少条数据

注意limit前面没有where
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}
func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from tb_article " + sql
	rows, err := common.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}
//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	// select tags from article
	rows, err := common.QueryDB(fmt.Sprintf("select %s from tb_articles", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}
//----------查询文章-------------
func QueryArticleWithId(id int) Article {
	row := common.QueryRowDB("select id,title,tags,short,content,author,createtime from tb_article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

//--------------按照标签查询--------------
/*
通过标签查询首页的数据
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号

通过%去匹配任意多个字符，至少是一个
*/
func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	//sql: like
	// tags: http&web&socket&互联网&计算机
	//       http&web
	//       web&socket&互联网&计算机
	//       web

	// http://localhost:8080?tag=web

	// %&web&%   %代表任何内容都可以匹配
	// %&web
	// web&%
	// web
	return QueryArticlesWithCon(sql)
}