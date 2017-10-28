package alexa

const (
	version                   = "1.0"
	plainTextOutputSpeechType = "PlainText"
	ssmlOutputSpeechType      = "SSML"
	simpleCardType            = "Simple"
	standardCardType          = "Standard"
	linkAccountCardType       = "LinkAccount"
)

// Response allows a handler to update the response returned to Alexa.
type Response interface {
	PlainText(text string)
	SSML(ssml string)

	SimpleCard(title, content string)
	StandardCard(title, text, smallImageURL, largeImageURL string)
	LinkAccountCard()

	RepromptPlainText(text string)
	RepromptSSML(ssml string)

	ShouldEndSession(value bool)
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
	Card             *card         `json:"card,omitempty"`
	OutputSpeech     *outputSpeech `json:"outputSpeech,omitempty"`
	Reprompt         *reprompt     `json:"reprompt,omitempty"`
	ShouldEndSession *bool         `json:"shouldEndSession,omitempty"`
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
