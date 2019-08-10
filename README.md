# Die Mannschaft

Die Mannschaft is Germany's football association that over the past ten years has been including technology into the match field, keeping track of player's statistics, performance, predictions and overall metadata for their internal tournaments.

For a long time they have been keeping all of this data locally, without sharing it with any other organization. In December 2018, the local government has created new laws forcing Die Mannschaft to share all their data with other government institutions. But there's the risk of their infrastructure not being able to support the amount of external requests, they have estimated that at least 5000 organizations will be fetching their data continuously.

What approach would you suggest them to use for efficiently sharing their data with other organizations?

## Solution

**Data Sync Pattern** is useful pattern for data broadcasting that can help integration accross systems.
![C4 Model](https://github.com/rochaeinar/integration-pattern/blob/master/diagram.png)

## Prerequisites
* It requires [Golang](https://golang.org/dl/) programming language
* It requires [Dep](https://github.com/golang/dep) dependency management tool

## Installation

1. Clone or download the project

   ```bash
    git clone https://github.com/rochaeinar/integration-pattern.git
   ```
2. Open a terminal and navigate until the location of the cloned/downloaded project.

## Usage

1. Execute the receiver command in at least in two terminals:

   ```bash
    go run receiver.go
   ```
2. Execute the sender command in a new terminal:
    
   ```bash
    go run sender.go
   ```

3. Review the terminals that has the listener, you should see the messages that the sender has sent.
