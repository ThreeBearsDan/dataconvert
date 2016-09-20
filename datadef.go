package dataconvert

//store client UA info
type UAinfo struct {
	ip            string
	location      string
	osFamily      string
	broswerFamily string
}

//sotre client info eg: QQ name ..
type ClientCard struct {
	info string
}

//store client tag
type ClientTag struct {
	tagName  string
	tagColor string
}

//store client info
type Client struct {
	UAinfo
	ClientCard
	ClientTag
}

//store agent info
type Agent struct {
	accountNumber string
	realName      string
	staffId       int
}
