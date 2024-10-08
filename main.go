package main

import (
	"fms/internal/interactor/pkg/connect"
	"fms/internal/router/client"
	"fms/internal/router/driver"
	"fms/internal/router/fleet"
	"fms/internal/router/gps"
	"fms/internal/router/gps_device"
	"fms/internal/router/login"
	"fms/internal/router/policy"
	"fms/internal/router/role"
	"fms/internal/router/status"
	"fms/internal/router/status_configuration"
	"fms/internal/router/subscription"
	"fms/internal/router/trailer"
	"fms/internal/router/transport_order"
	"fms/internal/router/transport_task"
	"fms/internal/router/user"
	"fms/internal/router/vehicle"
	"fmt"
	"net/http"

	_ "fms/api"
	"fms/internal/interactor/pkg/util/log"
	"fms/internal/router"

	//"fms/internal/router/permission"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// main is run all api form localhost port 8080

//	@title			FMS APIs
//	@version		0.1
//	@description	FMS APIs
//	@termsOfService

//	@contact.name
//	@contact.url
//	@contact.email

//	@license.name	AGPL 3.0
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.en.html

// @host fmp.t.api.jinher-net.com
// @BasePath	/fms
// @schemes	https
func main() {
	db, err := connect.PostgresSQL()
	if err != nil {
		log.Error(err)
		return
	}

	engine := router.Default()
	user.GetRouter(engine, db)
	login.GetRouter(engine, db)
	policy.GetRouter(engine, db)
	role.GetRouter(engine, db)
	fleet.GetRouter(engine, db)
	vehicle.GetRouter(engine, db)
	gps.GetRouter(engine, db)
	gps_device.GetRouter(engine, db)
	subscription.GetRouter(engine, db)
	driver.GetRouter(engine, db)
	status_configuration.GetRouter(engine, db)
	status.GetRouter(engine, db)
	trailer.GetRouter(engine, db)
	client.GetRouter(engine, db)
	transport_task.GetRouter(engine, db)
	transport_order.GetRouter(engine, db)
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8080", engine))
}
