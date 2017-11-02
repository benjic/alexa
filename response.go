package alexa

import (
	"encoding/json"
)

const (
	version                   = "1.0"
	plainTextOutputSpeechType = "PlainText"
	ssmlOutputSpeechType      = "SSML"
	simpleCardType            = "Simple"
	standardCardType          = "Standard"
	linkAccountCardType       = "LinkAccount"
)

type Response interface {
	PlainText(text string)
	SSML(ssml string)
	SimpleCard(title, content string)
	StandardCard(title, text, smallImageURL, largeImageURL string)
	LinkAccountCard()
	RepromptPlainText(text string)
	RepromptSSML(ssml string)
	ShouldEndSession(value bool)

	AudioPlayerDirectives
}

type playDirective struct {
	Type         string                 `json:"type"`
	PlayBehavior string                 `json:"playBehavior"`
	AudioItem    playDirectiveAudioItem `json:"audioItem"`
}

type playDirectiveAudioItem struct {
	Stream playDirectiveAudioStream `json:"stream"`
}

type playDirectiveAudioStream struct {
	URL                   string  `json:"url"`
	Token                 string  `json:"token"`
	ExpectedPreviousToken *string `json:"expectedPreviousToken,omitempty"`
	OffsetInMilliseconds  int     `json:"offsetInMilliseconds"`
}

type stopDirective struct {
	Type string `json:"type"`
}

type clearAudioQueueDirective struct {
	Type          string `json:"type"`
	ClearBehavior string `json:"clearBehavior"`
}

type outputSpeech struct {
	SSML *string `json:"ssml,omitempty"`
	Text *string `json:"text,omitempty"`
	Type string  `json:"type"`
}

type image struct {
	LargeImageURL string `json:"largeImageUrl"`
	SmallImageURL string `json:"smallImageUrl"`
}

type card struct {
	Content *string `json:"content,omitempty"`
	Image   *image  `json:"image,omitempty"`
	Text    *string `json:"text,omitempty"`
	Title   *string `json:"title,omitempty"`
	Type    string  `json:"type"`
}

// Response passes data back to Alexa.
type responseBuilder struct {
	Version  string    `json:"version"`
	Response *response `json:"response"`
}

type response struct {
	Card             *card               `json:"card,omitempty"`
	OutputSpeech     *outputSpeech       `json:"outputSpeech,omitempty"`
	Reprompt         *reprompt           `json:"reprompt,omitempty"`
	Directives       *responseDirectives `json:"directives,omitempty"`
	ShouldEndSession *bool               `json:"shouldEndSession,omitempty"`
}

type responseDirectives struct {
	playDirective            *playDirective
	stopAudioDirective       *stopDirective
	clearAudioQueueDirective *clearAudioQueueDirective
}

func (d responseDirectives) MarshalJSON() ([]byte, error) {
	ds := []interface{}{}

	if d.playDirective != nil {
		ds = append(ds, d.playDirective)
	}

	if d.stopAudioDirective != nil {
		ds = append(ds, d.stopAudioDirective)
	}

	if d.clearAudioQueueDirective != nil {
		ds = append(ds, d.clearAudioQueueDirective)
	}

	return json.Marshal(ds)
}

type reprompt struct {
	OutputSpeech *outputSpeech `json:"outputSpeech"`
}

func (b *responseBuilder) PlainText(text string) {
	if b.Response.OutputSpeech == nil {
		b.Response.OutputSpeech = &outputSpeech{}
	}
	b.Response.OutputSpeech.SSML = nil
	b.Response.OutputSpeech.Text = &text
	b.Response.OutputSpeech.Type = plainTextOutputSpeechType
}

func (b *responseBuilder) SSML(ssml string) {
	if b.Response.OutputSpeech == nil {
		b.Response.OutputSpeech = &outputSpeech{}
	}
	b.Response.OutputSpeech.SSML = &ssml
	b.Response.OutputSpeech.Text = nil
	b.Response.OutputSpeech.Type = ssmlOutputSpeechType
}

func (b *responseBuilder) SimpleCard(title, content string) {
	if b.Response.Card == nil {
		b.Response.Card = &card{}
	}

	b.Response.Card.Content = &content
	b.Response.Card.Image = nil
	b.Response.Card.Text = nil
	b.Response.Card.Title = &title
	b.Response.Card.Type = simpleCardType
}

