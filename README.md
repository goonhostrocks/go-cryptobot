# **go-cryptobot**

An idiomatic, lightweight, and type-safe Go API wrapper for the [Crypto Pay API](https://help.send.tg/ru/articles/10279948-crypto-pay-api) (Crypto Bot on Telegram).

Unlike other bloated wrappers, go-cryptobot has zero external dependencies, provides native structures matching the API, and offers a robust, framework-agnostic webhook verification suite.

## **Features**

* **No Bloat:** Built purely using Go's robust standard library.  
* **Type-Safe:** No manual map parsing. Everything from configuration options to responses maps to clean Go structs.  
* **Unified Package Layout:** No circular dependencies. The core client and models live in the root package, while webhooks live in their own lightweight sub-package.  
* **Framework-Agnostic Webhooks:** Easy integrations with **Fiber, Gin, Chi, Echo**, or standard net/http handlers.  
* **Secure Webhook Verification:** Employs the required double-SHA256 HMAC signature verification schema natively.

## **Installation**

Run this command inside your Go project directory:

go get oss.goonhost.rocks/go-cryptobot

## **Usage**

### **1\. Initializing the Client & Getting Info**

```go
package main

import (  
	"fmt"  
	"log"

	gocryptobot "oss.goonhost.rocks/go-cryptobot"  
)

func main() {  
	// Use gocryptobot.NewClient(token, baseURL)  
	// BaseURL should be "\[https://testnet-pay.crypt.bot/api\](https://testnet-pay.crypt.bot/api)" for testnet   
	// or "\[https://pay.crypt.bot/api\](https://pay.crypt.bot/api)" for mainnet.  
	client := gocryptobot.NewClient("YOUR\_API\_TOKEN", "\[https://testnet-pay.crypt.bot/api\](https://testnet-pay.crypt.bot/api)")

	me, err := client.GetMe()  
	if err \!= nil {  
		log.Fatalf("Error getting bot info: %v", err)  
	}

	fmt.Printf("Bot Connected\! Name: %s (ID: %d)\\n", me.Result.Name, me.Result.AppID)  
}
```

### **2\. Creating and Deleting Invoices**

```go
package main

import (  
	"fmt"  
	"log"

	gocryptobot "oss.goonhost.rocks/go-cryptobot"  
)

func main() {  
	client := gocryptobot.NewClient("YOUR\_API\_TOKEN", "\[https://testnet-pay.crypt.bot/api\](https://testnet-pay.crypt.bot/api)")

	// Create an optional boolean setting  
	allowComments := true

	// Build the request payload  
	req := \&gocryptobot.CreateInvoiceRequest{  
		CurrencyType:  "crypto",  
		Asset:         "TON",  
		Amount:        "1.50",  
		Description:   "Premium Subscription",  
		AllowComments: \&allowComments,  
	}

	// 1\. Create the invoice  
	invoice, err := client.CreateInvoice(req)  
	if err \!= nil {  
		log.Fatalf("Failed to create invoice: %v", err)  
	}  
	fmt.Printf("Invoice Created\! Pay here: %s\\n", invoice.Result.BotInvoiceURL)

	// 2\. Delete the invoice (if needed)  
	delRes, err := client.DeleteInvoice(\&gocryptobot.DeleteInvoiceRequest{  
		InvoiceID: invoice.Result.InvoiceID,  
	})  
	if err \!= nil {  
		log.Fatalf("Failed to delete invoice: %v", err)  
	}  
	  
	if delRes.Result {  
		fmt.Println("Invoice successfully deleted\!")  
	}  
}
```

### **3\. Verification & Processing Webhooks**

Because developers prefer custom routers like **Fiber** or **Gin** over Go's default http.ListenAndServe, our webhook utilities accept the raw request body bytes and signature header directly.

Here is an example using **Fiber**:

```
package main

import (  
	"log"

	"\[github.com/gofiber/fiber/v2\](https://github.com/gofiber/fiber/v2)"  
	"oss.goonhost.rocks/go-cryptobot/webhook"  
)

func main() {  
	app := fiber.New()

	app.Post("/crypto-webhook", func(c \*fiber.Ctx) error {  
		body := c.Body() // Extract raw request bytes  
		signature := c.Get("crypto-pay-api-signature") // Extract the signature header  
		token := "YOUR\_API\_TOKEN"

		// Safely parse and verify signature in one method  
		update, err := webhook.ParseAndVerify(body, token, signature)  
		if err \!= nil {  
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")  
		}

		// Ensure it's the expected update type  
		if update.UpdateType \== "invoice\_paid" {  
			log.Printf("Payment Received\! Invoice \#%d for %s %s was paid\!\\n",   
				update.Payload.InvoiceID,   
				update.Payload.Amount,   
				update.Payload.Asset,  
			)  
		}

		return c.SendStatus(fiber.StatusOK)  
	})

	log.Fatal(app.Listen(":3000"))  
}
```

## **Contributing**

Contributions are welcome\! Please feel free to open issues or submit Pull Requests for any missing endpoints.

Before submitting your PR, make sure your tests pass cleanly:

go test \-v ./...

## **License**

This project is licensed under the terms of the **GNU General Public License v3.0 (GPLv3)**. See the [LICENSE](http://docs.google.com/LICENSE) file for more details.

*The GPLv3 guarantees that this library remains open-source and that any derivative works or modifications also remain free and open to the developer community.*