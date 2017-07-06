## dbobj示例

```shell
go get github.com/hzwy23/dbobj-example

cd $GOPATH/src/github.com/hzwy23/dbobj-example

```

1. 修改数据库连接信息，打开dbobj-example/conf/asofdate.conf，修改完成后保存。
2. 将表导入数据库，表结构信息在sql_script.sql中。

在dbobj-example目录中，执行下边命令

```shell
go run main.go
```

如果conf目录与执行程序不在同一个目录，则需要设置HBIGDATA_HOME环境变量


```go
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

```
