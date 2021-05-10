package cache

// var (
// 	SET_PRODUCT_KEY = "Product:GO_PRODUCT"
// 	SET_CART_KEY    = "Cart:GO_CART/v3/cart?"
// )

// type ProductsCacheStruct struct {
// 	Products      []entities.GetAllProduct    `json:"products"`
// 	CountProducts entities.CountGetAllProduct `json:"count"`
// }

// type CartCacheStruct struct {
// 	Products []entities.GetAllCart `json:"products"`
// }
// type ProductsCacheEnStruct struct {
// 	Products      []entities.GetAllProductEn  `json:"products"`
// 	CountProducts entities.CountGetAllProduct `json:"count"`
// }

// func SetProductTest(key string, value string, expires time.Duration) {
// 	client := redisConnect()
// 	client.Set(key, value, expires*time.Second)
// }

// func SetProduct(key string, value ProductsCacheStruct, expires time.Duration) {
// 	client := redisConnect()

// 	json, err := json.Marshal(value)
// 	if err != nil {
// 		panic(err)
// 	}
// 	client.Set(key, json, expires*time.Second)
// }

// func GetProduct(key string) ProductsCacheStruct {
// 	client := redisConnect()

// 	val, err := client.Get(key).Result()
// 	if err != nil {
// 		return ProductsCacheStruct{}
// 	}

// 	product := ProductsCacheStruct{}
// 	err = json.Unmarshal([]byte(val), &product)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return product
// }
