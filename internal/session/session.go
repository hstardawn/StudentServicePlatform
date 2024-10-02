package session

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// func Init(r *gin.Engine) {
// 	config := getConfig()
// 	switch config.Driver {
// 	case string(Redis):
// 		setRedis(r, config.Name)
// 		break
// 	case string(Memory):
// 		setMemory(r, config.Name)
// 		break
// 	default:
// 		log.Fatal("ConfigError")
// 	}

// }