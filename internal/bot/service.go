package bot

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"text/template"
)

// Group of constants for bot messages
const (
	startMessage           = "👋 Привет, меня зовут Космос! \n Я бот, который поможет тебе купить футболку:)"
	feedbackMessage        = "Оцени магазин и качество вещей по пятибалльной шкале:"
	thxFeedbackMessage     = "Спасибо! <3"
	sorryFeedbackMessage   = "Нам очень жаль, что вам не понравилось, мы постараемся стать лучше!"
	heightFeedbackMessage  = "Чтобы \"Предмет\" смотрелся как задумано, выбери свой дипазон роста:"
	addressFeedbackMessage = "🇷🇺 Россия г. Санкт-Петербург"
	sorryHeightMessage     = "У нас пока что нет таких размеров, но мы уже стараемся исправить эту проблему!"
	itemsMessage           = "Выберите товар:"
	infoMessage            = "Средняя оценка: {{ .Avg }} ⭐️\n"
	itemMessage            = "{{ .Name }} \n{{ .Price }}р.\n{{ .Description }}"
)

// Group of constants for handling messages from user.
const (
	height        = "Рост"
	start         = "start"
	address       = "Адрес"
	feedBack      = "Оставить отзыв"
	thxFeedback   = "Спасибо!"
	sorryFeedback = "Мы стараемся!"
	sorryHeight   = "Неверный размер"
	size          = "размер"
	items         = "Предметы"
	item          = "item"
	info          = "info"
	rate          = "rate"
)

// itemButtons array of items. Automatically fulfilled from storage when bot starts.
var itemButtons = make([][]tgapi.InlineKeyboardButton, 0)

// Group of variables that are keyboard buttons.
var (
	startKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Купить 🛒", items),
			tgapi.NewInlineKeyboardButtonData("Адрес 📍", address),
			tgapi.NewInlineKeyboardButtonData("Отзыв ⭐️", feedBack),
			tgapi.NewInlineKeyboardButtonURL("VK 💙", "https://vk.com/ledda.store"),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Узнать размер ❔", height),
			tgapi.NewInlineKeyboardButtonData("О магазине ℹ️", info),
		),
	)

	addressKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	itemsKeyboard = tgapi.NewInlineKeyboardMarkup()

	feedBackKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("1", "rate::1"),
			tgapi.NewInlineKeyboardButtonData("2", "rate::2"),
			tgapi.NewInlineKeyboardButtonData("3", "rate::3"),
			tgapi.NewInlineKeyboardButtonData("4", "rate::4"),
			tgapi.NewInlineKeyboardButtonData("5", "rate::5"),
		),

		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	thxFeedbackKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("На главную", start),
		),
	)

	sorryFeedbackKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Изменить отзыв", feedBack),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("На главную", start),
		),
	)

	heightKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(" - 158", sorryHeight),
			tgapi.NewInlineKeyboardButtonData("159 - 170", size+"::S"),
			tgapi.NewInlineKeyboardButtonData("171 - 180", size+"::M"),
			tgapi.NewInlineKeyboardButtonData("181 - 188", size+"::L"),
			tgapi.NewInlineKeyboardButtonData("189 - ", sorryHeight),
		),

		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	soldKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Нет в наличии 💔", items),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", items),
		),
	)

	buyKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Купить 🛒", items),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)

	infoKeyboard = tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Оставить отзыв 💫", feedBack),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData("Назад", start),
		),
	)
)

// Group of templates for messages.
var (
	itemTemplate = template.Must(template.New("items").Parse(itemMessage))
	infoTemplate = template.Must(template.New("info").Parse(infoMessage))
)
