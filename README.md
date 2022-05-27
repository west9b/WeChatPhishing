# WeChatPhishing
go语言实现钓鱼抓取微信信息，搜索微信AccInfo.dat文件，从中正则匹配出当前微信名，手机号，邮箱，地区等信息，发送到cobaltstrike服务器。

# usage

### 01
CobaltStrike快速生成一个探针链接,把使用javaapplet获取信息勾选上
cobalt strike 选择Attacks ->Web Drive-by -> System Profiler
生成后替换wechat.go中的url->go build 编译exe
### 02
默认遍历"C:/users"可修改wechat.gol->dirname遍历指定盘符
当匹配到WeChat Files文件夹时才会遍历AccInfo.dat文件，缩短遍历时间
当有多个微信用户会一并正则将关键信息发送到CobaltStrike
### 03
缺点：exe后缀明显，投放方式各显神通
优点：正常文件遍历，不会报毒，配合cs返回IP地址，微信名，手机号，邮箱，等信息足够蓝队提交溯源报告
# Screenshots

# refer
本地提取电子取证直接使用已有的python脚本 https://github.com/ecat-sec/wechat_info_collect
