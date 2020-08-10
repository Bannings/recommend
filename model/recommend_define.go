package model

var (
	RuleRecommendNewUser = new(RecommendRuleData) //固定栏目候选集
	RuleRecommendOldUser = new(RecommendRuleData) //固定栏目候选集
	RuleRecommendWxapp   = new(RecommendRuleData) //固定栏目候选集
	NewID                = new(NewUserID)         //当天新用户对应ID
	RecommendMaps        = new(RecommendMap)
	FeedSeed             = new(FeedRecommendCandidates)
)

type UserType string

const (
	NewUser   UserType = "newer"
	OldUser   UserType = "older"
	WxappUser UserType = "wxapp"
)

func GetStaticRuleData(userType UserType) *RecommendRuleData {
	switch userType {
	case NewUser:
		return RuleRecommendNewUser
	case OldUser:
		return RuleRecommendOldUser
	case WxappUser:
		return RuleRecommendWxapp
	default:
		return RuleRecommendNewUser
	}
}

type Candidates struct {
	MaleCandidates   []string
	FemaleCandidates []string
}

type Candidate struct {
	ComicId string `gorm:"column:cid"`
}

type MaleFemaleTypeCandidate struct {
	MaleCandidateMap   map[string]string
	FemaleCandidateMap map[string]string
}

type TypeCandidates struct {
	MaleTypeCandidates   []TypeCandidate
	FemaleTypeCandidates []TypeCandidate
}

type TypeCandidate struct {
	ComicType string `gorm:"column:comic_type"`
	ComicList string `gorm:"column:comic_list"`
}

type FeedRecommendCandidates struct {
	Male   []MaleFeedCandidates
	Female []FemaleFeedCandidates
}

type MaleFeedCandidates struct {
	ComicId string `gorm:"column:male_recalls"`
}

type FemaleFeedCandidates struct {
	ComicId string `gorm:"column:female_recalls"`
}

func (feed *FeedRecommendCandidates) GetMaleCandidates() []string {
	var candidates []string
	for _, candidate := range feed.Male {
		candidates = append(candidates, candidate.ComicId)
	}
	return candidates
}

func (feed *FeedRecommendCandidates) GetFemaleCandidates() []string {
	var candidates []string
	for _, candidate := range feed.Female {
		candidates = append(candidates, candidate.ComicId)
	}
	return candidates
}

type RecommendRuleData struct {
	HighlyCandidate   Candidates
	NewCandidate      Candidates
	VipCandidate      Candidates
	FreeCandidate     Candidates
	PersonalCandidate Candidates
	ABCandidate       Candidates
}

func (t *TypeCandidates) GetMaleCandidatesMap() map[string]string {
	maleMap := make(map[string]string)
	for _, typeCandidate := range t.MaleTypeCandidates {
		maleMap[typeCandidate.ComicType] = typeCandidate.ComicList
	}
	return maleMap
}

func (t *TypeCandidates) GetFemaleCandidatesMap() map[string]string {
	femaleMap := make(map[string]string)
	for _, typeCandidate := range t.FemaleTypeCandidates {
		femaleMap[typeCandidate.ComicType] = typeCandidate.ComicList
	}
	return femaleMap
}

type HotRecommend struct {
	Male   []string
	Female []string
}

type RecommendMap struct {
	CartoonTagMap  map[int]TagInfo
	DescriptionMap map[string]string
	ComicFreeMap   map[string]int
}
