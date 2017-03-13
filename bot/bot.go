package bot

import (
	"fmt"
	"time"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/bot/commands"
	"github.com/bwmarrin/discordgo"
)

const (
	// https://github.com/izy521/discord.io/blob/master/docs/colors.md
	discordColorBlack      = 0
	discordColorAqua       = 1752220
	discordColorGreen      = 3066993
	discordColorBlue       = 3447003
	discordColorPurple     = 10181046
	discordColorGold       = 15844367
	discordColorOrange     = 15105570
	discordColorRed        = 15158332
	discordColorGrey       = 9807270
	discordColorDarkerGrey = 8359053
	discordColorNavy       = 3426654
	discordColorDarkAqua   = 1146986
	discordColorDarkGreen  = 2067276
	discordColorDarkBlue   = 2123412
	discordColorDarkPurple = 7419530
	discordColorDarkGold   = 12745742
	discordColorDarkOrange = 11027200
	discordColorDarkRed    = 10038562
	discordColorDarkGrey   = 9936031
	discordColorLightGrey  = 12370112
	discordColorDarkNavy   = 2899536
)

// Bot encompasses a DiscordGo Bot
type Bot struct {
	Start time.Time
	*discordgo.Session
}

// NewBot creates a new Bot
func NewBot(token string) (b *Bot, err error) {
	discord, err := discordgo.New(token)
	if err != nil {
		return
	}
	b = &Bot{Start: time.Now(), Session: discord}
	b.routeHandlers()
	return
}

func (b *Bot) routeHandlers() {
	b.AddHandler(DiceRollHandler)
	b.AddHandler(NewHandler(".about", commands.AboutCommand))
}

// Self returns the User struct associated with the bot user
func (b *Bot) Self() (*discordgo.User, error) {
	return b.User("@me")
}

func (b *Bot) String() string {
	return fmt.Sprintf("Bot[%v] - started at %v", b.Token, b.Start)
}
