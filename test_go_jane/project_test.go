package tese_go_jane

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"bytes"
	//"encoding/json"
)

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// )

func TestAddCampaign(t *testing.T) { //Testing for correct response for add campaign
	router := gin.Default()  

	var jsonStr = []byte(`{"start_capaign":1650790621,"end_campaign":1650813934,"max_impressions":3,"cpm":30,"keywords": ["5G", "4G"]}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/campaign", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t,[]byte`campaign_id:52` , w.Body)
}

func TestAddCampaignWrongBody(t *testing.T) { //Testing for correct response for add campaign
	router := gin.Default()  
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/campaign", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestAddDecision (t *testing.T) {
	router := gin.Default()  
	var jsonStr = []byte(`{"keywords":["5G"]}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/addecision", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t,[]byte`"campaign_id:51","impression_url":"http://localhost:8080/A32C52"` , w.Body)
}

func TestAddDecisionWrongBody (t *testing.T) {
	router := gin.Default()  
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/addecision", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestGetURL (t *testing.T) {
	router := gin.Default()  
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/A31C48", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t,[]byte`"{}"` , w.Body)
}

func TestGetURLNoURL (t *testing.T) {
	router := gin.Default()  
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/abcdefg", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t,[]byte`"{}"` , w.Body)
}

func TestGetCampaign (t *testing.T) {
	router := gin.Default()  
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/campaign/52", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t,"0" , w.Body.String())
}

func TestGetCampaignNoCampaign (t *testing.T) {
	router := gin.Default()  
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/campaign/52hhh", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t,nil , w.Body)
}