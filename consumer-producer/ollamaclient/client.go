// ollamaclient/client.go
package ollamaclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// ---------- Конфигурация ----------

// Model – название модели по умолчанию. Можно переопределить через переменную окружения OLLAMA_MODEL.
var Model = "llama3.1"

// Host – адрес Ollama‑API. По умолчанию берём из OLLAMA_HOST.
// Если переменная не задана – используем localhost:11434.
func getHost() string {
	if h := os.Getenv("OLLAMA_HOST"); h != "" {
		return h
	}
	return "localhost:11434"
}

// ---------- Внутренние типы ----------

// requestPayload – структура тела POST‑запроса к /api/generate
type requestPayload struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"` // мы работаем в non‑stream режиме
}

// responsePayload – минимум, который нам нужен из ответа
type responsePayload struct {
	Response string `json:"response"` // сгенерированный текст
}

// ---------- Публичная функция ----------

// Generate отправляет prompt к Ollama и возвращает сгенерированный текст.
// Если что‑то пошло не так – возвращается ошибка.
func Generate(prompt string) (string, error) {
	// Выбираем модель
	model := Model
	if envModel := os.Getenv("OLLAMA_MODEL"); envModel != "" {
		model = envModel
	}

	// Формируем тело запроса
	reqBody := requestPayload{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("json marshal: %w", err)
	}

	// Создаём HTTP‑запрос
	url := fmt.Sprintf("http://%s/api/generate", getHost())
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	// Проверяем HTTP‑статус
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non‑200 status %d: %s", resp.StatusCode, string(respBytes))
	}

	// Декодируем JSON‑ответ
	var payload responsePayload
	if err := json.Unmarshal(respBytes, &payload); err != nil {
		return "", fmt.Errorf("json unmarshal: %w", err)
	}

	// Возвращаем сгенерированный текст
	return payload.Response, nil
}
