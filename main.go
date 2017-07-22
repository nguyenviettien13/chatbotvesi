package main

import (
	"github.com/nguyenviettien13/chatbotvesi/michlabs/fbbot"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	PAGEACCESS_TOKEN="EAAEo7uCJDOEBAG8KM7xR1HzMZCkmZA2gUoghyYtcANiaGaKaNXOhtbhNWA6Oxko9GOZAnCh4fyZATZAhTLZA9Jy5CIGmBXDp27sgiZB8otFADsAUlvZCEZAD2iTSqajBZBRy1IkZBHxU39avIxQHQSspVzMpGOrke4yayReyoLHx56D0VUZAdDuOymUz"
	VERIFY_TOKEN ="ratbimat"
	PORT =8081

)

func main()  {
	var a Announce
	bot:=fbbot.New(PORT,VERIFY_TOKEN,PAGEACCESS_TOKEN)
	bot.AddMessageHandler(a)
	bot.Run()

}

type Announce struct {

}

func(a Announce) HandleMessage(bot *fbbot.Bot, msg *fbbot.Message)  {

	//khoi tao request
	user_url:=msg.Text
	fmt.Println(user_url)
	req, err := http.NewRequest("GET", "http://api.openfpt.vn/cyradar?url="+user_url,nil)
	if err!=nil {
		fmt.Sprint("tao request hong")
	} else {
		fmt.Sprint("da tao xong request")
	}
	req.Header.Set("api_key","6631fdd937b547479fe036c5420863fc")

	//day request len server
	client := &http.Client{}
	res, err :=client.Do(req)



	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	//giai ma file json tra ve
	var data Tracks

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal("Failed to parse json", err.Error())
	}else{
		fmt.Println("parse xong du lieu")
		// fmt.Println(string(body))
		fmt.Println(data.Conclusion)
	}


	//tra loi nguoi dung
	var reply string
	switch data.Conclusion {
	case "danger":
		reply="nguy hiem"
	default:
		reply="an toan"
	}

	m:=fbbot.NewTextMessage(reply)
	bot.Send(msg.Sender,m)

}

type Tracks struct {

	Conclusion string   `json: "conclusion"`
	Domain     string   `json: "domain"`
	Threat     []string `json: "threat"`
	Uri        string   `json: "uri`
}

