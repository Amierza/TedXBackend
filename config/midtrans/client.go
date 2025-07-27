package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient snap.Client

func InitMidtransClient(serverKey string, isProduction bool) {
	env := midtrans.Sandbox
	if isProduction {
		env = midtrans.Production
	}

	SnapClient.New(serverKey, env)
}
