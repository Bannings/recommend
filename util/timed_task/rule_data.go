package timed_task

import (
	"database/sql"
	"recommend/golbal"
	"recommend/log"
	"recommend/model"
	"recommend/util/db"
)

var (
	tableIndex = map[string]string{
		golbal.INDEX_HIGHLY:   "highly",
		golbal.INDEX_NEW:      "new",
		golbal.INDEX_FREE:     "free",
		golbal.INDEX_VIP:      "vip",
		golbal.INDEX_PERSONAL: "personal",
	}
)

func queryRecommendCandidate(index string, userType model.UserType, errMsg string) *model.Candidates {
	var err error
	var maleTable, femaleTable string
	minLen := model.DisplayLenMap[model.GetSection(index).DisplayType]
	maleTable = "rule_nrt_homepage_" + tableIndex[index] + "_" + string(userType) + "_male"
	femaleTable = "rule_nrt_homepage_" + tableIndex[index] + "_" + string(userType) + "_female"

	var maleCandidates []string
	rows, err := db.RecallDB().Table(maleTable).Select("cid").Limit(golbal.MaxRecommendLen).Rows()
	if err != nil {
		log.Error(err)
	}
	for rows.Next() {
		var cid sql.NullString
		if err = rows.Scan(&cid); err == nil {
			maleCandidates = append(maleCandidates, cid.String)
		} else {
			log.Error(err)
			continue
		}
	}

	var femaleCandidates []string
	rows, err = db.RecallDB().Table(femaleTable).Select("cid").Limit(golbal.MaxRecommendLen).Rows()
	if err != nil {
		log.Error(err)
	}
	for rows.Next() {
		var cid sql.NullString
		if err = rows.Scan(&cid); err == nil {
			femaleCandidates = append(femaleCandidates, cid.String)
		} else {
			log.Error(err)
			continue
		}
	}
	if len(maleCandidates) < minLen || len(femaleCandidates) < minLen {
		log.Error(errMsg)
		return nil
	}
	tmpCandidate := model.Candidates{MaleCandidates: maleCandidates, FemaleCandidates: femaleCandidates}
	return &tmpCandidate
}

func getNewUserPersonalRecall(table string) []string {
	var candidates []string
	rows, err := db.RecallDB().Table(table).Select("cid").Limit(golbal.MaxRecommendLen).Rows()
	if err != nil {
		log.Error(err)
		return candidates
	}

	for rows.Next() {
		var cid sql.NullString
		if err = rows.Scan(&cid); err == nil {
			candidates = append(candidates, cid.String)
		} else {
			log.Error(err)
			continue
		}
	}
	return candidates
}

func getCandidates() {
	var maleCandidates []model.MaleFeedCandidates
	var femaleCandidates []model.FemaleFeedCandidates
	err := db.RecallDB().Table("algo_nrt_ts_longtail_recall").Select("male_recalls").Where("male_recalls IS NOT NULL").Find(&maleCandidates).Limit(300).Error
	if err != nil {
		log.Errorf("Get male recall fail,err:%v", err)
	}
	model.FeedSeed.Male = maleCandidates
	err = db.RecallDB().Table("algo_nrt_ts_longtail_recall").Select("female_recalls").Where("male_recalls IS NOT NULL").Find(&femaleCandidates).Limit(300).Error
	if err != nil {
		log.Errorf("Get female recall fail,err:%v", err)
	}
	model.FeedSeed.Female = femaleCandidates
	return
}

