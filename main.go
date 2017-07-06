package main

import (
	"fmt"

	"github.com/hzwy23/dbobj"
)

type UserInfo struct {
	UserId   string
	UserName string
}

func GetUserDetails(userId int) ([]UserInfo, error) {
	rows, err := dbobj.Query("select user_id,user_name from dbobj_test_table where age = ?", userId)
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

func main() {
	tmp1, err := GetUserDetails(12)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success", tmp1)
}
