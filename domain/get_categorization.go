package domain

import (
	"strconv"
	"strings"
)

// https://sitereview.bluecoat.com/rest/categoryDetails?id=$NUM$
var list = map[int]string{
	1:   "Adult/Mature Content",
	3:   "Pornography",
	4:   "Sex Education",
	5:   "Intimate Apparel/Swimsuit",
	6:   "Nudity",
	7:   "Extreme",
	8:   "Alcohol/Tobacco",
	9:   "Scam/Questionable/Illegal",
	11:  "Gambling",
	13:  "Server Blacklist AFS",
	14:  "Violence/Hate/Racism",
	15:  "Weapons ",
	16:  "Abortion",
	17:  "Hacking",
	18:  "Phishing",
	20:  "Entertainment",
	21:  "Business/Economy",
	22:  "Alternative Spirituality/Belief",
	23:  "Alcohol",
	24:  "Tobacco",
	25:  "Controlled Substances",
	26:  "Child Pornography",
	27:  "Education",
	29:  "Charitable Organizations",
	30:  "Art/Culture",
	31:  "Financial Services",
	32:  "Brokerage/Trading",
	33:  "Games",
	34:  "Government/Legal",
	35:  "Military",
	36:  "Political/ Social Advocacy",
	37:  "Health",
	38:  "Technology/Internet",
	39:  "Hacking/Proxy Avoidance",
	40:  "Search Engines/Portals",
	43:  "Malicious Sources/Malnets",
	44:  "Malicious Outbound Data/Botnets",
	45:  "Job Search/Careers",
	46:  "News/Media",
	47:  "Personals/Dating",
	49:  "Refer ence",
	50:  "Mixed Content/Potentially Adult",
	51:  "Chat (IM)/SMS",
	52:  "Email",
	53:  "Newsgroups/Forums",
	54:  "Religion",
	55:  "Social Ne tworking",
	56:  "File Storage/Sharing",
	57:  "Remote Access Tools",
	58:  "Shopping",
	59:  "Auctions",
	60:  "Real Estate",
	61:  "Society/Daily Living",
	62:  "Sexuality/Alternative Lifestyles",
	63:  "Personal Sites",
	64:  "Restaurants/Dining/Food",
	65:  "Sports/Recreation",
	66:  "Tr avel",
	67:  "Vehicles",
	68:  "Humor/Jokes",
	69:  "Streaming Media/MP3/P2P",
	71:  "Software Downloads",
	72:  "Pay to Surf",
	83:  "Peer-to-Peer (P2P)",
	84:  "Audio/Video Clips",
	85:  "Office/Business Applications",
	86:  "Proxy Avoidance",
	87:  "For Kids",
	88:  "Web Ads/Analytics",
	89:  "Web Hosting",
	90:  "Uncategorized",
	91:  "Miscellaneous",
	92:  "Suspicious",
	93:  "Sexual Expression",
	94:  "LGBT",
	95:  "Translation",
	96:  "Non-Viewable/Infrastructure",
	97:  "Content Servers",
	98:  "Placeholders",
	99:  "Insufficient Content to Classify (ICC)",
	100: "Hobbies",
	101: "Spam",
	102: "Potentially Unwanted Software",
	103: "Dynamic DNS Host",
	104: "URL Redirector/Alias",
	106: "E-Card/Invitations",
	107: "Informational",
	108: "Computer/Information Security",
	109: "Internet Connected Devices",
	110: "Internet Telephony",
	111: "Online Meetings",
	112: "Media Sharing",
	113: "Radio/Audio Streams",
	114: "TV/Video Streams",
	118: "Piracy/Copyright Concerns",
	121: "Marijuana",
	123: "APT",
	126: "Potential Security Risk",
}

func getCategorization(categorization string) string {
	output := ""
	for i := 0; i < 150; i++ {
		if strings.Contains(categorization,
			"'catdesc.jsp?catnum="+strconv.Itoa(i)+"'") {
			output += list[i] + " and "
		}
	}

	return strings.TrimRight(output,
		" and ")
}