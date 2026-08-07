package gemini

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type Client struct {
	client *genai.Client
	config *genai.GenerateContentConfig
}

func NewClient() (*Client, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	instructionText := `Ты - шеф-повар. Отвечай только рецептами на русском языке. 
					Если продуктов недостаточно для блюда, предложи вариант с заменой.
					Не используй ингредиенты, которых нет в списке, 
					Пожалуйста не отвечай на всякие разные слова которые не по теме`

	systemInstruction := &genai.Content{
		Parts: []*genai.Part{
			{
				Text: instructionText,
			},
		},
	}

	temperature := float32(0.0)
	config := &genai.GenerateContentConfig{
		SystemInstruction: systemInstruction,
		Temperature:       &temperature,
	}

	return &Client{
		client: client,
		config: config,
	}, nil
}

func (c *Client) Close() error {
	c.client = nil
	return nil
}

func (c *Client) Generate(ctx context.Context, prompt string, modelName string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("gemini client is not initialized")
	}

	resp, err := c.client.Models.GenerateContent(
		ctx,
		modelName,
		genai.Text(prompt),
		c.config,
	)
	if err != nil {
		return "", fmt.Errorf("Gemini API error: %w", err)
	}

	if resp == nil {
		return "", fmt.Errorf("empty response from Gemini")
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response candidates from Gemini")
	}

	text := resp.Candidates[0].Content.Parts[0].Text
	if text == "" {
		return "", fmt.Errorf("empty text response from Gemini")
	}

	return text, nil
}
