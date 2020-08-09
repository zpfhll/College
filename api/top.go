package api

import (
	"college/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

const getSpecialityPath = "/speciality/:type"

// //専攻の取得のタイプ
// type TopRequestType string
// const (
//     byCollegeName   TopRequestType = "college_name" //大学の名で
// )

//トップの初期化　InitTop
func InitTop(router *gin.Engine) {
	top := router.Group("/top")
	{
		top.POST(getSpecialityPath, getSpeciality)
	}
}

func getSpeciality(c *gin.Context) {
	requestType := c.Param("type")
	if requestType == "college_name" {
		collegeName := c.PostForm("college_name")
		fmt.Println("college_name:", collegeName)
		specialitys, err := database.GetSpecialityByCollegeName(collegeName)
		if err != nil {
			c.JSON(501, gin.H{
				"code":    "ER00000001",
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{"code": 200, "msg": "SUCCESS", "data": specialitys})
		}
	}

}
