package gin

import (
	"context"
	a "github.com/core-go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type PrivilegesByEntityHandler struct {
	Load     func(ctx context.Context, id string) ([]a.Privilege, error)
	Error    func(context.Context, string)
	Offset   int
	Log      func(ctx context.Context, resource string, action string, success bool, desc string) error
	Resource string
	Action   string
}

func NewPrivilegesByEntityHandler(load func(ctx context.Context, id string) ([]a.Privilege, error), options ...func(context.Context, string)) *PrivilegesByEntityHandler {
	var logError func(context.Context, string)
	if len(options) >= 1 {
		logError = options[0]
	}
	return NewPrivilegesByEntityHandlerWithLog(load, logError, 0, nil)
}
func NewPrivilegesByEntityHandlerWithLog(load func(ctx context.Context, id string) ([]a.Privilege, error), logError func(context.Context, string), offset int, writeLog func(context.Context, string, string, bool, string) error, options ...string) *PrivilegesByEntityHandler {
	var resource, action string
	if len(options) >= 1 {
		resource = options[0]
	} else {
		resource = "privilege"
	}
	if len(options) >= 2 {
		action = options[1]
	} else {
		action = "all"
	}
	h := PrivilegesByEntityHandler{Load: load, Error: logError, Resource: resource, Action: action, Offset: offset, Log: writeLog}
	return &h
}
func (c *PrivilegesByEntityHandler) Privileges() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := ctx.Request
		id := ""
		if c.Offset <= 0 {
			i := strings.LastIndex(r.RequestURI, "/")
			if i >= 0 {
				id = r.RequestURI[i+1:]
			}
		} else {
			s := strings.Split(r.RequestURI, "/")
			if len(s)-c.Offset-1 >= 0 {
				id = s[len(s)-c.Offset-1]
			} else {
				ctx.String(http.StatusBadRequest, "URL is not valid")
				return
			}
		}
		privileges, err := c.Load(r.Context(), id)
		if err != nil {
			if c.Error != nil {
				c.Error(r.Context(), err.Error())
			}
			respond(ctx, http.StatusInternalServerError, internalServerError, c.Log, c.Resource, c.Action, false, err.Error())
		} else {
			respond(ctx, http.StatusOK, privileges, c.Log, c.Resource, c.Action, true, "")
		}
	}
}
