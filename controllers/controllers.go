package controllers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/derhabicht/rose-park/database"
	"github.com/derhabicht/rose-park/models"
)

type Controller interface {
	ResourceModel() models.Model
	Create(c *gin.Context)
	List(c *gin.Context)
	Fetch(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type controllerResourceInfo struct {
	ResourceType     reflect.Type
	IDFieldName      string
	IDFieldType      reflect.Type
	IDFieldZeroValue reflect.Value
}

func getControllerResourceInfo(m models.Model) controllerResourceInfo {
	return controllerResourceInfo{
		ResourceType:     reflect.TypeOf(m),
		IDFieldName:      m.GetIDFieldName(),
		IDFieldType:      reflect.TypeOf(m.GetID()),
		IDFieldZeroValue: reflect.Zero(reflect.TypeOf(m.GetID())),
	}
}

func CreateRecord(ctx *gin.Context, ctl Controller) {
	logrus.WithFields(logrus.Fields{
		"resource_method": "create",
		"controller":      reflect.TypeOf(ctl).String(),
	}).Info("Creating new resource in database")

	ri := getControllerResourceInfo(ctl.ResourceModel())
	r := reflect.New(ri.ResourceType)

	err := ctx.Bind(r)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "request body is malformed",
		})
	}
}

// ListRecords uses reflection to fetch a list of model instances from the database.
func ListRecords(ctx *gin.Context, ctl Controller, restrict bool) {
	// TODO: Restrict records by ownership
	// TODO: Panic if the model does not contain an ownership field and restrict is true
	// TODO: Respond with 204 for empty result sets

	logrus.WithFields(logrus.Fields{
		"resource_method": "list",
		"controller":      reflect.TypeOf(ctl).String(),
	}).Info("Fetching list of resources from database")

	ri := getControllerResourceInfo(ctl.ResourceModel())
	s := reflect.MakeSlice(reflect.SliceOf(ri.ResourceType), 0, 0)
	p := reflect.New(s.Type())
	p.Elem().Set(s)
	r := p.Interface()

	logrus.WithFields(logrus.Fields{
		"resource_type":      reflect.TypeOf(ctl.ResourceModel()).String(),
		"resource_list_type": reflect.TypeOf(r).String(),
	}).Debug("Resource reflection info")

	database.DB.Find(r)

	ctx.JSON(http.StatusOK, r)
}

func FetchRecord(c *gin.Context, ctl Controller, param string, preload []string) {
	logrus.WithFields(logrus.Fields{
		"resource_method": "fetch",
		"controller":      reflect.TypeOf(ctl).String(),
	}).Info("Fetching resource from database")

	tx := database.DB.Where(fmt.Sprintf("%s = ?", param), c.Param(param))

	for _, v := range preload {
		tx = tx.Preload(v)
	}

	ri := getControllerResourceInfo(ctl.ResourceModel())
	r := reflect.New(ri.ResourceType)

	tx.First(r.Interface())

	logrus.WithFields(logrus.Fields{
		"resource_type": ri.ResourceType.String(),
		"lookup_param":  param,
		"lookup_value":  c.Param(param),
		"is_valid":      r.IsValid(),
	}).Debug("Resource reflection info")

	logrus.WithFields(logrus.Fields{
		"resource_id_value":      r.Elem().FieldByName(ri.IDFieldName).Interface(),
		"resource_id_type":       ri.IDFieldType.String(),
		"resource_id_name":       ri.IDFieldName,
		"resource_id_zero_value": ri.IDFieldZeroValue,
	}).Debug("Resource ID field reflection info")

	if r.Elem().FieldByName(ri.IDFieldName).Interface() == ri.IDFieldZeroValue.Interface() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "resource not found",
			"id":    c.Param(param),
		})
		return
	}

	c.JSON(http.StatusOK, r.Interface())
}

func UpdateRecord(c *gin.Context) {

}

func DeleteRecord(c *gin.Context) {

}
