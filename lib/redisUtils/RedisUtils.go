package redisUtils

const (
	StockCacheLua = "local stock = KEYS[1]" +
		"local stock_lock = KEYS[2]" +
		"local stock_val = tonumber(ARGV[1])" +
		"local stock_lock_val = tonumber(ARGV[2])" +
		"local is_exists = redis.call('EXISTS', stock)" +
		"if is_exists == 1 then" +
		"    return 0 \n" +
		"else" +
		"    redis.call('set', stock, stock_val)" +
		"    redis.call('set', stock_lock, stock_lock_val)" +
		"    return 1 \n" +
		"end"

	StockReduceLua = "local stock = KEYS[1] \n" +
		"local stock_lock = KEYS[2] \n" +
		"local stock_change = tonumber(ARGV[1]) \n" +
		"local stock_lock_change = tonumber(ARGV[2]) \n" +
		"local is_exists = redis.call('EXISTS', stock) \n" +
		"if is_exists == 1 then" +
		"	local stockAftChange = redis.call('incrby', stock, stock_change) \n" +
		"	if(stockAftChange < 0) then \n" +
		"		redis.call('decrby', stock, stock_change) \n" +
		"		return -1" +
		"	else" +
		"		local stockLockAftChange = redis.call('incrby', stock_lock, stock_lock_change) \n" +
		"		return {stockAftChange,stockLockAftChange} \n" +
		"	end \n" +
		"else \n" +
		"	return 0 \n" +
		"end"
)

/**
缓存sku库存 以及锁定库存
*/
//func SkuStockInit(stockKey string, stockLockKey string, stockVal interface{}, stockLockVal interface{}) (int, error) {
//	Script := redis.NewScript(StockCacheLua)
//	ret, err := Script.Run(RedisClient, []string{stockKey, stockLockKey}, stockVal, stockLockVal).Result()
//	if err != nil {
//		return 0, err
//	}
//	return int(ret.(int64)), nil
//}

/**
减库存 0 不存在 -1 不足  list 成功
*/
//func ReduceSkuStock(stockKey string, stockLockKey string, stockVal interface{}, stockLockVal interface{}) (interface{}, error) {
//	Script := redis.NewScript(StockReduceLua)
//	ret, err := Script.Run(RedisClient, []string{stockKey, stockLockKey}, stockVal, stockLockVal).Result()
//	if err != nil {
//		return 0, err
//	}
//	return ret, nil
//}

/**
设置订单过期
*/
//func CacheOrderExpired(orderId string) error {
//	err := RedisClient.Set("order:"+orderId, orderId, time.Minute*30).Err()
//	if err != nil {
//		err := RedisClient.Set("order:"+orderId, orderId, time.Minute*30).Err()
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func DelOrderExpired(orderId string) error {
//	err := RedisClient.Del("order:" + orderId).Err()
//	if err != nil {
//		err := RedisClient.Del("order:" + orderId).Err()
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

// SetLocalPriceToRedis 设置当前价格进redis
//func SetLineToRedis(stockCode, timeLine string) error {
//	err := RedisClient.Set(stockCode, timeLine, time.Minute*30).Err()
//	if err != nil {
//		err := RedisClient.Set(stockCode, timeLine, time.Minute*30).Err()
//		if err != nil {
//			log.Println("设置股票分钟K线数据进缓存失败----->", err.Error())
//			return err
//		}
//	}
//	return nil
//}

// GetLocalPriceToRedis 从redis获取当前价格
//func GetLineToRedis(stockCode string) (string, error) {
//	timeLine, err := RedisClient.Get(stockCode).Result()
//	if err != nil {
//		return "", fmt.Errorf("获取股票分钟K线数据进缓存失败,股票代码%s,错误--->%s", stockCode, err.Error())
//	}
//	return timeLine, nil
//
//}
