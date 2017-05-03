## beego mysql 语法糖


## 使用

- go get github.com/lazywe/sql
- import "github.com/lazywe/sql"

#### where 条件
````golang
 where := []sql.SqlWhere{
	sql.SqlWhere{"a","eq","666",And},
 	sql.SqlWhere{"b","neq","44",Or},
 	sql.SqlWhere{"b","in","44,33,5,1",Or},
	sql.SqlWhere{"c","find_in_set","6",Nil},
	sql.SqlWhere{"c","find_in_set_or","6,7,8",Nil},
 }
````
- where目前支持
````golang
    exp = map[string]string{
    	"eq":              "=",
    	"neq":             "!=",
    	"gt":              ">",
    	"egt":             ">=",
    	"lt":              "<",
    	"elt":             "<=",
    	"like":            "LIKE",
    	"notlike":         "NOT LIKE",
    	"in":              "IN",
    	"notin":           "NOT IN",
    	"find_in_set":     "FIND_IN_SET",
    	"find_in_set_and": "FIND_IN_SET_AND",
    	"find_in_set_or":  "FIND_IN_SET_OR",
    }
````

#### order
````golang
order := map[string]string{"id":"desc","sort":"desc"}
// id倒序，sort倒序
````


#### 获取总数
````golang
err,result := sql.table("default").Where([]SqlWhere{{"id","eq","1",Nil}}).GetCount("id")
// err bool
// result int64 总数量
````

#### 获取列表

````golang
err,result := sql.table("default").Where(where).Order(order).Limit("0,20").GetAll("*")

// err bool
// result []map[string]interface{}
````

#### 获取单条数据

````golang
err,result := sql.table("default").Where([]SqlWhere{{"id","eq","1",Nil}}).GetOne("*")

// err bool
// result map[string]interface{} 
````

#### 插入

````golang
datas := map[string]string{
    "title" : "孙中山",
    "age":"18",
}
err,result := sql.table("default").Insert(datas)
// err bool
// result int64 主键
````

#### 修改
````golang
datas := map[string]string{
    "title" : "孙中山",
    "age":"18",
}
 where := []sql.SqlWhere{
	sql.SqlWhere{"id","eq","1",Nil},
 }
/// 不要忘记添加where条件哦，后果很严重
err := sql.table("default").Where(where).Updates(datas)
// err bool
````

#### 删除
````golang

 where := []sql.SqlWhere{
	sql.SqlWhere{"id","eq","1",Nil},
 }
/// 不要忘记添加where条件哦，后果很严重
err := sql.table("default").Where(where).Delete()
// err bool
````


#### 鸣谢
- 要是真的感觉不错，给个star thanks
- qq：156919881