package global

import (
	"strconv"
	"io"
	"log"

	"github.com/5112100070/Trek/src/app/order"
	"github.com/5112100070/Trek/src/app/public"
	"github.com/5112100070/Trek/src/app/session"
	"github.com/5112100070/Trek/src/app/user"
)

func InitLogError(errorHandle io.Writer) {
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitRepoBundle(dbBundle DBBundle) {
	// init order status
	orderService := order.InitOrderRepo()
	orderStatus, errInitOrderStatus := orderService.GetListOrderStatusCGX()
	if errInitOrderStatus != nil {
		Error.Fatalln("func InitRepoBundle failed when init list order status. Error: ", errInitOrderStatus)
		return
	}

	var mapOrderStatus = make(map[int]string)
	for k,v := range orderStatus {
		key,_ := strconv.Atoi(k) 
		mapOrderStatus[key] = v
	}

	pickupStatus, errInitPickupStatus := orderService.GetListPickupStatusCGX()
	if errInitPickupStatus != nil {
		Error.Fatalln("func InitRepoBundle failed when init list pickup status. Error: ", errInitPickupStatus)
		return
	}

	var mapPickupStatus = make(map[int]string)
	for k,v := range pickupStatus {
		key,_ := strconv.Atoi(k) 
		mapPickupStatus[key] = v
	}	

	Services = RepoBundle{
		Session:        session.InitSessionRepo(dbBundle.RedisSession),
		Public:         public.InitPublicRepo(dbBundle.DB),
		User:           user.InitUserRepo(),
		Order:          orderService,
		MapOrderStatus: mapOrderStatus,
		MapPickupStatus:mapPickupStatus,
	}
}
