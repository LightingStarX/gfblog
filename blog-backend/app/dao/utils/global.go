package utils

const (
	TableArticleId                = 6601
	TableUserId                   = 6604
	TableCommentId                = 6610
	TableRoleId                   = 6617
	TableArticleCategoryId        = 6622
	TableArticleExtId             = 6629
	TableUserArticleMappingId     = 6638
	TableArticleCategoryMappingId = 6641
	TableUserProfileId            = 6679
	TableUserRoleMappingId        = 6742
	TableUserRelation             = 6833
)

var (
	TableMap = map[string]int{
		"article":                  TableArticleId,
		"user":                     TableUserId,
		"role":                     TableRoleId,
		"article_category":         TableArticleCategoryId,
		"comment":                  TableCommentId,
		"article_ext":              TableArticleExtId,
		"user_article_mapping":     TableUserArticleMappingId,
		"article_category_mapping": TableArticleCategoryMappingId,
		"user_profile":             TableUserProfileId,
		"user_role_mapping":        TableUserRoleMappingId,
		"user_relation":            TableUserRelation,
	}
)
