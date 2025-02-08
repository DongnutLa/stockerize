package constants

type UserRole string

const (
	SudoRole    UserRole = "SUDO"
	AdminRole   UserRole = "ADMIN"
	ManagerRole UserRole = "MANAGER"
)

var ROLES = []UserRole{"SUDO", "ADMIN", "MANAGER"}
