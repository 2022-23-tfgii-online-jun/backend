package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MapRepository is the interface that defines the methods for accessing the map data store.
type MapRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)
	CreateWithOmit(omit string, value interface{}) error
	Find(out interface{}, conditions ...interface{}) error
	Update(value interface{}) error

	Delete(out interface{}) error
}

// MapService is the interface that defines the methods for managing maps in the application.
type MapService interface {
	CreateMap(c *gin.Context, createReq *entity.RequestCreateUpdateMap) (*entity.Map, int, error)
	UpdateMap(c *gin.Context, mapUUID uuid.UUID, updateReq *entity.RequestCreateUpdateMap) (*entity.Map, int, error)
	DeleteMap(c *gin.Context, mapUUID uuid.UUID) (int, error)
	GetAllMaps() ([]*entity.Map, error)
}
