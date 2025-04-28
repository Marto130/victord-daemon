package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"victord/daemon/internal/dto"
	indexEntity "victord/daemon/internal/entity/index"
	"victord/daemon/internal/index/models"
	"victord/daemon/internal/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateIndexHandler(t *testing.T) {
	type mocksIndex struct {
		indexMock *mocks.MockIndexService
	}

	type args struct {
		request string
	}

	tests := []struct {
		name       string
		args       args
		setupMocks func(*mocksIndex)
		wantBody   *dto.CreateIndexResponse
		wantError  string
		wantStatus int
	}{
		{
			name: "test_create_index_bad_request",
			args: args{
				request: `{`,
			},
			setupMocks: func(_ *mocksIndex) {},
			wantError:  "Invalid request",
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "test_create_index_service_error",
			args: args{
				request: `{"index_type":0,"method":0,"dims":5}`,
			},
			setupMocks: func(m *mocksIndex) {
				m.indexMock.EXPECT().CreateIndex(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			wantError:  "error",
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "test_create_index_service_ok",
			args: args{
				request: `{"index_type":0,"method":0,"dims":5}`,
			},
			setupMocks: func(m *mocksIndex) {
				expectedIndex := &models.IndexResource{IndexName: "index_1", IndexID: "1", Dims: 5, IndexType: 1, Method: 2}
				m.indexMock.EXPECT().CreateIndex(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedIndex, nil)
			},
			wantBody: &dto.CreateIndexResponse{
				Status:  "success",
				Message: "Index created successfully",
				Results: indexEntity.CreateIndexResult{
					IndexName: "index_1",
					ID:        "1",
					Dims:      5,
					IndexType: 1,
					Method:    2,
				},
			},
			wantStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			indexService := mocks.NewMockIndexService(ctrl)
			mocks := &mocksIndex{
				indexMock: indexService,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(mocks)
			}

			h := &Handler{
				IndexService: indexService,
			}

			request := httptest.NewRequest(http.MethodPost, "/api/index/index_1", strings.NewReader(tt.args.request))
			response := httptest.NewRecorder()

			h.CreateIndexHandler(response, request)
			result := response.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.wantStatus, result.StatusCode)

			currentBody, err := io.ReadAll(result.Body)
			assert.NoError(t, err)

			if tt.wantBody != nil {
				var expectedBody dto.CreateIndexResponse
				err := json.Unmarshal(currentBody, &expectedBody)
				assert.NoError(t, err)
				assert.Equal(t, tt.wantBody, &expectedBody)
			}
			if tt.wantError != "" {
				assert.Contains(t, string(tt.wantError), strings.TrimSpace(string(currentBody)))
			}
		})
	}
}
