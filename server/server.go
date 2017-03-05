package server

import (
	"net/http"

	"github.com/anevsky/cachego/memory"
	"github.com/anevsky/cachego/util"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server with cache
type SERVER struct {
	cache memory.CACHE
}

// Allocate server instance
func Create() SERVER {
	server := SERVER{
		cache: memory.Alloc(),
	}

	return server
}

// Setup and start a server
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
	api.GET("/key/:key", server.hasKey)
	api.POST("/list/element/:key", server.getListElement)
	api.POST("/dict/element/:key", server.getDictElement)
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
	api.DELETE("/dict/element/:key", server.removeFromDict)

	// Serve it like a boss
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

// Tramsform error object to JSON response
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

// Get total number of objects
// curl -i -w "\n" --user alex:secret localhost:8027/v1/len
func (server *SERVER) len(c echo.Context) error {
	return c.JSON(http.StatusOK, util.LenDTO{Length: server.cache.Len()})
}

// Get list of keys
// curl -i -w "\n" --user alex:secret localhost:8027/v1/keys
func (server *SERVER) keys(c echo.Context) error {
	return c.JSON(http.StatusOK, util.KeysDTO{Keys: server.cache.Keys()})
}

// Get cache stats
// curl -i -w "\n" --user alex:secret localhost:8027/v1/stats
func (server *SERVER) stats(c echo.Context) error {
	return c.JSON(http.StatusOK, util.StatsDTO{Stats: server.cache.Stats()})
}

// Get value from cache by key
// Auto type conversion
// Returns value or ErrorWrongType if not supported value type
// curl -i -w "\n" --user alex:secret localhost:8027/v1/get/vvv
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

// Get element from list by index
// curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":1}' localhost:8027/v1/list/element/lll
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

	return c.JSON(http.StatusOK, util.StringDTO{Value: v})
}

// Get element from dict by key
// curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":"k1"}' localhost:8027/v1/dict/element/ddd
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

	return c.JSON(http.StatusOK, util.StringDTO{Value: v})
}

// Check if object exists in cache by key
// curl -i -w "\n" --user alex:secret localhost:8027/v1/key/lll
func (server *SERVER) hasKey(c echo.Context) error {
	key := c.Param("key")

	v, err := server.cache.HasKey(key)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BoolDTO{Value: v})
}

// Set string
// curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":"s1"}' localhost:8027/v1/string/sss
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

// Set int
// curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":121}' localhost:8027/v1/int/iii
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

// Set list
// curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":["aa", "bb"]}' localhost:8027/v1/list/lll
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

// Set dict
// curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":{"k1": "v1", "k2": "v2"}}' localhost:8027/v1/dict/ddd
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

// Update string by key
// Returns old value or ErrorKeyNotFound
// curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":"s2"}' localhost:8027/v1/string/sss
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

// Update int by key
// Returns old value or ErrorKeyNotFound
// curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":123}' localhost:8027/v1/int/iii
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

// Update list by key
// Returns old value or ErrorKeyNotFound
// curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":["aa2", "bb2"]}' localhost:8027/v1/list/lll
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

// Update dict by key
// Returns old value or ErrorKeyNotFound
// curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":{"k12": "v12", "k22": "v22"}}' localhost:8027/v1/dict/ddd
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

// Append to list a string element
// Might return ErrorKeyNotFound or ErrorWrongType
// curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":"aa3"}' localhost:8027/v1/list/element/lll
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

// Increment an integer value by key
// Returns new value
// curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json'  localhost:8027/v1/int/increment/iii
func (server *SERVER) increment(c echo.Context) error {
	key := c.Param("key")

	v, err := server.cache.Increment(key)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.IntDTO{Value: v})
}

// Remove object from cache by key
// curl -i -w "\n" -X DELETE --user alex:secret -H 'Content-Type: application/json'  localhost:8027/v1/remove/iii
func (server *SERVER) remove(c echo.Context) error {
	key := c.Param("key")

	err := server.cache.Remove(key)
	if err != nil {
		return makeJSONError(c, err)
	}

	return c.JSON(http.StatusOK, util.BasicDTO{})
}

// Remove object from list by value
// Returns the index of removed element or -1 if value not found in list
// Might return ErrorKeyNotFound, ErrorWrongType
// curl -i -w "\n" -X DELETE --user alex:secret -H 'Content-Type: application/json' -d '{"value":"aa3"}' localhost:8027/v1/list/element/lll
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

// Remove object from dict by key
// curl -i -w "\n" -X DELETE --user alex:secret -H 'Content-Type: application/json' -d '{"value":"k12"}' localhost:8027/v1/dict/element/ddd
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

// Set TTL (time-to-live) in nanoseconds for object by key
// curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":5211}' localhost:8027/v1/ttl/iii
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
