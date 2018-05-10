package payment

// func TestAliyPay_GenderPayUrl(t *testing.T) {
// 	aliyConfig := map[string]interface{}{
// 		"app_id": "2088821583944105",
// 		"key":    "dn7tqiai8c1tut7wiqxh82j4t51k7jil",
// 	}
// 	aliy, err := NewPayment().Init(2, aliyConfig)
// 	oreder := Order{OrderID: "23232132322221", ProudctName: "I am test1", PriceTotal: 1, ProductID: 1000, IP: "127.0.0.1"}
// 	url, err := aliy.GenderPayUrl(oreder)
// 	fmt.Println(url, err)
// }

// func TestAliyPay_PayNotify(t *testing.T) {
// 	aliyConfig := map[string]interface{}{
// 		"partner":    "",
// 		"seller_id":  "",
// 		"key":        "",
// 		"notify_url": "",
// 	}
// 	notifyStr := ""
// 	aliy, err := NewPayment().Init(2, aliyConfig)
// 	res, err := aliy.PayNotify([]byte(notifyStr))
// 	fmt.Println("ok", err, string(res.ReturnData))
// }
