package alexa

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
)

// LambdaHandler interface which a lambda handler must fulfil.
type LambdaHandler func(ctx context.Context, event interface{}) (interface{}, error)

// GetLambdaSkillHandler provides a handler which can be used in a lambda function.
func (skill *Skill) GetLambdaSkillHandler() LambdaHandler {
	return func(ctx context.Context, event interface{}) (interface{}, error) {
		bodyBytes, _ := json.Marshal(event)

		var requestEnvelope *RequestEnvelope
		err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&requestEnvelope)

		if skill.Verbose {
			log.Println("--> Request: ", string(bodyBytes))
		}

		if err != nil {
			return nil, err
		}
		if !skill.SkipValidation {
			if err = requestEnvelope.isRequestValid(skill.ApplicationID); err != nil {
				return nil, err
			}
		}

		response, err := requestEnvelope.handleRequest(skill)

		if err != nil {
			return nil, err
		}

		if skill.Verbose {
			json, err := json.Marshal(response)
			if err != nil {
				return nil, err
			}
			log.Println("--> Response: ", string(json))
		}

		return response, nil
	}
}
