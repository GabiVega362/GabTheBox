package routes

import "github.com/gin-gonic/gin"

func GetSessionId(gctx *gin.Context) (string, bool) {
	sessid, err := gctx.Cookie(("GTBSESSID"))
	if err != nil {
		return "", false
	}
	return sessid, true
}

func IsAuthenticated(gctx *gin.Context) bool {
	_, ok := GetSessionId(gctx)
	return ok
}