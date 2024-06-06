package router

import "gopkg.in/telebot.v3"

type InlineButton struct{}

type InlineButtons interface {
	LanguageSelection(bot *telebot.Bot, c telebot.Context) [][]telebot.InlineButton
	CountrySelection(bot *telebot.Bot) [][]telebot.InlineButton
}
type KeyboardButton struct{}

type KeyboardButtons interface {
	WeatherSel(bot *telebot.Bot) [][]telebot.ReplyButton
	hourly(c telebot.Context) error
	today(c telebot.Context) error
	tomorrow(c telebot.Context) error
	afterTomorrow(c telebot.Context) error
	current(c telebot.Context) error
}
