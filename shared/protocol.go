package shared

type TypeMsg string

const (
	WelcomeMessage    TypeMsg = "Welcome"
	SettingsMessage   TypeMsg = "Settings"
	GameEventMessage  TypeMsg = "GameEvent"
	ErrorMessage      TypeMsg = "Error"
	DisconnectMessage TypeMsg = "Disconnect"
)

type ServerMessage struct {
	Type    TypeMsg `json:"type"`
	ID      uint64  `json:"id"`
	Content string  `json:"content"`
	FScreen `json:"f_screen"`
}

type ClientMessage struct {
	Type    TypeMsg `json:"type"`
	ID      uint64  `json:"id"`
	Content string  `json:"content"`
	Vector
	IsShooting bool `json:"is_shooting"`
}

type mystr struct {
	arr []int
}
