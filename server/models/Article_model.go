/*
	Article表，存储文章
*/

package models

import (
	"Blog/server/database"
	"fmt"
	"time"
)

type Article struct {
	Id        int    //文章ID
	Title     string //文章标题
	Content   string //文章内容
	Ceatetime int64  //创建时间

}

//--------------数据库操作-----------------

//查询所有文章标题及内容
func QueryArticle_content() [][]string {
	sql := fmt.Sprintf("select * from article")
	fmt.Println(sql)
	rows := database.QueryAllRowDB(sql)
	var article_content [][]string
	for rows.Next() {
		id := ""
		title := ""
		content := ""
		createtime := ""
		rows.Scan(&id, &title, &content, &createtime)
		article_content = append(article_content, []string{id, title, content, createtime})
	}
	return article_content

}

//根据文章ID查询内容
func QueryArticle_contentWightId(id int) string {
	sql := fmt.Sprintf("select content from article where id='%d'", id)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	article_content := ""
	row.Scan(&article_content)
	return article_content
}

//插入一篇文章
func InsertArticle(article Article) (int64, error) {
	return database.ModifyDB("insert into aerticle(title,content,createtime) value (?,?,?)", article.Title, article.Content, time.Now().Unix())
}

//删除文章
func DeleteArticleWithID(id int) (int64, error) {
	return database.ModifyDB("delete from article where id = '%s'", id)
}
