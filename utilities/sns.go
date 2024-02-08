package utilities

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendMessage(message string) {

	statserr := StatsdClient.Inc("endpoint.assignments.PostSubmission", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to load statsdclient at sns.SendMessage")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{}))

	svc := sns.New(sess)
	// topics, err := svc.ListTopics(nil)
	// if err != nil {
	// 	slog.LogAttrs(context.Background(), slog.LevelError, "Failed to list topics at Utilities.SendMessage")
	// }

	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(os.Getenv("AWS_SNS_TOPIC_ARN")),
	})

	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to publish message at Utilities.SendMessage "+err.Error())
	}

	slog.LogAttrs(context.Background(), slog.LevelInfo, "SNS Message Id: "+*result.MessageId)

	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to publish SNS notification at Utilities.SendMessage"+err.Error())
	}
}
