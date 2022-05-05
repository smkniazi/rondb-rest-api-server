/*
 * This file is part of the RonDB REST API Server
 * Copyright (c) 2022 Hopsworks AB
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package stat

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"hopsworks.ai/rdrs/version"

	"hopsworks.ai/rdrs/internal/common"
	ds "hopsworks.ai/rdrs/internal/datastructs"
	"hopsworks.ai/rdrs/internal/router/handler/pkread"
	tu "hopsworks.ai/rdrs/internal/router/handler/utils"
)

func TestPing(t *testing.T) {
	router := gin.Default()
	group := router.Group("/" + version.API_VERSION)
	group.GET(PATH, StatHandler)
	req, _ := http.NewRequest("GET", group.BasePath()+PATH, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Errorf("Test failed. Expected: %d, Got: %d", http.StatusOK, resp.Code)
	} else {
		t.Logf("Correct response received from the server")
	}
}

func TestStat(t *testing.T) {

	db := "DB004"
	table := "int_table"

	ch := make(chan int)

	numOps := 10

	tu.WithDBs(t, [][][]string{common.Database(db)},
		[]tu.RegisterTestHandler{pkread.RegisterPKTestHandler, RegisterStatTestHandler}, func(router *gin.Engine) {
			for i := 0; i < numOps; i++ {
				go performPkOp(t, router, db, table, ch)
			}
			for i := 0; i < numOps; i++ {
				<-ch
			}

			// get stats
			stats := getStats(t, router)
			if stats.NativeBufferStats.AllocationsCount != uint64(2*numOps) || stats.NativeBufferStats.BuffersCount != uint64(2*numOps) || stats.NativeBufferStats.FreeBuffers != uint64(2*numOps) {
				t.Fatalf("Native buffer stats do not match")
			}

			if stats.RonDBStats.NdbObjectsCreationCount != uint64(numOps) || stats.RonDBStats.NdbObjectsTotalCount != uint64(numOps) || stats.RonDBStats.NdbObjectsFreeCount != uint64(numOps) {
				t.Fatalf("RonDB stats do not match")
			}

		})
}

func performPkOp(t *testing.T, router *gin.Engine, db string, table string, ch chan int) {
	param := ds.PKReadBody{
		Filters:     tu.NewFiltersKVs(t, "id0", 0, "id1", 0),
		ReadColumns: tu.NewReadColumn(t, "col0"),
	}
	body, _ := json.MarshalIndent(param, "", "\t")

	url := tu.NewPKReadURL(db, table)
	tu.ProcessRequest(t, router, ds.PK_HTTP_VERB, url, string(body), http.StatusOK, "")

	ch <- 0
}

func getStats(t *testing.T, router *gin.Engine) ds.StatInfo {
	body := ""
	url := tu.NewStatURL()
	_, respBody := tu.ProcessRequest(t, router, ds.STAT_HTTP_VERB, url, string(body), http.StatusOK, "")

	var stats ds.StatInfo
	err := json.Unmarshal([]byte(respBody), &stats)
	if err != nil {
		t.Fatalf("%v", err)
	}
	return stats
}
