# OpenAI Go

## APIs

- Chat Completion
- Audio Transcription
- Text To Speech

## Examples

### Fiber

```go
package main

import (
	"os"

	"github.com/gabrielluizsf/openai-go/pkg/openai"
	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	System  string `json:"system"`
	Message string `json:"message"`
}

func main() {
	app := fiber.New()
	apiKey := os.Getenv("OPENAI_KEY")
	app.Post("/", func(c *fiber.Ctx) error {
		body := new(RequestBody)
		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errMessage": "Invalid Request Body",
			})
		}
		openai := openai.WithContext(c.Context(), apiKey)
		response, err := openai.ChatGPT(
			"gpt-3.5-turbo",
			[]chat.Message{
				{Role: "system", Content: body.System},
				{Role: "user", Content: body.Message},
			},
		)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errMessage": err.Error(),
			})
		}
		return c.JSON(response)
	})

	app.Listen(":3000")
}

```

### Echo

```go
package main

import (
	"net/http"
	"os"

	"github.com/gabrielluizsf/openai-go/pkg/openai"
	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestBody struct {
	System  string `json:"system"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiKey := os.Getenv("OPENAI_KEY")

	e.POST("/", func(c echo.Context) error {
		body := new(RequestBody)
		if err := c.Bind(body); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"errMessage": "Invalid Request Body"})
		}

		openaiClient := openai.WithContext(c.Request().Context(), apiKey)
		response, err := openaiClient.ChatGPT(
			"gpt-3.5-turbo",
			[]chat.Message{
				{Role: "system", Content: body.System},
				{Role: "user", Content: body.Message},
			},
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"errMessage": err.Error()})
		}

		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
```