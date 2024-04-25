mepost go sdk
=============

This Go library provides a simple way to interact with the Mepost API, allowing users to send emails, manage scheduled messages, and more using the Mepost platform.

Features
--------

*   Send emails directly and using templates
    
*   Retrieve and cancel scheduled messages
    
*   Query information about specific messages
    

Installation
------------

To install the Mepost Go SDK, you need to have Go installed on your machine (version 1.13+ is required). You can install this package by running:

```bash
go get github.com/yourusername/mepost-go-sdk
```
Usage
-----

Here's how to get started with the Mepost Go SDK:

### Initializing a Client

To use the SDK, you first need to create a client instance with your API key:

```go
package main

import (
    "fmt"
    "github.com/yourusername/mepost-go-sdk"
)

func main() {
    client := mepost.NewClient("your_api_key_here")

    // Now you can use `client` to access Mepost API methods
}
```
### Sending an Email

```go
emailData := map[string]interface{}{
    "from_email": "info@example.com",
    "subject": "Test Email",
    "html": "This is a test email sent from the Mepost Go SDK.",
    "to": []interface{}{
        map[string]interface{}{"email": "recipient1@example.com"},
        map[string]interface{}{"email": "recipient2@example.com"},
    },
}

response, err := client.SendEmail(emailData)
if err != nil {
    fmt.Println("Error sending email:", err)
    return
}
fmt.Println("Email sent successfully:", response)
```

### Handling Scheduled Messages

To retrieve a scheduled message:

```go
scheduleID := "your_schedule_id"
email, err := client.GetScheduledMessage(scheduleID)
if err != nil {
    fmt.Println("Error retrieving message:", err)
    return
}
fmt.Println("Scheduled Message:", email)
```
### API Methods

# **NewClient(apiKey string) \*Client**

*   **Description**: Creates a new client instance with the provided API key.
    
*   **Parameters**:
    
    *   **apiKey**: Your Mepost API key as a string.
        

# **SendEmail(emailData map\[string\]interface{}) (map\[string\]interface{}, error)**

*   **Description**: Sends an email with the provided email data.
    
*   **Parameters**:
    
    *   **emailData**: A map containing the email details such as **from\_email**, **subject**, **text**, and **to**.
        

# **SendEmailByTemplate(emailData map\[string\]interface{}, templateID string) (map\[string\]interface{}, error)**

*   **Description**: Sends an email using a specified template.
    
*   **Parameters**:
    
    *   **emailData**: A map containing the email details.
        
    *   **templateID**: The ID of the template you want to use for sending the email.
        

# **GetInfo(scheduleID string, email string) (map\[string\]interface{}, error)**

*   **Description**: Retrieves information about a specific scheduled message.
    
*   **Parameters**:
    
    *   **scheduleID**: The ID of the scheduled message.
        
    *   **email**: The email address associated with the message.
        

# **CancelScheduledMessage(scheduledMessageID string) (map\[string\]interface{}, error)**

*   **Description**: Cancels a scheduled message.
    
*   **Parameters**:
    
    *   **scheduledMessageID**: The ID of the message you want to cancel.
        

# **GetScheduledMessage(scheduleID string) (map\[string\]interface{}, error)**

*   **Description**: Retrieves a scheduled message.
    
*   **Parameters**:
    
    *   **scheduleID**: The ID of the scheduled message.

Contributing
------------

Contributions to the Mepost Go SDK are welcome! Please feel free to fork the repository, make changes, and submit pull requests.

License
-------

This SDK is distributed under the MIT License, see LICENSE for more information.
