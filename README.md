ShoppingCart 購物車後端系統 golang
===

[![](https://i.imgur.com/F1ZGUqG.png "Swagger")](https://shoppingcart.up.railway.app/swagger/index.html)

[REST API DEMO(使用SwaggerUI展示)](https://shoppingcart.up.railway.app/swagger/index.html)



## 簡介
本系統實現了電商購物的核心功能。
內容包括：會員註冊登入、jwt驗證、商品分類、商品管理、訂單管理、購物車管理。


## 使用技術

* mysql+gorm
* web框架gin
* JWT驗證
* swagger api 
* Viper配置管理
* 分頁功能
* docker部署 

## 專案架構
採用mvc設計模式，有五大模組：

1. user 使用者
2. category 商品分類
3. product 商品
4. order 訂單
5. cart 購物車