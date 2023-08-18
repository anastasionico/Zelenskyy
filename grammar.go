package main

func getPersonalPronouns() map[string]string {
	return map[string]string{
		"я ya (pronounced: yah)":          "I",
		"ти ty (pronounced: tee)":         "you",
		"він vin (pronounced: veen)":      "he",
		"вона vona (pronounced: voh-nah)": "she",
		"воно vono (pronounced: voh-no)":  "it",
		"ми my (pronounced: mih)":         "we",
		"ви vy (pronounced: vee)":         "you",
		"вони vony (pronounced: voh-nih)": "they",
		"хтось khtos (pronounced: kh-tos)":"someone",
		"ніхто nikhhto (pronounced: neekh-toh)":"no one",
		"всі vsi (pronounced: vsee)":"everyone",
	}
}

func getDemonstrativePronouns() map[string]string {
	return map[string]string{
		"це tse (pronounced: tseh)":  "this",
		"той toy (pronounced: toy)":  "that",
		"тут tut (pronounced: toot)": "here",
		"там tam (pronounced: tahm)": "there",
	}
}

func getQuestionWords() map[string]string {
	return map[string]string{
		"хто khto (pronounced: kh-toh)":    "who",
		"що shcho (pronounced: sh-choh)":   "what",
		"де de (pronounced: deh)":          "where",
		"коли koly (pronounced: koh-ly)":   "when",
		"чому chomu (pronounced: cho-moo)": "why",
		"як yak (pronounced: yak)":         "how",
	}
}

func getPresentVerbs() map[string]string {
	return map[string]string{
		"Я буваю buvayu (pronounced: boo-vah-yoo)": "I am",
		"Я маю mayu (pronounced: mah-yoo)": "I have",
		"Я роблю roblu (pronounced: roh-bloo)": "I do",
		"Я говорю hovoryu (pronounced: hoh-voh-ryoo)": "I speak",
		"Я йду ydu (pronounced: eedoo)": "I go",
		"Я бачу bachu (pronounced: ba-choo)": "I see",
		"Я знаю znayu (pronounced: znah-yoo)": "I know",
		"Я даю dayu (pronounced: dah-yoo)": "I give",
		"Я беру beru (pronounced: beh-roo)": "I take",
		"Я працюю pratsyuyu (pronounced: pra-tsoo-yoo-yoo)": "I work",
		"Я думаю dumayu (pronounced: doo-mah-yoo)": "I think",
		"Я чую chuyu (pronounced: choo-yoo)": "I feel",
		"Я хочу khochu (pronounced: khoh-choo)": "I want",
		"Я розумію rozumiyu (pronounced: ro-zoo-mee-yoo)": "I understand",
		"Я граю hrayu (pronounced: hra-yoo)": "I play",
		"Я їм yim (pronounced: yim)": "I eat",
		"Я п'pyu (pronounced: pyoo)":  "I drink",
		"Я сплю splyu (pronounced: sploo)": "I sleep",
		"Я приходжу prykhodzhu (pronounced: pri-khoh-dzho)": "I come",
		"Я вчуся vchusya (pronounced: vchoo-sya)": "I learn",
		"Я чекаю chekayu (pronounced: cheh-kah-yoo)": "I wait",
		"Я рухаюся rukhayusya (pronounced: roo-khah-yoo-sya)": "I move",
		"Я допомагаю dopomahayu (pronounced: doh-po-ma-ha-yoo)": "I help",
		"Я розважаюся rozvazhayusya (pronounced: roz-va-zha-yoo-sya)": "I have fun",
	}
}

