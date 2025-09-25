// backend/internal/services/llm_service.go
package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/yourname/csvxlchart/backend/internal/models"
)

type LLMService interface {
	Insight(colX, colY string, xs []string, ys []float64, stats models.Stats) string
	InsightGrouped(colX, colY, groupBy, agg string, xLabels []string, series []models.Series, stats models.Stats) string
}

type llmService struct {
	apiKey string
	client *http.Client
	model  string
}

func NewLLMService(apiKey string) LLMService {
	return &llmService{
		apiKey: apiKey,
		client: &http.Client{Timeout: 20 * time.Second},
		model:  "gpt-4o-mini", // adjust to your OpenAI model
	}
}

func (l *llmService) Insight(colX, colY string, xs []string, ys []float64, stats models.Stats) string {
	if l.apiKey == "" {
		return "Sample insight: On the X-axis (" + colX + ") and Y-axis (" + colY + "), the mean ≈ " +
			f2(stats.Mean) + " (min " + f2(stats.Min) + ", max " + f2(stats.Max) + ")."
	}
	return l.callLLM(
		"You are a helpful data analyst. Reply in English, concise (2–3 sentences). Explicitly reference the X-axis ("+colX+") and the Y-axis ("+colY+").",
		buildPrompt(colX, colY, xs, ys, stats),
	)
}

func (l *llmService) InsightGrouped(colX, colY, groupBy, agg string, xLabels []string, series []models.Series, stats models.Stats) string {
	if l.apiKey == "" {
		return "Sample insight (grouped by " + groupBy + "): Considering X-axis (" + colX + ") and Y-axis (" + colY + "), the overall mean ≈ " +
			f2(stats.Mean) + " with visible differences across groups."
	}

	names := make([]string, 0, len(series))
	for _, s := range series {
		names = append(names, s.Name)
	}

	userMsg := "We plotted the Y-axis (" + colY + ") against the X-axis (" + colX + "), grouped by " + groupBy +
		" using the " + agg + " aggregation.\n" +
		"X labels: " + strings.Join(xLabels, ", ") + ".\n" +
		"Groups: " + strings.Join(names, ", ") + ".\n" +
		"Overall Y summary: N=" + i2(stats.N) + ", mean=" + f2(stats.Mean) + ", min=" + f2(stats.Min) + ", max=" + f2(stats.Max) + ".\n" +
		"Write 2–3 sentences in English that explicitly reference the X-axis ("+colX+") and Y-axis ("+colY+"), highlighting group differences, peaks/lows, or notable patterns."

	return l.callLLM(
		"You are a helpful data analyst. Reply in English, concise (2–3 sentences).",
		userMsg,
	)
}

func (l *llmService) callLLM(sys, usr string) string {
	type message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	payload := map[string]any{
		"model": l.model,
		"messages": []message{
			{Role: "system", Content: sys},
			{Role: "user", Content: usr},
		},
		"temperature": 0.2,
	}
	b, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.apiKey)

	resp, err := l.client.Do(req)
	if err != nil {
		return "Failed to call LLM: " + err.Error()
	}
	defer resp.Body.Close()

	var out struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil || len(out.Choices) == 0 {
		return "No response from LLM."
	}
	return strings.TrimSpace(out.Choices[0].Message.Content)
}

// --- helpers ---

func buildPrompt(colX, colY string, xs []string, ys []float64, s models.Stats) string {
	return "We are visualizing Y-axis (" + colY + ") vs X-axis (" + colX + "). " +
		"Y summary: N=" + i2(s.N) + ", mean=" + f2(s.Mean) +
		", min=" + f2(s.Min) + ", max=" + f2(s.Max) + ". " +
		"Write 2–3 sentences in English, explicitly referencing the X-axis ("+colX+") and Y-axis ("+colY+"), focusing on trends, outliers, or patterns."
}

func f2(f float64) string {
	s := strconv.FormatFloat(f, 'f', 4, 64)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	if s == "" || s == "-0" {
		return "0"
	}
	return s
}

func i2(i int) string {
	return strconv.Itoa(i)
}
