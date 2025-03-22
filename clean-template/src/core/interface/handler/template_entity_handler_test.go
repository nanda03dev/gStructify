package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTemplateEntityService mocks TemplateEntityService
type MockTemplateEntityService struct {
	mock.Mock
}

func (m *MockTemplateEntityService) Create(templateEntityDTO dto.CreateTemplateEntityDTO) (*aggregate.TemplateEntity, error) {
	args := m.Called(templateEntityDTO)
	return args.Get(0).(*aggregate.TemplateEntity), args.Error(1)
}

func (m *MockTemplateEntityService) GetById(id string) (*aggregate.TemplateEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*aggregate.TemplateEntity), args.Error(1)
}

func (m *MockTemplateEntityService) FindWithFilter(filter common.FilterQuery) ([]*aggregate.TemplateEntity, error) {
	args := m.Called(filter)
	return args.Get(0).([]*aggregate.TemplateEntity), args.Error(1)
}

func (m *MockTemplateEntityService) Update(id string, templateEntityDTO dto.UpdateTemplateEntityDTO) (*aggregate.TemplateEntity, error) {
	args := m.Called(id, templateEntityDTO)
	return args.Get(0).(*aggregate.TemplateEntity), args.Error(1)
}

func (m *MockTemplateEntityService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestTemplateEntityHandler_Create(t *testing.T) {
	app := fiber.New()

	mockService := new(MockTemplateEntityService)
	handler := NewTemplateEntityHandler(mockService)

	mockTemplateEntity := &aggregate.TemplateEntity{ID: "123"}
	mockService.On("Create", mock.Anything).Return(mockTemplateEntity, nil)

	app.Post("/templateEntity", handler.CreateTemplateEntity)

	payload, _ := json.Marshal(dto.CreateTemplateEntityDTO{})
	req := httptest.NewRequest(http.MethodPost, "/templateEntity", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read and parse response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()

	var response struct {
		Data dto.TemplateEntityResponseDTO `json:"data"`
	}

	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, "123", response.Data.ID)
	mockService.AssertExpectations(t)
}

func TestTemplateEntityHandler_GetTemplateEntityByID(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTemplateEntityService)
	handler := NewTemplateEntityHandler(mockService)

	mockTemplateEntity := &aggregate.TemplateEntity{ID: "123"}
	mockService.On("GetById", "123").Return(mockTemplateEntity, nil)

	app.Get("/templateEntity/:id", handler.GetTemplateEntityByID)

	req := httptest.NewRequest(http.MethodGet, "/templateEntity/123", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response struct {
		Data dto.TemplateEntityResponseDTO `json:"data"`
	}
	json.Unmarshal(body, &response)

	assert.Equal(t, "123", response.Data.ID)

	mockService.AssertExpectations(t)
}

func TestTemplateEntityHandler_FindTemplateEntityWithFilter(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTemplateEntityService)
	handler := NewTemplateEntityHandler(mockService)

	app.Post("/templateEntity/filter", handler.FindTemplateEntityWithFilter)

	mockTemplateEntitys := []*aggregate.TemplateEntity{
		{ID: "123"},
		{ID: "124"},
	}
	mockService.On("FindWithFilter", mock.Anything).Return(mockTemplateEntitys, nil)

	payload, _ := json.Marshal(common.FilterQuery{})
	req := httptest.NewRequest(http.MethodPost, "/templateEntity/filter", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response struct {
		Data []dto.TemplateEntityResponseDTO `json:"data"`
	}
	json.Unmarshal(body, &response)

	assert.Len(t, response.Data, 2)
	assert.Equal(t, "123", response.Data[0].ID)
	assert.Equal(t, "124", response.Data[1].ID)

	mockService.AssertExpectations(t)
}

func TestTemplateEntityHandler_UpdateTemplateEntityById(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTemplateEntityService)
	handler := NewTemplateEntityHandler(mockService)

	app.Put("/templateEntity/:id", handler.UpdateTemplateEntityById)

	mockTemplateEntity := &aggregate.TemplateEntity{ID: "123"}
	mockService.On("GetById", "123").Return(mockTemplateEntity, nil)
	mockService.On("Update", "123", mock.Anything).Return(mockTemplateEntity, nil)

	payload, _ := json.Marshal(dto.UpdateTemplateEntityDTO{})
	req := httptest.NewRequest(http.MethodPut, "/templateEntity/123", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response struct {
		Data dto.TemplateEntityResponseDTO `json:"data"`
	}
	json.Unmarshal(body, &response)

	assert.Equal(t, "123", response.Data.ID)

	mockService.AssertExpectations(t)
}

// 5️⃣ Test: Delete TemplateEntity
func TestTemplateEntityHandler_DeleteTemplateEntityById(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTemplateEntityService)
	handler := NewTemplateEntityHandler(mockService)

	app.Delete("/templateEntity/:id", handler.DeleteTemplateEntityById)

	mockService.On("GetById", "123").Return(&aggregate.TemplateEntity{ID: "123"}, nil)
	mockService.On("Delete", "123").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/templateEntity/123", nil)

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response struct {
		Data string `json:"data"`
	}
	json.Unmarshal(body, &response)

	assert.Equal(t, common.DataDeletedSuccessfully, response.Data)

	mockService.AssertExpectations(t)
}