func getAdjectives() map[string]string {
	return map[string]string{
		"добре dobre (pronounced: doh-bre)":                        "good",
		"погано pohano (pronounced: poh-hah-no)":                   "bad",
		"великий velykyy (pronounced: veh-lyk-yih)":                "big",
		"малий malyy (pronounced: mah-lyih)":                       "small",
		"новий novyy (pronounced: noh-vyih)":                       "new",
		"старий staryy (pronounced: stah-ryih)":                    "old",
		"довгий dovhyy (pronounced: dohv-yih)":                     "long",
		"короткий korotkyy (pronounced: koh-roht-kyih)":            "short",
		"сильний sylnyy (pronounced: sil-nyih)":                    "strong",
		"слабкий slabkyy (pronounced: slahb-kyih)":                 "weak",
		"гарний harnyy (pronounced: hahr-nyih)":                    "beautiful",
		"поганий pohannyy (pronounced: poh-hahn-nyih)":             "ugly",
		"правильний pravylnyy (pronounced: prah-vyl-nyih)":         "right",
		"неправильний nepravylnyy (pronounced: neh-prah-vyl-nyih)": "wrong",
		"близько blyzko (pronounced: blyz-koh)":                    "near",
		"далеко daleko (pronounced: dah-leh-koh)":                  "far",
	}
}

func getPreposition() map[string]string {
	return map[string]string{
		"на na (pronounced: nah)":                           "on, at",
		"в v (pronounced: v)":                               "in",
		"під pid (pronounced: pid)":                         "under",
		"над nad (pronounced: nad)":                         "over",
		"поруч poruch (pronounced: poh-rooch)":              "beside",
		"всередині vseredyni (pronounced: vseh-reh-dy-nee)": "inside",
		"зовні zovni (pronounced: zohv-nee)":                "outside",
		"перед pered (pronounced: peh-rehd)":                "before, in front of",
		"після pislya (pronounced: pees-lya)":               "after",
		"між mizh (pronounced: meezh)":                      "between",
		"з z (pronounced: z)":                               "with",
		"без bez (pronounced: behz)":                        "without",
		"для dlya (pronounced: dlyah)":                      "for",
		"від vid (pronounced: veed)":                        "from",
		"до do (pronounced: doh)":                           "to until",
		"про pro (pronounced: proh)":                        "about",
		"за za (pronounced: zah)":                           "behind",
		"навколо navkolo (pronounced: nahv-koh-lo)":         "around",
		"через cherez (pronounced: cheh-rehz)":              "through",
	}
}

func getPossessiveDeterminers() map[string]string {
	return map[string]string{
		"мій miy (pronounced: mee-yih)":            "my",
		"твій tviy (pronounced: tvee-yih)":         "your",
		"його yoho (pronounced: yoh-ho)":           "his",
		"її yiyi (pronounced: yee-yee)":            "her",
		"наш nash (pronounced: nash)":              "our",
		"ваш vash (pronounced: vash)":              "your",
		"їхній yikhniy (pronounced: yikh-nee-yih)": "their",
	}
}

func getAdverb() map[string]string {
	return map[string]string{
		"потім potim (pronounced: poh-teem)":              "then",
		"завжди zavzhdy (pronounced: zahvzh-dy)":          "always",
		"іноді inodi (pronounced: ee-noh-dy)":             "sometimes",
		"часто chasto (pronounced: chah-stoh)":            "often",
		"майже mayzhe (pronounced: mai-zheh)":             "almost",
		"тільки tilky (pronounced: teel-kyih)":            "only",
		"ще shche (pronounced: sh-cheh)":                  "still yet",
		"дуже duzhe (pronounced: doo-zheh)":               "very",
		"трохи trokhy (pronounced: troh-khy)":             "a little",
		"багато bahato (pronounced: bah-hah-toh)":         "a lot",
		"небагато nebahato (pronounced: neh-bah-hah-toh)": "not much",
	}
}

