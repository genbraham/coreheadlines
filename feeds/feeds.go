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
		URL:             "https://www.reddit.com/r/C_Programming/.rss",
		Header:          "r/c_programming",
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
		URL:             "https://seekingalpha.com/market_currents.xml",
		Header:          "seekingalphabreaking",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
	{
		URL:             "https://seekingalpha.com/feed.xml",
		Header:          "seekingalphaarticles",
		Agent:           "bot",
		EnhancedHeaders: false,
	},
}
