package slack

type Config struct {
	APIToken string
}

type SlackClient struct {
	APIToken	string
}
//Client configures and returns a fully initialized client for accessing the Slack API.
func (c* Config) Client() (interface{}, error) {
	client := &SlackClient{}
	return client, nil
}