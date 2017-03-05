package server

import (
	"net/http"

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
	e.Server.Addr = ":8027"

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Network!\n")
	})

	// Group level middleware
	api := e.Group("/v1", middleware.BasicAuth(
		func(username, password string, c echo.Context) bool {
			if username == "alex" && password == "secret" {
				return true
			}
			return false
		}))

	// core
	api.GET("/len", server.len)
	api.GET("/keys", server.keys)
	api.GET("/stats", server.stats)
	// accessors - read
	api.GET("/get/:key", server.get)
	api.GET("/list/element/:key", server.getListElement)
	api.GET("/dict/element/:key", server.getDictElement)
	api.GET("/key/:key", server.hasKey)
	// mutators - create
	api.POST("/string/:key", server.setString)
	api.POST("/int/:key", server.setInt)
	api.POST("/list/:key", server.setList)
	api.POST("/dict/:key", server.setDict)
	api.POST("/ttl/:key", server.setTTL)
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

	return c.JSON(http.StatusBadRequest,
		util.BasicDTO{ErrorCode: errorCode, ErrorMessage: err.Error()})
}

///////////////////////////////////////
// Handlers
///////////////////////////////////////

/*
curl -i -w "\n" --user alex:juno localhost:1323/v1/len
result: {"error_code":0,"length":0}
*/
func (server *SERVER) len(c echo.Context) error {
	return c.JSON(http.StatusOK, util.LenDTO{Length: server.cache.Len()})
}

func (server *SERVER) keys(c echo.Context) error {
	return c.JSON(http.StatusOK, util.KeysDTO{Keys: server.cache.Keys()})
}

func (server *SERVER) stats(c echo.Context) error {
	return c.JSON(http.StatusOK, util.StatsDTO{Stats: server.cache.Stats()})
}

func (server *SERVER) get(c echo.Context) error {
	key := c.Param("key")
	value, err := server.cache.Get(key)

	if err != nil {
		return makeJSONError(c, err)
	}

	switch v := value.(type) {
	case int:
		return c.JSON(http.StatusOK, util.IntDTO{Value: v})
	case string:
		return c.JSON(http.StatusOK, util.StringDTO{Value: v})
	case util.List:
		return c.JSON(http.StatusOK, util.ListDTO{Value: v})
	case util.Dict:
		return c.JSON(http.StatusOK, util.DictDTO{Value: v})
	default:
		return makeJSONError(c, util.ErrorWrongType)
	}
}

func (server *SERVER) getListElement(c echo.Context) error {
	key := c.Param("key")

	value := new(util.IntDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	v, err := server.cache.GetListElement(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.StringDTO{Value: v.(string)})
}

func (server *SERVER) getDictElement(c echo.Context) error {
	key := c.Param("key")

	value := new(util.StringDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	v, err := server.cache.GetDictElement(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.StringDTO{Value: v.(string)})
}

func (server *SERVER) hasKey(c echo.Context) error {
	key := c.Param("key")

	v, err := server.cache.HasKey(key)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BoolDTO{Value: v})
}

func (server *SERVER) setString(c echo.Context) error {
	key := c.Param("key")

	value := new(util.StringDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	err := server.cache.SetString(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

func (server *SERVER) setInt(c echo.Context) error {
	key := c.Param("key")

	value := new(util.IntDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	err := server.cache.SetInt(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

func (server *SERVER) setList(c echo.Context) error {
	key := c.Param("key")

	value := new(util.ListDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	err := server.cache.SetList(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

func (server *SERVER) setDict(c echo.Context) error {
	key := c.Param("key")

	value := new(util.DictDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	err := server.cache.SetDict(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

func (server *SERVER) updateString(c echo.Context) error {
	key := c.Param("key")

	value := new(util.StringDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	v, err := server.cache.UpdateString(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.StringDTO{Value: v})
}

func (server *SERVER) updateInt(c echo.Context) error {
	key := c.Param("key")

	value := new(util.IntDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	v, err := server.cache.UpdateInt(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.IntDTO{Value: v})
}

func (server *SERVER) updateList(c echo.Context) error {
	key := c.Param("key")

	value := new(util.ListDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	v, err := server.cache.UpdateList(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.ListDTO{Value: v})
}

func (server *SERVER) updateDict(c echo.Context) error {
	key := c.Param("key")

	value := new(util.DictDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	v, err := server.cache.UpdateDict(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.DictDTO{Value: v})
}

func (server *SERVER) appendToList(c echo.Context) error {
	key := c.Param("key")

	value := new(util.StringDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	err := server.cache.AppendToList(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

func (server *SERVER) increment(c echo.Context) error {
	key := c.Param("key")

	v, err := server.cache.Increment(key)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.IntDTO{Value: v})
}

func (server *SERVER) remove(c echo.Context) error {
	key := c.Param("key")

	err := server.cache.Remove(key)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

func (server *SERVER) removeFromList(c echo.Context) error {
	key := c.Param("key")

	value := new(util.StringDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	v, err := server.cache.RemoveFromList(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.IntDTO{Value: v})
}

func (server *SERVER) removeFromDict(c echo.Context) error {
	key := c.Param("key")

	value := new(util.StringDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	err := server.cache.RemoveFromDict(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

func (server *SERVER) setTTL(c echo.Context) error {
	key := c.Param("key")

	value := new(util.IntDTO)
	if err := c.Bind(value); err != nil {
		return makeJSONError(c, err)
	}

	err := server.cache.SetTTL(key, value.Value)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}
