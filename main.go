package main

import (
	"gin-july/common"
	"gin-july/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	r := gin.Default()

	// config
	common.InitConfig()

	// logging
	common.InitLogrus()

	// db
	dbs := common.InitDB()
	for _, db := range dbs {
		defer db.Close()
	}

	// router
	r = router.CollectApiRoute(r)

	// port
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	// run
	//panic(r.Run()) // listen and serve on 0.0.0.0:8080

	if err := endless.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
