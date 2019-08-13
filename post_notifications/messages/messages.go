package messages

const (
	TwitterMsg = `<b>New Tweet from </b><a href="https://www.twitter.com/sentinel_co">@sentinel_co</a>

%s

link: https://twitter.com/%s/status/%s`
	RetweetMsg = "<b>New Retweet from %s</b>\n\n%s\nhttps://twitter.com/%s/status/%s"
	MediumPost = "<b>New Medium Publication from %s</b>\n\n<b>%s</b>\n\nlink:%s"
	RedditPost = "*New Reddit post from %s*\n\n*%s*\n%s\n\nlink:[%s]"
)
