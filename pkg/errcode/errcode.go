package errcode

const Success = 0 // 成功

// 用户模块错误码（1xxx）
const (
	ErrUserNotFound = iota + 1000
	ErrUserInvalidPassword
)

// 文章模块错误码（2xxx）
const (
	ErrPostsNotFound = iota + 2000
)
