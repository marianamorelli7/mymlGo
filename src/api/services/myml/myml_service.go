package myml

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"sync"
)
var (
	router = gin.Default()
)


func GetUser(userID int64) (*myml.User, *apierrors.ApiError) {
	if userID == 0 {
		return nil, &apierrors.ApiError{
			"userId invalido",
			http.StatusBadRequest,
		}
	}
	user := &myml.User{ID: userID}
	if apiErr := user.GetUser();
		apiErr != nil {
		return nil, (*apierrors.ApiError)(apiErr)
	}
	return user, nil
}

func GetSite(siteId string) (*myml.Site, *apierrors.ApiError) {
	site := &myml.Site{ID: siteId}
	if apiErr := site.GetSite()
		apiErr != nil {
		return nil, apiErr
	}
	return site, nil
}

func GetCategory(siteId string) (*myml.Category, *apierrors.ApiError)  {
	category  := new(myml.Category)
	if apiErr := category.GetCategory(siteId); apiErr != nil {
		return nil, apiErr
	}
	return category, nil
}

func Get(user *myml.User) (*myml.Myml, *apierrors.ApiError){

	var wg sync.WaitGroup
	c := make(chan *myml.Category,1)
	s := make(chan *myml.Site,1)
	/*
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
*/

	wg.Add(2)
	go func() {
		site, _ := GetSite(user.SiteID)
		s <- site
		wg.Done()
	}()
	go func() {
		category, _ := GetCategory(user.SiteID)
		c <- category
		wg.Done()
	}()
	wg.Wait()

 	response := &myml.Myml{Categories: *<-c, Sites: *<-s}

	return response, nil

}