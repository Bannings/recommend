package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"recommend/handler"
	"recommend/router/middleware"
)

func InitRouter() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middleware.ServerHeader)
	g.Use(middleware.Log())
	g.Use(middleware.UploadUserInfo())
	g.Use(gzip.Gzip(gzip.BestCompression))
	api := g.Group("top")
	{
		// 最初一版的接口，访问量已经很少
		api.GET("/homepage", handler.GetHomepageTopRecommend)
		api.GET("/highly", handler.GetFromHighlyRecommend)
		api.GET("/quality", handler.GetFromQualityRecommend)
		api.GET("/unique", handler.GetFromUniqueRecommend)
		api.GET("/new", handler.GetFromNewRecommend)
		api.GET("/vip", handler.GetFromVipRecommend)
		api.GET("/personal", handler.GetFromPersonalRecommend)
		//banner内部接口，用于banner重载
		api.GET("/internal/banner_reload", handler.BannerReload)

		v2api := api.Group("v2")
		{
			v2api.GET("/homepage", handler.HomepageTopRecommend)
			v2api.GET("/homepage_default", handler.HomepageTopRecommend)
			v2api.GET("/billboard", handler.GetBillBoard)
			v2api.GET("/relate", handler.GetRelateRecommend)
			v2api.GET("/book_list", handler.GetBookList)
			//
			v2api.GET("/more", handler.HomepageTopMoreV2)
			v2api.GET("/more_default", handler.HomepageTopMoreV2)
			v2api.GET("/renew", handler.HomepageTopRenew)
			//
			v2api.GET("/recommend_stream", handler.HomepageRecommendStreamWrapper)
			v2api.GET("/recommend_stream_default", handler.HomepageRecommendStreamWrapper)

			// 用于记录新用户兴趣标签的接口
			v2api.POST("/user_interest_comic_types", handler.NewUserInterestCtypesRecordApi)
			v2api.GET("/user_interest_comic_types", handler.NewUserInterestCtypesApi)
		}

		v3api := api.Group("v3")
		{
			v3api.GET("/homepage", handler.HomepageTopRecommendV2)
			v3api.GET("/book_list", handler.GetBookListV2)
			v3api.GET("/feed_stream", handler.GetFeedStream)
			v3api.GET("/feed_stream_relate", handler.GetFeedStreamRelateRecommend)
			v3api.GET("/recommend_stream", handler.HomepageRecommendStreamV3Wrapper)
			v3api.GET("/recommend_stream_default", handler.HomepageRecommendStreamV3Wrapper)
			v3api.GET("/billboard", handler.GetBillBoardV2)
		}
		v4api := api.Group("v4")
		{
			v4api.GET("/homepage", handler.HomepageTopRecommendV3)
		}
	}
	g.GET("/recommend/type", handler.GetFromTypeRecommendDefault)

	// 微信小程序接口
	wxappApi := g.Group("wxapp/top")
	{
		wxappApi.GET("homepage", handler.GetHomepageTopRecommendV2Wxapp)
		wxappApi.GET("more", handler.GetHomepageTopMoreV2Default)
		wxappApi.GET("recommend_stream", handler.HomepageRecommendStreamWrapper)
	}

	return g
}