func (b *responseBuilder) StandardCard(title, text, smallImageURL, largeImageURL string) {
	if b.Response.Card == nil {
		b.Response.Card = &card{}
	}

	b.Response.Card.Content = nil
	b.Response.Card.Image = &image{smallImageURL, largeImageURL}
	b.Response.Card.Text = &text
	b.Response.Card.Title = &title
	b.Response.Card.Type = standardCardType
}

func (b *responseBuilder) LinkAccountCard() {
	if b.Response.Card == nil {
		b.Response.Card = &card{}
	}

	b.Response.Card.Content = nil
	b.Response.Card.Image = nil
	b.Response.Card.Text = nil
	b.Response.Card.Title = nil
	b.Response.Card.Type = linkAccountCardType
}

func (b *responseBuilder) RepromptPlainText(text string) {
	if b.Response.Reprompt == nil {
		b.Response.Reprompt = &reprompt{&outputSpeech{}}
	}
	b.Response.Reprompt.OutputSpeech.SSML = nil
	b.Response.Reprompt.OutputSpeech.Text = &text
	b.Response.Reprompt.OutputSpeech.Type = plainTextOutputSpeechType
}

func (b *responseBuilder) RepromptSSML(ssml string) {
	if b.Response.Reprompt == nil {
		b.Response.Reprompt = &reprompt{&outputSpeech{}}
	}
	b.Response.Reprompt.OutputSpeech.SSML = &ssml
	b.Response.Reprompt.OutputSpeech.Text = nil
	b.Response.Reprompt.OutputSpeech.Type = ssmlOutputSpeechType
}

func (b *responseBuilder) ShouldEndSession(value bool) {
	b.Response.ShouldEndSession = &value
}

func (b *responseBuilder) ReplaceAllAudio(token, url string, offsetInMilliseconds int) {
	if b.Response.Directives == nil {
		b.Response.Directives = &responseDirectives{}
	}

	b.Response.Directives.playDirective = &playDirective{
		Type:         "AudioPlayer.Play",
		PlayBehavior: "REPLACE_ALL",
		AudioItem: playDirectiveAudioItem{
			playDirectiveAudioStream{
				OffsetInMilliseconds: offsetInMilliseconds,
				Token:                token,
				URL:                  url,
			},
		},
	}
}

func (b *responseBuilder) EnqueueAudio(token, expectedPreviousToken, url string, offsetInMilliseconds int) {
	if b.Response.Directives == nil {
		b.Response.Directives = &responseDirectives{}
	}

	b.Response.Directives.playDirective = &playDirective{
		Type:         "AudioPlayer.Play",
		PlayBehavior: "ENQUEUE",
		AudioItem: playDirectiveAudioItem{
			playDirectiveAudioStream{
				ExpectedPreviousToken: &expectedPreviousToken,
				OffsetInMilliseconds:  offsetInMilliseconds,
				Token:                 token,
				URL:                   url,
			},
		},
	}
}

func (b *responseBuilder) ReplacedEnqueuedAudio(token, url string, offsetInMilliseconds int) {
	if b.Response.Directives == nil {
		b.Response.Directives = &responseDirectives{}
	}

	b.Response.Directives.playDirective = &playDirective{
		Type:         "AudioPlayer.Play",
		PlayBehavior: "REPLACE_ENQUEUED",
		AudioItem: playDirectiveAudioItem{
			playDirectiveAudioStream{
				OffsetInMilliseconds: offsetInMilliseconds,
				Token:                token,
				URL:                  url,
			},
		},
	}
}

func (b responseBuilder) StopAudio() {
	b.Response.Directives.stopAudioDirective = &stopDirective{Type: "AudioPlayer.Stop"}
}

func (b responseBuilder) ClearEnqueuedAudio() {
	b.Response.Directives.clearAudioQueueDirective = &clearAudioQueueDirective{
		Type:          "AudioPlayer.ClearQueue",
		ClearBehavior: "CLEAR_ENQUEUED",
	}
}

func (b responseBuilder) ClearAllAudio() {
	b.Response.Directives.clearAudioQueueDirective = &clearAudioQueueDirective{
		Type:          "AudioPlayer.ClearQueue",
		ClearBehavior: "CLEAR_ALL",
	}
}
