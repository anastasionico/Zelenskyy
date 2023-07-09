package main

import (
	"fmt"
	"strings"
)

func main() {
	point := float64(0)
	var words = getWords()
	
	for ukr, eng := range words {
		var userInput string
		fmt.Printf("Guess the meaning of this word: %v \n", strings.ToUpper(ukr))
		fmt.Scan(&userInput)

		if strings.ToLower(userInput) == strings.ToLower(eng) {
			point++
			fmt.Println("Great Job!")
			fmt.Printf("Current point %v \n", point)
		} else {
			fmt.Printf("Error, the real meaning is : %v \n", strings.ToUpper(eng))
			fmt.Printf("Current point %v \n", point)
		}
	}
	getResult(point, words)
}

func getWords() map[string]string {
	var words = map[string]string{
		"я ya (pronounced: yah)":     "I",
		"ти ty (pronounced: tee)":    "you",
		"він vin (pronounced: veen)": "he",
		// "вона vona (pronounced: voh-nah)":"she",
		// "воно vono (pronounced: voh-no)":"it",
		// "ми my (pronounced: mih)":"we",
		// "ви vy (pronounced: vee)":"you",
		// "вони vony (pronounced: voh-nih)":"they",
		// "це tse (pronounced: tseh)":"this",
		// "той toy (pronounced: toy)":"that",
		// "тут tut (pronounced: toot)":"here",
		// "там tam (pronounced: tahm)":"there",
		// "хто khto (pronounced: kh-toh)":"who",
		// "що shcho (pronounced: sh-choh)":"what",
		// "де de (pronounced: deh)":"where",
		// "коли koly (pronounced: koh-ly)":"when",
		// "чому chomu (pronounced: cho-moo)":"why",
		// "як yak (pronounced: yak)":"how",
		// "бути buty (pronounced: boo-tih)":"be",
		// "мати maty (pronounced: mah-ty)":"have",
		// "робити robyty (pronounced: roh-bytih)":"do",
		// "говорити hovoryty (pronounced: hoh-vo-rytih)":"speak",
		// "йти yty (pronounced: ee-ty)":"go",
		// "бачити bachyty (pronounced: ba-chih-ty)":"see",
		// "знати znaty (pronounced: znah-ty)":"know",
		// "давати davaty (pronounced: dah-va-ty)":"give",
		// "брати braty (pronounced: brah-ty)":"take",
		// "робота robota (pronounced: ro-boh-tah)":"work",
		// "день den (pronounced: den)":"day",
		// "час chas (pronounced: chas)":"time",
		// "рік rik (pronounced: rik)":"year",
		// "добре dobre (pronounced: doh-bre)":"good",
		// "погано pohano (pronounced: poh-hah-no)":"bad",
		// "великий velykyy (pronounced: veh-lyk-yih)":"big",
		// "малий malyy (pronounced: mah-lyih)":"small",
		// "новий novyy (pronounced: noh-vyih)":"new",
		// "старий staryy (pronounced: stah-ryih)":"old",
		// "довгий dovhyy (pronounced: dohv-yih)":"long",
		// "короткий korotkyy (pronounced: koh-roht-kyih)":"short",
		// "сильний sylnyy (pronounced: sil-nyih)":"strong",
		// "слабкий slabkyy (pronounced: slahb-kyih)":"weak",
		// "гарний harnyy (pronounced: hahr-nyih)":"beautiful",
		// "поганий pohannyy (pronounced: poh-hahn-nyih)":"ugly",
		// "правильний pravylnyy (pronounced: prah-vyl-nyih)":"right",
		// "неправильний nepravylnyy (pronounced: neh-prah-vyl-nyih)":"wrong",
		// "близько blyzko (pronounced: blyz-koh)":"near",
		// "далеко daleko (pronounced: dah-leh-koh)":"far",
		// "на na (pronounced: nah)":"on, at",
		// "в v (pronounced: v)":"in",
		// "під pid (pronounced: pid)":"under",
		// "над nad (pronounced: nad)":"over",
		// "поруч poruch (pronounced: poh-rooch)":"beside",
		// "всередині vseredyni (pronounced: vseh-reh-dy-nee)":"inside",
		// "зовні zovni (pronounced: zohv-nee)":"outside",
		// "перед pered (pronounced: peh-rehd)":"before, in front of",
		// "після pislya (pronounced: pees-lya)":"after",
		// "між mizh (pronounced: meezh)":"between",
		// "з z (pronounced: z)":"with",
		// "без bez (pronounced: behz)":"without",
		// "для dlya (pronounced: dlyah)":"for",
		// "від vid (pronounced: veed)":"from",
		// "до do (pronounced: doh)":"to until",
		// "про pro (pronounced: proh)":"about",
		// "за za (pronounced: zah)":"behind",
		// "навколо navkolo (pronounced: nahv-koh-lo)":"around",
		// "через cherez (pronounced: cheh-rehz)":"through",
		// "наприклад napryklad (pronounced: nah-pryk-lahd)":"for example",
		// "тому tomu shcho (pronounced: toh-moo sh-choh)":"because",
		// "якщо yakshcho (pronounced: yak-sh-choh)":"if",
		// "колишній kolyshniy (pronounced: koh-lysh-nyih)":"former",
		// "потім potim (pronounced: poh-teem)":"then",
		// "завжди zavzhdy (pronounced: zahvzh-dy)":"always",
		// "іноді inodi (pronounced: ee-noh-dy)":"sometimes",
		// "часто chasto (pronounced: chah-stoh)":"often",
		// "майже mayzhe (pronounced: mai-zheh)":"almost",
		// "тільки tilky (pronounced: teel-kyih)":"only",
		// "ще shche (pronounced: sh-cheh)": "still yet",
		// "дуже duzhe (pronounced: doo-zheh)":"very",
		// "трохи trokhy (pronounced: troh-khy)":"a little",
		// "багато bahato (pronounced: bah-hah-toh)":"a lot",
		// "небагато nebahato (pronounced: neh-bah-hah-toh)":"not much",
		// "все vse (pronounced: vseh)":"everything",
		// "ніщо nishcho (pronounced: neesh-choh)":"nothing",
		// "хтось khtos (pronounced: kh-tos)":"someone",
		// "ніхто nikhhto (pronounced: neekh-toh)":"no one",
		// "всі vsi (pronounced: vsee)":"everyone",
		// "цей tsey (pronounced: tseh-yih)":"this",
		// "мій miy (pronounced: mee-yih)":"my",
		// "твій tviy (pronounced: tvee-yih)":"your",
		// "його yoho (pronounced: yoh-ho)":"his",
		// "її yiyi (pronounced: yee-yee)":"her",
		// "наш nash (pronounced: nash)":"our",
		// "ваш vash (pronounced: vash)":"your",
		// "їхній yikhniy (pronounced: yikh-nee-yih)":"their",
		// "усі usi (pronounced: oo-see)":"all",
		// "інший inshyy (pronounced: een-sh-yyih)":"other",
	}

	return words
}

func getResult(point float64, words map[string]string) {
	fmt.Println("\nGame over!")
	fmt.Printf("You made %v point on %v elements \n", point, len(words))

	var result = point / float64(len(words)) * 100
	fmt.Printf("You was right %v%% of the time \n", int(result))
}