package service

import (
	"errors"
	"fmt"
	"victord/daemon/internal/dto"
	vectorEntity "victord/daemon/internal/entity/vector"
	"victord/daemon/internal/store/service"
	storeService "victord/daemon/internal/store/service"
)

func InsertVector(vectorData *dto.InsertVectorRequest, idxName string) (*uint64, error) {

	indexResource, exists := storeService.GetIndex(idxName)
	if !exists {
		return nil, errors.New("index not found")
	}

	vIndex := indexResource.VIndex
	vector := vectorData.Vector
	vectId := vectorData.ID

	dims, dimsExists := storeService.GetIndexDims(idxName)
	if !dimsExists {
		return nil, errors.New("Index dimensions not found")
	}

	if len(vector) != int(dims) {
		return nil, errors.New("Vector dimensions do not match index dimensions")

	}

	if err := vIndex.Insert(vectId, vector); err != nil {
		return nil, errors.New(err.Error())
	}

	return &vectId, nil

}

func DeleteVector(vectorId uint64, idxName string) (*uint64, error) {
	indexResource, exists := storeService.GetIndex(idxName)
	if !exists {
		return nil, errors.New("Index not found")
	}

	vIndex := indexResource.VIndex

	err := vIndex.Delete(vectorId)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &vectorId, nil
}

func SearchVector(vector []*float32, idxName string, topK int) (*vectorEntity.SearchVectorResult, error) {
	fmt.Println("Index name:", idxName)

	indexResource, exists := service.GetIndex(idxName)
	if !exists {
		return nil, errors.New("Index not found")
	}

	vIndex := indexResource.VIndex

	flatVector := make([]float32, len(vector))
	for i, f := range vector {
		if f == nil {
			return nil, errors.New("nil value in vector")
		}
		flatVector[i] = *f
	}

	fmt.Println("Vector to search:", vector)
	fmt.Println("topK:", topK)

	result, err := vIndex.Search(flatVector, topK)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	res := &vectorEntity.SearchVectorResult{
		ID:       result.ID,
		Distance: result.Distance,
	}

	return res, nil
}
