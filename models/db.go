package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"userApi/utils"
)

func OpenDB() *sql.DB {
	//db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.10:3306)/gotest?charset=utf8")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gotest?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

type User struct {
	Id   int
	Name string
	Type string
}
type Relationship struct {
	Id       int
	Uid      int
	OtherUid int
	State    string
}

type Relation struct {
	Id    int
	State string
	Type  string
}

func UserAll() []User {
	db := OpenDB()
	defer db.Close()
	//userList := list.New()
	rows, err := db.Query("SELECT * FROM users")
	db.query
	userArr := make([]User, 0)

	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.Id, &user.Name)
		utils.CheckErr(err)
		user.Type = "user"
		//sigleUser := make([]User, 1)
		//userList.PushBack(user)
		//sigleUser[0] = *user
		userArr = append(userArr, *user)
		//fmt.Println(userArr)
	}
	return userArr
}

func InsertUser(user User) User {
	db := OpenDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT users (name) values (?)")
	defer stmt.Close()
	utils.CheckErr(err)
	res, err := stmt.Exec(user.Name)
	utils.CheckErr(err)
	id, err := res.LastInsertId()
	utils.CheckErr(err)
	fmt.Println(id)
	user.Id = int(id)
	user.Type = "user"
	return user
}

func RelationsByUid(uid int) []Relation {
	db := OpenDB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT u.id,r.state FROM relations r INNER JOIN users u ON r.uid=u.id WHERE r.otherUid=?")
	defer stmt.Close()
	utils.CheckErr(err)
	rows, err := stmt.Query(uid)
	utils.CheckErr(err)
	relations := make([]Relation, 0)
	for rows.Next() {
		relation := new(Relation)
		err = rows.Scan(&relation.Id, &relation.State)
		utils.CheckErr(err)
		relation.Type = "relationship"
		relations = append(relations, *relation)
	}
	return relations
}

func PutUserRelations(user_id int, other_user_id int, state string) Relation {
	db := OpenDB()
	defer db.Close()

	//step1 判断like 还是disliked
	if state == "liked" {
		fmt.Println("liked")
		stmp, err := db.Prepare("select r.state from relations r where r.uid=? and r.otherUid=? order by id desc limit 1 ")
		defer stmp.Close()
		utils.CheckErr(err)
		row := stmp.QueryRow(other_user_id, user_id)

		var stateResult string
		row.Scan(&stateResult)
		fmt.Println("liked stateresult:" + stateResult)
		if stateResult == "liked" {
			state = "matched"
			stmp, err = db.Prepare("update relations set state='matched' where uid=? and otherUid=?")
			utils.CheckErr(err)
			stmp.Exec(other_user_id, user_id)
		}
		fmt.Println("liked end")
	} else if state == "disliked" {
		fmt.Println("disliked")
		stmp, err := db.Prepare("select r.state from relations r where r.uid=? and r.otherUid=? order by id desc limit 1")
		stmp.Close()
		row := stmp.QueryRow(other_user_id, user_id)
		utils.CheckErr(err)
		var stateResult string
		row.Scan(&stateResult)
		fmt.Println("disliked stateresult:" + stateResult)
		if stateResult == "matched" {
			stmp, err = db.Prepare("update relations set state='liked' where uid=? and otherUid=?")
			utils.CheckErr(err)
			stmp.Exec(other_user_id, user_id)
		}
		fmt.Println("disliked end")
	}

	//step2 插入关系表

	stmp, err := db.Prepare("select r.id from relations r where r.uid=? and r.otherUid=? ")
	defer stmp.Close()
	row := stmp.QueryRow(user_id, other_user_id)
	utils.CheckErr(err)
	var objId int
	row.Scan(&objId)

	if objId > 0 {
		fmt.Println("objId:" + string(objId))
		//update
		stmp, err = db.Prepare("update relations set state=? where id=?")
		utils.CheckErr(err)
		stmp.Exec(state, objId)
	} else {
		stmt, err := db.Prepare("INSERT relations (uid,otherUid,state) values (?,?,?)")
		utils.CheckErr(err)
		res, err := stmt.Exec(user_id, other_user_id, state)
		id, err := res.LastInsertId()
		utils.CheckErr(err)
		fmt.Println(id)
	}

	//step 3构建结果result
	relation := Relation{
		Id:    other_user_id,
		State: state,
		Type:  "relationship",
	}
	return relation
}
