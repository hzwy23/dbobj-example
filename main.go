package main

import (
	"fmt"

	"github.com/hzwy23/dbobj"
)

type UserInfo struct {
	UserId   string
	UserName string
}

func GetUserDetails(age int) ([]UserInfo, error) {
	rows, err := dbobj.Query("select user_id,user_name from dbobj_test_table where age = ?", age)
	if err != nil {
		fmt.Println("query table failed", err)
		return nil, err
	}
	var rst []UserInfo
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		fmt.Println("scan table failed", err)
		return nil, err
	}
	return rst, nil
}

func GetUserDetails2(age int, rst *[]UserInfo) error {
	return dbobj.QueryForSlice("select user_id,user_name from dbobj_test_table where age = ?", rst, age)
}

func GetUserDetails3(userId string, rst *UserInfo) error {
	return dbobj.QueryForStruct("select user_id,user_name from dbobj_test_table where user_id = ?", rst, userId)
}

func GetUserDetails4(userId string, args ...interface{}) error {
	rows, err := dbobj.Query("select user_id,user_name from dbobj_test_table where user_id = ?", userId)
	if err != nil {
		fmt.Println("query table failed", err)
		return err
	}
	return dbobj.ScanRow(rows, args...)
}

func GetUserDetails5(args []interface{}, result ...interface{}) error {
	return dbobj.QueryForObject("select user_id,user_name from dbobj_test_table where user_id = ? and age = ?", args, result...)
}

func main() {
	tmp1, err := GetUserDetails(12)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success", tmp1)

	var rst []UserInfo
	err = GetUserDetails2(12, &rst)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok,", rst)

	var obj UserInfo
	err = GetUserDetails3("China", &obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("three:", obj)

	var id = ""
	var name = ""
	err = GetUserDetails4("China", &id, &name)
	fmt.Println(err, id, name)
	fmt.Println("***********************************")

	id = ""
	name = ""
	err = GetUserDetails5(dbobj.PackArgs("China", 12), &id, &name)
	fmt.Println(err, id, name)
}
