package telegram

import (
	"bytes"
	"fmt"
	"strings"

	"gopkg.in/tucnak/telebot.v2"

	"github.com/moira-alert/moira"
)

type messageType string

const (
	// Album type used if notification has plots
	Album messageType = "album"
	// Message type used if notification has not plot
	Message messageType = "message"
)

const (
	albumCaptionMaxCharacters     = 1024
	messageMaxCharacters          = 4096
	additionalInfoCharactersCount = 400
)

var characterLimits = map[messageType]int{
	Message: messageMaxCharacters,
	Album:   albumCaptionMaxCharacters,
}

// SendEvents implements Sender interface Send
func (sender *Sender) SendEvents(events moira.NotificationEvents, contact moira.ContactData, trigger moira.TriggerData, plots [][]byte, throttled bool) error {
	msgType := getMessageType(plots)
	message := sender.buildMessage(events, trigger, throttled, characterLimits[msgType])
	sender.logger.Debugf("Calling telegram api with chat_id %s and message body %s", contact.Value, message)
	chat, err := sender.getChat(contact.Value)
	if err != nil {
		if ok, errorBrokenContact := checkBrokenContactError(sender.logger, err); ok {
			return errorBrokenContact
		}
		return err
	}
	if err := sender.talk(chat, message, plots, msgType); err != nil {
		if ok, errorBrokenContact := checkBrokenContactError(sender.logger, err); ok {
			return errorBrokenContact
		}
		return fmt.Errorf("failed to send message to telegram contact %s: %s. ", contact.Value, err)
	}
	return nil
}

func (sender *Sender) buildMessage(events moira.NotificationEvents, trigger moira.TriggerData, throttled bool, maxChars int) string {
	var buffer bytes.Buffer
	state := events.GetSubjectState()
	tags := trigger.GetTags()
	emoji := emojiStates[state]

	title := fmt.Sprintf("%s%s %s %s (%d)\n", emoji, state, trigger.Name, tags, len(events))
	buffer.WriteString(title)

	var messageCharsCount, printEventsCount int
	messageCharsCount += len([]rune(title))
	messageLimitReached := false

	for _, event := range events {
		line := fmt.Sprintf("\n%s: %s = %s (%s to %s)", event.FormatTimestamp(sender.location), event.Metric, event.GetMetricsValues(), event.OldState, event.State)
		if msg := event.CreateMessage(sender.location); len(msg) > 0 {
			line += fmt.Sprintf(". %s", msg)
		}
		lineCharsCount := len([]rune(line))
		if messageCharsCount+lineCharsCount > maxChars-additionalInfoCharactersCount {
			messageLimitReached = true
			break
		}
		buffer.WriteString(line)
		messageCharsCount += lineCharsCount
		printEventsCount++
	}

	if messageLimitReached {
		buffer.WriteString(fmt.Sprintf("\n\n...and %d more events.", len(events)-printEventsCount))
	}
	url := trigger.GetTriggerURI(sender.frontURI)
	if url != "" {
		buffer.WriteString(fmt.Sprintf("\n\n%s\n", url))
	}

	if throttled {
		buffer.WriteString("\nPlease, fix your system or tune this trigger to generate less events.")
	}
	return buffer.String()
}

func (sender *Sender) getChatUID(username string) (string, error) {
	var uid string
	if strings.HasPrefix(username, "%") {
		uid = "-100" + username[1:]
	} else {
		var err error
		uid, err = sender.DataBase.GetIDByUsername(messenger, username)
		if err != nil {
			return "", fmt.Errorf("failed to get username uuid: %s", err.Error())
		}
	}
	return uid, nil
}

func (sender *Sender) getChat(username string) (*telebot.Chat, error) {
	uid, err := sender.getChatUID(username)
	if err != nil {
		return nil, err
	}
	chat, err := sender.bot.ChatByID(uid)
	if err != nil {
		return nil, fmt.Errorf("can't find recipient %s: %s", uid, err.Error())
	}
	return chat, nil
}

// talk processes one talk
func (sender *Sender) talk(chat *telebot.Chat, message string, plots [][]byte, messageType messageType) error {
	if messageType == Album {
		sender.logger.Debug("talk as album")
		return sender.sendAsAlbum(chat, plots, message)
	}
	sender.logger.Debug("talk as send message")
	return sender.sendAsMessage(chat, message)
}

func (sender *Sender) sendAsMessage(chat *telebot.Chat, message string) error {
	_, err := sender.bot.Send(chat, message)
	if err != nil {
		sender.logger.Debugf("can't send event message [%s] to %v: %s", message, chat.ID, err.Error())
	}
	return err
}

func checkBrokenContactError(logger moira.Logger, err error) (bool, error) {
	logger.Info("Check broken contact")
	if err == nil {
		return false, nil
	}
	if e, ok := err.(*telebot.APIError); ok {
		logger.Debug("It's telebot.APIError from talk(): code = %d, msg = %s, desc = %s", e.Code, e.Message, e.Description)
		if e.Code == telebot.ErrUnauthorized.Code { // all forbid errors
			return true, moira.NewSenderBrokenContactError(err)
		}
	}
	if strings.HasPrefix(err.Error(), "failed to get username uuid") {
		logger.Debug("It's error from getChat(): ", err)
		return true, moira.NewSenderBrokenContactError(err)
	}
	return false, nil
}

func prepareAlbum(plots [][]byte, caption string) telebot.Album {
	var album telebot.Album
	for _, plot := range plots {
		photo := &telebot.Photo{File: telebot.FromReader(bytes.NewReader(plot)), Caption: caption}
		album = append(album, photo)
		caption = "" // Caption should be defined only for first photo
	}
	return album
}

func (sender *Sender) sendAsAlbum(chat *telebot.Chat, plots [][]byte, caption string) error {
	album := prepareAlbum(plots, caption)

	_, err := sender.bot.SendAlbum(chat, album)
	if err != nil {
		sender.logger.Debugf("can't send event plots to %v: %s", chat.ID, err.Error())
	}
	return err
}

func getMessageType(plots [][]byte) messageType {
	if len(plots) > 0 {
		return Album
	}
	return Message
}
