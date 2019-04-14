# Go home weather station

I was spending too much on my AWS MySQL RDS instance with my flask raspberry pi  home weather station which saved a temp reading every 15 minutes with a Celery heartbeat. This project serves to re-write that functionality to log constantly, but only save once a day (since you're billed by the hour on the RDS instance) and an excuse for me to get to know vim a bit better.

## Package Overviews

* **config**: environment based configuration items
* **gin**: core REST API package with all route handlers
* **gorm**: persistence layer with connection setup for mysql
* **gin_suite_test**: ginkgo packaged test suite for running integration style tests on API endpoints

## Configuration

The `config/localConfig.go` is used as the local configuration file.  This file is gitignored and should stay up to date with the `config/localConfig_sample.go` file.

## Adding Routes

Routes are defined in `gin/routes.go` are initialized in the router on startup of the application.  

## Adding Models

All models are defined in the `gorm` package.  Each model should have its own file and will be auto-migrated on initialization of the database.

## yo generator-gin-api

Project structure generated with Sezzle's Yeoman generator-gin-api: https://github.com/sezzle/generator-gin-api
