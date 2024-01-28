package main

import (
    "bytes"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    // Set up AWS S3 configuration
    s3Config := &aws.Config{
        Credentials: credentials.NewStaticCredentials("<DO_ACCESS_KEY>", "<DO_SECRET_KEY>", ""),
        Endpoint:    aws.String("https://nyc3.digitaloceanspaces.com"),
        Region:      aws.String("nyc3"), // e.g., "nyc3"
    }

    // Create a new session
    newSession, err := session.NewSession(s3Config)
    if err != nil {
        log.Fatalf("Failed to create session: %s", err)
    }

    // Create a new S3 service client
    s3Client := s3.New(newSession)

    // Define the file and bucket
    file := bytes.NewReader([]byte("Hello, DigitalOcean Spaces! This is a S3 Put object example via go!"))
    bucket := "<BUCKET_NAME>"
    key := "example.txt"

    // Put object into the bucket
    _, err = s3Client.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   file,
        ACL:    aws.String("public-read"), // or other ACL settings
    })

    if err != nil {
        log.Fatalf("Failed to put object: %s", err)
    } else {
        fmt.Printf("Successfully uploaded %s to %s/%s\n", key, bucket, key)
    }
}

