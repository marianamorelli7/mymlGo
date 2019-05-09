package myml

import (
	"github.com/gin-gonic/gin"
	services "github.com/mercadolibre/myml/src/api/services/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
)


var (
	router = gin.Default()
)

func Get(context *gin.Context) {
	userID := context.Param("userID")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		context.JSON(apiErr.Status, apiErr)
		return
	}
	user, apiErr := services.GetUser(id)
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}

	response, apiErr := services.Get(user)
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}
	context.JSON(http.StatusOK, response)
}


/*
func Get(context *gin.Context) {
	var wg sync.WaitGroup
	c := make(chan *domain.Category,1)
	s := make(chan *domain.Site,1)
	userID := context.Param("userID")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		context.JSON(apiErr.Status, apiErr)
		return
	}
	user, apiErr := services.GetUser(id)
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}

	wg.Add(2)
	go func() {
		site, _ := services.GetSite(user.SiteID)
		s <- site
		wg.Done()
	}()
	go func() {
		category, _ := services.GetCategory(user.SiteID)
		c <- category
		wg.Done()
	}()
	wg.Wait()

	response := &domain.Myml{Categories: *<-c, Sites: *<-s}

	context.JSON(http.StatusOK, response)
}
*/
