package main

// BotConfig struct is used to store run-time configuration
// information for the bot application.
type BotConfig struct {
	BotToken       string `json:"bot_token"`
	ChatID         int64  `json:"chat_id"`
	BotDebug       bool   `json:"botdebug"`
	ClientFile     string `json:"clientfile"`
	MonitorRunning bool   `json:"monitorrunning"`
	HeartBeat      bool   `json:"heartbeat"`
}

// TOML Config
type swwcConfig struct {
	Title string
	Bot   botConfig
}

type botConfig struct {
	Token string
	Debug bool
}