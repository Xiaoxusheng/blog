// Code generated by goctl. DO NOT EDIT.
package types

type UserLoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserLoginResp struct {
	Respone
	Data map[string]string `json:"data,default="" "`
}

type UserRigisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
	NickName string `from:"nickName,optional"`
	QQ       int    `from:"qq,optional"`
}

type UserRigisterResp struct {
	Respone
}

type UpdateRoleReq struct {
	Id   string `json:"id"`
	Role int    `json:"role"`
}

type UpdateRoleResp struct {
	Respone
}

type GetUserListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetUserListResp struct {
	Respone
	Data map[string][]*Userlist `json:"data"`
}

type AdminRigisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
	NickName string `form:"nickName,optional"`
	QQ       int    `form:"qq,optional"`
}

type AdminRigisterResp struct {
	Respone
}

type UpdatePasswordReq struct {
	Password string `json:"password"`
}

type UpdatePasswordResp struct {
	Respone
}

type UpdateIpReq struct {
	Id string `json:"id"`
	Ip string `json:"ip"`
}

type UpdateIpResp struct {
	Respone
}

type GetUserInfoByIdReq struct {
}

type GetUserInfoByIdResp struct {
	Respone
	Data map[string]Userlist `json:"data"`
}

type UpdateUserInfoReq struct {
	Avatar   string `json:"avatar,optional"`
	NickName string `json:"nickName,optional"`
	QQ       int    `json:"qq,optional"`
}

type UpdateUserInfoResp struct {
	Respone
}

type PublishTalkReq struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
	Imglist string `json:"imglist"`
}

type PublishTalkResp struct {
	Respone
	Data map[string]string `json:"data"`
}

type GetTalkListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetTalkListResp struct {
	Respone
	Data map[string][]*TalkList `json:"data"`
}

type UpdateTalkReq struct {
	Id      string `json:"identity"`
	Content string `json:"content,optional"`
	Imglist string `json:"imglist,optional"`
	Status  int    `json:"status,optional"`
	Istop   int    `json:"istop,optional"`
}

type UpdateTalkResp struct {
	Respone
}

type DeleteTalkReq struct {
	Id string `json:"identity"`
}

type DeleteTalkResp struct {
	Respone
}

type RevertTalkReq struct {
	Id string `json:"identity"`
}

type RevertTalkResp struct {
	Respone
}

type TalkLikeReq struct {
	Id string `json:"id"`
}

type TalkLikeResp struct {
	Respone
}

type CancelTalkLikeReq struct {
	Id string `json:"id"`
}

type CancelTalkLikeResp struct {
	Respone
}

type GetBlogtTalkListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetBlogtTalkListResp struct {
	Respone
	Data map[string][]*TalkList `json:"data"`
}

type GetTalkDetailReq struct {
	Id string `json:"id"`
}

type GetTalkDetailResp struct {
	Respone
	Data map[string]*TalkList `json:"data"`
}

type AddTagReq struct {
	TagName string `json:"tagName"`
}

type AddTagResp struct {
	Respone
}

type UpdateTagReq struct {
	Id      string `json:"id"`
	TagName string `json:"tagName"`
}

type UpdateTagResp struct {
	Respone
}

type DeleteTagReq struct {
	Id string `json:"id"`
}

type DeleteTagResp struct {
	Respone
}

type GetTagListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetTagListResp struct {
	Respone
	Data map[string][]*TagList `json:"data"`
}

type AddArticleReq struct {
	Article Articlelist `json:"article"`
}

type AddArticleResp struct {
	Respone
}

type UpdateArticleReq struct {
	Article Articlelist `json:"article"`
}

type UpdateArticleResp struct {
	Respone
}

type UpdateArticleTopReq struct {
	Id    string `json:"id"`
	IsTop int    `json:"isTop"`
}

type UpdateArticleTopResp struct {
	Respone
}

type DeleteArticleReq struct {
	Id string `json:"id"`
}

type DeleteArticleResp struct {
	Respone
}

type RecoveArticleReq struct {
	Id string `json:"id"`
}

type RecoveArticleResp struct {
	Respone
}

type IsPublicArticleReq struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}

type IsPublicArticleResp struct {
	Respone
}

type GetArticleListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetArticleListResp struct {
	Respone
	Data map[string][]*ArticleList `json:"data"`
}

type BlogArticleListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type BlogArticleListResp struct {
	Respone
	Data map[string][]*ArticleList `json:"data"`
}

type BlogTimeArticleListReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type BlogTimeArticleListResp struct {
	Respone
	Data map[string][]*ArticleList `json:"data"`
}

type GetArticleByTagReq struct {
	Tag string `json:"tag"`
}

type GetArticleByTagResp struct {
	Respone
	Data map[string][]*ArticleList `json:"data"`
}

type GetArticleByIdReq struct {
	Id string `json:"id"`
}

type GetArticleByIdResp struct {
	Respone
	Data map[string][]*ArticleList `json:"data"`
}

type GetRecommendArticleByIdReq struct {
	Id string `json:"id"`
}

