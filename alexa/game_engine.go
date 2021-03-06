package alexa

// GameEngineStartInputDirective directive to start the game engine.
type GameEngineStartInputDirective struct {
	Type                 string                                  `json:"type,omitempty"`
	Timeout              int                                     `json:"timeout"`
	MaximumHistoryLength int                                     `json:"maximumHistoryLength,omitempty"`
	Proxies              []interface{}                           `json:"proxies,omitempty"`
	Recognizers          map[string]interface{}                  `json:"recognizers"`
	Events               map[string]*GameEngineRegistrationEvent `json:"events"`
}

// GameEngineStopInputHandlerDirective stops Echo Button events from being sent to your skill.
type GameEngineStopInputHandlerDirective struct {
	Type                 string `json:"type,omitempty"`
	OriginatingRequestID string `json:"originatingRequestId"`
}

// GameEnginePatternRecognizer is true when all of the specified events have occurred in the specified order.
type GameEnginePatternRecognizer struct {
	// Must be match
	Type      string              `json:"type"`
	Anchor    string              `json:"anchor,omitempty"`
	Fuzzy     bool                `json:"fuzzy"`
	GadgetIds []string            `json:"gadgetIds,omitempty"`
	Actions   []interface{}       `json:"actions,omitempty"`
	Pattern   []GameEnginePattern `json:"pattern"`
}

// GameEnginePattern is an object that provides all of the events that need to occur, in a specific order, for this recognizer to be true
type GameEnginePattern struct {
	GadgetIds []string `json:"gadgetIds,omitempty"`
	Colors    []string `json:"colors,omitempty"`
	Action    string   `json:"action,omitempty"`
}

// GameEngineDeviationRecognizer returns true when another specified recognizer reports that the player has deviated from its expected pattern.
type GameEngineDeviationRecognizer struct {
	// Must be deviation
	Type       string `json:"type"`
	Recognizer string `json:"recognizer"`
}

// GameEngineProgressRecognizer consults another recognizer for the degree of completion, and is true if that degree is above the specified threshold. The completion parameter is specified as a decimal percentage.
type GameEngineProgressRecognizer struct {
	// Must be progress
	Type       string `json:"type"`
	Recognizer string `json:"recognizer"`
	Completion int    `json:"completion"`
}

// GameEngineRegistrationEvent object is where you define the conditions that must be met for your skill to be notified of Echo Button input. You must define at least one event.
type GameEngineRegistrationEvent struct {
	Meets []string `json:"meets"`
	Fails []string `json:"fails,omitempty"`
	// Possible values: history, matches
	Reports                 string `json:"reports,omitempty"`
	ShouldEndInputHandler   bool   `json:"shouldEndInputHandler"`
	MaximumInvocations      int    `json:"maximumInvocations,omitempty"`
	TriggerTimeMilliseconds int    `json:"triggerTimeMilliseconds,omitempty"`
}

// GameEngineInputEvent contains list of events sent from the Input Handler. Each event that you specify will be sent only once to your skill as it becomes true. Note that in any InputHandlerEvent request one or more events may have become true at the same time.
type GameEngineInputEvent struct {
	Name        string `json:"name"`
	InputEvents []struct {
		GadgetID  string `json:"gadgetId"`
		Timestamp string `json:"timestamp"`
		Action    string `json:"action"`
		Color     string `json:"color"`
		Feature   string `json:"feature"`
	} `json:"inputEvents"`
}

// GameEngineInputHandlerEventRequest is send by GameEngine to notify your skill about Echo Button events
type GameEngineInputHandlerEventRequest struct {
	CommonRequest
	// From GamEngine.InputHandlerEvent
	OriginatingRequestID string                 `json:"originatingRequestId"`
	Events               []GameEngineInputEvent `json:"events"`
}

// AddGameEngineStartInputDirective creates a new directive with StartInputerHandler Type and adds it to the response.
func (r *Response) AddGameEngineStartInputDirective(timeout int) *GameEngineStartInputDirective {
	d := &GameEngineStartInputDirective{
		Type:        "GameEngine.StartInputHandler",
		Timeout:     timeout,
		Recognizers: make(map[string]interface{}),
	}
	r.AddDirective(d)
	return d
}

// AddGameEngineStopInputHandlerDirective creates a new directive to stop listening for input events and adds it to the response.
func (r *Response) AddGameEngineStopInputHandlerDirective(originatingRequestID string) *GameEngineStopInputHandlerDirective {
	d := &GameEngineStopInputHandlerDirective{
		Type:                 "GameEngine.StopInputHandler",
		OriginatingRequestID: originatingRequestID,
	}
	r.AddDirective(d)
	return d
}

// AddPatternRecognizer adds a recognizer with the given name and returns the reference. The recognizer is true when all of the specified events have occurred in the specified order.
func (sid *GameEngineStartInputDirective) AddPatternRecognizer(name string) *GameEnginePatternRecognizer {
	recognizer := &GameEnginePatternRecognizer{
		Type:    "match",
		Pattern: make([]GameEnginePattern, 0),
	}
	sid.Recognizers[name] = recognizer
	return recognizer
}

// AddPattern adds a pattern object. All patterns must occur in order for the recognizer to be true
func (gep *GameEnginePatternRecognizer) AddPattern(gadgetIds, colors []string, action string) {
	pattern := GameEnginePattern{
		GadgetIds: gadgetIds,
		Colors:    colors,
		Action:    action,
	}
	gep.Pattern = append(gep.Pattern, pattern)
}

// AddDeviationRecognizer adds a recognizer with the given name and returns the reference.
func (sid *GameEngineStartInputDirective) AddDeviationRecognizer(name string, recognizerName string) *GameEngineDeviationRecognizer {
	recognizer := &GameEngineDeviationRecognizer{
		Type:       "deviation",
		Recognizer: recognizerName,
	}
	sid.Recognizers[name] = recognizer
	return recognizer
}

// AddEvent adds a GameEngine Event registration to the directive.
func (sid *GameEngineStartInputDirective) AddEvent(name string, shouldEndInputHandler bool, meetsRecognizers []string) *GameEngineRegistrationEvent {
	event := &GameEngineRegistrationEvent{
		Meets: meetsRecognizers,
		ShouldEndInputHandler: shouldEndInputHandler,
	}
	if sid.Events == nil {
		sid.Events = make(map[string]*GameEngineRegistrationEvent)
	}
	sid.Events[name] = event
	return event
}

// AddProgressRecognizer adds a recognizer with the given name and returns the reference. The recognizer is true when all of the specified events have occurred in the specified order.
func (sid *GameEngineStartInputDirective) AddProgressRecognizer(name, recognizerName string, completion int) *GameEngineProgressRecognizer {
	recognizer := &GameEngineProgressRecognizer{
		Type:       "progress",
		Recognizer: recognizerName,
		Completion: completion,
	}
	sid.Recognizers[name] = recognizer
	return recognizer
}
