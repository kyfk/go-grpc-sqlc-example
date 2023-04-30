package jwt

const UserIDClaimKey = "USER_ID"

type userIDCtxKey struct{}

var UserIDCtxKey = userIDCtxKey{}
