package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

type CPIPPI struct {
	Qgcpi []float64
	Hncpi []float64
	Qgppi []float64
	Hnppi []float64
	Date  []string
}

type Topic struct {
	Title   string
	Lead    string
	Content string
}

type Article struct {
	Title   string
	Date    string
	Author  string
	Clink   int
	Lead    string
	Content string
}

type Navbar struct {
	Active bool
	Href   string
	Name   string
}

func Getpagination(pagination string) int64 {
	i := beego.AppConfig.String(pagination)
	ipi, _ := strconv.ParseInt(i, 10, 64)
	return ipi
}

func (navbar *Navbar) Getnavbar(active string) (Navbar, []Navbar, []Navbar) {
	navob1 := Navbar{false, "/", "资 讯"}
	navob2 := Navbar{false, "/", "政 策"}
	navob3 := Navbar{false, "/", "分 析"}
	navob4 := Navbar{false, "/", "旅 游"}
	navob5 := Navbar{false, "/", "美 食"}
	navob6 := Navbar{false, "/", "教 育"}
	navob7 := Navbar{false, "http://www.hnyyjg.com", "医 药"}
	navobarray := []Navbar{navob1, navob2, navob3, navob4, navob5, navob6, navob7}
	for i := range navobarray {
		if navobarray[i].Name == active {
			navobarray[i].Active = true
		}
	}
	navfirst := navobarray[0]
	navfive := navobarray[1:5]
	navmore := navobarray[5:len(navobarray)]
	return navfirst, navfive, navmore
}

func (c *MainController) Get() {
	c.Data["Title"] = "首页"
	c.Data["Intro"] = "湖南省最权威的价格信息发布平台！"

	var navbar Navbar
	nav1, nav2, nav3 := navbar.Getnavbar("资 讯")
	c.Data["Nav1"] = nav1
	c.Data["Nav2"] = nav2
	c.Data["Nav3"] = nav3

	var cp CPIPPI
	cps := cp.GetCPI_PPI()
	c.Data["Dates"] = cps.Date
	c.Data["Qgcpi"] = cps.Qgcpi
	c.Data["Hncpi"] = cps.Hncpi
	c.Data["Qgppi"] = cps.Qgppi
	c.Data["Hnppi"] = cps.Hnppi

	ina := Getpagination("indexpagination")
	var article Article
	articles := article.GetArticles(ina)
	c.Data["Articles"] = articles
	var topic Topic
	itopic := topic.GetTopic()
	c.Data["Topic"] = itopic
	c.TplNames = "home.html"
}

func (c *MainController) Topic() {
	c.Data["Title"] = "专题"
	var navbar Navbar
	nav1, nav2, nav3 := navbar.Getnavbar("")
	c.Data["Nav1"] = nav1
	c.Data["Nav2"] = nav2
	c.Data["Nav3"] = nav3
	c.TplNames = "topic.html"
}

func (cps *CPIPPI) GetCPI_PPI() *CPIPPI {
	cps.Qgcpi = []float64{3.2, 3.0, 2.5, 2.0, 2.5, 2.4, 1.8, 2.5, 2.2, 2.3, 2.0, 1.6}
	cps.Hncpi = []float64{3.7, 3.6, 3.3, 2.9, 3.4, 2.5, 1.6, 1.8, 2.0, 2.0, 1.6, 1.2}
	cps.Qgppi = []float64{-1.5, -1.4, -1.4, -2.0, -1.6, -2.3, -2.0, -1.4, -2.1, -0.9, -1.2, -1.8}
	cps.Hnppi = []float64{-1.2, -1.2, -1.0, -1.7, -1.5, -2.2, -2.0, -1.2, -0.9, -0.6, -0.9, -1.4}
	cps.Date = []string{"2013-10", "2013-11", "2013-12", "2014-01", "2014-02", "2014-03", "2014-04", "2014-05", "2014-06", "2014-07", "2014-08", "2014-09"}
	return cps
}

func (topic *Topic) GetTopic() *Topic {
	topic.Title = ""
	topic.Lead = "专题文章测试一下！嘎嘎！！"
	topic.Content = ""
	return topic
}

