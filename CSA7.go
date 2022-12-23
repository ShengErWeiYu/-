package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Number   int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
}

type qas struct {
	QuestionID int    `gorm:"column:questionID;primaryKey;autoIncrement:true" json:"questionID"`
	Question   string `gorm:"column:question" json:"question"`
	Questioner string `gorm:"column:questioner" json:"questioner"`
	Answer     string `gorm:"column:answer" json:"answer"`
	Respondent string `gorm:"column:respondent" json:"respondent"`
}

var db *sql.DB

func initdb() {
	var err error
	dsn := "root:haoting39872256@tcp(127.0.0.1:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connect success")
}
func userRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	sqlStr := "select username from users where number>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u Users
		err := rows.Scan(&u.Username)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if username == u.Username {
			c.JSON(http.StatusOK, gin.H{"message": "该用户名已被注册"})
			return
		}
	}
	sqlStr = "insert into users(username,password) values (?,?)"
	_, err = db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
	}
}
func userLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	sqlStr := "select username,password from users where number>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "若页面无反应，则无该用户"})
	defer rows.Close()
	for rows.Next() {
		var u Users
		err := rows.Scan(&u.Username, &u.Password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if username == u.Username {
			if password == u.Password {
				c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "密码不正确,请重新登录"})
				return
			}
		}
	}
}
func userQuestion(c *gin.Context) {
	username := c.Query("username")
	question := c.Query("question")
	sqlStr := "insert into qas(questioner,question) values (?,?)"
	_, err := db.Exec(sqlStr, username, question)
	if err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "提问成功"})
		return
	}
}

func userAnswer(c *gin.Context) {
	username := c.Query("username")
	answer := c.Query("answer")
	questionID := c.Query("questionID")
	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	sqlStr := "update qas set answer = ?,respondent = ? where questionID = ?"
	_, err := db.Exec(sqlStr, answer, username, questionID)
	if err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "回答问题成功"})
		return
	}
}

func userGetQuestion(c *gin.Context) {
	username := c.Query("username")
	sqlStr := "select questioner,question,questionID from qas where questionID >?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var q qas
		err = rows.Scan(&q.Questioner, &q.Question, &q.QuestionID)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if username == q.Questioner {
			c.JSON(http.StatusOK, gin.H{
				"question":   q.Question,
				"questionID": q.QuestionID,
			})
		}
	}
}
func userGetAnswer(c *gin.Context) {
	username := c.Query("username")
	sqlStr := "select respondent,question,questionID,answer from qas where questionID >?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var q qas
		err = rows.Scan(&q.Respondent, &q.Question, &q.QuestionID, &q.Answer)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if username == q.Respondent {
			c.JSON(http.StatusOK, gin.H{
				"question":   q.Question,
				"questionID": q.QuestionID,
				"answer":     q.Answer,
			})
		}
	}
}
func userQuestionAlter(c *gin.Context) {
	username := c.Query("username")
	questionID := c.Query("questionID")
	newquestion := c.Query("newquestion")
	sqlStr := "update qas set question=? where questionID=? and questioner=?"
	_, err := db.Exec(sqlStr, newquestion, questionID, username)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "问题更改完毕!"})
}

func userAnswerAlter(c *gin.Context) {
	username := c.Query("username")
	questionID := c.Query("questionID")
	newanswer := c.Query("newanswer")
	sqlStr := "update qas set answer=? where questionID=? and respondent=?"
	_, err := db.Exec(sqlStr, newanswer, questionID, username)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "答案更改完毕!"})
}

func userAnswerDelete(c *gin.Context) {
	username := c.Query("username")
	questionID := c.Query("questionID")
	sqlStr := "update qas set answer='' where questionID=? and respondent=?"
	_, err := db.Exec(sqlStr, questionID, username)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "答案删除完毕!"})
}
func userQuestionDelete(c *gin.Context) {
	username := c.Query("username")
	questionID := c.Query("questionID")
	sqlStr := "update qas set question='' where questionID=? and questioner=?"
	_, err := db.Exec(sqlStr, questionID, username)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "问题删除完毕!"})
}
func main() {
	r := gin.Default()
	initdb()
	user := r.Group("/user")
	{
		user.GET("/register", userRegister)
		user.GET("/login", userLogin)
	}
	qa := r.Group("/qa")
	{
		qa.GET("/question", userQuestion)
		qa.GET("/answer", userAnswer)
		qa.GET("/getquestion", userGetQuestion)
		qa.GET("/getanswer", userGetAnswer)
		qa.GET("/questiondelete", userQuestionDelete)
		qa.GET("/answerdelete", userAnswerDelete)
		qa.GET("/questionalter", userQuestionAlter)
		qa.GET("/answeralter", userAnswerAlter)
	}
	r.Run(":8080")

}
