# Lambda-GO-MongoDB-S3-Backup

This is a Lambda function to backup your MongoDB database as an archive to S3. This pipes the command instead of creating a temp file. The hope is to overcome the Lambda ephemeral disk capacity.

## Getting Started

Download the lambda-function.zip

Upload the zip to your Lambda function.

Fill out the required environment variables.

```
FILENAME
BUCKETNAME
MONGOHOST
MONGOUSERNAME
MONGOPW
AUTHDB
BACKUPDB
```

Add S3 policy to lambda role to allow Lambda to access S3.

Configure Lambda timeout according to DB size and network speed.

## Tests

Currently I have only tested with fairly small databases from MongoDB's cloud Atlas to AWS US-East-1.

## Development

- Edit main.go.
- Run below.
```
GOOS=linux GOARCH=amd64 go build -o main main.go
```
- Overwrite the "main" in the executables folder with your new "main".
- Zip contents.
- Upload to Lambda function.

## Authors

* **Ben May** - [TheBenMay](https://github.com/TheBenMay)

