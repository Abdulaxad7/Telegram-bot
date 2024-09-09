package router

import "gopkg.in/telebot.v3"

type InlineButton struct{}

type InlineButtons interface {
	LanguageSelection(bot *telebot.Bot, c telebot.Context) [][]telebot.InlineButton
	CountrySelection(bot *telebot.Bot) [][]telebot.InlineButton
}
type KeyboardButton struct{}
type Context telebot.Context

type KeyboardButtons interface {
	WeatherSel(bot *telebot.Bot) [][]telebot.ReplyButton
	hourly(c Context) error
	today(c Context) error
	tomorrow(c Context) error
	afterTomorrow(c Context) error
	current(c Context) error
}
