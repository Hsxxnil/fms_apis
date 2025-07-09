# FMS APIs

一套以 **Golang** + **PostgreSQL** 為後端、**Angular** 為前端框架開發的 **車隊管理系統**，系統整合多項核心功能，協助企業有效管理車輛資源、提升運營效率並降低營運成本。
系統功能涵蓋：
* 即時車輛追蹤：透過定位系統掌握每輛車的即時位置與行駛狀況
* 智慧路線優化：根據目的地、自訂條件與交通狀況自動建議最佳路線
* 車輛維護提醒：依照行駛里程與保養週期，自動提示維修與保養時機，降低突發故障風險
* 數據視覺化報表：彙整車輛使用、油耗、維修紀錄等資料，協助管理者做出精準決策

系統採用前後端分離架構，具備良好的可擴充性與維護彈性，適用於物流、配送、租賃等多種車隊場景，全面強化車隊管理與營運品質。

#Golang #Gin #PostgreSQL #Angular #PrimeNG #Swagger #GoogleMapAPI #藍新金流API #IoT

## 專案連結

* 前端畫面：[點我查看](https://hsxxnil.notion.site/FMS-11c5b51f95f581669849fe01de74b605)
* Swagger API 文件：[點我查看](https://hsxxnil.github.io/swagger-ui/?urls.primaryName=FMS)

## 安裝
1. 下載專案

```bash
git clone https://github.com/Hsxxnil/fms_apis.git
cd fms_apis
```

2. 建立 Makefile

> 請根據您的作業系統選擇對應的範本進行複製：
* Linux / macOS
```bash
cp Makefile.example.linux Makefile
```

* Windows
```bash
copy Makefile.example.windows Makefile
```

3. 初始化

> 如為初次建立開發環境，請先根據您的作業系統安裝必要套件：
* Linux / macOS
```bash
brew install golang-migrate golangci-lint protobuf
```

* Windows（建議使用 Scoop，或手動安裝以下套件）：
```bash
scoop install golang-migrate golangci-lint protobuf
```

> 執行以下指令將自動安裝依賴套件並建立必要的目錄結構：
```bash
make setup
```

4. 設定環境參數

> 開啟並編輯以下檔案，填入資料庫連線資訊、JWT 金鑰等必要參數：
```file
config/debug_config.go
```

5. 更新套件

>執行以下指令升級相關套件
```bash
make update_lib
```

## 資料庫遷移

> 執行以下指令使用[golang-migrate](https://github.com/golang-migrate/migrate)做資料庫遷移及做資料表版控：
```bash
make migration
```

## 執行
> 執行以下指令在本地端啟動伺服器並自動重載：
```bash
make air
```

## License

本專案使用的 [Vodka](https://github.com/dylanlyu/vodka) 採用 [MIT License](https://opensource.org/licenses/MIT) 授權。
