# Upfluence Challenge

To solve the challenge presented by Upfluence, I have created a simple web server that receives a request with 
a duration and a dimension, and returns the average of the dimension in the last n seconds.

The solution is written in Go and is using the standard library to handle the requests. I have divided the code into
two main files, the main.go file that contains the server logic and the analysis.go file that contains the logic to
calculate the average of the dimension.

I used Regex to parse the `timestamp` and `likes` from the data-stream. I would have changed the text stream to a JSON format
to make it easier to parse the data if I had more time, but I wasn't able to do it in the time frame.
I also would have liked to simplify the analysis.go into different components to make it more modular. And created more
tests to cover the code.

Also I know the output asked was in the terminal, but for some reason I wasn't able to connect to the localhost:8080 with 
`curl` so I opted to do it in the browser. It does indeed give the result in the browser and I have put the URL and 
instructions in the Usage section.

## Installation

In the project directory, you can run the following to install the dependencies:

```bash
go mod download
```

## Usage
First run the following command in the terminal to start the server:

```bash
go run .
```
And then go to your browser and paste the following URL:

```bash
http://localhost:8080/analysis?duration=30s&dimension=likes
```
After waiting for 30 seconds, you will see the result in the browser.