func (article *Article) GetArticles(ipi int64) []*Article {
	var articles []*Article
	articles = make([]*Article, ipi)
	var i int64
	for i = 0; i < ipi; i++ {
		article.Title = "王石川：养老保险破除“双轨制”之后"
		article.Date = "2014-12-31"
		article.Author = "admin"
		article.Clink = 20
		article.Lead = "机关事业单位与城镇职工的养老金并轨了，其他群体要不要并轨？中国社科院社会政策研究中心秘书长、研究员唐钧表示，我国现有养老金实行七轨制，“国家公务员，事业单位人员，城市居民，农民，企业职工，军人，还有一块农民工，农民工虽然讲可以包含在城市职工这个范畴里，但实质上是不同的。"
		article.Content = "【财经网记者 陈君】这两天互联网上出大事了，淘宝小二叫板国家工商总局，一个企业和国家主管部委直接厮杀，应该是史上第一次。 事情相关经过大致如下： 1月23日，工商总局抽检购物网站：淘宝网正品率最低仅为37.25%。 1月27日，80后淘宝小二致信工商总局市场司司长：抽检程序存疑。 1月28日上午，工商总局发布白皮书：阿里系网络交易平台存5大问题。 1月28日下午，淘宝回应工商总局：正式投诉网监司司长刘红亮。 1月28日晚，白皮书被工商总局悄然撤下。 1月29日上午，工商总局回应淘宝投诉：暂未收到投诉流程。淘宝回应，正在准备材料。 财经网记者就此事采访了社科院财经战略研究院的学者王雪峰，他对实体商品交易市场和网络商品交易市场的发展情况以及相关的法律法规情况都比较了解。 王雪峰 王雪峰 工商总局意在强化网络市场监管，促进网络平台依规履责 记者：您如何理解工商总局这次的行动？他们是基于什么样的出发点？ 王雪峰：从大背景来说，市场在我国经济中的地位不断提升，已经由补充性、基础性地位提升为决定性地位。伴随市场地位的提升，我国市场结构也日益复杂，表现形式也日益多样化和高科技化。特别是近10多年来，网络交易市场迅猛发展，已经成为我国市场体系的重要组成部分。依据市场形成和演进的动力基础和国内外规范市场发展、维护市场秩序的实践经验，我国网络市场到了强化监管和规范市场秩序的发展阶段。 从大态势和要求来看，我国经济正在步入“新常态”，经济增速在减速换挡，经济结构在加速调整，增长动力正在由投资、出口向创新和消费转变。“新常态”下的发展模式要求充分发挥市场的决定作用和政府的调控职能。市场决定作用的有效发挥就需要规范、诚信，安全、高效的市场环境和社会环境；也需要政府市场监管部门强力的监管执法队伍，创新监管手段、维护市场秩序。 国家工商总局作为我国法定的市场监管部门，监管市场主体、规范市场秩序、维护市场环境是其法定的职责。在市场监管体制改革方面，近年来，工商系统改革的力度较大，步子较快，无论是企业注册制改革还是企业注册资本认缴制改革，都是在放宽准入、降低门槛的举措，都是释放由准入监管向后续行为监管的政策信号。准入放宽必然阶段性引起鱼龙混杂的局面， 在此背景下，工商总局如想更好履行市场监管职能，规范市场，维护市场秩序，履行2008年国务院“三定”方案的职责（编者注：“三定”工作的指导思想是：贯彻落实深化行政管理体制改革的意见和国务院机构改革方案，着力转变职能、理顺关系、优化结构、提高效能，以构建权责一致、分工合理、决策科学、执行顺畅、监督有力的行政管理体制。），就需要以市场主体的诚信为重要抓手，理清政府和市场平台企业的职责权限，促使企业平台履行相应的职责和义务。鉴于此，我们认为，这次工商总局“定向监测”和向社会公示监测结果行为应该是对网络市场监管的一次创新和尝试，并可能会成为以后市场监管的常态举措。 在平台企业层面，不管是实体市场的举办主体企业还是网络市场的平台企业（比如淘宝），他们掌握着大量的经营主体数据，对入场或上线商户情况的了解比政府更清楚，对平台内经营主体的行为具有较强的掌控能力；因而，平台企业具有配合政府监管部门监管场内经营商户的职责、义务和能力。这一次淘宝回应说希望工商总局入场监管，显然这是推责行为，该管的要管起来，该负的责任负起来，才是真正负责任的企业行为，推责给政府有失担当。再者，驻场监管是我国市场发展不成熟的产物，目前，在工商监管体制上已经取消了这种“替企业履责”的监管方式，职责边界清晰、依法监管、依法履责是市场经济条件下市场监管的必然要求。市场举办方和平台企业以后要逐步强化责任意识和担当能力，不要再延续过去把责任推给政府，依赖政府管理的模式，这才是企业长大成熟的标志，才符合市场经济的要求。 总之，无论你是一个经营户还是一个平台举办方或者管理方，都应该履责守法，才有资格享有市场经济和法律赋予你的权力，都应该跟政府监管部门积极配合，对政府监管部门提出的问题，相关方应该积极正面回应，与政府共同营造和谐有序的市场环境。这次淘宝的反映有些让人失望，显得有些任性自大，有失网络市场老大的风范；与国家监管市场、强化网络市场监管、改善网络市场经营和消费环境的导向不符。这叫小聪明，不是大智慧，应引以为鉴。 淘宝卖了假货应该认错改正 对抗表态有失风范 记者：您认为官方思维在改变，企业没有跟上，是因为能力问题、思维问题还是双方交流出了问题？ 王雪峰：个人判断是思维惯性、判断失误和企业不成熟。"
		articles[i] = article
	}
	return articles
}

func (c *MainController) GetAjaxArticles() {
	ipi := Getpagination("indexpagination")
	var article Article
	articles := article.GetArticles(ipi)
	c.Data["Articles"] = articles
	id := c.GetString("Id")
	beego.Info(id)
	c.TplNames = "ajax.tpl"
}
