package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/pipe.v2"
)

func runBackup() (string, error) {
	//Set Path for Lambda to call local executables
	os.Setenv("PATH", os.Getenv("PATH")+":"+os.Getenv("LAMBDA_TASK_ROOT"))

	//Set Vars
	var exitMessage string

	// Get Env Variables and assign
	fileName := os.Getenv("FILENAME")
	bucketName := os.Getenv("BUCKETNAME")
	mongoHost := os.Getenv("MONGOHOST")
	mongoUsername := os.Getenv("MONGOUSERNAME")
	mongoPW := os.Getenv("MONGOPW")
	authDB := os.Getenv("AUTHDB")
	backupDB := os.Getenv("BACKUPDB")

	// Check if all Env Variables are valid
	if fileName == "" {
		return "", errors.New("No FILENAME Env Var")
	}
	if bucketName == "" {
		return "", errors.New("No BUCKETNAME Env Var")
	}
	if mongoHost == "" {
		return "", errors.New("No MONGOHOST Env Var")
	}
	if mongoUsername == "" {
		return "", errors.New("No MONGOUSERNAME Env Var")
	}
	if mongoPW == "" {
		return "", errors.New("No MONGOPW Env Var")
	}
	if authDB == "" {
		return "", errors.New("No AUTHDB Env Var")
	}
	if backupDB == "" {
		return "", errors.New("No BACKUPDB Env Var")
	}

	//Setup Command 1
	cmd1Name := "mongodump"
	cmd1Args := []string{
		"--host",
		mongoHost,
		"--ssl",
		"--username",
		mongoUsername,
		"--password",
		mongoPW,
		"--authenticationDatabase",
		authDB,
		"--db",
		backupDB,
		"--archive"}

	//Setup Command 2
	cmd2Name := "aws"
	cmd2Args := []string{
		"s3",
		"cp",
		"-",
		"s3://" + bucketName + "/" + fileName}

	//Assign pipe of the two commands
	p := pipe.Line(
		pipe.Exec(cmd1Name, cmd1Args...),
		pipe.Exec(cmd2Name, cmd2Args...),
	)
	//Run the pip and get the output
	output, err := pipe.CombinedOutput(p)
	fmt.Printf("%s", output)
	if err != nil {
		fmt.Printf("%v\n", err)
		return "%v\n", err
	} else {
		fmt.Println("Successfully uploaded to S3")
		exitMessage = "Backup Run Successfully"
	}

	return exitMessage, err
}

func main() {
	lambda.Start(runBackup)
}
