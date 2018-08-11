package app

// env vars settings.yaml

type Specification struct {
	// Common App variables
	AppName         string `split_words:"true"  default:"LocalOrderNumberGenerator"`
	GitHash         string `split_words:"true"  default:"Github-Hash"`
	Branch          string `split_words:"true"  default:"Github-Branch"`
	BuildNumber     string `split_words:"true"  default:"Jenkins-BuildNumber"`
	Environment     string `split_words:"true"  default:"LOCAL"`
	NewRelicLicense string `split_words:"true"`
	DatacenterId	string `split_words:"true"`
	MongoHost       string `split_words:"true"  default:"localhost:27017"`
}

var Configuration *Specification
