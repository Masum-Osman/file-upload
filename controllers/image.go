package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// ImageController operations for Image
type ImageController struct {
	beego.Controller
}

type Image struct {
	Id       string
	Url      string
	ParamId  string
	MetaData MetaData
}

type MetaData struct {
	Category    string
	Author      string
	LastUpdated string
}

// URLMapping ...
func (c *ImageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Image
// @Param	body		body 	models.Image	true		"body for Image content"
// @Success 201 {object} models.Image
// @Failure 403 body is empty
// @router / [post]
func (c *ImageController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Image by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Image
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ImageController) GetOne() {
	image_id := c.GetString(":id")

	if image_id != "" {
		image := Image{}

		image.Id = "o785uc9e5e32veav656vgd_custom"
		image.ParamId = image_id
		image.Url = "https://s3-ap-southeast-1.amazonaws.com/merchantlogo.pay/merchant-default-logo.png"
		image.MetaData.Author = "ami nije"
		image.MetaData.Category = "human"
		image.MetaData.LastUpdated = "10-1-2501"

		c.Data["json"] = image
		c.ServeJSON()
	}
}

// GetAll ...
// @Title GetAll
// @Description get Image
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Image
// @Failure 403
// @router / [get]
func (c *ImageController) GetAll() {
	image := Image{}

	image.Id = "o785uc9e5e32veav656vgd"
	image.Url = "https://s3-ap-southeast-1.amazonaws.com/merchantlogo.pay/merchant-default-logo.png"
	image.MetaData.Author = "ami nije"
	image.MetaData.Category = "human"
	image.MetaData.LastUpdated = "10-1-2501"

	c.Data["json"] = image
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Image
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Image	true		"body for Image content"
// @Success 200 {object} models.Image
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ImageController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Image
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ImageController) Delete() {

}
