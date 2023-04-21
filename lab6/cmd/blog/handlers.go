package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Header             []headerData
	TopBlock 		   []topBlockData
	NavMenu            []navMenuData
	FeaturedPostsTitle string
	FeaturedPostsLine  string
	FeaturedPosts      []featuredPostsData
	MostRecentTitle    string
	MostRecentLine     string
	MostRecent         []mostRecentData
	Footer             []footerData
}

type headerData struct {
	Escape      string
	HeaderItems []headerItemsData
}

type headerItemsData struct {
	Home       string
	Categories string
	About      string
	Contact    string
}

type topBlockData struct {
	Background string
	Title      string
	Pretitle   string
	Button     string
}

type navMenuData struct {
	Nature      string
	Photography string
	Relaxation  string
	Vacation    string
	Travel      string
	Adventure   string
}

type featuredPostsData struct {
	ImageUrl       string `db:"image_url"`
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	AuthorUrl      string `db:"author_url"`
	Author         string `db:"author"`
	PublishDate    string `db:"publish_date"`
}

type mostRecentData struct {
	ImageUrl       string `db:"image_url"`
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	Line 	       string ``
	AuthorUrl      string `db:"author_url"`
	Author         string `db:"author"`
	PublishDate    string `db:"publish_date"`
}

type footerData struct {
	Background string
	Title      string
	Line       string
	Form       string
	BotBlock   []botBlockData
}

type botBlockData struct {
	Escape      string
	HeaderItems []headerItemsData
}

type postPage struct {
	Header    []postHeaderData
	MainBlock []mainBlockData
	Footer    []footerData
}

type postHeaderData struct {
	Escape      string
	HeaderItems []headerItemsData
}

type mainBlockData struct {
	Title     string
	Pretitle  string
	PostImage string
	Text      []textData
}

type textData struct {
	FirstPar  string
	SecondPar string
	ThirdPar  string
	FourthPar string
}

// func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		featuredposts, err := featuredPosts(db)
// 		if err != nil {
// 			http.Error(w, "Error", 500) // В случае ошибки парсинга - возвращаем 500
// 			log.Println(err)
// 			return // Не забыва	ем завершить выполнение ф-ии
// 		}

// 		mostrecent, err := mostRecent(db)
// 		if err != nil {
// 			http.Error(w, "Error", 500) // В случае ошибки парсинга - возвращаем 500
// 			log.Println(err)
// 			return // Не забываем завершить выполнение ф-ии
// 		}

// 		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
// 		if err != nil {
// 			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
// 			log.Println(err.Error())
// 			return // Не забываем завершить выполнение ф-ии
// 		}

// 		data := indexPage{
// 			Header:   header(),
// 			TopBlock: topBlock(),
// 			NavMenu:  navMenu(),
// 			FeaturedPostsTitle: "Featured Posts",
// 			FeaturedPostsLine: "../static/images/upper-line.png",
// 			FeaturedPosts: featuredposts,
// 			MostRecentTitle: "Most Recent",
// 			MostRecentLine: "../static/images/line.png",
// 			MostRecent: mostrecent,
// 			Footer:   footer(),
// 		}