type GetRecommendArticleByIdResp struct {
	Respone
	Data map[string][]*ArticleList `json:"data"`
}

type GetArticleListByContentReq struct {
	Content string `json:"content"`
}

type GetArticleListByContentResp struct {
	Respone
	Data map[string]ArticleList `json:"data"`
}

type GetHotArticleReq struct {
}

type GetHotArticleResp struct {
	Respone
	Data map[string]ArticleList `json:"data"`
}

type LikeArticleReq struct {
	Id string `json:"id"`
}

type LikeArticleResp struct {
	Respone
}

type UnlikeArticleReq struct {
	Id string `json:"id"`
}

type UnlikeArticleResp struct {
	Respone
}

type AddReadingDurationReq struct {
	Id       string  `json:"id"`
	Duration float64 `json:"duration"`
}

type AddReadingDurationResp struct {
	Respone
}

type AddCommentReq struct {
	Id         string `json:"id"`
	Type       int    `json:"type"`
	ForId      string `json:"forId"`
	FromId     string `json:"fromId"`
	FromName   string `json:"fromName"`
	FromAvatar string `json:"fromAvatar"`
	Content    string `json:"content"`
	Ip         string `json:"ip"`
}

type AddCommentResp struct {
	Respone
}

type DeleteCommentReq struct {
	Id string `json:"id"`
}

type DeleteCommentResp struct {
	Respone
}

type ApplyCommentReq struct {
	Id         string `json:"id"`
	Type       int    `json:"type"`
	ForId      string `json:"forId"`
	FromId     string `json:"fromId"`
	FromName   string `json:"fromName"`
	FromAvatar string `json:"fromAvatar"`
	Content    string `json:"content"`
	Ip         string `json:"ip"`
	ToId       string `json:"toId"`
	ToName     string `json:"toName"`
	ToAvatar   string `json:"toAvatar"`
}

type ApplyCommentResp struct {
	Respone
}

type CategoryReq struct {
	CategoryName string `json:"categoryName"`
}

type CategoryResp struct {
	Respone
}

type UpdatecategoryReq struct {
	Id           string `json:"id"`
	CategoryName string `json:"categoryName"`
}

type UpdatecategoryResp struct {
	Respone
}

type DeleteCategoryReq struct {
	Id string `json:"id"`
}

type DeleteCategoryResp struct {
	Respone
}

type GetCategoryListReq struct {
}

type GetCategoryListResp struct {
	Respone
	Data map[string][]*CategoryList `json:"data"`
}

type GetCategoryDictionaryReq struct {
}

type GetCategoryDictionaryResp struct {
	Respone
	Data map[string]string `json:"data"`
}

type Respone struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Userlist struct {
	Identity  string `json:"identity"`
	Username  string `json:"username"`
	Role      int    `json:"role"`
	Nick_name string `json:"nick_Name"`
	Avatar    string `json:"avatar"`
	Qq        int    `json:"qq"`
	Ip        string `json:"ip"`
}

type TalkList struct {
	Identity  string `json:"identity"`
	UserId    string `json:"userId,omitempty"`
	Content   string `json:"content,omitempty"`
	Status    int    `json:"status,omitempty"`
	IsTop     int    `json:"isTop,omitempty"`
	LikeTimes int    `json:"likeTimes" json:"likeTimes,omitempty"`
	Url       string `json:"url,omitempty"`
}

type TagList struct {
	Identity string `json:"identity"`
	TagName  string `json:"tag_Name"`
	Status   int    `json:"status"`
}

type Articlelist struct {
	Identity           string `json:"identity"`
	ArticleTitle       string `json:"articleTitle"`
	AuthorId           int    `json:"authorId"`
	CategoryId         int    `json:"categoryId"`
	ArticleContent     string `json:"articleContent"`
	ArticleCover       string `json:"articleCover"`
	IsTop              int    `json:"isTop"`
	Status             int    `json:"status"`
	Type               int    `json:"type"`
	OriginUrl          string `json:"originUrl"`
	ViewTimes          int    `json:"viewTimes"`
	ArticleDescription string `json:"articleDescription"`
	ThumbsUpTimes      int    `json:"thumbsUpTimes"`
	ReadingDuration    int    `json:"readingDuration"`
}

type ArticleList struct {
	Identity           string `json:"identity"`
	ArticleTitle       string `json:"articleTitle"`
	AuthorId           int    `json:"authorId"`
	CategoryId         int    `json:"categoryId"`
	ArticleContent     string `json:"articleContent"`
	ArticleCover       string `json:"articleCover"`
	IsTop              int    `json:"isTop"`
	Status             int    `json:"status"`
	Type               int    `json:"type"`
	OriginUrl          string `json:"originUrl"`
	ViewTimes          int    `json:"viewTimes"`
	ArticleDescription string `json:"articleDescription"`
	ThumbsUpTimes      int    `json:"thumbsUpTimes"`
	ReadingDuration    int    `json:"readingDuration"`
}

type CategoryList struct {
	Identity     string `json:"identity"`
	Status       int    `json:"status"`
	CategoryName string `json:"categoryName"`
}
