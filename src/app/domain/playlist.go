package domain

type Playlist struct {
	Kind	string	`json:"kind"`
	Etag	string	`json:"etag"`
	Pageinfo
	Item
}

type Pageinfo struct{
	Totalresults	int	`json:"totalResults"`
	Resultsperpage	int	`json:"resultsPerPage"`
}

type Item struct{
	Kind	string	`json:"kind"`
	Etag	string	`json:"etag"`
	Id		string	`json:"id"`
	Snippet
}

type Snippet struct{
	Publishedat	string	`json:"publishedAt"`
	ChannelId	string	`json:"channelId"`
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	Thumbnails
}

type Thumbnails struct{
	Default
	Medium
	High
	Standard
	Maxres
	Channeltitle	string	`json:"channelTitle"`
	Defaultlanguage	string	`json:"defaultLanguage"`
	Localized
}

type Default struct{
	Url	string	`json:"url"`
	Width	int	`json:"width"`
	Height	int	`json:"height"`
}

type Medium struct{
	Url	string	`json:"url"`
	Width	int	`json:"width"`
	Height	int	`json:"height"`
}

type High struct{
	Url	string	`json:"url"`
	Width	int	`json:"width"`
	Height	int	`json:"height"`
}

type Standard struct{
	Url	string	`json:"url"`
	Width	int	`json:"width"`
	Height	int	`json:"height"`
}

type Maxres struct{
	Url	string	`json:"url"`
	Width	int	`json:"width"`
	Height	int	`json:"height"`
}

type Localized struct{
	Title	string	`json:"title"`
	Description	string	`json:"description"`
}
