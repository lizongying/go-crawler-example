package bnu_spider

type RespFind struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Desc    string `json:"desc"`
	Data    struct {
		JstjID   interface{} `json:"jstjId"`
		ZifujiID string      `json:"zifujiId"`
		List     []struct {
			Data []struct {
				ID              int           `json:"id"`
				Hanzi           string        `json:"hanzi"`
				BuShou          string        `json:"buShou"`
				ZBiHua          string        `json:"zBiHua"`
				BwBiHua         string        `json:"bwBiHua"`
				BiShun          string        `json:"biShun"`
				ZhuYinZiMu      string        `json:"zhuYinZiMu"`
				Unicode         string        `json:"unicode"`
				Rhz             int           `json:"rhz"`
				Wbmz            int           `json:"wbmz"`
				Pinyin          string        `json:"pinyin"`
				ZifujiChangyong int           `json:"zifuji_changyong"`
				ZifujiTongyong  int           `json:"zifuji_tongyong"`
				ZifujiGuji      int           `json:"zifuji_guji"`
				ZifujiLidai     int           `json:"zifuji_lidai"`
				PinYins         []interface{} `json:"pinYins"`
				ZiFuJis         []interface{} `json:"ziFuJis"`
			} `json:"data"`
			Zbihua string `json:"zbihua"`
			Count  int    `json:"count"`
			Bihua  int    `json:"bihua"`
		} `json:"list"`
		Bishun string `json:"bishun"`
	} `json:"data"`
}

type RespSearch struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Desc    string `json:"desc"`
	Data    struct {
		Fantizi []struct {
			Hanzi   string `json:"hanzi"`
			Jhfs    string `json:"jhfs"`
			IsFanti string `json:"isFanti"`
		} `json:"fantizi"`
		Hanzi struct {
			ID               int      `json:"id"`
			Hanzi            string   `json:"hanzi"`
			BuShou           string   `json:"buShou"`
			ZBiHua           string   `json:"zBiHua"`
			BwBiHua          string   `json:"bwBiHua"`
			BiShun           string   `json:"biShun"`
			ZiJi             string   `json:"ziJi"`
			JieGouLeiXing    string   `json:"jieGouLeiXing"`
			JieGouTuShi      string   `json:"jieGouTuShi"`
			GouXingMoShi     string   `json:"gouXingMoShi"`
			IdsChaifen       string   `json:"ids_chaifen"`
			IdsJSON          string   `json:"idsJson"`
			IdsJSONFanti     []string `json:"idsJsonFanti"`
			DongTaiYanHua    string   `json:"dongTaiYanHua"`
			DongTaiYanHuaPic string   `json:"dongTaiYanHua_pic"`
			ZhuYinZiMu       string   `json:"zhuYinZiMu"`
			Unicode          string   `json:"unicode"`
			Sijiao           string   `json:"sijiao"`
			WuBiMa           string   `json:"wuBiMa"`
			ZhengMa          string   `json:"zhengMa"`
			Big5             string   `json:"big5"`
			YiTiZi           string   `json:"yiTiZi"`
			YtTree           struct {
				ID       string `json:"id"`
				Unicode  string `json:"unicode"`
				Hanzi    string `json:"hanzi"`
				Yiti     string `json:"yiti"`
				Children []struct {
					ID       string      `json:"id"`
					Unicode  string      `json:"unicode"`
					Hanzi    string      `json:"hanzi"`
					Yiti     interface{} `json:"yiti"`
					Children interface{} `json:"children"`
				} `json:"children"`
			} `json:"ytTree"`
			Rhz             int           `json:"rhz"`
			Wbmz            int           `json:"wbmz"`
			Pinyin          string        `json:"pinyin"`
			Pinyindiao      string        `json:"pinyindiao"`
			ZifujiChangyong int           `json:"zifuji_changyong"`
			ZifujiTongyong  int           `json:"zifuji_tongyong"`
			ZifujiGuji      int           `json:"zifuji_guji"`
			ZifujiLidai     int           `json:"zifuji_lidai"`
			PinYins         []interface{} `json:"pinYins"`
			ZiFuJis         []interface{} `json:"ziFuJis"`
		} `json:"hanzi"`
		Pinyins   []string    `json:"pinyins"`
		ShowWrite int         `json:"showWrite"`
		Zhengzi   interface{} `json:"zhengzi"`
		ZfjID     string      `json:"zfjId"`
		Zfj       string      `json:"zfj"`
	} `json:"data"`
}

type ExtraFind struct {
	Bishun string
}

type ExtraSearch struct {
	Word string
}

type DataWord struct {
	Id      int `bson:"_id" json:"id"`
	Fantizi []struct {
		Hanzi   string `json:"hanzi"`
		Jhfs    string `json:"jhfs"`
		IsFanti string `json:"isFanti"`
	} `json:"fantizi"`
	Hanzi struct {
		ID               int      `json:"id"`
		Hanzi            string   `json:"hanzi"`
		BuShou           string   `json:"buShou"`
		ZBiHua           string   `json:"zBiHua"`
		BwBiHua          string   `json:"bwBiHua"`
		BiShun           string   `json:"biShun"`
		ZiJi             string   `json:"ziJi"`
		JieGouLeiXing    string   `json:"jieGouLeiXing"`
		JieGouTuShi      string   `json:"jieGouTuShi"`
		GouXingMoShi     string   `json:"gouXingMoShi"`
		IdsChaifen       string   `json:"ids_chaifen"`
		IdsJSON          string   `json:"idsJson"`
		IdsJSONFanti     []string `json:"idsJsonFanti"`
		DongTaiYanHua    string   `json:"dongTaiYanHua"`
		DongTaiYanHuaPic string   `json:"dongTaiYanHua_pic"`
		ZhuYinZiMu       string   `json:"zhuYinZiMu"`
		Unicode          string   `json:"unicode"`
		Sijiao           string   `json:"sijiao"`
		WuBiMa           string   `json:"wuBiMa"`
		ZhengMa          string   `json:"zhengMa"`
		Big5             string   `json:"big5"`
		YiTiZi           string   `json:"yiTiZi"`
		YtTree           struct {
			ID       string `json:"id"`
			Unicode  string `json:"unicode"`
			Hanzi    string `json:"hanzi"`
			Yiti     string `json:"yiti"`
			Children []struct {
				ID       string      `json:"id"`
				Unicode  string      `json:"unicode"`
				Hanzi    string      `json:"hanzi"`
				Yiti     interface{} `json:"yiti"`
				Children interface{} `json:"children"`
			} `json:"children"`
		} `json:"ytTree"`
		Rhz             int           `json:"rhz"`
		Wbmz            int           `json:"wbmz"`
		Pinyin          string        `json:"pinyin"`
		Pinyindiao      string        `json:"pinyindiao"`
		ZifujiChangyong int           `json:"zifuji_changyong"`
		ZifujiTongyong  int           `json:"zifuji_tongyong"`
		ZifujiGuji      int           `json:"zifuji_guji"`
		ZifujiLidai     int           `json:"zifuji_lidai"`
		PinYins         []interface{} `json:"pinYins"`
		ZiFuJis         []interface{} `json:"ziFuJis"`
	} `json:"hanzi"`
	Pinyins   []string    `json:"pinyins"`
	ShowWrite int         `json:"showWrite"`
	Zhengzi   interface{} `json:"zhengzi"`
	ZfjID     string      `json:"zfjId"`
	Zfj       string      `json:"zfj"`
}
