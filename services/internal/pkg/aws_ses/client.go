package aws_ses

type Client struct {
	Region    string
	AccessKey string
	SecretKey string
	From      string
}

func NewClient(region string, accessKey string, secretKey string, from string) *Client {
	return &Client{
		Region:    region,
		AccessKey: accessKey,
		SecretKey: secretKey,
		From:      from,
	}
}
