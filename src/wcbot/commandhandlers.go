// Copyright © 2018 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BigOokie/skywire-wing-commander/src/utils"
	"github.com/BigOokie/skywire-wing-commander/src/wcconst"

	log "github.com/sirupsen/logrus"
)

// Handler for help command
func (bot *Bot) handleCommandHelp(ctx *BotContext, command, args string) error {
	log.Debug("Handle command /help")
	return bot.Send(ctx, "whisper", "markdown", fmt.Sprintf(wcconst.MsgHelp, bot.config.Telegram.Admin))
}

// Handler for about command
func (bot *Bot) handleCommandAbout(ctx *BotContext, command, args string) error {
	log.Debug("Handle command /about")
	return bot.Send(ctx, "whisper", "markdown", wcconst.MsgAbout)
}

// Handler for showconfig command
func (bot *Bot) handleCommandShowConfig(ctx *BotContext, command, args string) error {
	log.Debug("Handle command /showconfig")
	err := bot.Send(ctx, "whisper", "markdown", fmt.Sprintf(wcconst.MsgShowConfig, bot.config.String()))
	if err != nil {
		log.Error("handleCommandShowConfig::Send: %s", err)
		log.Debug("handleCommandShowConfig::Send - Attempting to resend as text")
		err = bot.Send(ctx, "whisper", "text", fmt.Sprintf(wcconst.MsgShowConfig, bot.config.String()))
	}
	return err
}

// Handler for start command
func (bot *Bot) handleCommandStart(ctx *BotContext, command, args string) error {
	log.Debug("Handle command /start")

	if bot.skyMgrMonitor.IsRunning() {
		log.Debug(wcconst.MsgMonitorAlreadyStarted)
		return bot.Send(ctx, "whisper", "markdown", wcconst.MsgMonitorAlreadyStarted)
	}

	log.Debug(wcconst.MsgMonitorStart)
	cancelContext, cancelFunc := context.WithCancel(context.Background())
	bot.skyMgrMonitor.SetCancelFunc(cancelFunc)
	bot.skyMgrMonitor.monitorStatusMsgChan = make(chan string)

	// Start the Event Monitor - provide cancelContext
	go bot.monitorEventLoop(cancelContext, ctx, bot.skyMgrMonitor.monitorStatusMsgChan)
	// Start monitoring the local Manager - provide cancelContext
	go bot.skyMgrMonitor.RunManagerMonitor(cancelContext, bot.skyMgrMonitor.monitorStatusMsgChan, bot.config.Monitor.IntervalSec)
	// Start monitoring the local Manager - provide cancelContext
	//go bot.skyMgrMonitor.RunDiscoveryMonitor(cancelContext, bot.skyMgrMonitor.monitorStatusMsgChan, bot.config.Monitor.DiscoveryMonitorIntMin)

	return bot.Send(ctx, "whisper", "markdown", wcconst.MsgMonitorStart)
}

// Handler for stop command
func (bot *Bot) handleCommandStop(ctx *BotContext, command, args string) error {
	log.Debug("Handle command /stop")

	if bot.skyMgrMonitor.IsRunning() {
		log.Debug(wcconst.MsgMonitorStop)
		bot.skyMgrMonitor.DoCancelFunc()
		bot.skyMgrMonitor.SetCancelFunc(nil)
		close(bot.skyMgrMonitor.monitorStatusMsgChan)
		bot.skyMgrMonitor.monitorStatusMsgChan = nil
		log.Debug(wcconst.MsgMonitorStopped)
		return bot.Send(ctx, "whisper", "markdown", wcconst.MsgMonitorStop)
	}

	log.Debug(wcconst.MsgMonitorNotRunning)
	return bot.Send(ctx, "whisper", "markdown", wcconst.MsgMonitorNotRunning)
}

