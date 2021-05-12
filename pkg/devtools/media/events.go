package media

// PlayerPropertiesChanged asynchronous event. This can be called multiple times, and can be used to set / override /
// remove player properties. A null propValue indicates removal.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#event-playerPropertiesChanged
type PlayerPropertiesChanged struct {
	PlayerID   string           `json:"playerId"`
	Properties []PlayerProperty `json:"properties"`
}

// PlayerEventsAdded asynchronous event. Send events as a list, allowing them to be batched on the browser for less
// congestion. If batched, events must ALWAYS be in chronological order.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#event-playerEventsAdded
type PlayerEventsAdded struct {
	PlayerID string        `json:"playerId"`
	Events   []PlayerEvent `json:"events"`
}

// PlayerMessagesLogged asynchronous event. Send a list of any messages that need to be delivered.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#event-playerMessagesLogged
type PlayerMessagesLogged struct {
	PlayerID string          `json:"playerId"`
	Messages []PlayerMessage `json:"messages"`
}

// PlayerErrorsRaised asynchronous event. Send a list of any errors that need to be delivered.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#event-playerErrorsRaised
type PlayerErrorsRaised struct {
	PlayerID string        `json:"playerId"`
	Errors   []PlayerError `json:"errors"`
}

// PlayersCreated asynchronous event. Called whenever a player is created, or when a new agent joins and receives
// a list of active players. If an agent is restored, it will receive the full
// list of player ids and all events again.
//
// https://chromedevtools.github.io/devtools-protocol/tot/Media/#event-playersCreated
type PlayersCreated struct {
	Players []string `json:"players"`
}
