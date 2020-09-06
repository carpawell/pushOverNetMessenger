package pushApp

import "github.com/gregdel/pushover"

type PushApp struct {
	App       *pushover.Pushover
	Recipient *pushover.Recipient
}

func New(token string, user string) *PushApp {
	return &PushApp{
		App:       pushover.New(token),
		Recipient: pushover.NewRecipient(user),
	}
}

func (pa PushApp) SendMessage(message string) (*pushover.Response, error) {
	pushMessage := pushover.NewMessage(message)
	response, err := pa.App.SendMessage(pushMessage, pa.Recipient)
	if err != nil {
		return nil, err
	}

	return response, nil
}
