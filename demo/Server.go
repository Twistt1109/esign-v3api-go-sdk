package demo

import "github.com/Twistt1109/esign-v3api-go-sdk"

func main() {

	client := NewESign(esign.NewClient(
		ESIGN_APP_ID,
		ESIGN_APP_SECRET,
		ESIGN_HOST,
	), "")

	client.GetVipSignUrl("", "", "", "", "")

	client.Download("")
}