// Handler for status command
func (bot *Bot) handleCommandStatus(ctx *BotContext, command, args string) error {
	log.Debug("Handle command /status")

	if !bot.skyMgrMonitor.IsRunning() {
		// Monitor not running
		return bot.Send(ctx, "whisper", "markdown", wcconst.MsgMonitorNotRunning)
	}

	// Monitor is running
	discConnNodes, err := bot.skyMgrMonitor.ConnectedDiscNodeCount()

	// Everything is ok
	status := "👍"
	statusmsg := ""
	if err != nil {
		// Error connecting to Discovery Server
		status = "⚠️"
		statusmsg = wcconst.MsgErrorGetDiscNodes
	} else if bot.skyMgrMonitor.GetConnectedNodeCount() != discConnNodes {
		// We connected but not all nodes are reported as connected
		status = "⚠️"
		statusmsg = wcconst.MsgDiscSomeNodes
	}

	return bot.Send(ctx, "whisper", "markdown",
		fmt.Sprintf(wcconst.MsgStatus, status, bot.skyMgrMonitor.GetConnectedNodeCount(), discConnNodes, statusmsg))
}

// Handler for help CheckUpdate
func (bot *Bot) handleCommandCheckUpdate(ctx *BotContext, command, args string) error {
	log.Debug("Handle command /checkupdate")
	bot.Send(ctx, "whisper", "markdown", "Checking for updates...")

	updateAvailable, updateMsg := utils.UpdateAvailable("BigOokie", "skywire-wing-commander", wcconst.BotVersion)
	if updateAvailable {
		return bot.Send(ctx, "whisper", "markdown",
			fmt.Sprintf("*Update available:* %s", updateMsg))
	} else {
		return bot.Send(ctx, "whisper", "markdown",
			fmt.Sprintf("*Up to date:* %s", updateMsg))
	}
}

func (bot *Bot) handleDirectMessageFallback(ctx *BotContext, text string) (bool, error) {
	errmsg := fmt.Sprintf("Sorry, I only take commands. '%s' is not a command.\n\n%s", text, wcconst.MsgHelpShort)
	log.Debugf(errmsg)
	return true, bot.Reply(ctx, "markdown", errmsg)
}

// AddPrivateMessageHandler adds a private MessageHandler to the Bot
func (bot *Bot) AddPrivateMessageHandler(handler MessageHandler) {
	bot.privateMessageHandlers = append(bot.privateMessageHandlers, handler)
}

// AddGroupMessageHandler adds a group MessageHandler to the Bot
func (bot *Bot) AddGroupMessageHandler(handler MessageHandler) {
	bot.groupMessageHandlers = append(bot.groupMessageHandlers, handler)
}

// monitorEventLoop monitors for event messages from the SkyMgrMonitor (when running).
// Its also responsible for managing the Heartbeat (if configured)
func (bot *Bot) monitorEventLoop(runctx context.Context, botctx *BotContext, statusMsgChan <-chan string) {
	tickerHB := time.NewTicker(bot.config.Monitor.HeartbeatIntMin)

	for {
		select {
		// Monitor Status Message
		case msg := <-statusMsgChan:
			log.Debugf("Bot.monitorEventLoop: Status event: %s", msg)
			bot.Send(botctx, "whisper", "markdown", msg)

		// Heartbeat ticker event
		case <-tickerHB.C:
			log.Debug("Bot.monitorEventLoop - Heartbeat event")

			discConnNodes, err := bot.skyMgrMonitor.ConnectedDiscNodeCount()
			var msg string
			if err != nil {
				// Discover Server Error
				msg = fmt.Sprintf(wcconst.MsgHeartbeat+"\n"+wcconst.MsgErrorGetDiscNodes,
					bot.skyMgrMonitor.GetConnectedNodeCount(), discConnNodes)
			} else {
				// Ok
				msg = fmt.Sprintf(wcconst.MsgHeartbeat,
					bot.skyMgrMonitor.GetConnectedNodeCount(), discConnNodes)
			}

			bot.Send(botctx, "whisper", "markdown", msg)

		// Context has been cancelled. Shutdown
		case <-runctx.Done():
			log.Debugln("Bot.monitorEventLoop - Done event.")
			return
		}
	}
}
