package models

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/lexkong/log"

	"pppobear.cn/jxc-backend/module/handler"
	_ "pppobear.cn/jxc-backend/module/logger"

	"pppobear.cn/jxc-backend/config"
)

var Model *gorm.DB

func init() {
	var err error
	dbConfig := config.GetEnv()
	Model, err = gorm.Open(
		"mssql",
		fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			dbConfig.DatabaseUsername,
			dbConfig.DatabasePassword,
			dbConfig.DatabaseIP,
			dbConfig.DatabasePort,
			dbConfig.DatabaseName,
		),
	)

	if err != nil {
		log.Fatal("Database fail to connect.", err)
	}
	Model.LogMode(config.GetEnv().Debug)
	Model.DB().SetMaxIdleConns(0)
}

type Filter interface {
	Filter(db *gorm.DB)
}

func List(data interface{}, model interface{}, req *handler.ListRequest, filter Filter,
	user *map[string]interface{}) (handler.ListResponse, error) {
	var (
		total,
		totalPage uint64
		resp handler.ListResponse
		err  error
	)

	sql := Model.Model(model)
	if filter != nil {
		filter.Filter(sql)
	}

	switch sl := data.(type) {
	case []*PurchasesModel:
		HideOthers(sql, "buyer_id", user)
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, PurchasesSearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page-1).Limit(req.PageSize).Order(req.OrderBy).
			Set("gorm:auto_preload", true).Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl

	case []*SalesModel:
		HideOthers(sql, "sales_staff_id", user)
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, SalesSearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page-1).Limit(req.PageSize).Order(req.OrderBy).
			Set("gorm:auto_preload", true).Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl

	case []*CustomerModel:
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, CustomerSearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page - 1).Limit(req.PageSize).Order(req.OrderBy).
			Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl

	case []*GoodsModel:
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, GoodsSearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page - 1).Limit(req.PageSize).Order(req.OrderBy).
			Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl

	case []*InventoryModel:
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, InventorySearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page-1).Limit(req.PageSize).Order(req.OrderBy).
			Set("gorm:auto_preload", true).Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl

	case []*StaffModel:
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, StaffSearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page-1).Limit(req.PageSize).Order(req.OrderBy).
			Set("gorm:auto_preload", true).Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl
	case []*DepartmentModel:
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, DepSearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page - 1).Limit(req.PageSize).Order(req.OrderBy).
			Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl
	case []*SpecialtyModel:
		if s := strings.Trim(req.Search, " "); s != "" {
			Search(sql, SpeSearchFields, s)
		}
		total, totalPage, err = formatPage(sql, req)
		if err = sql.Offset(req.Page - 1).Limit(req.PageSize).Order(req.OrderBy).
			Find(&sl).Error; err != nil {
			return resp, err
		}
		resp.Data = sl
	}

	resp.Page = uint64(req.Page)
	resp.PageSize = uint64(req.PageSize)
	resp.Total = total
	resp.TotalPage = uint64(totalPage)
	return resp, nil
}

func HideOthers(sql *gorm.DB, userIdCol string, user *map[string]interface{}) {
	if (*user)["permission"] != 1 && (*user)["id"] != "" {
		*sql = *sql.Where(userIdCol+" = ?", (*user)["id"])
	}
}

func Search(sql *gorm.DB, searchColumn []string, search string) {
	var searchSql strings.Builder
	searchSql.Grow(200)
	searchSql.WriteString("(")
	search = regexp.MustCompile(" +").ReplaceAllString(search, " ")
	search = regexp.MustCompile(`[+.!/_,$%^*("']|[?！，。？、~@#￥%…&*（）]`).ReplaceAllString(search, " ") // 防止SQL注入
	words := strings.Split(search, " ")
	var fk, pk, tb, cols string
	for i, column := range searchColumn {
		if i != 0 {
			searchSql.WriteString("OR ")
		}
		searchSql.WriteString("(")
		otherTablePat := regexp.MustCompile(`(.*?)--(.*?):(.*?)\.(.*)`)
		if otherTablePat.MatchString(column) {
			res := otherTablePat.FindStringSubmatch(column)
			fk, pk, tb, cols = res[1], res[2], res[3], res[4]
			colStrs := strings.Split(cols, "&")
			searchSql.WriteString(fk + " IN (SELECT " + pk + " FROM [" + tb + "] WHERE (")
			for j, col := range colStrs {
				if j != 0 {
					searchSql.WriteString(" OR")
				}
				searchSql.WriteString(" (")
				for k, word := range words {
					if k != 0 {
						searchSql.WriteString(" AND ")
					}
					searchSql.WriteString(col + " LIKE '%" + word + "%'")
				}
				searchSql.WriteString(")")
			}
			searchSql.WriteString(")")
			searchSql.WriteString(")")
		} else {
			searchSql.WriteString(" (")
			for j, word := range words {
				if j != 0 {
					searchSql.WriteString(" AND ")
				}
				searchSql.WriteString(column + " LIKE '%" + word + "%'")
			}
			searchSql.WriteString(")")
		}
		searchSql.WriteString(")")
	}
	searchSql.WriteString(")")
	*sql = *sql.Where(searchSql.String())
}

func formatPage(sql *gorm.DB, req *handler.ListRequest) (total, totalPage uint64, err error) {
	err = sql.Count(&total).Error
	if req.PageSize < 1 {
		req.PageSize = int(config.GetPagination().DefaultPageSize)
	}
	totalPage = uint64(math.Ceil(float64(total) / float64(req.PageSize)))
	if req.Page > int(totalPage) {
		req.Page = int(totalPage)
	} else if req.Page < 1 {
		req.Page = 1
	}

	orderBy := strings.Trim(req.OrderBy, " ")
	if len(orderBy) == 0 {
		orderBy = "id"
	}
	if orderBy[0] == '-' {
		req.OrderBy = string(orderBy[1:]) + " desc"
	} else {
		req.OrderBy = orderBy
	}
	return
}
