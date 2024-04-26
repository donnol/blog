package from

import (
	"testing"

	"github.com/donnol/do"
)

func TestFromTable(t *testing.T) {
	item := UserTable{
		Id:         1,
		Name:       "jd",
		CreateTime: "2024-04-26",
		Phone:      "115",
	}
	roleIds := []do.Id{1}
	orgId := do.Id(1)
	postId := do.Id(1)

	var e User
	e.FromTable(item)
	e.WithRelation(roleIds, orgId, postId)
	do.Assert(t, e.Id, item.Id)
	do.Assert(t, e.Name, item.Name)
	do.Assert(t, e.Created, item.CreateTime)
	do.Assert(t, e.Phone, item.Phone)
	do.AssertSlice(t, e.RoleIds, roleIds)
	do.Assert(t, e.OrgId, orgId)
	do.Assert(t, e.PostId, postId)

	{
		item := UserTable{
			Id:         2,
			Name:       "jc",
			CreateTime: "2024-04-27",
			Phone:      "105",
		}
		roleIds := []do.Id{2}
		orgId := do.Id(2)
		postId := do.Id(2)

		var e User
		(&e).FromTable(item)
		(&e).WithRelation(roleIds, orgId, postId)
		do.Assert(t, e.Id, item.Id)
		do.Assert(t, e.Name, item.Name)
		do.Assert(t, e.Created, item.CreateTime)
		do.Assert(t, e.Phone, item.Phone)
		do.AssertSlice(t, e.RoleIds, roleIds)
		do.Assert(t, e.OrgId, orgId)
		do.Assert(t, e.PostId, postId)

	}
}
