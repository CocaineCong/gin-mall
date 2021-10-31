package service

import (
	"FanOneMall/cache"
	"FanOneMall/model"
	"FanOneMall/pkg/e"
	"FanOneMall/pkg/logging"
	"FanOneMall/serializer"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/tealeg/xlsx"
	"strconv"
	"time"
)

//展示商品详情的服务
type ShowProductService struct {
}

type Product model.Product

//删除商品的服务
type DeleteProductService struct {
}

//更新商品的服务
type UpdateProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale string `form:"on_sale" json:"on_sale"`
	Num string `form:"num" json:"num"`
}

//上传商品
type UpProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale string `form:"on_sale" json:"on_sale"`
	Num string `form:"num" json:"num"`
}

//搜索商品的服务
type SearchProductsService struct {
	Search string `form:"search" json:"search"`
}

// 商品
func (service *ShowProductService) Show(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

//创建商品
func (service *CreateProductService) Create() serializer.Response {
	boss_id,_ := strconv.Atoi(service.BossID)
	product := model.Product{
		Name:          service.Name,
		CategoryID:    service.CategortID,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       service.ImgPath,
		Price:         service.Price,
		DiscountPrice: service.DiscoutPrice,
		BossID:        boss_id,
		BossName:      service.BossName,
		BossAvatar:    service.BossAvatar,
	}
	code := e.SUCCESS
	err := model.DB.Create(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListProductsService) List() serializer.Response {
	var products []model.Product
	total := 0
	code := e.SUCCESS
	if service.Limit == 0 {
		service.Limit = 15
	}
	if service.CategoryID == 0 {
		if err := model.DB.Model(model.Product{}).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).
			Offset(service.Start).Find(&products).
			Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(model.Product{}).
			Where("category_id=?", service.CategoryID).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("category_id=?", service.CategoryID).
			Limit(service.Limit).
			Offset(service.Start).
			Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	//res,_ := json.Marshal(data)
	//fmt.Println(string(res))
	var productList  []serializer.Product
	for _, item := range products {
		products := serializer.BuildProduct(item)
		productList = append(productList, products)
	}
	//for _,productL := range productList{
	//	fmt.Println(productL.Num)
	//}
	// 生成一个新的文件
	//file := xlsx.NewFile()
	file := xlsx.File{}
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	titleRow.AddCell().Value = "product"
	titleRow.AddCell().Value = "title"
	titleRow.AddCell().Value = "info"
	titleRow.AddCell().Value = "img_path"
	titleRow.AddCell().Value = "price"
	titleRow.AddCell().Value = "discount"
	titleRow.AddCell().Value = "boss_id"
	titleRow.AddCell().Value = "boss_name"
	// 插入内容
	for a := 0; a <= len(productList); a++ {
		row := sheet.AddRow()
		row.WriteSlice(productList[0],a)
	}
	// 保存文件
	Time := time.Now()
	year := Time.Year()
	month := Time.Month()
	fileName := fmt.Sprintf("%d年%d月",year,month)
	err := file.Save(fileName+".xlsx")
	if err != nil {
		panic(err)
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}


func (service *UpProductService) UpProduct() serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product,service.ID).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product.OnSale = service.OnSale
	err = model.DB.Save(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

//删除商品
func (service *DeleteProductService) Delete(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//更新商品
func (service *UpdateProductService) Update() serializer.Response {
	product := model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryID,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       service.ImgPath,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale: 	   service.OnSale,
	}
	product.ID = service.ID
	code := e.SUCCESS
	err := model.DB.Save(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//搜索商品
func (service *SearchProductsService) Show() serializer.Response {
	products := []model.Product{}
	code := e.SUCCESS
	err := model.DB.Where("name LIKE ?", "%s"+service.Search+"%").Find(&products).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var productsTemp []model.Product
	err = model.DB.Where("info LIKE ?", "%s"+service.Search+"%").Find(&productsTemp).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products = append(products, productsTemp...)
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProducts(products),
		Msg:    e.GetMsg(code),
	}
}

func ListenOrder() {
	go func() {
		for {
			opt := redis.ZRangeBy{
				Min:    strconv.Itoa(0),
				Max:    strconv.Itoa(int(time.Now().Unix())),
				Offset: 0,
				Count:  10,
			}
			orderList, err := cache.RedisClient.ZRangeByScore("SOMETHING", opt).Result()
			if err != nil {
				logging.Info("redis err: ", err)
			}
			if len(orderList) != 0 {
				var numList []int
				for _, v := range orderList {
					i, err := strconv.Atoi(v)
					if err != nil {
						logging.Info("Atoi err", err)
					}
					numList = append(numList, i)
				}
				if err := model.DB.Delete(&model.Order{}, "order_num IN (?)", numList).Error; err != nil {
					logging.Info("myql err:", err)
				}
				if err := cache.RedisClient.ZRem("SOMETHING", orderList).Err(); err != nil {
					logging.Info("redis err:", err)
				}
			}
		}
	}()
}


func (Product *Product) View() uint64 {
	//增加视频点击数
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(Product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//AddView 视频游览
func (Product *Product) AddView() {
	//增加视频点击数
	cache.RedisClient.Incr(cache.ProductViewKey(Product.ID))
	//增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Product.ID)))
}

//AddElecRank 增加加点排行点击数
func (Product *Product) AddElecRank() {
	//增加家电排汗点击数
	cache.RedisClient.ZIncrBy(cache.ElectricalRank, 1, strconv.Itoa(int(Product.ID)))
}

//AddAcceRank 增加配件排行点击数
func (Product *Product) AddAcceRank() {
	//增加配件排行点击数
	cache.RedisClient.ZIncrBy(cache.AccessoryRank, 1, strconv.Itoa(int(Product.ID)))
}

type CreateProductService struct {
	Name         string `form:"name" json:"name"`
	CategortID   int    `form:"category_id" json:"categort_id"`
	Title        string `form:"title" json:"title" bind:"required,min=2,max=100"`
	Info         string `form:"info" json:"info" bind:"max=1000"`
	ImgPath      string `form:"img_path" json:"img_path"`
	Price        string `form:"price" json:"price"`
	DiscoutPrice string `form:"discount_price" json:"discout_price"`
	BossID       string    `form:"boss_id" json:"boss_id" bind:"required"`
	BossName     string `form:"boss_name" json:"boss_name"`
	BossAvatar   string `form:"boss_avatar" json:"boss_avatar"`
}

type ListProductsService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
	CategoryID uint `form:"category_id" json:"category_id"`
}

