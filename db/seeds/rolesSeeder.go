package seeds

import (
	"go-auth/entity"
	"go-auth/models/struct/role"
)

func (s Seed) RoleSeed() {

	s.db.FirstOrCreate(&entity.Role{}, entity.Role{Name: role.Admin.String()})
	s.db.FirstOrCreate(&entity.Role{}, entity.Role{Name: role.Client.String()})
	s.db.FirstOrCreate(&entity.Role{}, entity.Role{Name: role.Moderator.String()})
	s.db.FirstOrCreate(&entity.Role{}, entity.Role{Name: role.Unknown.String()})
}
