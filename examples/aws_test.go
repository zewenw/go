package examples

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestAws(t *testing.T) {

	t.Run("Aws Batch Delete", func(t *testing.T) {
		ctx := context.Background()
		cfg, err := config.LoadDefaultConfig(
			ctx,
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
			config.WithEndpointResolver(aws.EndpointResolverFunc(
				func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:               "http://api.dev.am/localstack/",
						SigningRegion:     "eu-west-2",
						HostnameImmutable: true,
					}, nil
				},
			)),
		)
		if err != nil {
			panic(err)
		}

		s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.HTTPClient = &http.Client{}
		})

		action := S3Actions{
			s3Client,
			nil,
		}
		keyOne := "24 Apr.md"
		keyTwo := "2"
		keyThree := "3"
		objects := []types.ObjectIdentifier{
			{
				Key: &keyOne,
			}, {
				Key: &keyTwo,
			}, {
				Key: &keyThree,
			},
		}
		err = action.DeleteObjects(ctx, "test", objects, false)
		fmt.Println(err)
	})

}

// S3Actions wraps S3 service actions.
type S3Actions struct {
	S3Client  *s3.Client
	S3Manager *manager.Uploader
}

// DeleteObjects deletes a list of objects from a bucket.
func (actor S3Actions) DeleteObjects(ctx context.Context, bucket string, objects []types.ObjectIdentifier, bypassGovernance bool) error {
	if len(objects) == 0 {
		return nil
	}

	input := s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &types.Delete{
			Objects: objects,
			Quiet:   aws.Bool(false),
		},
	}
	if bypassGovernance {
		input.BypassGovernanceRetention = aws.Bool(true)
	}
	delOut, err := actor.S3Client.DeleteObjects(ctx, &input)
	for _, d := range delOut.Deleted {
		fmt.Printf("Successfully deleted: %s\n", *d.Key)
	}
	for _, e := range delOut.Errors {
		fmt.Printf("Failed to delete: %s\nReason: %s\n", *e.Key, *e.Message)
	}
	if err != nil || len(delOut.Errors) > 0 {
		log.Printf("Error deleting objects from bucket %s.\n", bucket)
		if err != nil {
			var noBucket *types.NoSuchBucket
			if errors.As(err, &noBucket) {
				log.Printf("Bucket %s does not exist.\n", bucket)
				err = noBucket
			}
		} else if len(delOut.Errors) > 0 {
			for _, outErr := range delOut.Errors {
				log.Printf("%s: %s\n", *outErr.Key, *outErr.Message)
			}
			err = fmt.Errorf("%s", *delOut.Errors[0].Message)
		}
	} else {
		for _, delObjs := range delOut.Deleted {
			err = s3.NewObjectNotExistsWaiter(actor.S3Client).Wait(
				ctx, &s3.HeadObjectInput{Bucket: aws.String(bucket), Key: delObjs.Key}, time.Minute)
			if err != nil {
				log.Printf("Failed attempt to wait for object %s to be deleted.\n", *delObjs.Key)
			} else {
				log.Printf("Deleted %s.\n", *delObjs.Key)
			}
		}
	}
	return err
}
