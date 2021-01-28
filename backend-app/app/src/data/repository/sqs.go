package repository

import (
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	. "github.com/wataru-morioka/webauthn-project/backend-app/app/src/config"
)

var publisher *sqs.SQS

type SqsRepository struct {}

func NewSqsRepository() *SqsRepository {
	var once sync.Once
	once.Do(func() {
		log.Print("AWSセッション生成")
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		publisher = sqs.New(sess)
	})
	return &SqsRepository{}
}

func (s SqsRepository) Publish(message *string) error {
	env := NewEnv()
	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(*message),
		QueueUrl:     aws.String(env.QueueEndpoint),
		MessageGroupId: aws.String(Project),
		MessageDeduplicationId: aws.String(Project),
	}

	res, err := publisher.SendMessage(params)
	if err != nil {
		log.Printf("キュー登録エラー %+v", err)
		return err
	}

	log.Printf("sent to queue response: %+v", res)

	return nil
}