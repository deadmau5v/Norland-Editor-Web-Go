package apis

import (
	"encoding/json"
	"math"
	"os"
	"reflect"

	"log/slog"

	"github.com/gin-gonic/gin"
)

var GameData []byte

// round 保留 3 位小数
func round(object map[string]interface{}) interface{} {
	for key, value := range object {
		if value == nil {
			continue
		}
		if reflect.TypeOf(value).Kind() == reflect.Map {
			object[key] = round(value.(map[string]interface{}))
		}
		if reflect.TypeOf(value).Kind() == reflect.Float64 ||
			reflect.TypeOf(value).Kind() == reflect.Float32 ||
			reflect.TypeOf(value).Kind() == reflect.Int ||
			reflect.TypeOf(value).Kind() == reflect.Int8 ||
			reflect.TypeOf(value).Kind() == reflect.Int16 ||
			reflect.TypeOf(value).Kind() == reflect.Int32 ||
			reflect.TypeOf(value).Kind() == reflect.Int64 {
			object[key] = math.Round(value.(float64)*1000) / 1000
		}
	}
	return object
}

// GetDefaultGameData 获取默认游戏配置
func GetDefaultGameData(c *gin.Context) {
	var response map[string]interface{}
	json.Unmarshal(GameData, &response)
	response = round(response).(map[string]interface{})
	c.JSON(200, response)
}

// init 初始化 加载默认游戏配置
func init() {
	var err error
	GameData, err = os.ReadFile("defualt.json")
	if err != nil {
		slog.Error("Error opening game data", "error", err)
	}
}