// 		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
// 		if err != nil {
// 			http.Error(w, "Internal Server Error", 500)
// 			log.Println(err.Error())
// 			return
// 		}
// 	}
// }

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredposts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err)
			return
		}

		mostRecent, err := mostRecent(db)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/test.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := indexPage{
			Header:             header(),
			TopBlock:           topBlock(),
			NavMenu:            navMenu(),
			FeaturedPostsTitle: "Featured Posts",
 			FeaturedPostsLine: "../static/images/upper-line.png",
 			FeaturedPosts: featuredposts,
 			MostRecentTitle: "Most Recent",
 			MostRecentLine: "../static/images/line.png",
 			MostRecent: mostRecent,
			Footer:   footer(),
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func header() []headerData {
	return []headerData{
		{
			Escape:      "../static/images/escape.svg",
			HeaderItems: headerItems(),
		},
	}
}

func headerItems() []headerItemsData {
	return []headerItemsData{
		{
			Home:       "HOME",
			Categories: "CATEGORIES",
			About:      "ABOUT",
			Contact:    "CONTACT",
		},
	}
}

func topBlock() []topBlockData {
	return []topBlockData{
		{
			Background: "../static/images/header.png",
			Title:      "Let's do it together",
			Pretitle:   "We travel in the world in search of stories. Come along for the ride.",
			Button:     "View Latest Posts",
		},
	}
}

func navMenu() []navMenuData {
	return []navMenuData{
		{
			Nature:      "Nature",
			Photography: "Photography",
			Relaxation:  "Relaxation",
			Vacation:    "Vacation",
			Travel:      "Travel",
			Adventure:   "Adventure",
		},
	}
}


func featuredPosts(db *sqlx.DB) ([]featuredPostsData, error) {
	const query = `
		SELECT
		  title,
		  subtitle,
		  author,
		  author_url,
		  publish_date,
		  image_url
		FROM
		  post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts
	var featuredposts []featuredPostsData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&featuredposts, query) // Делаем запрос в базу данных
	if err != nil {     // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return featuredposts, nil
}

func mostRecent(db *sqlx.DB) ([]mostRecentData, error) {
	const query = `
		SELECT
		  title,
		  subtitle,
		  author,
		  author_url,
		  publish_date,
		  image_url
		FROM
		  post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции most-posts
	var mostrecent []mostRecentData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&mostrecent, query) // Делаем запрос в базу данных
	if err != nil {     // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return mostrecent, nil

}

func footer() []footerData {
	return []footerData{
		{
			Background: "../static/images/footer.png",
			Title:      "Stay in Touch",
			Line:       "../static/images/line.png",
			Form:       "Submit",
			BotBlock:   botBlock(),
		},
	}
}

func botBlock() []botBlockData {
	return []botBlockData{
		{
			Escape:      "../static/images/escape.svg",
			HeaderItems: headerItems(),
		},
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := postPage{
		Header:    postHeader(),
		MainBlock: mainBlock(),
		Footer:    footer(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func postHeader() []postHeaderData {
	return []postHeaderData{
		{
			Escape:      "../static/images/escape-black.svg",
			HeaderItems: headerItems(),
		},
	}
}

func mainBlock() []mainBlockData {
	return []mainBlockData{
		{
			Title:     "The Road Ahead",
			Pretitle:  "The road ahead might be paved - it might not be.",
			PostImage: "../static/images/the-road-ahead.png",
			Text:      text(),
		},
	}
}

func text() []textData {
	return []textData{
		{
			FirstPar:  "Dark spruce forest frowned on either side the frozen waterway. The trees had been stripped by a recent wind of their white covering of frost, and they seemed to lean towards each other, black and ominous, in the fading light. A vast silence reigned over the land. The land itself was a desolation, lifeless, without movement, so lone and cold that the spirit of it was not even that of sadness. There was a hint in it of laughter, but of a laughter more terrible than any sadness—a laughter that was mirthless as the smile of the sphinx, a laughter cold as the frost and partaking of the grimness of infallibility. It was the masterful and incommunicable wisdom of eternity laughing at the futility of life and the effort of life. It was the Wild, the savage, frozen-hearted Northland Wild.",
			SecondPar: "But there was life, abroad in the land and defiant. Down the frozen waterway toiled a string of wolfish dogs. Their bristly fur was rimed with frost. Their breath froze in the air as it left their mouths, spouting forth in spumes of vapour that settled upon the hair of their bodies and formed into crystals of frost. Leather harness was on the dogs, and leather traces attached them to a sled which dragged along behind. The sled was without runners. It was made of stout birch-bark, and its full surface rested on the snow. The front end of the sled was turned up, like a scroll, in order to force down and under the bore of soft snow that surged like a wave before it. On the sled, securely lashed, was a long and narrow oblong box. There were other things on the sled—blankets, an axe, and a coffee-pot and frying-pan; but prominent, occupying most of the space, was the long and narrow oblong box.",
			ThirdPar:  "In advance of the dogs, on wide snowshoes, toiled a man. At the rear of the sled toiled a second man. On the sled, in the box, lay a third man whose toil was over,—a man whom the Wild had conquered and beaten down until he would never move nor struggle again. It is not the way of the Wild to like movement. Life is an offence to it, for life is movement; and the Wild aims always to destroy movement. It freezes the water to prevent it running to the sea; it drives the sap out of the trees till they are frozen to their mighty hearts; and most ferociously and terribly of all does the Wild harry and crush into submission man—man who is the most restless of life, ever in revolt against the dictum that all movement must in the end come to the cessation of movement.",
			FourthPar: "But at front and rear, unawed and indomitable, toiled the two men who were not yet dead. bodies were covered with fur and soft-tanned leather. Eyelashes and cheeks and lips were so coated with the crystals from their frozen breath that their faces were not discernible. This gave them the seeming of ghostly masques, undertakers in a spectral world at the funeral of some ghost. But under it all they were men, penetrating the land of desolation and mockery and silence, puny adventurers bent on colossal adventure, pitting themselves against the might of a world as remote and alien and pulseless as the abysses of space.",
		},
	}
}
