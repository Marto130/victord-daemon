package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"victord/daemon/pkg/models"

	"victord/daemon/internal/services"

	"github.com/gorilla/mux"
)

func InsertVectorHandler(w http.ResponseWriter, r *http.Request) {
	var req *models.InsertVectorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error decoding request:", err)
		http.Error(w, "Invalid insert vector request payload", http.StatusBadRequest)
		return
	}

	urlParams := mux.Vars(r)
	indexNameParam := urlParams["indexName"]

	vecId, err := services.InsertVector(req, indexNameParam)
	if err != nil {
		fmt.Println("Error inserting vector:", err)
		http.Error(w, "Failed to insert vector", http.StatusInternalServerError)
		return
	}

	res := models.InsertVectorResponse{
		Status:  "success",
		Message: "Vector inserted successfully",
		Results: models.InsertVectorResult{
			ID:     *vecId,
			Vector: req.Vector,
		},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}

func DeleteVectorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	indexName := vars["indexName"]
	vectorID := vars["vectorID"]

	vectId, err := strconv.ParseUint(vectorID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid vector ID", http.StatusBadRequest)
		return
	}

	services.DeleteVector(vectId, indexName)

	res := models.DeleteVectorResponse{
		Status:  "200",
		Message: "Vector deleted successfully",
		Results: models.DeleteVectorResult{
			ID: vectId,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func SearchVectorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--SEARCH HANDLER--")
	vars := mux.Vars(r)
	indexName := vars["indexName"]

	vectorParam := r.URL.Query().Get("vector")
	if vectorParam == "" {
		http.Error(w, "Missing vector parameter", http.StatusBadRequest)
		return
	}

	vectorChunks := strings.Split(vectorParam, ",")
	var vector []*float32

	for _, p := range vectorChunks {
		val, err := strconv.ParseFloat(strings.TrimSpace(p), 32)
		if err != nil {
			http.Error(w, "Invalid vector value: "+p, http.StatusBadRequest)
			return
		}
		v := float32(val)
		vector = append(vector, &v)
	}

	fmt.Println("Vector to search:", vector)

	kParam := r.URL.Query().Get("top_k")
	var k int
	var err error
	if kParam != "" {
		k, err = strconv.Atoi(kParam)
		if err != nil {
			http.Error(w, "Invalid k parameter", http.StatusBadRequest)
			return
		}
		if k <= 0 {
			http.Error(w, "k must be greater than 0", http.StatusBadRequest)
			return
		}
	} else {
		k = 5
	}

	fmt.Println("topK:", k)

	result, err := services.SearchVector(vector, indexName, k)
	if err != nil {
		fmt.Println("Error searching vector:", err)
		http.Error(w, "Failed to search vector", http.StatusInternalServerError)
		return
	}

	res := models.SearchVectorResponse{
		Status:  "success",
		Message: "Vector search completed successfully",
		Results: models.SearchVectorResult{
			ID:       result.ID,
			Distance: result.Distance,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
