package maps

import (
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repo ports.MapRepository
}

// NewService returns a new instance of the map service with the given map repository.
func NewService(mapRepo ports.MapRepository) ports.MapService {
	return &service{
		repo: mapRepo,
	}
}

// CreateMap is the service for creating a map and saving it in the database.
func (s *service) CreateMap(c *gin.Context, createReq *entity.RequestCreateUpdateMap) (*entity.Map, int, error) {
	// Create a new map entity from the request data.
	newMap := &entity.Map{
		Name:              createReq.Name,
		Latitude:          createReq.Latitude,
		Longitude:         createReq.Longitude,
		Type:              createReq.Type,
		HoursAvailability: createReq.HoursAvailability,
		Phone:             createReq.Phone,
		IsPublished:       createReq.IsPublished,
	}

	// Save the map to the database.
	err := s.repo.CreateWithOmit("uuid", newMap)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error creating map: %s", err)
	}

	return newMap, http.StatusOK, nil
}

// UpdateMap is the service for updating a map in the database.
func (s *service) UpdateMap(c *gin.Context, mapUUID uuid.UUID, updateReq *entity.RequestCreateUpdateMap) (*entity.Map, int, error) {
	// Find the existing map by UUID.
	mapEntity := &entity.Map{}
	foundMap, err := s.repo.FindByUUID(mapUUID, mapEntity)
	if err != nil {
		// Return an error if the map is not found.
		return nil, http.StatusNotFound, err
	}
	mapEntity, ok := foundMap.(*entity.Map)
	if !ok {
		return nil, http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Update the map fields with the new data from the update request.
	mapEntity.Name = updateReq.Name
	mapEntity.Latitude = updateReq.Latitude
	mapEntity.Longitude = updateReq.Longitude
	mapEntity.Type = updateReq.Type
	mapEntity.HoursAvailability = updateReq.HoursAvailability
	mapEntity.Phone = updateReq.Phone
	mapEntity.IsPublished = updateReq.IsPublished

	// Update the map in the database.
	err = s.repo.Update(mapEntity)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error updating map: %s", err)
	}

	// Return the updated map.
	return mapEntity, http.StatusOK, nil
}

// DeleteMap deletes a map from the database by its UUID.
func (s *service) DeleteMap(c *gin.Context, mapUUID uuid.UUID) (int, error) {
	// Find the existing map by UUID.
	mapEntity := &entity.Map{}
	foundMap, err := s.repo.FindByUUID(mapUUID, mapEntity)
	if err != nil {
		// Return an error if the map is not found.
		return http.StatusNotFound, err
	}
	mapEntity, ok := foundMap.(*entity.Map)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Delete the map from the repository.
	err = s.repo.Delete(mapEntity)
	if err != nil {
		// Return an error response if there was an issue deleting the map.
		return http.StatusInternalServerError, fmt.Errorf("failed to delete map")
	}

	// Return a success response.
	return http.StatusOK, nil
}

// GetAllMaps returns all maps stored in the database.
func (s *service) GetAllMaps() ([]*entity.Map, error) {
	// Get all maps from the database.
	var maps []*entity.Map
	if err := s.repo.Find(&maps); err != nil {
		return nil, err
	}
	return maps, nil
}
