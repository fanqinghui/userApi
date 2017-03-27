package controllers

import (
	_ "container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"userApi/models"
	"userApi/utils"
)

/**
请求方法分析：get:用户列表json展示，post:接收data数据，解析json，进行用户添加
**/
func UserHandle(rw http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		userArr := models.UserAll()
		userJson, err := json.Marshal(userArr)
		utils.CheckErr(err)
		os.Stdout.Write(userJson)
		fmt.Fprintln(rw, string(userJson))
	} else if req.Method == "POST" {

		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("params data error")
			return
		}
		jsonName := utils.DoJsonData(data)
		fmt.Println(jsonName)

		nameByte := []byte(jsonName)
		var user models.User
		err = json.Unmarshal(nameByte, &user)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintln(rw, "Unmarshal json error,may be post data error!")
			return
		}
		fmt.Println(user)
		user = models.InsertUser(user)
		fmt.Println(user)
		userJson, err := json.Marshal(user)
		utils.CheckErr(err)
		fmt.Fprintln(rw, string(userJson))
	}
}

/**
get：获取某个user所有关系用户
**/
func RelationGet(rw http.ResponseWriter, req *http.Request) {

	uid := req.FormValue(":user_id")
	id, err := strconv.Atoi(uid)
	if err != nil {
		fmt.Fprintln(rw, "uid must be int")
		return
	}
	relationArr := models.RelationsByUid(id)
	//查询所有的关系。组装成json
	relationJson, err := json.Marshal(relationArr)
	utils.CheckErr(err)
	os.Stdout.Write(relationJson)
	fmt.Fprintln(rw, string(relationJson))
}

/*
put，处理2用户关系
If two users have "liked" each other, then the state of the relationship is "matched"
*/
func RelationPut(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println(err)
	}
	//处理
	bodyStr := utils.DoJsonData(body)
	fmt.Println("body:" + bodyStr)
	body = []byte(bodyStr)
	var relationVo models.Relation
	err = json.Unmarshal(body, &relationVo)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(rw, "state json unmarshal err")
		return
	}
	state := relationVo.State //liked or disliked 校验
	fmt.Println(state)
	//	if state == "liked" || state == "disliked" {
	//	fmt.Fprintln(rw, "state can only be 'liked' or 'disliked'")
	//return
	//}

	uid := req.FormValue(":user_id")
	otherUid := req.FormValue(":other_user_id")

	user_id, err := strconv.Atoi(uid)
	if err != nil {
		fmt.Fprintln(rw, "user_id must be int")
		return
	}
	other_user_id, err := strconv.Atoi(otherUid)
	if err != nil {
		fmt.Fprintln(rw, "other_user_id must be int")
		return
	}

	fmt.Println("put请求-uid" + uid + " otherUid:" + otherUid + " state" + state)
	relation := models.PutUserRelations(user_id, other_user_id, state)
	relationJson, err := json.Marshal(relation)
	utils.CheckErr(err)
	os.Stdout.Write(relationJson)
	fmt.Fprintln(rw, string(relationJson))
}
