package enum

type RoleType int8

const (
	_ RoleType = iota
	AdminRole
	UserRole
	VisitorRole
)