func getColors() map[string]string {
	return map[string]string{
		"Червоний Chervonyy (pronounced: cher-vo-neeh)":               "Red",
		"Синій Syniy (pronounced: see-nee)":                           "Blue",
		"Зелений Zelenyy (pronounced: zeh-leh-neeh)":                  "Green",
		"Жовтий Zhovtyy (pronounced: zho-vtyeey)":                     "Yellow",
		"Рожевий Rozhovyy (pronounced: ro-zho-vyeey)":                 "Pink",
		"Помаранчевий Pomaranchevyy (pronounced: po-ma-rahn-chev-yy)": "Orange",
		"Фіолетовий Fioletovyy (pronounced: fyoh-leh-toh-vyy)":        "Purple",
		"Білий Bilyy (pronounced: bee-ly)":                            "White",
		"Чорний Chornyy (pronounced: chorn-yy)":                       "Black",
		"Сірий Siryy (pronounced: see-ryy)":                           "Gray",
		"Коричневий Korychnevyy (pronounced: ko-rych-nev-yy)":         "Brown",
		"Блакитний Blakytnyy (pronounced: bla-kyt-nyy)":               "Light Blue",
		"Золотий Zolotyy (pronounced: zo-lo-tyy)":                     "Gold",
		"Срібний Sribnyy (pronounced: sreeb-nyy)":                     "Silver",
	}
}

func getDates() map[string]string {
	return map[string]string{
		"Понеділок Ponedilok (pronounced: po-ne-dih-lok)":        "Monday",
		"Вівторок Vivtorok (pronounced: veev-toh-rok)":           "Tuesday",
		"Середа Sereda (pronounced: se-reh-da)":                  "Wednesday",
		"Четвер Chetver (pronounced: chet-ver)":                  "Thursday",
		"П'ятниця P'yatnytsya (pronounced: pyat-nit-sya)":        "Friday",
		"Субота Subota (pronounced: soo-boh-ta)":                 "Saturday",
		"Неділя Nedilya (pronounced: ne-di-lya)":                 "Sunday",
		"Січень Sichen' (pronounced: see-chen)":                  "January",
		"Лютий Lyutyi (pronounced: lyu-tiy)":                     "February",
		"Березень Berezen' (pronounced: beh-reh-zen)":            "March",
		"Квітень Kviten' (pronounced: kvi-ten)":                  "April",
		"Травень Traven' (pronounced: tra-ven)":                  "May",
		"Червень Cherven' (pronounced: cher-ven)":                "June",
		"Липень Lypen' (pronounced: ly-pen)":                     "July",
		"Серпень Serpen' (pronounced: ser-pen)":                  "August",
		"Вересень Veresen' (pronounced: ve-re-sen)":              "September",
		"Жовтень Zhovten' (pronounced: zho-vten)":                "October",
		"Листопад Lystopad (pronounced: lysto-pad)":              "November",
		"Грудень Hrudeny' (pronounced: hru-den)":                 "December",
		"Позавчора Pozavchora (pronounced: po-za-vcho-ra)":       "The day before yesterday",
		"Вчора Vchora (pronounced: vcho-ra)":                     "Yesterday",
		"Сьогодні S'ohodni (pronounced: s'oh-hod-ni)":            "Today",
		"Завтра Zavtra (pronounced: zav-tra)":                    "Tomorrow",
		"Післязавтра Pislyazavtra (pronounced: pis-lya-zav-tra)": "The day after tomorrow",
		"Весна Vesna (pronounced: ves-na)":                       "Spring",
		"Літо Lito (pronounced: lee-to)":                         "Summer",
		"Осінь Osin' (pronounced: oh-seen)":                      "Autumn",
		"Зима Zyma (pronounced: zy-ma)":                          "Winter",
	}
}

func getConjunctions() map[string]string {
	return map[string]string{
		"наприклад napryklad (pronounced: nah-pryk-lahd)":"for example",
		"тому tomu shcho (pronounced: toh-moo sh-choh)":"because",
		"якщо yakshcho (pronounced: yak-sh-choh)":"if",
		"колишній kolyshniy (pronounced: koh-lysh-nyih)":"former",
		"все vse (pronounced: vseh)":"everything",
		"ніщо nishcho (pronounced: neesh-choh)":"nothing",
		"цей tsey (pronounced: tseh-yih)":"this",
		"усі usi (pronounced: oo-see)":"all",
		"інший inshyy (pronounced: een-sh-yyih)":"other",
	}
}