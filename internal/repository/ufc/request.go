package ufc

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/ibra-bybuy/wsports-parser/pkg/constants"
	"github.com/ibra-bybuy/wsports-parser/pkg/model"
	"github.com/ibra-bybuy/wsports-parser/pkg/utils/datetime"
)

const eventsURL = "https://ufc.ru/events"

const cookieEng = "sa-user-id=s%253A0-50f9024e-9159-41a0-75b3-4578556d5134.1fjnijFDWy%252Bi8%252Fd3AKHYqGC2BFSNQgiolDzZX1DnSNY; sa-user-id-v2=s%253AUPkCTpFZQaB1s0V4VW1RNLwAqWU.GnV4HdftNJXvXRYAAfsFU4Gi9ynkFZFGXJU%252FE026WQw; _gid=GA1.2.8560177.1670679498; _gcl_au=1.1.719167758.1670679498; optimizelyEndUserId=oeu1670679498432r0.6225648374523445; _schn=_qn0xn3; _scid=dc72e294-8cd5-4106-9db5-7d682a6544e2; _tt_enable_cookie=1; _ttp=8853276e-987c-4465-8b67-0277c1cb676a; _sctr=1|1670619600000; OptanonAlertBoxClosed=2022-12-10T13:38:56.880Z; _ga_10DX6TQH49=GS1.1.1670679355.1.1.1670680209.0.0.0; _ga=GA1.1.1700728375.1670679498; OptanonConsent=isGpcEnabled=0&datestamp=Sat+Dec+10+2022+16%3A50%3A10+GMT%2B0300+(Moscow+Standard+Time)&version=202208.1.0&isIABGlobal=false&hosts=&consentId=84bd52e3-7823-4889-92ca-7a04cd109122&interactionCount=1&landingPath=NotLandingPage&groups=1%3A1%2C2%3A1%2C3%3A1%2C4%3A1&geolocation=RU%3BCE&AwaitingReconsent=false; STYXKEY_region=RUSSIA.RU.en.Default"
const cookieRu = "_gid=GA1.2.1096680601.1674668097; _gcl_au=1.1.1180467523.1674668098; _schn=_2xdhxv; _scid=9b50f141-55a8-458b-9d03-446b3cd3293b; _ym_uid=1674668099195977218; _ym_d=1674668099; _ym_isad=1; optimizelyEndUserId=oeu1674668099273r0.8093758870085648; _ym_visorc=w; sa-user-id=s%253A0-d88387f2-da1e-431c-57b6-350c0264f153.KFZXTQB9jiTXYUPWqoH0YsiE%252FYsjqBSnwASH3C7KwzU; sa-user-id-v2=s%253A2IOH8toeQxxXtjUMAmTxU7wAqTo.vQUzmGi%252B6ZOzV6m2U90yDur20lzrMjq7NpR1QfxMXQc; _tt_enable_cookie=1; _ttp=9V3En5-QdDRZMNdVZmVd9pxNtcU; _sctr=1|1674594000000; _ga_10DX6TQH49=GS1.1.1674668098.1.1.1674668105.0.0.0; _ga=GA1.2.420726224.1674668097"

func (r *repository) parseEvents(events *[]model.Event, lang *model.Lang) {

	r.c.OnRequest(func(r *colly.Request) {
		if lang.Code == model.LangRu.Code {
			r.Headers.Set("cookie", cookieRu)
		} else {
			r.Headers.Set("cookie", cookieEng)
		}
	})

	r.c.OnHTML(".block-region-ds-content", func(e *colly.HTMLElement) {
		bannerImage := e.ChildAttr(".c-hero__image img", "src")
		bannerFightName := e.ChildText(".c-hero--full__headline")

		e.ForEach(".l-listing__item", func(i int, e *colly.HTMLElement) {
			mainEvent := e.ChildText(".c-card-event--result__headline a")

			link := e.ChildAttr(".c-card-event--result__logo a", "href")
			eventName := strings.ToUpper("UFC FIGHT NIGHT " + mainEvent)
			if strings.Contains(link, "night") == false {
				linkSplit := strings.Split(link, "-")
				if len(linkSplit) == 2 {
					eventName = "UFC " + linkSplit[1]
				}
			}

			mainCardTimeStamp := e.ChildAttr(".c-card-event--result__date", "data-main-card-timestamp")
			prelimsTimeStamp := e.ChildAttr(".c-card-event--result__date", "data-prelims-card-timestamp")
			earlyCardTimeStamp := e.ChildAttr(".c-card-event--result__date", "data-early-card-timestamp")
			address := strings.ReplaceAll(e.ChildText(".address"), "\n", " ")

			startTime, err := datetime.StrUnixToTime(mainCardTimeStamp)

			if earlyCardTimeStamp != "" {
				startTime, err = datetime.StrUnixToTime(earlyCardTimeStamp)
			} else if prelimsTimeStamp != "" {
				startTime, err = datetime.StrUnixToTime(prelimsTimeStamp)
			}

			avatar := ""

			if mainEvent == bannerFightName {
				avatar = bannerImage
			}

			if err == nil {

				e.ForEach(".fight-card-tickets", func(i int, e *colly.HTMLElement) {
					fighters := strings.Split(e.Attr("data-fight-label"), "vs")
					if len(fighters) >= 2 {
						ID := fmt.Sprintf("ufc_%s_%d", mainCardTimeStamp, i)
						fighter1Name := strings.TrimSpace(fighters[0])
						fighter2Name := strings.TrimSpace(fighters[1])
						fighter1Ava := e.ChildAttr(".field--name-red-corner img", "src")
						fighter2Ava := e.ChildAttr(".field--name-blue-corner img", "src")

						event := model.Event{
							ID:   ID,
							Name: eventName,
							Teams: []model.Team{
								{
									ID:        fighter1Ava,
									Name:      fighter1Name,
									AvatarURL: fighter1Ava,
									Lang:      *lang,
								},
								{
									ID:        fighter2Ava,
									Name:      fighter2Name,
									AvatarURL: fighter2Ava,
									Lang:      *lang,
								},
							},
							StartAt:   datetime.Full(startTime),
							EndAt:     datetime.Full(startTime.Add(time.Hour * 9)),
							AvatarURL: avatar,
							Address:   address,
							Lang:      *lang,
							Sport:     constants.MMA_TYPE,
						}

						*events = append(*events, event)
					}
				})
			}
		})
	})

	r.c.Visit(eventsURL)
}
