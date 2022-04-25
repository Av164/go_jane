package handlers

import(
	"fmt"
	"context"
	"go_trial_2/db"
	"go_trial_2/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"strconv"
	"strings"
	
)

func AddCampaign(c *gin.Context){ //Adding new campaign - /campaign
	var newCampaign models.CampaignNew  
	var CampaignId int


	//Binding CampaignNew object to Json Request
	if err := c.BindJSON(&newCampaign); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		fmt.Println("ddsdsdssdgaaaaetrt")
		return
    }

	dbobj := db.SetupDB() //connecting to database

	fmt.Println("Adding campaign with  Start Time: " + strconv.Itoa(newCampaign.StartCampaign) + 
	" End Time: " + strconv.Itoa(newCampaign.EndCampaign) + " Max Impressions: " +strconv.Itoa(newCampaign.MaxImpressions)+ 
	" CPM: "+strconv.Itoa(newCampaign.Cpm))


	ctx := context.Background()
    tx, err1 := dbobj.BeginTx(ctx, nil)  //starting a transaction

	if(err1!=nil){
		fmt.Println("Error in setting up context: "+err1.Error()) //if setting up contet fails
		return
	}


	//Inserting new campaign within context of transaction
	insert_stmnt := "Insert into tbl_campaign (start_campaign, end_campaign, max_impressions, cpm) values(to_timestamp($1),to_timestamp($2),$3,$4) returning campaign_id"
	CampaignIdRes, err1 := tx.QueryContext(ctx,insert_stmnt, newCampaign.StartCampaign, newCampaign.EndCampaign, newCampaign.MaxImpressions, newCampaign.Cpm)//.Scan(&CampaignId)

	if err1 != nil{ //If query fails
		fmt.Println("Error in query: " +insert_stmnt)
		fmt.Println(err1)
		return
	}

	if(CampaignIdRes.Next()){ //If query returns value(It mostly will)
		err1 = CampaignIdRes.Scan(&CampaignId)
		if err1 != nil{ //If scanning fai.s
			fmt.Println("Error in scanning")
			fmt.Println(err1)
			return
		}
	}

	CampaignIdRes.Close() //Closing cursor

	if err1 != nil {
        // Incase we find any error in the query execution, rollback the transaction
        tx.Rollback()
        fmt.Println("\n", (err1), "\n ....Transaction rollback!\n")
        return
    }
	insert_stmnt1 := "Insert into tbl_keywords (campaign_id, keyword) values ($1, $2)"
	for i:=0; i<len(newCampaign.Keywords); i++{
		
		rows,err1 := tx.QueryContext(ctx,insert_stmnt1, CampaignId, newCampaign.Keywords[i])
		fmt.Printf("t1: %T\n", rows)
		if(err1 != nil){
			fmt.Println(newCampaign.Keywords[i])
			tx.Rollback()
			fmt.Println(err1)
			return
		}
		rows.Close()
	}


	
	err1 = tx.Commit()


	if err1 != nil {
        // Incase we find any error in the query execution, rollback the transaction
        tx.Rollback()
        fmt.Println("\n", (err1), "\n ....Transaction rollback!\n")
        return
    }

	dbobj.Close()

	m := map[string]int{"campaign_id":CampaignId}
	fmt.Println(m)

	c.IndentedJSON(http.StatusOK, m)
}


func AddDecision(c *gin.Context){
	var newAddecision models.Addecision

	var SearchArray []string

	if err := c.BindJSON(&newAddecision); err != nil {
        m := map[string]string{"error":"Invalid Json"}
		c.IndentedJSON(http.StatusBadRequest, m)
    }

	for i:=0 ; i<len(newAddecision.Keywords);i++{ //converting input to lowercase array
		var Lower string = strings.ToLower(newAddecision.Keywords[i])
		SearchArray = append(SearchArray, Lower)
	}

	dbobj  :=db.SetupDB()
	fmt.Println("Getting campaign id for keywords: ")
	fmt.Println(newAddecision.Keywords, SearchArray)

	//Sorting tbl_campaigns 
	rows,err:= dbobj.Query(`select campaign_id from tbl_campaign where campaign_id in (select campaign_id from tbl_keywords where lower(keyword) = any ($1)) and end_campaign>NOW() and impressions<max_impressions order by cpm desc,end_campaign asc, campaign_id asc limit 1`, pq.Array(SearchArray));

	if err!=nil{
		fmt.Printf("Error in query: select campaign_id from tbl_campaign where campaign_id in (select campaign_id from tbl_keywords where lower(keyword) = any ($1)) and end_campaign>NOW() and impressions<max_impressions order by cpm desc,end_campaign asc, campaign_id asc limit 1")
		fmt.Println(err)
		return
	}



	var m models.AddecisionRes
	if(rows.Next()){

		err = rows.Scan(&m.CampaignId)
		if err!=nil{
			fmt.Println("Error in Scanning")
			fmt.Println(err)
			return
		}	

		//Creating ad
		insert_stmnt := "insert into tbl_ad (url_id, campaign_id) values (concat('A',(select last_value from tbl_ad_ad_id_seq),'C',$1::text), $2) returning url_id";

		var Impression string
		err = dbobj.QueryRow(insert_stmnt, strconv.Itoa(m.CampaignId),m.CampaignId).Scan(&Impression)

		m.URL = "http://localhost:8080/"+Impression

		if(err!=nil){
			fmt.Println("Error in: "+insert_stmnt)
			fmt.Println(err)
			return
		}
		
		fmt.Println(m)
		c.IndentedJSON(http.StatusOK, m)	
	} else{

		m:=  map[string]int{}
		fmt.Println(m)
		c.IndentedJSON(http.StatusOK, m)	
	}
	dbobj.Close()
}


func GetURL(c *gin.Context){
	url := c.Param("url")
	dbobj := db.SetupDB()

	rows,err:= dbobj.Query("select url_id from tbl_ad where url_id=$1 and impressions<max_impressions", url) //Looking for valid campaign
	if(err!=nil){
		fmt.Println("Error in Query: "+"select url_id from tbl_ad where url_id=$1")
		fmt.Println(err)
		return
	}
	
	m:= map[string]int{}
	if(rows.Next()){
		var Impression string
		
		err=rows.Scan(&Impression)

		if(err!=nil){
			fmt.Println("Error in Scanning")
			fmt.Println(err)
			return
		}

		update_stmnt:="update tbl_campaign set impressions=impressions+1 where campaign_id=(select campaign_id from tbl_ad where url_id=$1) and impressions<max_impressions"

		_,err:= dbobj.Exec(update_stmnt, Impression)
		
		if(err!=nil){
			fmt.Println("Error in Query: "+update_stmnt)
			fmt.Println(err)
			return
		}

		c.IndentedJSON(http.StatusOK, m)
	} else{
		c.IndentedJSON(http.StatusBadRequest, m)
	}
	dbobj.Close()
}

func GetCampaign(c *gin.Context){
	id := c.Param("id")
	dbobj := db.SetupDB()

	get_impressions := "select impressions from tbl_campaign where campaign_id=($1)"

	row,err:= dbobj.Query(get_impressions, id)

	if(err!=nil){
		fmt.Println("Error in Query: "+get_impressions)
		fmt.Println(err)
		return
	}

	if(row.Next()){
		var Impressions int
		err=row.Scan(&Impressions)

		if(err!=nil){
			fmt.Println("Error in Scanning")
			fmt.Println(err)
			return
		}

		c.IndentedJSON(http.StatusOK, Impressions)
	} else{
		c.IndentedJSON(http.StatusBadRequest, nil)
	}
}

