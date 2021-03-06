<img src="https://bit.ly/2GmKJqS" alt="urbn" width="100"/>

#  URBN Backend Engineering Summer 2018 Intern Project
##  Sterling Order Number Generator
###  v0
* Ava Patterson, <i>Scrum Master</i>
* Adam Holsinger, <i>Developer</i>
* Billy Groble, <i>Developer</i>
* Megan Paradowski, <i>Developer</i>
* Michelle Monteith, <i>Developer</i>
* Zoe Colbert, <i>QA Engineer</i>

#### A Note About Commits
This is a cleaned version of our project so we were able to publish it, therefore there's no commits to credit our individual work.
#### Personal Contributions
* Mongo client & Mock Mongo client
    - Mongo interfacing
    - MongoDocument object
    - Error object
* OrderNumber DAO & Mock OrderNumber DAO
* Unit Tests
    - Mongo
    - OrderNumber DAO
    - config


##  Quick Start
*  Clone the project in your $GOPATH/src/github.com/urbn/ directory
*  Start Goland and open the project
*  From the command line, run the command(s): `dep ensure`

*  Create a run configuration for the go file app/cmd/main.go and in that run configuration make sure you have the environment below.
*  You should now be all set you can run the app by click in the '>' button in the toolbar or debug by clicking the little 'bug' like button

## MongoDB set up
* Database name: `OrderNumberGen`
* Collection name: `ordernums`

* To initialize the database, manually insert the following documents:
```
{
    "prefix" : "AN",
    "brandId": "an",
    "dataCenterId": "US-NV",
    "orderNumber" : 0
}
{
    "prefix" : "AP",
    "brandId": "an",
    "dataCenterId": "US-PA",
    "orderNumber" : 0
}
{
    "prefix" : "TN",
    "brandId": "uo",
    "dataCenterId": "US-NV",
    "orderNumber" : 0
}
{
    "prefix" : "TP",
    "brandId": "uo",
    "dataCenterId": "US-PA",
    "orderNumber" : 0
}
{
    "prefix" : "FN",
    "brandId": "fp",
    "dataCenterId": "US-NV",
    "orderNumber" : 0
}
{
    "prefix" : "FP",
    "brandId": "fp",
    "dataCenterId": "US-PA",
    "orderNumber" : 0
}
```
