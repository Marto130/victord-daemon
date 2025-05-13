package service

/*import (
	"errors"
	"reflect"
	"testing"
	"victord/daemon/internal/dto"
	vectorEntity "victord/daemon/internal/entity/vector"
	"victord/daemon/internal/index/models"
	"victord/daemon/internal/mocks"
	"victord/daemon/platform/types"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type mocksIndex struct {
	store       *mocks.MockIndexStore
	vectorIndex *mocks.MockVectorOps
}

func Test_vectorService_InsertVector(t *testing.T) {
	type args struct {
		vectorData *dto.InsertVectorRequest
		idxName    string
	}

	tests := []struct {
		name       string
		setupMocks func(m *mocksIndex)
		args       args
		want       *uint64
		wantErr    bool
	}{
		{
			name: "index_not_found",
			args: args{
				idxName: "index_1",
				vectorData: &dto.InsertVectorRequest{
					ID:     1,
					Vector: []float32{0.1, 0.2, 0.3, 0.4, 0.5},
				},
			},
			setupMocks: func(m *mocksIndex) {
				m.store.EXPECT().GetIndex("index_1").Return(nil, false)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "dimension_does_not_match",
			args: args{
				idxName: "index_1",
				vectorData: &dto.InsertVectorRequest{
					ID:     1,
					Vector: []float32{0.1, 0.2, 0.3, 0.4, 0.5},
				},
			},
			setupMocks: func(m *mocksIndex) {
				index := &models.IndexResource{IndexName: "index_1", Dims: 3}
				m.store.EXPECT().GetIndex("index_1").Return(index, true)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "index_created_ok",
			args: args{
				idxName: "index_1",
				vectorData: &dto.InsertVectorRequest{
					ID:     1,
					Vector: []float32{0.1, 0.2, 0.3, 0.4, 0.5},
				},
			},
			setupMocks: func(m *mocksIndex) {
				index := &models.IndexResource{IndexName: "index_1", Dims: 5, VIndex: m.vectorIndex}
				m.store.EXPECT().GetIndex("index_1").Return(index, true)
				m.vectorIndex.EXPECT().Insert(uint64(1), gomock.Any()).Return(nil)
			},
			want: func() *uint64 {
				val := uint64(1)
				return &val
			}(),
			wantErr: false,
		},
		{
			name: "index_created_error",
			args: args{
				idxName: "index_1",
				vectorData: &dto.InsertVectorRequest{
					ID:     1,
					Vector: []float32{0.1, 0.2, 0.3, 0.4, 0.5},
				},
			},
			setupMocks: func(m *mocksIndex) {
				index := &models.IndexResource{IndexName: "index_1", Dims: 5, VIndex: m.vectorIndex}
				m.store.EXPECT().GetIndex("index_1").Return(index, true)
				m.vectorIndex.EXPECT().Insert(uint64(1), gomock.Any()).Return(errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := mocks.NewMockIndexStore(ctrl)
			vectorIndexMock := mocks.NewMockVectorOps(ctrl)

			mocks := &mocksIndex{
				store:       storeMock,
				vectorIndex: vectorIndexMock,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(mocks)
			}

			vectorService := NewVectorService(storeMock)
			got, err := vectorService.InsertVector(tt.args.vectorData, tt.args.idxName)
			if (err != nil) != tt.wantErr {
				t.Errorf("vectorService.InsertVector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				require.NotNil(t, got)
				require.Equal(t, *tt.want, *got)
			} else {
				require.Nil(t, got)
			}
		})
	}
}

func Test_vectorService_DeleteVector(t *testing.T) {
	type args struct {
		vectorId uint64
		idxName  string
	}
	tests := []struct {
		name       string
		args       args
		setupMocks func(m *mocksIndex)
		wantErr    bool
	}{
		{
			name: "index_not_found",
			args: args{
				idxName:  "index_1",
				vectorId: 1,
			},
			setupMocks: func(m *mocksIndex) {
				m.store.EXPECT().GetIndex("index_1").Return(nil, false)
			},
			wantErr: true,
		},

		{
			name: "index_deleted_ok",
			args: args{
				idxName:  "index_1",
				vectorId: 1,
			},
			setupMocks: func(m *mocksIndex) {
				index := &models.IndexResource{IndexName: "index_1", VIndex: m.vectorIndex}
				m.store.EXPECT().GetIndex("index_1").Return(index, true)
				m.vectorIndex.EXPECT().Delete(uint64(1)).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "index_deleted_error",
			args: args{
				idxName:  "index_1",
				vectorId: 1,
			},
			setupMocks: func(m *mocksIndex) {
				index := &models.IndexResource{IndexName: "index_1", VIndex: m.vectorIndex}
				m.store.EXPECT().GetIndex("index_1").Return(index, true)
				m.vectorIndex.EXPECT().Delete(uint64(1)).Return(errors.New("error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := mocks.NewMockIndexStore(ctrl)
			vectorIndexMock := mocks.NewMockVectorOps(ctrl)

			mocks := &mocksIndex{
				store:       storeMock,
				vectorIndex: vectorIndexMock,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(mocks)
			}

			vectorService := NewVectorService(storeMock)
			if err := vectorService.DeleteVector(tt.args.vectorId, tt.args.idxName); (err != nil) != tt.wantErr {
				t.Errorf("vectorService.DeleteVector() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_vectorService_SearchVector(t *testing.T) {
	type args struct {
		vector  []float32
		idxName string
		topK    int
	}
	tests := []struct {
		name       string
		args       args
		setupMocks func(m *mocksIndex)
		want       *vectorEntity.SearchVectorResult
		wantErr    bool
	}{
		{
			name: "index_not_found",
			args: args{
				idxName: "index_1",
			},
			setupMocks: func(m *mocksIndex) {
				m.store.EXPECT().GetIndex("index_1").Return(nil, false)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "index_search_ok",
			args: args{
				idxName: "index_1",
				vector:  []float32{0.1, 0.2, 0.3},
				topK:    3,
			},
			setupMocks: func(m *mocksIndex) {
				index := &models.IndexResource{IndexName: "index_1", VIndex: m.vectorIndex}
				m.store.EXPECT().GetIndex("index_1").Return(index, true)
				m.vectorIndex.EXPECT().
					Search(gomock.Any(), 3).
					Return(&types.MatchResult{
						ID:       123,
						Distance: 0.456,
					}, nil)
			},
			want: &vectorEntity.SearchVectorResult{
				ID:       123,
				Distance: 0.456,
			},
			wantErr: false,
		},
		{
			name: "index_search_error",
			args: args{
				idxName: "index_1",
				vector:  []float32{0.1, 0.2, 0.3},
				topK:    3,
			},
			setupMocks: func(m *mocksIndex) {
				index := &models.IndexResource{IndexName: "index_1", VIndex: m.vectorIndex}
				m.store.EXPECT().GetIndex("index_1").Return(index, true)
				m.vectorIndex.EXPECT().Search(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := mocks.NewMockIndexStore(ctrl)
			vectorIndexMock := mocks.NewMockVectorOps(ctrl)

			mocks := &mocksIndex{
				store:       storeMock,
				vectorIndex: vectorIndexMock,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(mocks)
			}

			vectorService := NewVectorService(storeMock)
			got, err := vectorService.SearchVector(tt.args.vector, tt.args.idxName, tt.args.topK)
			if (err != nil) != tt.wantErr {
				t.Errorf("vectorService.SearchVector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vectorService.SearchVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
