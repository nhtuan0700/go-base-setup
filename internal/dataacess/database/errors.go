package database

type ErrorCode uint32

const (
	DBOK           ErrorCode = 0
	DBDataNotFound ErrorCode = 10
	DBGetFailed    ErrorCode = 20
	DBUpdateFailed ErrorCode = 21
	DBInsertFailed ErrorCode = 22
	DBDeleteFailed ErrorCode = 23
)
