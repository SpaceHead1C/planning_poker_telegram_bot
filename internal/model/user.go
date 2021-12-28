package model

const (
	userStatusUndefind  = iota
	userStatusVisitor   // not in room
	userStatusGamer     // in room
	userStatusInGame    // wait estimation process
	userStatusEstimator // estimates
)

const (
	userRoleUndefind = iota
	userRoleGamer
	userRoleScrumMaster
)

type User struct {
	ID           uint
	Name         string
	ActivityMark uint64
}
