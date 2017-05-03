package sql

import (
	"testing"

	"fmt"
)

var (
	c = new(Sql)
)

func Test_Init(t *testing.T) {
	fmt.Println("sdf")
}

func Test_BuildWhere(t *testing.T) {
	where := []SqlWhere{
		SqlWhere{"a", "eq", "666", And},
		SqlWhere{"b", "neq", "44", Or},
		SqlWhere{"b", "gt", "44", Or},
		SqlWhere{"b", "lt", "44", Or},
		SqlWhere{"b", "elt", "44", Or},
		SqlWhere{"b", "egt", "44", Or},
		SqlWhere{"h", "like", "44", Or},
		SqlWhere{"c", "find_in_set", "6", And},
		SqlWhere{"dc", "in", "6,5,6,87,8", Or},
		SqlWhere{"c", "find_in_set", "7", Nil},
	}
	result := c.Where(where)
	t.Log(result)
}

func Test_GetAll(t *testing.T) {
	where := []SqlWhere{
		SqlWhere{"area", "eq", "2", And},
		SqlWhere{"cate_gory", "find_in_set_or", "5,4", Nil},
	}
	order := map[string]string{
		"sort": "desc",
	}
	err, result := c.Table("mi_video").Where(where).Order(order).GetAll("id")
	t.Log(err)
	t.Log(result)
}

func Test_GetOne(t *testing.T) {
	where := []SqlWhere{
		SqlWhere{"area", "eq", "2", And},
		SqlWhere{"cate_gory", "find_in_set_or", "5,4", Nil},
	}
	err, result := c.Table("mi_video").Where(where).GetOne("id")
	t.Log(err)
	t.Log(result)
}

func Test_GetCount(t *testing.T) {
	where := []SqlWhere{
		SqlWhere{"area", "eq", "2", And},
		SqlWhere{"cate_gory", "find_in_set_or", "5,4", Nil},
	}
	err, result := c.Table("mi_video").Where(where).GetCount("id")
	t.Log(err)
	t.Log(result)
}

func Test_buildCondition(t *testing.T) {

	result := c.buildCondition("a", "and", "1")
	t.Log(result)
}
