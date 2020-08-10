## 首页推荐接口说明

### 首页V2
- 测试URL
    http://test.samh.xndm.tech/top/v3/homepage
- 请求方式
    GET
- 示例
    http://test.samh.xndm.tech/top/v2/homepage?uid=77948219&gender=0&device_id=864429030653832&platformname=android
- 请求参数

| 请求参数 | 参数类型 | 是否必填 | 参数说明 |
|:------ |:------ |:------ |:------ |
| uid | string | 是|用户ID |
| udid | string | 是| 唯一标识一台硬件设备ID |
| gender | int | 否 | 用户性别，0为男，1为女（默认为男） |
| platform | string | 是 | 平台，目前支持android, ios |
| platformname | string | 是 | 平台名称，目前支持android, iphone, ipad |

### 首页栏目更多接口V2
- 测试URL
    http://test.samh.xndm.tech/top/v2/more
- 请求方式
    GET
- 示例
    http://test.samh.xndm.tech/top/v2/more?uid=77948219&gender=0&device_id=864429030653832&platformname=android&section_id=1&page_num=1&page_size=10
- 请求参数

| 请求参数 | 参数类型 | 是否必填 | 参数说明 |
|:------ |:------ |:------ |:------ |
| uid | string | 是|用户ID |
| udid | string | 是| 唯一标识一台硬件设备ID |
| gender | int | 否 | 用户性别，0为男，1为女（默认为男） |
| platform | string | 是 | 平台，目前支持android, ios |
| platformname | string | 是 | 平台名称，目前支持android, iphone, ipad |
| section_id | string | 是 | 首页栏目id |
| page_num | int | 是 | 页数（为空或为0时表示换一换，生成随机页数） |
| page_size | int | 是 | 返回漫画个数 |

### 首页栏目编辑精选V2（推荐流，老版本）
- 测试URL
    http://test.samh.xndm.tech/top/v2/recommend_stream
- 请求方式
    GET
- 示例
    http://test.samh.xndm.tech/top/v2/recommend_stream?uid=77948219&gender=0&device_id=864429030653832&platformname=android&page_size=18
- 请求参数

| 请求参数 | 参数类型 | 是否必填 | 参数说明 |
|:------ |:------ |:------ |:------ |
| uid | string | 是|用户ID |
| udid | string | 是| 唯一标识一台硬件设备ID |
| gender | int | 否 | 用户性别，0为男，1为女（默认为男） |
| platform | string | 是 | 平台，目前支持android, ios |
| platformname | string | 是 | 平台名称，目前支持android, iphone, ipad |
| page_size | int | 是 | 返回漫画个数 |


### 首页栏目编辑精选V3（推荐流,老版本）
- 测试URL
    http://test.samh.xndm.tech/top/v3/recommend_stream
- 请求方式
    GET
- 示例
    http://test.samh.xndm.tech/top/v3/recommend_stream?uid=77948219&gender=0&device_id=864429030653832&platformname=android&page_size=12&page_num=1
- 请求参数

| 请求参数 | 参数类型 | 是否必填 | 参数说明 |
|:------ |:------ |:------ |:------ |
| uid | string | 是|用户ID |
| udid | string | 是| 唯一标识一台硬件设备ID |
| gender | int | 否 | 用户性别，0为男，1为女（默认为男） |
| platform | string | 是 | 平台，目前支持android, ios |
| platformname | string | 是 | 平台名称，目前支持android, iphone, ipad |
| page_size | int | 是 | 返回漫画个数 |
| page_num | int | 是 | 页数为1时返回display_type固定为7，之后为8 |

### 获取新用户感兴趣的漫画标签
- 测试URL
    http://test.samh.xndm.tech/top/v2/user_interest_comic_types
- 请求方式
    GET
- 示例
    http://test.samh.xndm.tech/top/v2/user_interest_comic_types

### 上报新用户感兴趣的漫画类型
- 请求URL
    https://recommend.samh.xndm.tech:6443/top/v2/user_interest_comic_types
- 测试URL
    http://test.samh.xndm.tech/top/v2/user_interest_comic_types
- 请求方式
    POST
- 示例
    同时可接受json和form参数
    `curl -X POST "http://test.samh.xndm.tech/top/v2/user_interest_comic_types" --data '{"uid":"123445", "comic_types":"玄幻|热血|校园"}' -H "Content-Type:application/json"`
    `curl -X POST "http://test.samh.xndm.tech/top/v2/user_interest_comic_types" --data 'uid=123465&comic_types=玄幻|热血|校园' -H "Content-Type:application/x-www-form-urlencoded"`


### 首页看了又看

- 测试URL
    http://test.samh.xndm.tech/top/v2/relate
- 请求方式
    GET
- 示例
    http://test.samh.xndm.tech/top/v2/relate?uid=77948219&gender=0&device_id=864429030653832&platformname=android&page_size=18
- 请求参数

| 请求参数 | 参数类型 | 是否必填 | 参数说明 |
|:------ |:------ |:------ |:------ |
| uid | string | 是|用户ID |
| udid | string | 是| 唯一标识一台硬件设备ID |
| gender | int | 否 | 用户性别，0为男，1为女（默认为男） |
| platform | string | 是 | 平台，目前支持android, ios |
| platformname | string | 是 | 平台名称，目前支持android, iphone, ipad |
| page_size | int | 是 | 返回漫画个数 |
