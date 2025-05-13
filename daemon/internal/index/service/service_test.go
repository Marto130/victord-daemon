package service

/*import (
	"context"
	"errors"
	"reflect"
	"testing"
	"victord/daemon/internal/dto"
	"victord/daemon/internal/index/models"
	"victord/daemon/internal/mocks"

	gm "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_indexService_CreateIndex(t *testing.T) {
	type mocksIndex struct {
		store       *mocks.MockIndexStore
		indexOps    *mocks.MockIndexOps
		vectorIndex *mocks.MockVectorOps
	}

	type args struct {
		ctx  context.Context
		idx  *dto.CreateIndexRequest
		name string
	}
	tests := []struct {
		name       string
		args       args
		want       func(*mocksIndex) *models.IndexResource
		wantErr    bool
		setupMocks func(*mocksIndex)
	}{
		{
			name: "test_alloc_index_error",
			args: args{
				ctx:  context.TODO(),
				name: "index_test",
				idx:  &dto.CreateIndexRequest{},
			},
			setupMocks: func(m *mocksIndex) {
				m.indexOps.EXPECT().AllocIndex(gm.Any(), gm.Any(), gm.Any()).
					Return(nil, errors.New("error"))

			},
			want: func(_ *mocksIndex) *models.IndexResource {
				return nil
			},
			wantErr: true,
		},
		{
			name: "test_alloc_index_ok",
			args: args{
				ctx:  context.TODO(),
				name: "index_test",
				idx: &dto.CreateIndexRequest{
					IndexType: 1,
					Method:    1,
					Dims:      uint16(2),
				},
			},
			setupMocks: func(m *mocksIndex) {
				m.indexOps.EXPECT().AllocIndex(gm.Any(), gm.Any(), gm.Any()).
					Return(m.vectorIndex, nil)
				m.store.EXPECT().StoreIndex(gm.Any())
			},
			want: func(mi *mocksIndex) *models.IndexResource {
				return &models.IndexResource{
					IndexType: 1,
					Method:    1,
					Dims:      uint16(2),
					VIndex:    mi.vectorIndex,
					IndexName: "index_test",
					IndexID:   "",
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gm.NewController(t)
			defer ctrl.Finish()

			storeMock := mocks.NewMockIndexStore(ctrl)
			indexMock := mocks.NewMockIndexOps(ctrl)
			vectorIndexMock := mocks.NewMockVectorOps(ctrl)

			mocks := &mocksIndex{
				store:       storeMock,
				indexOps:    indexMock,
				vectorIndex: vectorIndexMock,
			}

			if tt.setupMocks != nil {
				tt.setupMocks(mocks)
			}

			i := NewIndexService(storeMock, indexMock)
			got, err := i.CreateIndex(tt.args.ctx, tt.args.idx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("indexService.CreateIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := tt.want(mocks)
			if want != nil {
				assert.NotNil(t, want.IndexID, "indexService.CreateIndex() generated vectorID should not be nil")
				got.IndexID = ""
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("indexService.CreateIndex() = %v, want %v", got, want)
			}
		})
	}
}
*/
