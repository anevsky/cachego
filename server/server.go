package server

import (
	"net/http"
	"strconv"

	"github.com/anevsky/cachego/memory"
	"github.com/anevsky/cachego/util"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type SERVER struct {
	cache memory.CACHE
}

func Create() SERVER {
	server := SERVER{
		cache: memory.Alloc(),
	}

	return server
}

func (server *SERVER) StartUp() {
	// Setup
	e := echo.New()
	e.Server.Addr = ":1323"

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Group level middleware
	api := e.Group("/v1", middleware.BasicAuth(func(username, password string, c echo.Context) bool {
		if username == "alex" && password == "juno" {
			return true
		}
		return false
	}))

	// core
	api.GET("/len", server.len)
	api.GET("/keys", server.keys)
	api.GET("/stats", server.stats)
	// accessors - read
	api.GET("/get", server.get)
	api.GET("/list/element/:key", server.getListElement)
	api.GET("/dict/element/:key", server.getDictElement)
	api.GET("/key", server.hasKey)
	// mutators - create
	api.POST("/string/:key", server.setString)
	api.POST("/int/:key", server.setInt)
	api.POST("/list/:key", server.setList)
	api.POST("/dict/:key", server.setDict)
	// mutators - update
	api.PUT("/string/:key", server.updateString)
	api.PUT("/int/:key", server.updateInt)
	api.PUT("/list/:key", server.updateList)
	api.PUT("/dict/:key", server.updateDict)
	api.PUT("/list/element/:key", server.appendToList)
	api.PUT("/int/increment/:key", server.increment)
	// mutators - delete
	api.DELETE("/remove/:key", server.remove)
	api.DELETE("/list/element/:key", server.removeFromList)
	api.DELETE("/dict/element:key", server.removeFromDict)

	// Serve it like a boss
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

func makeJSONError(c echo.Context, err error) error {
	var errorCode int
	switch err := err.(type) {
	case util.CacheError:
		errorCode = err.Code
	default:
		errorCode = util.ErrorBadRequest.Code
	}

	return c.JSON(http.StatusBadRequest, util.ResponseBasic{ErrorCode: errorCode, ErrorMessage: err.Error()})
}

///////////////////////////////////////
// Handlers
///////////////////////////////////////

/*
curl -i -w "\n" --user alex:juno localhost:1323/v1/len
result: {"error_code":0,"length":0}
*/
func (server *SERVER) len(c echo.Context) error {
	return c.JSON(http.StatusOK, util.ResponseLen{Length: server.cache.Len()})
}

func (server *SERVER) keys(c echo.Context) error {
	return c.JSON(http.StatusOK, util.ResponseKeys{Keys: server.cache.Keys()})
}

func (server *SERVER) stats(c echo.Context) error {
	return c.JSON(http.StatusOK, util.ResponseStats{Stats: server.cache.Stats()})
}

func (server *SERVER) get(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

/*
curl -i -w "\n" --user alex:juno localhost:1323/v1/list/element/testList?index=1
*/
func (server *SERVER) getListElement(c echo.Context) error {
	key := c.Param("key")
	index, _ := strconv.Atoi(c.QueryParam("index"))

	v, err := server.cache.GetListElement(key, index)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.ResponseString{Value: v.(string)})
}

func (server *SERVER) getDictElement(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) hasKey(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) setString(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) setInt(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) setList(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) setDict(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) updateString(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) updateInt(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) updateList(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) updateDict(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) appendToList(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) increment(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) remove(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) removeFromList(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}

func (server *SERVER) removeFromDict(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "n/a")
}
