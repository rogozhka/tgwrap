package tgwrap

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBot_GetWebhookInfo(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	whi, err := bot.GetWebhookInfo()
	is.Nil(err)
	is.NotNil(whi)
}

func TestBot_GetWebhookInfoContext(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)

	_, err := bot.GetWebhookInfoContext(ctx)
	is.NotNil(err)
}
