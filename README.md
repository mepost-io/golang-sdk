
mepost Go SDK
==========

The `mepost-sdk` is a Go library designed to simplify interactions with the Mepost API. It provides convenient methods to send and manage messages efficiently. This SDK is perfect for developers looking to integrate Mepost messaging capabilities into their Go applications.

Features
--------

-   Send emails directly through the Mepost API.
-   Schedule and manage email deliveries.
-   Retrieve detailed information about scheduled messages.
-   Cancel scheduled messages.
-   Send emails using predefined templates.

Installation
------------

Install the `mepost-sdk` by running:

```bash
go get github.com/yourusername/mepost-sdk
```

Usage
-----

Here is a quick example to get you started:

```go
package main

import (
    "fmt"
    "github.com/yourusername/mepost-sdk"
)

func main() {
    // Create an instance of Mepost Client with your API key
    client := mepost.NewClient("your_api_key_here")

    // Send an email
    emailData := map[string]interface{}{
        "from_email": "info@example.com",
        "from_name": "Example Company",
        "html": "This is a test email sent from the Mepost Go SDK.",
        "subject": "Example Subject",
        "to": []map[string]string{
            {"email": "recipient1@example.com"},
            {"email": "recipient2@example.com"},
        },
    }

    response, err := client.SendEmail(emailData)
    if err != nil {
        fmt.Println("Error sending email:", err)
        return
    }

    fmt.Println("Email sent successfully:", response)
}
```

API Methods
-----------

### `NewClient(apiKey string)`

Initializes and returns a new instance of MepostClient.

-   Parameters
    -   `apiKey`: Your Mepost API key.

### Company Endpoints

#### `AddDomain(request AddDomainRequest)`

Adds a domain to the Mepost account.

-   Parameters
    -   `request`: An object containing the domain to add.

#### `GetDomainList()`

Retrieves a list of domains associated with the Mepost account.

-   No parameters.

#### `RemoveDomain(request RemoveDomainRequest)`

Removes a domain from the Mepost account.

-   Parameters
    -   `request`: An object containing the domain to remove.

### Groups Endpoints

#### `ListGroups(limit int, page int)`

Retrieves a list of email groups.

-   Parameters
    -   `limit`: The maximum number of groups to return (default: 10).
    -   `page`: The page number for pagination (default: 1).

#### `CreateGroup(request CreateNewGroupRequest)`

Creates a new email group.

-   Parameters
    -   `request`: An object containing the details of the new group.

#### `DeleteGroup(groupId string)`

Deletes an email group.

-   Parameters
    -   `groupId`: The ID of the group to delete.

#### `GetGroupById(groupId string)`

Retrieves information about a specific email group.

-   Parameters
    -   `groupId`: The ID of the group to retrieve.

#### `UpdateGroup(groupId string, request RenameGroupRequest)`

Updates the name of an email group.

-   Parameters
    -   `groupId`: The ID of the group to update.
    -   `request`: An object containing the new group name.

### Subscribers Endpoints

#### `ListSubscribers(groupId string, limit int, page int)`

Retrieves a list of subscribers in a group.

-   Parameters
    -   `groupId`: The ID of the group.
    -   `limit`: The maximum number of subscribers to return (default: 10).
    -   `page`: The page number for pagination (default: 1).

#### `AddSubscriber(groupId string, request CreateSubscriberRequest)`

Adds a subscriber to a group.

-   Parameters
    -   `groupId`: The ID of the group.
    -   `request`: An object containing subscriber details.

#### `DeleteSubscriber(groupId string, request DeleteSubscriberRequest)`

Deletes a subscriber from a group.

-   Parameters
    -   `groupId`: The ID of the group.
    -   `request`: An object containing the emails of subscribers to delete.

#### `GetSubscriberByEmail(groupId string, email string)`

Retrieves subscriber details by email.

-   Parameters
    -   `groupId`: The ID of the group.
    -   `email`: The email address of the subscriber.

### Messages Endpoints

#### `GetMessageInfo(scheduleId string, email string)`

Retrieves information about a specific scheduled message.

-   Parameters
    -   `scheduleId`: The ID of the scheduled message.
    -   `email`: The email address to which the message was sent.

#### `CancelScheduledMessage(request CancelScheduledMessageRequest)`

Cancels a scheduled message.

-   Parameters
    -   `request`: An object containing the scheduled message ID.

#### `SendMarketing(request SendMarketingRequest)`

Sends a marketing email.

-   Parameters
    -   `request`: An object for sending marketing emails.

#### `SendMessageByTemplate(request SendMessageByTemplateRequest)`

Sends an email using a template.

-   Parameters
    -   `request`: An object containing the message details and template ID.

#### `GetScheduleInfo(scheduleId string)`

Retrieves schedule information for a specific scheduled message.

-   Parameters
    -   `scheduleId`: The ID of the scheduled message.

#### `SendTransactional(request SendTransactionalRequest)`

Sends a transactional email.

-   Parameters
    -   `request`: An object for sending transactional emails.

#### `SendTransactionalByTemplate(request SendMessageByTemplateRequest)`

Sends a transactional email using a template.

-   Parameters
    -   `request`: An object containing the message details and template ID.

### Outbound IP Endpoints

#### `CreateIpGroup(request CreateIpGroupRequest)`

Creates a new IP group.

-   Parameters
    -   `request`: An object containing the IP group details.

#### `GetIpGroupInfo(name string)`

Retrieves information about a specific IP group.

-   Parameters
    -   `name`: The name of the IP group.

#### `ListIpGroups()`

Retrieves a list of all IP groups.

-   No parameters.

#### `CancelWarmup(request CancelWarmUpRequest)`

Cancels a warmup process for an IP address.

-   Parameters
    -   `request`: An object containing the IP address.

#### `GetIpInfo(ip string)`

Retrieves information about a specific IP address.

-   Parameters
    -   `ip`: The IP address to retrieve.

#### `ListIps()`

Retrieves a list of all IP addresses.

-   No parameters.

#### `SetIpGroup(request SetIpGroupRequest)`

Assigns an IP address to a specific IP group.

-   Parameters
    -   `request`: An object containing the IP address and group details.

#### `StartWarmup(request StartWarmUpRequest)`

Starts a warmup process for an IP address.

-   Parameters
    -   `request`: An object containing the IP address.

Contributing
------------

Contributions are always welcome! Please read the contributing guide for ways to contribute to this project.

License
-------

`mepost-sdk` is released under the MIT License. See the LICENSE file for more details.
