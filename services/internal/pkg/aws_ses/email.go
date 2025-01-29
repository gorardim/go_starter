package aws_ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Email struct {
	To      []string
	Subject string
	Body    string
}

func (c *Client) SendEmail(email *Email) error {
	creds := credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	awsConfig := &aws.Config{
		Region:      aws.String(c.Region),
		Credentials: creds,
	}
	sess := session.Must(session.NewSession(awsConfig))
	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice(email.To),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(email.Body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(email.Subject),
			},
		},
		Source: aws.String(c.From),
	}

	_, err := svc.SendEmail(input)
	if err != nil {
		return err
	}
	return nil
}
