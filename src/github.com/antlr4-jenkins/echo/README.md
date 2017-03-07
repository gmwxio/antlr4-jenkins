# Alexa config for Jenkins

Invocation Name:
```
jenkins
```

Schema
```
{
    "intents": [
        { "intent": "AMAZON.PauseIntent"  },
        { "intent": "AMAZON.ResumeIntent" },
     	{ "intent" : "Status",      "slots": []},
		{ "intent" : "Menu",        "slots": []},
		{ "intent" : "Bye"}
    ]
}
```

Utterances
```
Status status
Menu menu
Bye bye
```

Endpoint
```
https://jenkins.wx.io/echo/jenkins
```
Nginx proxy forwarding to Go http server

