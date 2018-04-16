package helloworld

import (
	"github.com/gorilla/mux"
	"github.com/patst/alexa-skills-kit-for-go/alexa"
	"log"
	"net/http"
	"time"
)

var Skill alexa.Skill

func main() {
	router := mux.NewRouter()
	Skill = alexa.Skill{
		ApplicationId:  "FILL WITH SKILL ID", // Echo App ID from Amazon Dashboard
		OnIntent:       intentDispatchHandler,
		OnLaunch:       launchRequestHandler,
		OnSessionEnded: sessionEndedRequestHandler,
	}
	skillHandler := Skill.GetSkillHandler()

	router.Handle("/echo/api/trueorfalse", skillHandler).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting webserver")
	log.Fatal(srv.ListenAndServe())
}

func intentDispatchHandler(requestEnvelope *alexa.RequestEnvelope, request *alexa.IntentRequest, response *alexa.OutgoingResponse) {
	switch request.Intent.Name {
	case "HelloWorldIntent":
		helloWorldIntentHandler(request, response)
	case "AMAZON.StopIntent":
		cancelAndStopIntentHandler(request, response)
	case "AMAZON.CancelIntent":
		cancelAndStopIntentHandler(request, response)
	case "AMAZON.HelpIntent":
		helpIntentHandler(request, response)
	default:
		log.Println("Unknown intent!", request.Intent.Name)
	}
}

// HelloWorldIntent
func helloWorldIntentHandler(request *alexa.IntentRequest, outgoingResponse *alexa.OutgoingResponse) {
	speechText := "Hello world"
	outgoingResponse.Response.SetOutputSpeech(speechText)
	outgoingResponse.Response.SimpleCard("HelloWorld", speechText)
}

func cancelAndStopIntentHandler(request *alexa.IntentRequest, outgoingResponse *alexa.OutgoingResponse) {
	speechText := "Goodbye"
	outgoingResponse.Response.SetOutputSpeech(speechText)
	outgoingResponse.Response.SimpleCard("HelloWorld", speechText)
}

func helpIntentHandler(request *alexa.IntentRequest, outgoingResponse *alexa.OutgoingResponse) {
	speechText := "You can say hello to me!"
	outgoingResponse.Response.SetOutputSpeech(speechText)
	outgoingResponse.Response.SetReprompt(speechText)
	outgoingResponse.Response.SimpleCard("HelloWorld", speechText)
}

func launchRequestHandler(requestEnvelope *alexa.RequestEnvelope, request *alexa.LaunchRequest, outgoingResponse *alexa.OutgoingResponse) {
	speechText := "Welcome to the Alexa Skills Kit, you can say hello"
	outgoingResponse.Response.SetOutputSpeech(speechText)
	outgoingResponse.Response.SetReprompt(speechText)
	outgoingResponse.Response.SimpleCard("HelloWorld", speechText)
}

func sessionEndedRequestHandler(requestEnvelope *alexa.RequestEnvelope, request *alexa.SessionEndedRequest, outgoingResponse *alexa.OutgoingResponse) {
	// cleanup stuff
}