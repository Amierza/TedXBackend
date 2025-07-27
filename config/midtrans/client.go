package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient snap.Client

func InitMidtransClient(serverKey string, envType string) {
	env := midtrans.Sandbox
	if envType == "production" {
		env = midtrans.Production
	}

	SnapClient.New(serverKey, env)
}
