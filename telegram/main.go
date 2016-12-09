package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ahmdrz/nahjolbalaghe/telegram/database"
	"github.com/tucnak/telebot"
)

var db *database.Database
var admin int = 83919508
var token string = "314141809:AAEzok7Ra7fWYXK_q2fvIJVspHvV8Pue3N4"

func main() {
	log.Println("Opening database")
	db = database.Open()
	log.Println("Reading config file")
	Texts, err := Decode("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on telegram")
	bot, err := telebot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	callbacks := make(chan telebot.Callback)
	messages := make(chan telebot.Message)
	bot.Callbacks = callbacks
	bot.Messages = messages

	go func() {
		for message := range messages {
			text := message.Text
			log.Println(text)

			if text == "/start" {
				bot.ForwardMessage(telebot.User{ID: admin}, message)
				bot.SendMessage(message.Sender, Texts["start"], &telebot.SendOptions{
					ParseMode: telebot.ModeMarkdown,
					ReplyMarkup: telebot.ReplyMarkup{
						InlineKeyboard: [][]telebot.KeyboardButton{
							[]telebot.KeyboardButton{
								telebot.KeyboardButton{
									Text: "کانال نهج البلاغه",
									URL:  "https://telegram.me/joinchat/BQCClEFtXu0fkWE3Yu3XEg",
								},
							},
							[]telebot.KeyboardButton{
								telebot.KeyboardButton{
									Text: "حمایت از ما",
									URL:  "https://idpay.ir/balaghebot",
								},
							},
						},
					},
				})
			} else {
				list := db.SearchWisdom(text)
				if len(list) > 10 {
					bot.SendMessage(message.Sender, Texts["many"], nil)
				} else {
					if len(list) > 0 {
						keyboard := make([][]telebot.KeyboardButton, 0)
						for _, u := range list {
							row := make([]telebot.KeyboardButton, 1)
							row[0] = telebot.KeyboardButton{
								Data: "wisdom" + strconv.Itoa(u.ID),
								Text: u.Title,
							}
							keyboard = append(keyboard, row)
						}
						bot.SendMessage(message.Sender, Texts["result"], &telebot.SendOptions{
							ReplyMarkup: telebot.ReplyMarkup{
								InlineKeyboard: keyboard,
							},
						})
					} else {
						bot.SendMessage(message.Sender, Texts["notfound"], nil)
					}
				}
			}
		}
	}()

	go func() {
		for callback := range callbacks {
			data := callback.Data
			message := callback.Message
			message.Sender = callback.Sender
			if len(data) > 0 {
				if strings.HasPrefix(data, "wisdom") {
					data = strings.Replace(data, "wisdom", "", -1)
					a, _ := strconv.Atoi(data)
					wisdom := db.GetWisdom(a)
					bot.SendMessage(message.Sender, wisdom.Title+"\n\n"+wisdom.Subject+"\n\n"+wisdom.Text, &telebot.SendOptions{
						ParseMode: telebot.ModeMarkdown,
						ReplyMarkup: telebot.ReplyMarkup{
							InlineKeyboard: [][]telebot.KeyboardButton{
								[]telebot.KeyboardButton{
									telebot.KeyboardButton{
										Text: "کانال نهج البلاغه",
										URL:  "https://telegram.me/joinchat/BQCClEFtXu0fkWE3Yu3XEg",
									},
								},
								[]telebot.KeyboardButton{
									telebot.KeyboardButton{
										Text: "حمایت از ما",
										URL:  "https://idpay.ir/balaghebot",
									},
								},
							},
						},
					})
				}
			}
		}
	}()

	bot.Start(1 * time.Second)
}
