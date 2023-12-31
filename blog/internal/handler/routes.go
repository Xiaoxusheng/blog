// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	article "blog/internal/handler/article"
	category "blog/internal/handler/category"
	comment "blog/internal/handler/comment"
	tag "blog/internal/handler/tag"
	user "blog/internal/handler/user"
	userl "blog/internal/handler/userl"
	"blog/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: userl.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/rigister",
				Handler: userl.UserRigisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/adminRigister",
				Handler: userl.AdminRigisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken, serverCtx.AdminAuth},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/user/updateRole",
					Handler: user.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getUserList",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/updateIp",
					Handler: user.UpdateIpHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/talk/publishTalk",
					Handler: user.PublishTalkHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/talk/getTalkList",
					Handler: user.GetTalkListHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/talk/updateTalk",
					Handler: user.UpdateTalkHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/talk/deleteTalk",
					Handler: user.DeleteTalkHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/talk/revertTalk",
					Handler: user.RevertTalkHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/user/updatePassword",
					Handler: user.UpdatePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getUserInfoById",
					Handler: user.GetUserInfoByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/updateUserInfo",
					Handler: user.UpdateUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/talklike",
					Handler: user.TalkLikeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/cancelTalklike",
					Handler: user.CancelTalkLikeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getblogTalkList",
					Handler: user.GetBlogtTalkListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getTalkDetail",
					Handler: user.GetTalkDetailHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken, serverCtx.AdminAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/tag/addTag",
					Handler: tag.AddTagHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/tag/updateTag",
					Handler: tag.UpdateTagHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/tag/deleteTag",
					Handler: tag.DeleteTagHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/tag/getTagList",
					Handler: tag.GetTagListHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken, serverCtx.AdminAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/article/addArticle",
					Handler: article.AddArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/article/updateArticle",
					Handler: article.UpdateArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/article/updateArticleTop",
					Handler: article.UpdateArticleTopHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/deleteArticle",
					Handler: article.DeleteArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/recoveArticle",
					Handler: article.RecoveArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/isPublicArticle",
					Handler: article.IsPublicArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/GetArticleList",
					Handler: article.GetArticleListHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/article/blogArticleList",
					Handler: article.BlogArticleListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/blogTimeArticleList",
					Handler: article.BlogTimeArticleListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/getArticleByTag",
					Handler: article.GetArticleByTagHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/getArticleById",
					Handler: article.GetArticleByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/getRecommendArticleById",
					Handler: article.GetRecommendArticleByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/getArticleListByContent",
					Handler: article.GetArticleListByContentHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/getHotArticle",
					Handler: article.GetHotArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/like",
					Handler: article.LikeArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/unlike",
					Handler: article.UnlikeArticleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/article/addReadingDuration",
					Handler: article.AddReadingDurationHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/comment/addComment",
					Handler: comment.AddCommentHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/comment/DeleteComment",
					Handler: comment.DeleteCommentHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/comment/applyComment",
					Handler: comment.ApplyCommentHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/category/addCategory",
					Handler: category.AddCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/category/updateCategory",
					Handler: category.UpdateCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/category/deleteCategory",
					Handler: category.DeleteCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/category/getCategoryList",
					Handler: category.GetCategoryListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/category/getCategoryDictionary",
					Handler: category.GetCategoryDictionaryHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)
}
