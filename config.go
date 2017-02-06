package mdrest

type Config struct {
	Watch    bool
	BasePath string        //the base path of you project, default is "", you can use "/" or "/docs/"
	SrcDir   string
	OutputType string      //you can put json or  html is default
	SiteMapDeep int
	ServerBaseDir  string  //if the markdown image in assets/xxx.png,
	DistDir  string
	NoLogging  bool
	NoIndex    bool
	NoSiteMap  bool
}
