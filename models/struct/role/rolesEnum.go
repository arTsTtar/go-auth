package role

type Role struct {
	slug string
}

func (r Role) String() string {
	return r.slug
}

var prefix = "ROLE_"

var (
	Admin     = Role{prefix + "ADMIN"}
	Client    = Role{prefix + "CLIENT"}
	Moderator = Role{prefix + "MODERATOR"}
	Unknown   = Role{prefix + "UNKNOWN"}
)
