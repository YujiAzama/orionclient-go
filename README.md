# FIWARE Orion Client for Go

This is a client library for calling FIWARE Orion API using Go.

## Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Overview

FIWARE Orion Client provides:

- Create Orion Subscription
- Get Orion Subscriptions
- Retrieve Orion Subscriptions
- Delete Orion Subscription
- Using Subscription as an object in Go

## Installing

Using FIWARE Orion Client is easy. First, use go get to install the latest version of the library.

```bash
go get -u github.com/YujiAzama/orionclient-go/orionclient
```

Next, include Orion Client in your application:

```go
import "github.com/YujiAzama/orionclient-go/orionclient"
```

## Getting Started

First, define the Configuration of OrionClient. 

```go
config := orionclient.ClientConfig{Host: "localhost", Port: 1026, Token: "access-token"}

client, err := orionclient.NewClient(config)

if err != nil {
        panic(err)
}

```

To get all subscriptions, call as follows:

```go
var service = "myservice"          // FIWARE Service
var servicepath = "myservicepath"  // FIWARE Service Path

subscriptions, err := client.GetSubscriptions(context.Background(), service, servicepath)

if err != nil {
        panic(err)
}

for _, subscription := range subscriptions {
        fmt.Println(subscription)
}
```

To get a subscription, call as follows:

```go
subscription, err := client.GetSubscription(context.Background(), "5f0be8bcd9d315f846e98f8d", service, servicepath)
fmt.Println(subscription)
```

To create a subscription, call as follows:

```go
jsondata := `{"description":"Go client test","subject":{"entities":[{"id":"Room1","type":"Room"}]},"notification":{"http":{"url":"http://test.com/v2/notify"}}}`
var subs orionclient.Subscription
json.Unmarshal([]byte(jsondata), &subs)
subscriptionId, _ := client.CreateSubscription(context.Background(), subs, service, servicepath)
fmt.Ptintln(subscriptionId)
```

To delete a subscription, call as follows:

```go
var subscriptionId = "5f0be8bcd9d315f846e98f8d"
client.DeleteSubscription(context.Background(), subscriptionId, service, servicepath)
```

## Contributing

1. Fork it
2. Download your fork to your PC (git clone https://github.com/your_username/orionclient-go)
3. Create your feature branch (git checkout -b my-new-feature)
4. Make changes and add them (git add .)
5. Commit your changes (git commit -m 'Add some feature')
6. Push to the branch (git push origin my-new-feature)
7. Create new pull request

## License

orionclient-go is released under the Apache 2.0 license. See [LICENSE](https://github.com/YujiAzama/orionclient-go/LICENSE)
