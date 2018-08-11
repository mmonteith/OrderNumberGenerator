package app

const (
	MongoDatabase         = "OrderNumberGen"
	MongoOrdersCollection = "ordernums"
)

const (
	// Brand constants
	ANAbbrev = "an"
	UOAbbrev = "uo"
	FPAbbrev = "fp"

	// Data center constants
	PennsylvaniaDC = "US-PA"
	NevadaDC       = "US-NV"

	// Prefix constants
	ANPrefix   = "A"
	UOPrefix   = "T"
	FPPrefix   = "F"
	PADCPrefix = "P"
	NVDCPrefix = "N"
)

//sterling-order-number Response output by the /sterling-order-number endpoint
type SterlingOrderNumberResponse struct {
	DataCenterId        string `json:"dataCenterId"`
	Brand               string `json:"brand"`
	SterlingOrderNumber string `json:"sterlingOrderNumber"`
}

//healthResponse Response output by the /health endpoint
type HealthResponse struct {
	TimeStamp   string `json:"timestamp"`
	AppName     string `json:"appName"`
	Branch      string `json:"githubBranch"`
	BuildNumber string `json:"jenkinsBuildNumber"`
	GitHash     string `json:"githubHash"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
}

//Mapping for brand ID to orderNumber prefix
var BrandPrefix = map[string]string{
	ANAbbrev: ANPrefix,
	UOAbbrev: UOPrefix,
	FPAbbrev: FPPrefix,
}

//Mapping for site-id to brand
var BrandId = map[string]string{
	"an-us": ANAbbrev,
	"an-uk": ANAbbrev,
	"an-de": ANAbbrev,
	"an-fr": ANAbbrev,
	"uo-us": UOAbbrev,
	"uo-ca": UOAbbrev,
	"uo-uk": UOAbbrev,
	"uo-de": UOAbbrev,
	"uo-fr": UOAbbrev,
	"fp-us": FPAbbrev,
	"fp-uk": FPAbbrev,
	"fp-cn": FPAbbrev,
}

//Mapping for datacenter ID
var DataCenterId = map[string]string{
	"US-NV": NevadaDC,
	"US-PA": PennsylvaniaDC,
}

//Mapping for datacenter ID to prefix
var DataCenterPrefix = map[string]string{
	PennsylvaniaDC: PADCPrefix,
	NevadaDC:       NVDCPrefix,
}

//Mongo Document object
type MongoDocument struct {
	Prefix       string `bson:"prefix"`
	BrandId      string `bson:"brandId"`
	DataCenterId string `bson:"dataCenterId"`
	OrderNumber  int    `bson:"orderNumber"`
}

//Error object
type Error struct {
	Code    int
	Status  int
	Message string
}
