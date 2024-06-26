package syncs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hansbala/gyncer/database"
)

// given a list of string ids starts a new sync
func StartSyncsHandler(c *gin.Context) {
	var sync database.StartSync
	if err := c.BindJSON(&sync); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "malformed input"})
		return
	}

	db, err := database.ConnectToDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "database connection error"})
		return
	}

	syncDatas, err := database.GetSyncDatas(db, sync.SyncIds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = StartSyncWrapper(syncDatas); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("success: finished %d syncs", len(sync.SyncIds))})
}

// creates a new sync in the database
func CreateSyncHandler(c *gin.Context) {
	var newSync database.Sync
	if err := c.BindJSON(&newSync); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "malformed input"})
		return
	}

	if ok := newSync.IsValidSync(); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid sync data"})
		return
	}

	db, err := database.ConnectToDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "database connection error"})
		return
	}

	err = database.InsertNewSync(db, &newSync)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "database insertion error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success: sync created"})
}
