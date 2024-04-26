package from

import "github.com/donnol/do"

type User struct {
	Id      do.Id
	Name    string
	Created string
	Phone   string
	RoleIds []do.Id
	OrgId   do.Id
	PostId  do.Id
}

type UserTable struct {
	Id         do.Id
	Name       string
	CreateTime string
	Phone      string
}

func (p User) ToTable() UserTable {
	return UserTable{Name: p.Name, Phone: p.Phone}
}

func (p *User) FromTable(item UserTable) {
	p.Id = do.Id(item.Id)
	p.Name = item.Name
	p.Created = item.CreateTime
	p.Phone = item.Phone
}

func (p *User) WithRelation(roleIds []do.Id, orgId, postId do.Id) {
	p.RoleIds = roleIds
	p.OrgId = orgId
	p.PostId = postId
}

// ineffective assignment to field User.Name (SA4005)go-staticcheck
// func (p User) FromTable2(item UserTable) {
// 	p.Id = do.Id(item.Id)
// 	p.Name = item.Name
// 	p.Created = item.CreateTime
// 	p.Phone = item.Phone
// }
