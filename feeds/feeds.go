package feeds

type FeedConfig struct {
	URL             string
	Header          string
	Agent           string // // "bot", "chrome", "reader"
	EnhancedHeaders bool   // When true, use enhanced headers for the request
}

var Feeds = []FeedConfig{
	// *
	// **
	// ***
	// tech
	{
		URL:             "https://www.reddit.com/r/javascript/.rss",
		Header:          "r/javascript",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.reddit.com/r/hardware/.rss",
		Header:          "r/hardware",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.reddit.com/r/debian/.rss",
		Header:          "r/debian",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://bits.debian.org/feeds/feed.rss",
		Header:          "bitsdebian",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://techmeme.com/feed.xml",
		Header:          "techmeme",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://rss.slashdot.org/Slashdot/slashdotMain",
		Header:          "slashdot",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://hnrss.org/frontpage",
		Header:          "hackernews",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://lobste.rs/rss",
		Header:          "lobsters",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	// *
	// **
	// ***
	// geopolitics
	{
		URL:             "https://rss.nytimes.com/services/xml/rss/nyt/World.xml",
		Header:          "nyt",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.reddit.com/r/worldnews/.rss",
		Header:          "r/worldnews",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.reddit.com/r/geopolitics/.rss",
		Header:          "r/geopolitics",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.reddit.com/r/anime_titties/.rss",
		Header:          "r/anime_titties",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.foreignaffairs.com/rss.xml",
		Header:          "foreignaffairs",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.cgtn.com/subscribe/rss/section/politics.xml",
		Header:          "cgtn",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://valdaiclub.com/export/rss/feed.xml",
		Header:          "valdai",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.inss.org.il/publication/feed/",
		Header:          "inss",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.jpost.com/rss/rssfeedsfrontpage.aspx",
		Header:          "jpost",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.israelhayom.com/feed/",
		Header:          "hayom",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.israelnationalnews.com/Rss.aspx?act=.1&cat=25",
		Header:          "arutzsheva",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	// *
	// **
	// ***
	// finance
	{
		URL:             "https://www.federalreserve.gov/feeds/press_monetary.xml",
		Header:          "fed",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.ecb.europa.eu/rss/press.html",
		Header:          "ecb",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.ft.com/rss/home",
		Header:          "ft",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://www.ft.com/markets?format=rss",
		Header:          "ft-markets",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://search.cnbc.com/rs/search/combinedcms/view.xml?partnerId=wrss01&id=100727362",
		Header:          "cnbc-world",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://search.cnbc.com/rs/search/combinedcms/view.xml?partnerId=wrss01&id=100003114",
		Header:          "cnbc-top",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://search.cnbc.com/rs/search/combinedcms/view.xml?partnerId=wrss01&id=10000664",
		Header:          "cnbc-finance",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://search.cnbc.com/rs/search/combinedcms/view.xml?partnerId=wrss01&id=15839135",
		Header:          "cnbc-earnings",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://search.cnbc.com/rs/search/combinedcms/view.xml?partnerId=wrss01&id=15838381",
		Header:          "cnbc-squawkstreet",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://search.cnbc.com/rs/search/combinedcms/view.xml?partnerId=wrss01&id=15838368",
		Header:          "cnbc-squawkbox",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	// {
	// 	URL:             "https://seekingalpha.com/market_currents.xml",
	// 	Header:          "seekingalphabreaking",
	// 	Agent:           "bot",
	// 	EnhancedHeaders: false,
	// },
	// {
	// 	URL:             "https://seekingalpha.com/feed.xml",
	// 	Header:          "seekingalphaarticles",
	// 	Agent:           "bot",
	// 	EnhancedHeaders: false,
	// },
}
