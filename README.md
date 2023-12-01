# Toronto Time Database

## Docker Image
```
docker run -p 7575:7575 jasbirnetwork/torontotimedb:v2
```

## Access Toronto Time App

Open [http://localhost:7575/current-time](http://localhost:7575/current-time) with your browser to see the result.

## You should get reponse like below:

```
{
"current_time": "2023-12-01T00:05:59-05:00"
}
```

## Get data from database
Open [http://localhost:7575/all-times](http://localhost:7575/all-times) with your browser to see the result.

## You should get reponse like below:

```
[
    {
        "id": 1,
        "timestamp": "2023-12-01T04:17:34.862739262Z"
    },
    {
        "id": 2,
        "timestamp": "2023-12-01T04:20:50.733272626Z"
    },
    {
        "id": 3,
        "timestamp": "2023-12-01T04:20:51.363558233Z"
    },
    {
        "id": 4,
        "timestamp": "2023-12-01T04:20:51.764223001Z"
    },
    {
        "id": 5,
        "timestamp": "2023-12-01T04:26:34.851160864Z"
    }
]
```

## Explaination

##### We have use sqllite database
```
sqlite3 toronto_time.db
```
```
func init() {
	// Open database connection using gorm with SQLite
	var err error
	db, err = gorm.Open(sqlite.Open("file:toronto_time.db?cache=shared&_loc=auto"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Auto Migrate the time_log table
	db.AutoMigrate(&TimeLog{})
}
```
* The first we define create the sqllite file with the name `toronto_time.db`.
* Inside `init()` we connect to our database and run migration as per defined struct.

### Insert data
```
result := db.Create(&timeLog)
if result.Error != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log timestamp"})
	return
}
```
* Insert current time into the database using gorm `db.Create` method.

## Get Inserted response
```
func getAllTimes(c *gin.Context) {
	var timeLogs []TimeLog
	// Retrieve all entries from the time_log table
	result := db.Find(&timeLogs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve time logs"})
		return
	}
	// Respond with the retrieved time logs in JSON format
	c.JSON(http.StatusOK, timeLogs)
}
```
* API endpoint to get all times from the time_log table

### For test the app run below command.

```
go test
```
