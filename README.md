# Jane Technologies Inc. - Advertising Coding Challenge - Avdhut Joshi Solution

## Database Structure

The database that this project uses has 3 tables:
1. tbl_campaign
2. tbl_ad
3. tbl_keywords

The structure of each database is as follows:

tbl_campaign
     Column      |            Type             | Collation | Nullable |                      Default|
-----------------|-----------------------------|-----------|----------|---------------------------------------------------|
 campaign id     | integer - Primary Key       |           | not null | Auto Increment(Serial)|
 start campaign  | timestamp without time zone |           |          | |
 end c mpaign    | timestamp without time zone |           |          | |
 max impressions | integer                     |           |          | |
 cpm             | integer                     |           |          | |
 impressions     | integer                     |           |          | 0|

 
 
tbl_ad                                   
   Column    |         Type          | Collation | Nullable |                Default|
-------------|-----------------------|-----------|----------|---------------------------------------|
 ad id       | integer - Primary Key |           | not null | Auto Increment(Serial)|
 url_id      | character varying(45) |           |          | |
 campaign_id | integer               |           |          | |

tbl_keywords
   Column    |         Type          | Collation | Nullable | Default |
-------------+-----------------------+-----------+----------+---------|
 campaign_id | integer               |           | not null | |
 keyword     | character varying(45) |           | not null | |
 
Indexes: `"tbl_keywords_pkey" PRIMARY KEY, btree (campaign_id, keyword)`
    
    

## Endpoints:

### 1. POST /campaign
This creates a new campaign by adding an entry to tbl_campaign e.g:

 campaign id |   start campaign    |    end campaign     | max impressions | cpm | impressions |
-------------|---------------------|---------------------|-----------------|-----|-------------
2 | 2022-04-23 19:25:02 | 2022-04-23 21:36:39 |               1 |  20 |           0|

The column impressions is set to 0 by default

Along with multiple rows are added to tbl_keywords with where campaign_id is the campaign generated on the last api call and keyword is the word in the list passed in the "keyword" parameter in the request e.g.

 campaign_id | keyword |
-------------|---------|
2 | android |
2 | iOS |
          
          
Both these queries are wrapped in a transaction.

Endpoint returns `campaign_id` which is the last primary key added to `tbl_campaign


### 2. POST /addecision
A list of `keywords` is sent as input.

This endpoints checks for valid campaigns i.e. ones that are active and impressions are less than `max_impressions`and at least one of the input keywords apply to that keyword i.e a `campaign_id`, `keyword`pair exists in `tbl_keyword`. Campaigns are sorted by highest cpm, if there is a tie then earliest `end_timestamp` is selected is there is still a tie then the smallest `campaign_id` is selected. If a match is found then an add is created for that campaign by entering a row in `tbl_ad`. The naming convention that is followed is `A+add_id+C+campaign_id` e.g.:

 ad id | url id | campaign_id |
-------|--------|------------- |
10 | A10C43 | 43 |

The select campaign ID	and the URL for the ad is returned in the response.

`{"campaign_id":43, "impression_url":"http://localhost:8080/A10C43"}` 
`
### 3. GET [impression-url] 
This is used to call an add generated in  POST /addecision. When a URL with an existing add is called, it incrrements the `tbl_campaign.impressed` column. If the ad does not exist it returns a status 400.

e.g. 

`http://localhost:8080/A10C43` will return `{}` with status code 200

`http://localhost:8080/adcsed` will return `{}` 400

### GET /campaign/[campaign-id]
This returns the number of times an ad was impressed for that campaign i.e. the value in `tbl_campaign.impressed`. If the campaign does not exist it returns 400.

e.g. 

`http://localhost:8080/campaign/43` will return `'1'` 

`http://localhost:8080/campaign/adcsed` will return status code 400
