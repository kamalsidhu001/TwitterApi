﻿# My Twitter API 

## Introduction
This assignment demonstrates how to interact with the Twitter API using Go. It teaches you how to authenticate using OAuth1, post a new tweet, and delete an existing tweet. The program illustrates the practical use of HTTP requests in conjunction with API interactions, making it a valuable learning experience for anyone looking to work with RESTful APIs.

## Setup Instructions

### 1. Create a Twitter Developer Account
To use the Twitter API, you must have a Twitter Developer account. Follow these steps:
- Go to the [Twitter Developer Portal](https://developer.twitter.com/).
- Sign in with your Twitter account or create a new one.
- Apply for a developer account by filling out the required forms and providing necessary information about your intended use of the API.

### 2. Generate API Keys
Once you have a developer account:
- Navigate to the "Projects & Apps" section in the developer portal.
- Create a new project and app.
- After creating the app, you will find your API keys:
  - **API Key** (Consumer Key)
  - **API Secret Key** (Consumer Secret)
  - **Access Token**
  - **Access Token Secret**

### 3. Configure the Program
- Open the `main.go` file in your favorite editor.
- Provide your  `consumerKey`, `consumerSecret`, `accessToken`, and `accessSecret` generated keys.

### 4. Install Packages
For successfully running your code you have to install some packages for your project(gitHub.com/dghubble/oauth1)for authenticating your request.


### 4. Run the Program
To run the program, make sure you have Go installed on your machine. Then, execute the following commands in your terminal:
```bash
go run main.go

