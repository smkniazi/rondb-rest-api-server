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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"hopsworks.ai/rondb-rest-api-server/version"
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
