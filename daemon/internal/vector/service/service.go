package service

import (
	"errors"
	"fmt"
	"victord/daemon/internal/dto"
	vectorEntity "victord/daemon/internal/entity/vector"
	"victord/daemon/internal/store/service"
)

type vectorService struct {
	store service.IndexStore
}

func NewVectorService(store service.IndexStore) VectorService {
	return &vectorService{
		store: store,
	}
}

func (v *vectorService) InsertVector(vectorData *dto.InsertVectorRequest, idxName string) (*uint64, error) {

	indexResource, exists := v.store.GetIndex(idxName)
	if !exists {
		return nil, errors.New("index not found")
	}

	vector := vectorData.Vector
	vectId := vectorData.ID

	if len(vector) != int(indexResource.Dims) {
		return nil, errors.New("vector dimensions do not match index dimensions")

	}

	if err := indexResource.VIndex.Insert(vectId, vector); err != nil {
		return nil, errors.New(err.Error())
	}

	return &vectId, nil

}

func (v *vectorService) DeleteVector(vectorId uint64, idxName string) error {
	indexResource, exists := v.store.GetIndex(idxName)
	if !exists {
		return errors.New("index not found")
	}

	err := indexResource.VIndex.Delete(vectorId)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (v *vectorService) SearchVector(vector []float32, idxName string, topK int) (*vectorEntity.SearchVectorResult, error) {
	indexResource, exists := v.store.GetIndex(idxName)
	if !exists {
		return nil, errors.New("index not found")
	}

	fmt.Println("Vector to search: ", vector)
	fmt.Println("topK: ", topK)

	result, err := indexResource.VIndex.Search(vector, topK)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	res := &vectorEntity.SearchVectorResult{
		ID:       result.ID,
		Distance: result.Distance,
	}

	return res, nil
}
