package Routes

import (
	"Project2/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	p := router.Group("/person")
	{
		p.POST("/create", Controllers.CreatePerson)
		p.GET("/:person_id/info", Controllers.GetPersonInfo)
	}

	return router

}

// func getData(c *gin.Context) {
// 	fmt.Println("IN: getData")

// 	dbconnMySQL, dberr := db.GetSQLConnection()
// 	if dberr != nil {
// 		fmt.Println("ERROR in DB ", dberr)
// 		// logginghelper.LogError("ERROR in DB ", dberr)
// 		return false, dberr
// 	}

// 	session := dbconnMySQL.NewSession(nil)

// 	type Person struct {
// 		name int
// 		age  int
// 	}
// 	var person Person

// 	noOfRecords, err := session.SelectBySql(
// 		`SELECT g.id, g.applicationId, g.workflowID, g.currentSequenceNo, g.patchId, g.district, g.districtIdOld, g.tehsilId, g.tehsilIdOld, g.fileId, g.origanal_filename, g.area, g.remark, ST_AsGeoJSON(g.gisdata), g.intersection_count, g.intersection_ids, g.applicantName, g.areaType, g.claimType, g.isPolygonUploadedByIndividual, g.isPolygonAvailable, g.polygonAccuracy, g.gisMetaData, g.reasonText, g.optionalRemark, g.queryRaiseId, g.measuredOn, g.appVersion, g.createdOn, g.createdBy, g.isLandCombinedConfirmedOnGroundByFRC, g.combinedWithLandPatchId, g.status, g.isPolygonAreaMeasurable, g.createdByIP, g.longitude, g.latitude, g.applicantFatherName, g.casteCategory, g.dwellerID, g.gisDataManual, g.gisMetaDataManual, g.polygonSurveyedManually, b.batchProcessId
// 		FROM BatchProcessStatus b INNER JOIN gisrecords g
// 		ON b.applicationId = g.applicationId
// 			AND b.patchId = g.patchId
// 			AND b.areaType = g.areaType
// 			AND b.isPolygonUploadedByIndividual = g.isPolygonUploadedByIndividual
// 		WHERE b.batchProcessId >= ` + strconv.Itoa(start) +
// 			` AND b.batchProcessId <= ` + strconv.Itoa(end) +
// 			` AND b.isProcessed = 0;
// 		`).Load(&person)

// 	if err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}
// 	fmt.Println("noOfRecords: ", noOfRecords)
// 	fmt.Println("person: ", person)

// 	c.JSON(http.StatusOK, "Data is present")
// }