func reloadNewUserRuleData() {
	tmpHighlyCandidate := queryRecommendCandidate(golbal.INDEX_HIGHLY, model.NewUser, "新用户本周强推长度不足")
	if tmpHighlyCandidate != nil {
		model.RuleRecommendNewUser.HighlyCandidate = *tmpHighlyCandidate
	}

	// 人气新作推荐数据更新
	tmpNewCandidate := queryRecommendCandidate(golbal.INDEX_NEW, model.NewUser, "新用户人气新作长度不足")
	if tmpNewCandidate != nil {
		model.RuleRecommendNewUser.NewCandidate = *tmpNewCandidate
	}

	// VIP热销推荐数据更新
	tmpVipCandidate := queryRecommendCandidate(golbal.INDEX_VIP, model.NewUser, "新用户VIP热销长度不足")
	if tmpVipCandidate != nil {
		model.RuleRecommendNewUser.VipCandidate = *tmpVipCandidate
	}

	// 免费专区推荐数据更新
	tmpFreeCandidate := queryRecommendCandidate(golbal.INDEX_FREE, model.NewUser, "新用户免费专区长度不足")
	if tmpFreeCandidate != nil {
		model.RuleRecommendNewUser.FreeCandidate = *tmpFreeCandidate
	}

	// 猜你喜欢推荐数据更新
	tmpPersonalCandidate := queryRecommendCandidate(golbal.INDEX_PERSONAL, model.NewUser, "新用户猜你喜欢长度不足")
	if tmpPersonalCandidate != nil {
		model.RuleRecommendNewUser.PersonalCandidate = *tmpPersonalCandidate
	}

}

func reloadOldUserRuleData() {
	tmpHighlyCandidate := queryRecommendCandidate(golbal.INDEX_HIGHLY, model.OldUser, "老用户本周强推长度不足")
	if tmpHighlyCandidate != nil {
		model.RuleRecommendOldUser.HighlyCandidate = *tmpHighlyCandidate
	}

	// 人气新作推荐数据更新
	tmpNewCandidate := queryRecommendCandidate(golbal.INDEX_NEW, model.OldUser, "老用户人气新作长度不足")
	if tmpNewCandidate != nil {
		model.RuleRecommendOldUser.NewCandidate = *tmpNewCandidate
	}

	// VIP热销推荐数据更新
	tmpVipCandidate := queryRecommendCandidate(golbal.INDEX_VIP, model.OldUser, "老用户VIP热销长度不足")
	if tmpVipCandidate != nil {
		model.RuleRecommendOldUser.VipCandidate = *tmpVipCandidate
	}

	// 免费专区推荐数据更新
	tmpFreeCandidate := queryRecommendCandidate(golbal.INDEX_FREE, model.OldUser, "老用户免费专区长度不足")
	if tmpFreeCandidate != nil {
		model.RuleRecommendOldUser.FreeCandidate = *tmpFreeCandidate
	}

	// 猜你喜欢推荐数据更新
	tmpPersonalCandidate := queryRecommendCandidate(golbal.INDEX_PERSONAL, model.OldUser, "老用户猜你喜欢长度不足")
	if tmpPersonalCandidate != nil {
		model.RuleRecommendOldUser.PersonalCandidate = *tmpPersonalCandidate
	}

	getCandidates()
}

func ReloadRecommendRuleDataWxapp() {
	// 本周强推推荐数据更新

	tmpHighlyCandidate := queryRecommendCandidate(golbal.INDEX_HIGHLY, model.WxappUser, "微信小程序本周强推长度不足")
	if tmpHighlyCandidate != nil {
		model.RuleRecommendWxapp.HighlyCandidate = *tmpHighlyCandidate
	}
	// 人气新作推荐数据更新

	tmpNewCandidate := queryRecommendCandidate(golbal.INDEX_NEW, model.WxappUser, "微信小程序人气新作长度不足")
	if tmpNewCandidate != nil {
		model.RuleRecommendWxapp.NewCandidate = *tmpNewCandidate
	}
	// 猜你喜欢推荐数据更新

	tmpPersonalCandidate := queryRecommendCandidate(golbal.INDEX_PERSONAL, model.NewUser, "微信小程序猜你喜欢长度不足")
	if tmpPersonalCandidate != nil {
		model.RuleRecommendWxapp.PersonalCandidate = *tmpPersonalCandidate
	}
}
