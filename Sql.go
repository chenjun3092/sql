package sql

import (
	"strings"

	"fmt"

	"github.com/astaxie/beego/orm"
)

/**
 * 设置数据表
 * @table string 数据表名称
 * @return *Sql
 */
func (s *Sql) Table(table string) *Sql {
	s.table = table
	return s
}

/**
 * 获取所有数据
 * @param string field 字段，跟原生写法一样 id as id
 * @demo sql.table("default").Where([]SqlWhere{{"id","eq","1",Nil}}).GetAll("*")
 * @return bool,[]orm.Params
 */
func (s *Sql) GetAll(field string) (bool, []orm.Params) {
	o := orm.NewOrm()
	maps := []orm.Params{}
	/// sql 构造
	sql := "SELECT " + field + " FROM " + s.table + s.where + s.orderBy + s.limit
	_, err := o.Raw(sql).Values(&maps)
	return err != orm.ErrNoRows, maps
}

/**
 * 获取数据总条数
 * @param string field 字段，跟原生写法一样 id as id
 * @demo sql.table("default").Where([]SqlWhere{{"id","eq","1",Nil}}).GetCount("id")
 * @return bool,int64
 */
func (s *Sql) GetCount(field string) (bool, int64) {
	o := orm.NewOrm()
	/// sql 构造
	maps := []orm.Params{}
	sql := "SELECT " + field + " FROM " + s.table + s.where + s.limit
	num, err := o.Raw(sql).Values(&maps)
	return err != orm.ErrNoRows, num
}

/**
 * 获取单条数据
 * @param string field 字段，跟原生写法一样 id as id
 * @demo sql.table("default").Where([]SqlWhere{{"id","eq","1",Nil}}).GetOne("*")
 * @return bool,orm.Params
 */
func (s *Sql) GetOne(field string) (bool, orm.Params) {
	o := orm.NewOrm()
	maps := []orm.Params{}
	/// sql 构造
	sql := "SELECT " + field + " FROM " + s.table + s.where + s.orderBy + " limit 0,1"
	_, err := o.Raw(sql).Values(&maps)
	if len(maps) > 0 {
		return err != orm.ErrNoRows, maps[0]
	} else {
		return err != orm.ErrNoRows, nil
	}
}

/**
 * 修改
 * @param map[string]string datas 插入的字段
 * @return bool
 */
func (s *Sql) Update(datas map[string]string) bool {
	if len(datas) < 0 {
		return false
	}
	datastring := ""
	for k, v := range datas {
		if datastring != "" {
			datastring += ", "
		}
		// datastring += k + " = " + v
		datastring += fmt.Sprintf("%s = '%s'", k, v)
	}
	o := orm.NewOrm()
	sql := "UPDATE " + s.table + " set " + datastring + s.where
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return false
	}
	return true
}

/**
 * 字段增加数值
 *
 * @param string  key 增加的字段
 * @param int  var	  增加的值
 * @return bool
 */
func (s *Sql) Inc(key string, val int) bool {
	datastring := ""
	datastring += fmt.Sprintf("%s = %s+%d", key, key, val)
	o := orm.NewOrm()
	sql := "UPDATE " + s.table + " set " + datastring + s.where
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return false
	}
	return true
}

/**
 * 删除
 * @return bool
 */
func (s *Sql) Delete() bool {
	if s.where == "" {
		return false
	}
	o := orm.NewOrm()
	sql := fmt.Sprintf("DELETE from %s %s", s.table, s.where)
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return false
	}
	return true
}

/**
 * 新增
 * @param map[string]string datas key字段，value值
 * @return *Sql
 */
func (s *Sql) Insert(datas map[string]string) (bool, int64) {
	if len(datas) < 0 {
		return false, 0
	}
	datakey := ""
	datavalue := ""
	for k, v := range datas {
		if datakey != "" {
			datakey += ", "
			datavalue += ", "
		}
		datakey += k
		datavalue += fmt.Sprintf("'%s'", v)
	}
	o := orm.NewOrm()
	sql := fmt.Sprintf("Insert into %s (%s)"+" VALUES (%s)", s.table, datakey, datavalue)
	result, err := o.Raw(sql).Exec()
	if err != nil {
		return false, 0
	}
	num, err2 := result.LastInsertId()
	if err2 != nil {
		return false, 0
	}
	return true, num
}

/**
 * limit
 * @param string str limit字符 0,10       0偏移量，10查询数量
 * @return *Sql
 */
func (s *Sql) Limit(offset string, pagesize string) *Sql {
	if offset != "" && pagesize != "" {
		s.limit = fmt.Sprintf(" limit %s,%s", offset, pagesize)
	}
	return s
}

/**
 * order
 * @param map[string]string maps order
 * @demo map[string]string{"id":"desc","sort":"desc"}
 * @return *Sql
 */
func (s *Sql) Order(maps map[string]string) *Sql {
	order := ""
	if len(maps) > 0 {
		for k, v := range maps {
			if order != "" {
				order += ","
			}
			order += k + " " + v
		}
		order = " order by " + order
		s.orderBy = order
	}
	return s
}

/**
 * 根据数组构建sqlwhere*
 * @param []string 搜索数据接口
 * @demo
 * where := []SqlWhere{
 *		SqlWhere{"a","eq","666",And},
 *		SqlWhere{"b","neq","44",Or},
 *		SqlWhere{"c","find_in_set","6",Nil},
 * }
 * @return *Sql
 **/
func (s *Sql) Where(where []SqlWhere) *Sql {
	if where != nil && len(where) > 0 {
		wherestring := " where "
		for k, v := range where {
			result := s.buildCondition(v.Field, v.Condition, v.Value)
			wherestring += result
			if k+1 != len(where) {
				wherestring += v.Operation.String()
			}
			if v.Operation == Nil {
				break
			}
		}
		s.where = wherestring
	}
	return s
}

/**
 * 自定义条件类型
 * @param string field 条件的字段
 * @param string condition 条件类型
 * @param string value 条件的值，此处可扩展
 * @return string
 **/
func (s *Sql) buildCondition(field string, condition string, value string) string {
	result := ""
	switch condition {
	case "eq", "neq", "gt", "lt", "egt", "elt", "like", "notlike":
		result += fmt.Sprintf("%s %s '%s'", field, exp[condition], value)
	case "in", "notin":
		// result += field + " IN(" + value + ")"

		valuearr := strings.Split(value, ",")
		valuestring := strings.Join(valuearr, "','")
		result += fmt.Sprintf("%s %s('%s')", field, exp[condition], valuestring)

	case "find_in_set":
		result += "FIND_IN_SET('" + value + "'," + field + ")"
	case "find_in_set_or":
		arr := strings.Split(value, ",")
		result += "( "
		for k, v := range arr {
			if k != 0 {
				result += " or "
			}
			result += "FIND_IN_SET('" + v + "'," + field + ")"
		}
		result += ")"
	case "find_in_set_and":
		arr := strings.Split(value, ",")
		result += "( "
		for k, v := range arr {
			if k != 0 {
				result += " and "
			}
			result += "FIND_IN_SET('" + v + "'," + field + ")"
		}
		result += ")"
	default:
		return value
	}
	return result
}
