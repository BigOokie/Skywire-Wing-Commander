// Copyright © 2018 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package wcconst

// Define constants used by the application
const (
	BotVersion = "v0.2.0-beta.3"

	// Bot command messages:
	// Help message
	MsgHelpShort = "*Usage:*\n" +
		"- /help - show this message\n" +
		"- /about - show information and credits about my creator and any contributors\n" +
		"- /status - ask how I'm going.. and if I'm still running\n" +
		"- /showconfig - display runtime configuration (from config.toml)\n" +
		"- /start - start activly monitoring your Skyminer. Once started, notifications will be sent to you for events that occur. A heartbeat will also be initiated to let you know if the bot and the Miner are still running.\n" +
		"- /stop - stop monitoring your Skyminer. Once stopped, I won't send any more notifications\n"

	MsgHelp = "*Wing Commander* here. I will help you to manage and monitor your Skyminer and its Nodes.\n\n" +
		MsgHelpShort +
		"\n" +
		"\n" +
		"Note: I am bound to this chat. I will only respond to commands from my configured Admin (%s)."

	// About cmd message
	MsgAbout = "*Wing Commander (" + BotVersion + ")*\n" +
		"A Telegram bot written in *Go* designed to help the *Skyfleet* community monitor and manage their *Skyminers*.\n" +
		"\n" +
		"*Created by:* @BigOokie *2018*\n" +
		"*GitHub:* https://github.com/BigOokie/skywire-wing-commander\n" +
		"*Twitter:* https://twitter.com/BigOokie\n" +
		"\n" +
		"Issues and feature requests must be logged via [GitHub](https://github.com/BigOokie/skywire-wing-commander/issues/new/choose)\n" +
		"\n" +
		"*Skycoin*: https://www.skycoin.net/\n" +
		"\n" +
		"*Donations most welcome* 👍\n" +
		"*Skycoin:* ES5LccJDhBCK275APmW9tmQNEgiYwTFKQF"

	MsgShowConfig = "*Wing Commander Configuration*\n%s\n"

	MsgErrorGetNodes     = "⚠️ An error occurred getting the list of Nodes from the Manager."
	MsgErrorGetDiscNodes = "⚠️ An error occurred checking Discovery Node connections."

	MsgConnectedNodes = "*Connected Nodes:* %v"
	MsgDiscConnNodes  = "*Discovery Connected Nodes:* %v"
	// Status cmd message
	MsgStatus = "*Wing Commander* Ready and reporting for duty 👍\n" + MsgConnectedNodes + "\n" + MsgDiscConnNodes
	// Heartbeat message
	MsgHeartbeat = "*Wing Commander Heatbeat* ❣️\n" + MsgConnectedNodes + "\n" + MsgDiscConnNodes

	// Node Connect/Disconnect Event Messages
	MsgNodeConnected    = "*Node Connected:* %s\n\n" + MsgConnectedNodes
	MsgNodeDisconnected = "‼ *Node Disconnected:* %s\n\n" + MsgConnectedNodes

	// Start cmd messages
	MsgMonitorAlreadyStarted = "️️*Wing Commander* Monitoring has already been started."
	MsgMonitorStart          = "*Wing Commander* Monitoring starting..."

	// Stop cmd message
	MsgMonitorStop       = "*Wing Commander* Monitoring stopping..."
	MsgMonitorStopped    = "*Wing Commander* Monitoring stopped..."
	MsgMonitorNotRunning = "*Wing Commander* Monitoring is not running..."

	// Default cmd message (unhandled)
	msgDefault = "Sorry. I don't know that command."

	// OS Interupt Signals
	MsgOSInteruptSig = "*Wing Commander* OS Interupt Signal Received. Exiting."
	MsgOSKillSig     = "*Wing Commander* OS Kill Signal Received. Exiting."
)
