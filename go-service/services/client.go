package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"student-report-service/models"
)

var ErrStudentNotFound = errors.New("student not found")

type StudentClient interface {
	FetchStudent(ctx context.Context, id string) (*models.Student, error)
}

type HTTPStudentClient struct {
	BaseURL    string
	ServiceKey string
	client     *http.Client
}

func NewHTTPStudentClient(baseURL, serviceKey string) *HTTPStudentClient {
	return &HTTPStudentClient{
		BaseURL:    baseURL,
		ServiceKey: serviceKey,
		client:     &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *HTTPStudentClient) FetchStudent(ctx context.Context, id string) (*models.Student, error) {
	url := fmt.Sprintf("%s/api/v1/internal/students/%s", c.BaseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Service-Key", c.ServiceKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to reach backend service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrStudentNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("backend returned status %d: %s", resp.StatusCode, string(body))
	}

	var student models.Student
	if err := json.Unmarshal(body, &student); err != nil {
		return nil, fmt.Errorf("failed to parse student data: %w", err)
	}

	return &student, nil
}
