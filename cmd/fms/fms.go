package main

import (
	"fms/internal/interactor/pkg/connect"
	"fms/internal/interactor/pkg/util/log"
	"fms/internal/router"
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

	"github.com/apex/gateway"
)

func main() {
	db, err := connect.PostgresSQL()
	if err != nil {
		log.Error(err)
		return
	}
	engine := router.Default()
	engine = user.GetRouter(engine, db)
	engine = login.GetRouter(engine, db)
	engine = policy.GetRouter(engine, db)
	engine = role.GetRouter(engine, db)
	engine = vehicle.GetRouter(engine, db)
	engine = gps.GetRouter(engine, db)
	engine = fleet.GetRouter(engine, db)
	engine = gps_device.GetRouter(engine, db)
	engine = subscription.GetRouter(engine, db)
	engine = driver.GetRouter(engine, db)
	engine = status_configuration.GetRouter(engine, db)
	engine = status.GetRouter(engine, db)
	engine = trailer.GetRouter(engine, db)
	engine = client.GetRouter(engine, db)
	engine = transport_task.GetRouter(engine, db)
	engine = transport_order.GetRouter(engine, db)
	log.Fatal(gateway.ListenAndServe(":8080", engine))
}
