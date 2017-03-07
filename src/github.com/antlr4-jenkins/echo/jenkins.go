package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

type JenkinRest struct {
	Color string
}

var Applications = map[string]interface{}{
	"/echo/jenkins": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.660a97fd-8509-438a-b34d-ad2fab338841", // Echo App ID from Amazon Dashboard
		OnIntent: JenkinsIntentHandler,
		OnLaunch: JenkinsLaunchHandler,
		// OnSessionEnded: EchoSessionEndHandler,
		// Handler:  EchoIntentHandler2,
	},
}

func main() {
	alexa.Run(Applications, "3001")
}

func JenkinsLaunchHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	fmt.Printf("EchoLaunchHandler %v\n", echoReq)
	echoResp.OutputSpeechSSML("<speak>You <say-as interpret-as='interjection'>rang</say-as>. Awaiting your <w role='ivona:VB'>command</w> or say <w role='ivona:NN'>menu</w> for help </speak>").EndSession(false)
}

func JenkinsIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	fmt.Printf("EchoIntentHandler %v\n", echoReq)
	intent := echoReq.GetIntentName()
	switch intent {
	case "Status":
		f, err := http.Get("http://jenkins.wx.io:8080/job/antlr4-benchmark/api/json")
		// f, err := http.Get("http://jenkins.wx.io:8080/job/antlr4-benchmark/api/json?pretty=true")
		if err != nil {
			fmt.Printf("URL Get err %v\n", err)
			echoResp.OutputSpeech("can't reach build server").EndSession(false)
			return
		}
		d := json.NewDecoder(f.Body)
		resp := &JenkinRest{}
		err = d.Decode(resp)
		if err != nil {
			fmt.Printf("decode %v\n", err)
			echoResp.OutputSpeech("can't decode message").EndSession(false)
			return
		}
		echoResp.OutputSpeech("The last build was " + resp.Color).EndSession(false)
		fmt.Printf("color %v\n", resp.Color)
	case "Menu":
		echoResp.OutputSpeechSSML("<speak><say-as interpret-as='interjection'>Seriously</say-as> you actually thought that was implemented</speak>").
			EndSession(false)
	case "Bye":
		echoResp.OutputSpeech("Seeya later.").EndSession(true)
	}
}
