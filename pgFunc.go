package main

import (
	"pgFunc/spider"
)

func main() {

	html :=`
<!doctype html>
<html lang="zh-CN">

<head>
  <title>腾讯首页</title>
  <meta charset="gb2312">
  <meta http-equiv="X-UA-Compatible" content="IE=Edge" />
  <meta name="baidu-site-verification" content="cNitg6enc2" />
  <meta name="baidu_union_verify" content="4508fc7dced37cf569c36f88135276d2">
  <meta name="theme-color" content="#FFF" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="format-detection" content="telephone=no">
  <!-- <script src="//js.aq.qq.com/js/aq_common.js"></script> -->
  <script type="text/javascript">
try {
  if (location.search.indexOf('?pc') !== 0 && /Android|Windows Phone|iPhone|iPod/i.test(navigator.userAgent)) {
    window.location.href = 'https://xw.qq.com?f=qqcom';
  }
} catch (e) {}
</script><!--[if !IE]>|xGv00|2d5210e6c1b95e3bf4b8983f9cb00ab3<![endif]-->
  <meta content="资讯,新闻,财经,房产,视频,NBA,科技,腾讯网,腾讯,QQ,Tencent" name="Keywords">
  <meta name="description" content="腾讯网从2003年创立至今，已经成为集新闻信息，区域垂直生活服务、社会化媒体资讯和产品为一体的互联网媒体平台。腾讯网下设新闻、科技、财经、娱乐、体育、汽车、时尚等多个频道，充分满足用户对不同类型资讯的需求。同时专注不同领域内容，打造精品栏目，并顺应技术发展趋势，推出网络直播等创新形式，改变了用户获取资讯的方式和习惯。" />
  <link rel="shortcut icon" href="//mat1.gtimg.com/www/icon/favicon2.ico" />
  <link rel="stylesheet" href="//mat1.gtimg.com/qqcdn/qqindex2021/dist/css/qq_c4b4edd7.css" charset="utf-8">
  <style>
    #skin-close{
      display: none;
    }
    /* .qq-body-skin{
      background-color: #d20001!important;
    } */
    .qq-body-skin .layout{
      margin-bottom:0;
      padding-bottom: 20px;
    }
    .qq-body-skin .qq-skin,
    .qq-body-skin .qq-top{
      padding-bottom: 0;
    }
    .garyBody{
      filter: grayscale(100%);
      -webkit-filter: grayscale(100%);
      -moz-filter: grayscale(100%);
      -ms-filter: grayscale(100%);
      -o-filter: grayscale(100%);
      -webkit-filter: grayscale(1);
    }
    .iebrowser{
        background-color: #ddd ;
        opacity: 0.5;
        filter: Alpha(opacity=50);
    }
    .iebrowser .qq-nav .nav-mod{
        background-color: #b0b0b1 !important;
    }

    .iebrowser .searchBtn{
      background-color: #b0b0b1 !important;
    }

    .iebrowser .mod .hd .tit.active a{
      display: inline-block !important;
      color: #161617 !important;
      border-bottom: 4px solid #212121 !important;
    }
    .iebrowser .mod .bd .cate {
      color: #212222 !important;
    }
    .iebrowser  a:hover {
      color: #030303 !important;
    }
    .p1t {
      overflow: hidden;
    }
    
    .p1t img {
      width: auto;
      margin-left: -9px;
    }
  </style>
</head>

<body>

  <div class="global">

    <!-- 大皮肤 -->
    <div id="big-skin" class="layout qq-skin"></div>
    <!-- /大皮肤 -->

    <!-- 头部 -->
    <div class="layout qq-top cf" bossexpo="bg_top">

      <h1 class="top-logo fl">
        <a href="/" target="_blank" bosszone="top_logo">
          <img width="100%" src="//mat1.gtimg.com/pingjs/ext2020/qqindex2018/dist/img/qq_logo_2x.png" alt="腾讯网">
        </a>
      </h1>

      <!-- 小皮肤 -->
      <div id="small-skin" class="skin-min fl"></div>
      <!-- /小皮肤 -->

      <!-- 搜索 -->
      <div class="top-search fl" id="sosobar" role="search" bosszone="top_search">
        <form id="searchForm" method="get" name="soso_search_box" action="https://www.sogou.com/tx?hdq=sogou-wsse-3f7bcd0b3ea82268-0001&ie=utf-8&query="
          target="_blank">
          <div id="searchTxt" class="searchTxt">
            <input type="hidden" value="w.q.in.sb.web" name="cid" />
            <div class="searchMenu fl">
              <div class="searchSelected" id="searchSelected">网页</div>
              <div class="searchTab" id="searchTab">
                <ul></ul>
              </div>
            </div>
            <input id="sougouTxt" type="text" value="" name="w" aria-label="请输入搜索文字" />
            <div class="searchSmart" id="searchSmart" style="display:none;">
              <ul></ul>
            </div>
            <div class="fr">
              <button id="searchBtn" class="searchBtn" type="submit">搜狗搜索</button>
            </div>
          </div>
        </form>
      </div>
      <script type="text/javascript">
        function sogouShow() {}
        function sosoShow() {}
      </script>
      <!-- /搜索 -->

      <!-- 登录 -->
      <div id="top-login" class="top-login fr">
        <div class="item item-qzone fl">
          <a href="https://qzone.qq.com" class="q-icons l-qzone" target="_blank" bosszone="top_qzone">Qzone</a>
          <div class="pop">
            <i class="arr-icon"></i>
            <a class="txt" href="https://qzone.qq.com" target="_blank" bosszone="top_qzone">点击查看QQ空间</a>
          </div>
        </div>
        <div class="item item-qmail fl">
          <a href="https://mail.qq.com" class="q-icons l-qmail" target="_blank" bosszone="top_mail">Qmail</a>
          <div class="pop">
            <i class="arr-icon"></i>
            <a class="txt" href="https://mail.qq.com/cgi-bin/loginpage" target="_blank" bosszone="top_mail">点击查看QQ邮箱</a>
          </div>
        </div>
        <div class="item item-login fl">
          <a class="l-login" href="javascript:;" onclick="userLogin()" bosszone="top_login">登录</a>
          <div class="pop">
            <i class="arr-icon"></i>
            <div class="nick">你好，</div>
            <a class="loginout" href="javascript:;" onclick="login.loginOut()" bosszone="top_login">[退出登录]</a>
          </div>
        </div>
      </div>
      <!-- /登录 --><!--515e8a9b2bd0d8267501d39b452aab86--><!--[if !IE]>|xGv00|40c170743488334a9b50c847f97eac0c<![endif]-->

    </div>
    <!-- /头部 -->

    <!-- 导航 -->
    <div class="layout qq-nav">
      <div class="nav-mod cf">
        <style type="text/css">
.qq-nav .nav-mod .nav-item{white-space: nowrap;}
</style>

<ul class="nav-main fl" bossexpo="bg_dh_1">
    <li class="nav-item">
    <a href="http://news.qq.com/" target="_blank" bosszone="dh_1">新闻</a>
  </li>
    <li class="nav-item">
    <a href="http://v.qq.com/" target="_blank" bosszone="dh_2">视频</a>
  </li>
    <li class="nav-item">
    <a href="http://new.qq.com/ch/photo/" target="_blank" bosszone="dh_3">图片</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/milite/" target="_blank" bosszone="dh_4">军事</a>
  </li>
    <li class="nav-item">
    <a href="https://2020.qq.com/" target="_blank" bosszone="dh_5">奥运会</a>
  </li>
    <li class="nav-item">
    <a href="https://sports.qq.com/" target="_blank" bosszone="dh_6">体育</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/ent/" target="_blank" bosszone="dh_7">娱乐</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/finance/" target="_blank" bosszone="dh_8">财经</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/tech/" target="_blank" bosszone="dh_9">科技</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/fashion/" target="_blank" bosszone="dh_10">时尚</a>
  </li>
    <li class="nav-item">
    <a href="http://auto.qq.com/" target="_blank" bosszone="dh_11">汽车</a>
  </li>
    <li class="nav-item">
    <a href="http://house.qq.com/" target="_blank" bosszone="dh_12">房产</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/edu/" target="_blank" bosszone="dh_13">教育</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/cul/" target="_blank" bosszone="dh_14">文化</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/astro/" target="_blank" bosszone="dh_15">星座</a>
  </li>
    <li class="nav-item">
    <a href="https://new.qq.com/ch/games/" target="_blank" bosszone="dh_16">游戏</a>
  </li>
  </ul><!--9673133db6a5f928e28dc974457a5081-->
        <div class="nav-more fl">
  <div class="more-txt" bosszone="dh_more">更多</div>
  <div class="nav-sub" bossexpo="bg_dh_2">
    <ul class="sub-list cf">
            <li class="nav-item">
        <a href="https://new.qq.com/ch/society/" target="_blank" bosszone="dh_1_2">法制</a>
      </li>
            <li class="nav-item">
        <a href="https://v.qq.com/tv/" target="_blank" bosszone="dh_2_2">热剧</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/antip/" target="_blank" bosszone="dh_3_2">抗肺炎</a>
      </li>
            <li class="nav-item">
        <a href="http://new.qq.com/ch/history/" target="_blank" bosszone="dh_4_2">历史</a>
      </li>
            <li class="nav-item">
        <a href="http://sports.qq.com/premierleague/" target="_blank" bosszone="dh_5_2">英超</a>
      </li>
            <li class="nav-item">
        <a href="http://sports.qq.com/cba/" target="_blank" bosszone="dh_6_2">CBA</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch2/star" target="_blank" bosszone="dh_7_2">明星</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/finance_licai/" target="_blank" bosszone="dh_8_2">理财</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/kepu/" target="_blank" bosszone="dh_9_2">科普</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/health/" target="_blank" bosszone="dh_10_2">健康</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/auto/" target="_blank" bosszone="dh_11_2">车型</a>
      </li>
            <li class="nav-item">
        <a href="http://www.jia360.com" target="_blank" bosszone="dh_12_2">家居</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/baby/" target="_blank" bosszone="dh_13_2">育儿</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/emotion/" target="_blank" bosszone="dh_14_2">情感</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/comic/" target="_blank" bosszone="dh_15_2">动漫</a>
      </li>
            <li class="nav-item">
        <a href="http://gongyi.qq.com/" target="_blank" bosszone="dh_16_2">公益</a>
      </li>
            <li class="nav-item">
        <a href="http://tianqi.qq.com/index.htm" target="_blank" bosszone="dh_17_2">天气</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/politics/" target="_blank" bosszone="dh_18_2">政务</a>
      </li>
            <li class="nav-item">
        <a href="https://v.qq.com/channel/variety" target="_blank" bosszone="dh_19_2">综艺</a>
      </li>
            <li class="nav-item">
        <a href="http://news.qq.com/photon/photoex.htm" target="_blank" bosszone="dh_20_2">影展</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/world/" target="_blank" bosszone="dh_21_2">国际</a>
      </li>
            <li class="nav-item">
        <a href="http://sports.qq.com/csocce/csl/" target="_blank" bosszone="dh_22_2">中超</a>
      </li>
            <li class="nav-item">
        <a href="http://fans.sports.qq.com/#/" target="_blank" bosszone="dh_23_2">社区</a>
      </li>
            <li class="nav-item">
        <a href="http://v.qq.com/movie/" target="_blank" bosszone="dh_24_2">电影</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/finance_stock/" target="_blank" bosszone="dh_25_2">证券</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/digi/" target="_blank" bosszone="dh_26_2">数码</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch2/makeup" target="_blank" bosszone="dh_27_2">美容</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/visit/" target="_blank" bosszone="dh_28_2">旅游</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/life/" target="_blank" bosszone="dh_29_2">生活</a>
      </li>
            <li class="nav-item">
        <a href="http://kid.qq.com/" target="_blank" bosszone="dh_30_2">儿童</a>
      </li>
            <li class="nav-item">
        <a href="http://book.qq.com/" target="_blank" bosszone="dh_31_2">文学</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/omv/" target="_blank" bosszone="dh_32_2">享看</a>
      </li>
            <li class="nav-item">
        <a href="https://new.qq.com/ch/cul_ru/" target="_blank" bosszone="dh_33_2">新国风</a>
      </li>
            <li class="nav-item">
        <a href="http://www.qq.com/map/" target="_blank" bosszone="dh_34_2">全部</a>
      </li>
          </ul>
  </div>
</div><!--47161ce01cd51874e1a4c0114e7043d2-->
      </div>
    </div>
    <!-- /导航 -->

    <!-- 广告1 -->
    <div class="layout qq-gg gg-1 cf">
      <div class="col-1 fl">
        <!--NEW_QQCOM_N_Width1_div AD begin...."l=NEW_QQCOM_N_Width1&log=off"--><div id="NEW_QQCOM_N_Width1" style="width:920px;height:75px;" class="l_qq_com"></div><!--NEW_QQCOM_N_Width1 AD end --><!--[if !IE]>|xGv00|b3d18b0084cd79d75eec6cc3e21077fc<![endif]-->
      </div>
      <div class="col-2 fr">
        <div id="gg-focus1" class="gg-focus">
  <ul class="list">
        <li class="item">
              <a href="http://www.shdf.gov.cn/shdf/channels/740.html" target="_blank">
          <img src="//img1.gtimg.com/ninja/2/2020/06/ninja159334882913158.png" alt="中国扫黄打非网举报入口">
        </a>
          </li>
        <li class="item">
              <a href="http://www.qq.com/jubaoxuzhi.htm" target="_blank">
          <img src="//mat1.gtimg.com/pingjs/ext2020/test2017/netwatch.png" alt="网络监控">
        </a>
          </li>
        <li class="item">
              <a href="http://www.piyao.org.cn/yybgt/index.htm" target="_blank">
          <img src="//img1.gtimg.com/ninja/2/2020/11/ninja160551721287830.jpg" alt="网络谣言曝光台">
        </a>
          </li>
        <li class="item">
              <a href="https://110.qq.com/article_detail.html?id=5F1B69AAE7DBFB2D2B90" target="_blank">
          <img src="https://mat1.gtimg.com/rain/apub2019/2b675d7fd05a.wcn.png" alt="未成年人不良内容举报入口">
        </a>
          </li>
        <li class="item">
              <a href="http://www.12377.cn/" target="_blank">
          <img src="//img1.gtimg.com/ninja/2/2018/10/ninja153907291410277.png" alt="网上有害信息举报">
        </a>
          </li>
        <li class="item">
              <a href="https://110.qq.com/" target="_blank">
          <img src="//img1.gtimg.com/ninja/2/2018/10/ninja153907290259802.png" alt="腾讯101">
        </a>
          </li>
      </ul>
  <div class="dot"></div>
</div><!--0c6e9ea2687a67d991eb45a564d3a3e9-->
      </div>
    </div>
    <!-- /广告1 -->

    <div class="layout qq-main cf">
      <div class="col col-1 fl">

        <div id="main-news" class="mod m-news">

          <div class="hd cf">
            <h2 class="tit active fl"><a href="//news.qq.com" target="_blank" bosszone="yw_logo">要闻</a></h2>
            <span class="tit-line fl"></span>
            <h2 class="tit fl"><a href="//new.qq.com/ch/antip/" target="_blank" bosszone="antip_logo">抗肺炎</a></h2>
            <span class="tit-line fl" style="display:none"></span>
            <h2 class="tit fl"></h2>
            <div id="m-weather" class="m-weather f14 fr">
  <a id="weaher-info" href="https://tianqi.qq.com/" target="_blank">
    <div id="ipWeather" class="w-city fl"></div>
    <div id="weatherIcon" class="w-icon fl"></div>
    <div id="weatherTemperature" class="w-du fl"></div>
  </a>

  <div id="weatherMore" class="weather-more">

    <!-- 天气详情 -->
    <div class="face front">
      <div class="weatherMoreTitle cf">
        <div class="weather-info fl">
          <a id="weatherMoreLink" href="https://tianqi.qq.com/" target="_blank">
            <span id="weatherMoreCity"></span>
            <span id="weatherMoreTxt"></span>
            <span id="weatherMoreTem"></span>
          </a>
        </div>
        <div class="weatherMoreSet fr" id="weatherMoreSet">
          <a href="javascript:void(0);">更换城市</a>
        </div>
      </div>
      <div class="weatherMoreAir">
        <a id="weatherMoreAirLink" href="https://tianqi.qq.com/" target="_blank">
          空气质量：<span id="weatherMoreAirTxt" style="padding-right:20px;"></span>
          PM2.5：<span id="weatherMoreAirPmTxt"></span>
        </a>
      </div>
      <a id="weatherMoreFuture" class="weatherMoreFuture cf" href="https://tianqi.qq.com/" target="_blank">
        <div class="weatherMoreFutureCon">
          <div class="weatherMoreIcon w-icon" id="weatherMoreTodayIcon"></div>
          <p><strong>今天</strong></p>
          <p>
            <span id="weatherMoreTodayHighest" class="weatherMoreMath">-</span>℃ -
            <span id="weatherMoreTodayLowest" class="weatherMoreMathLow">-</span>
            <span class="weatherMoreSign">℃</span>
          </p>
        </div>
        <div class="weatherMoreFutureCon">
          <div class="weatherMoreIcon w-icon" id="weatherMoreTomorrowIcon"></div>
          <p><strong>明天</strong></p>
          <p>
            <span id="weatherMoreTomorrowHighest" class="weatherMoreMath">-</span>℃ -
            <span id="weatherMoreTomorrowLowest" class="weatherMoreMathLow">-</span>
            <span class="weatherMoreSign">℃</span></p>
        </div>
        <div class="weatherMoreFutureCon">
          <div class="weatherMoreIcon w-icon" id="weatherMoreAfterTomorrowIcon"></div>
          <p><strong>后天</strong></p>
          <p>
            <span id="weatherMoreAfterTomorrowHighest" class="weatherMoreMath">-</span>℃ -
            <span id="weatherMoreAfterTomorrowLowest" class="weatherMoreMathLow">-</span>
            <span class="weatherMoreSign">℃</span>
          </p>
        </div>
      </a>
    </div>
    <!-- /天气详情 -->

    <!-- 城市设置 -->
    <div class="face back">
      <div class="weatherMoreTitle cf">
        <div class="fl">
          <span>设置城市</span>
        </div>
        <a href="javascript:void(0);" id="weatherMoreReset" class="weatherMoreReset">恢复默认城市</a>
      </div>
      <div class="weatherMoreSelectLayout cf">
        <div class="weatherMoreProviceLayout fl">
          <div class="weatherMoreProviceDefault" id="ipSetProvince">北京市</div>
          <div class="weatherMoreProviceSelect" id="weatherMoreProviceSelect">
            <ul>
              <li><a href="javascript:void(0);">北京市</a></li>
              <li><a href="javascript:void(0);">上海市</a></li>
              <li><a href="javascript:void(0);">天津市</a></li>
              <li><a href="javascript:void(0);">重庆市</a></li>
              <li><a href="javascript:void(0);">河北省</a></li>
              <li><a href="javascript:void(0);">山西省</a></li>
              <li><a href="javascript:void(0);">内蒙古</a></li>
              <li><a href="javascript:void(0);">江苏省</a></li>
              <li><a href="javascript:void(0);">安徽省</a></li>
              <li><a href="javascript:void(0);">山东省</a></li>
              <li><a href="javascript:void(0);">辽宁省</a></li>
              <li><a href="javascript:void(0);">吉林省</a></li>
              <li><a href="javascript:void(0);">黑龙江省</a></li>
              <li><a href="javascript:void(0);">浙江省</a></li>
              <li><a href="javascript:void(0);">江西省</a></li>
              <li><a href="javascript:void(0);">福建省</a></li>
              <li><a href="javascript:void(0);">湖北省</a></li>
              <li><a href="javascript:void(0);">湖南省</a></li>
              <li><a href="javascript:void(0);">河南省</a></li>
              <li><a href="javascript:void(0);">广东省</a></li>
              <li><a href="javascript:void(0);">广西</a></li>
              <li><a href="javascript:void(0);">海南省</a></li>
              <li><a href="javascript:void(0);">四川省</a></li>
              <li><a href="javascript:void(0);">贵州省</a></li>
              <li><a href="javascript:void(0);">云南省</a></li>
              <li><a href="javascript:void(0);">西藏</a></li>
              <li><a href="javascript:void(0);">陕西省</a></li>
              <li><a href="javascript:void(0);">甘肃省</a></li>
              <li><a href="javascript:void(0);">宁夏</a></li>
              <li><a href="javascript:void(0);">青海省</a></li>
              <li><a href="javascript:void(0);">新疆</a></li>
              <li><a href="javascript:void(0);">香港</a></li>
              <li><a href="javascript:void(0);">澳门</a></li>
              <li><a href="javascript:void(0);">台湾省</a></li>
            </ul>
          </div>
        </div>
        <div class="weatherMoreCityLayout fl">
          <div class="weatherMoreCityDefault" id="ipSetCity">北京市</div>
          <div class="weatherMoreCitySelect" id="weatherMoreCitySelect">
            <ul id="weatherMoreCitySelectUl">
              <li><a href="javascript:void(0);">北京市</a></li>
            </ul>
          </div>
        </div>
      </div>
      <div class="weatherMoreNews">
        <div id="weatherMoreNewsCheckbox" class="weatherMoreNewsCheckbox weatherMoreNewsYes" style="display:none">同时更新资讯所属地</div>
      </div>
      <div class="weatherMoreBtn">
        <input type="button" value="确定" id="weatherMoreSubmit" class="weatherMoreSubmit" />
        <input type="button" value="取消" id="weatherMoreCancel" class="weatherMoreCancel" />
      </div>
    </div>
    <!-- /城市设置 -->

  </div>
</div><!--22387c29e2e6565bc4936f977fba8cfc--><!--[if !IE]>|xGv00|e8fbb0435570dc4bd2db993b13cd0260<![endif]-->
          </div>
          <div class="bd">

            <!-- 要闻 -->
            <div id="tab-news-01" class="tab-news" bossexpo="bg_yw">
              <style>
.bgcolor1 {
    background: #f56300;
    color: #FFF;
    padding: 0 3px;
    border-radius: 3px;
}

.bgcolor1:hover {
    color: #FFF;
}
li.news-top-first {
  height:36px;
  text-indent: -2222px;
}
</style>
<ul class="yw-list" bosszone="yw_1">
<li class="news-top ">
  <a class=" bold" href="https://new.qq.com/omn/20210803/20210803A02WYB00.html" target="_blank" newsexpo="yw1_1" style="color:black">大党丨热血忠魂</a>
</li>
<li class=" ">
    <a style="color:black" class=" bold" href="https://new.qq.com/omn/20210803/20210803A00SBH00.html" target="_blank" data-icon="no-icon" newsexpo="yw2_1">迈好第一步 见到新气象</a>    <a style="color:black" class=" bold news_color_" href="https://new.qq.com/omn/AUS20210/AUS2021080300737500.html" target="_blank" data-icon="no-icon" newsexpo="yw2_2">60秒，看中国军队的强大能力</a>  </li>
<li class=" ">
    <a style="color:black" class=" bold" href="https://new.qq.com/omn/20210803/20210803A01A8G00.html" target="_blank" data-icon="no-icon" newsexpo="yw3_1">美国施压中国不可能成功</a>    <a style="color:black" class=" bold" href="https://new.qq.com/omn/20210803/20210803A05O0700.html" target="_blank" data-icon="no-icon" newsexpo="yw3_2">特工能搞病毒溯源要科学家干嘛</a>  </li>
<li class=" ">
  <a class=" news_color_2" href="https://view.inews.qq.com/a/TWF2021072400913800" target="_blank" newsexpo="yw4_1" style="color:black">坚决反对将新冠病毒溯源问题政治化</a>
</li>
<li class=" ">
  <a class="" href="https://new.qq.com/omn/20210730/20210730A03J1Z00.html" target="_blank" newsexpo="yw5_1" style="color:black">环球漫评：“五个一百”，在网络空间抒写奋进赞歌</a>
</li>
<li class=" ">
  <a class="" href="https://new.qq.com/omn/20210730/20210730A0E14B00.html" target="_blank" newsexpo="yw6_1" style="color:black">评论：假借奥运刺激民粹，民进党政客何其亢奋</a>
</li>
<li class=" ">
  <a class="" href="https://new.qq.com/omn/20210730/20210730A06I7F00.html" target="_blank" newsexpo="yw7_1" style="color:black">泪目！武警官兵转移后留下这些黑板报</a>
</li>
<li class=" ">
  <a class="" href="https://new.qq.com/omn/SJD20210/SJD2021080200560100.html" target="_blank" newsexpo="yw8_1" style="color:black">腾讯新闻进一步落实“清朗·‘饭圈’乱象整治”专项公告</a>
</li>
</ul><!--94c3a4b60d86e79103fe5f818a69aeec-->
              <style type="text/css">
.news_color_3{color:#0c82ff!important;}
.news_color_4{color:#df5147!important;}
</style>

<ul class="yw-list" bosszone="yw_2">
          <li class="news-pic-txt cf">
      <div class="pic fl">
        <a href="https://new.qq.com/omn/20210802/20210802V0EKXC00.html" target="_blank" newsexpo="yw9_1">
          <img src="//img1.gtimg.com/ninja/2/2021/08/ninja162794942637175.jpg" alt="江苏开启防疫“飞行模式” 无人机带你扫行程码">
        </a>
      </div>
      <div class="txt fl">
        <a href="https://new.qq.com/omn/20210802/20210802V0EKXC00.html" target="_blank" newsexpo="yw9_1">江苏开启防疫“飞行模式” 无人机带你扫行程码</a>
        <div class="info">
          <a href="https://new.qq.com/omn/20210802/20210802V0EKXC00.html" target="_blank">
            
          </a>
        </div>
      </div>
    </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210803/20210803A01Q4R00.html" target="_blank" newsexpo="yw10_1">江苏：南京聚集性疫情发生的根本原因是思想上的松懈</a>
          </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210802/20210802A0EJ6800.html" target="_blank" newsexpo="yw11_1">“天下粮仓”河南1450万亩农作物受灾，粮价会大涨吗？</a>
          </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210802/20210802A07YUJ00.html" target="_blank" newsexpo="yw12_1">外交部驻港公署：香港司法机关依法办案不容置喙</a>
          </li>
        <li>
              <a class="q-icons icon-video" href="https://new.qq.com/omn/20210802/20210802A0EE9I00.html" target="_blank" newsexpo="yw13_1">新华社记者“卧底”奈雪的茶：蟑螂乱爬 标签随意换 </a>
          </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210802/20210802A0FJYS00.html" target="_blank" newsexpo="yw14_1">常德游船和扬州棋牌室如何成为疫情传播重点场所？</a>
          </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210802/20210802A0DRO700.html" target="_blank" newsexpo="yw15_1">河南郑州遇难人数近300，公众需要一个答案</a>
          </li>
                    </ul><ul class="yw-list" bosszone="yw_3">
        <li class="news-pic-txt cf">
      <div class="pic fl">
        <a href="https://new.qq.com/omn/20210802/20210802A0AO3H00.html" target="_blank" newsexpo="yw16_1">
          <img src="//img1.gtimg.com/ninja/2/2021/08/ninja162794944331129.jpg" alt="再见，郎平指导，你辛苦了">
        </a>
      </div>
      <div class="txt fl">
        <a href="https://new.qq.com/omn/20210802/20210802A0AO3H00.html" target="_blank" newsexpo="yw16_1">再见，郎平指导，你辛苦了</a>
        <div class="info">
          <a href="https://new.qq.com/omn/20210802/20210802A0AO3H00.html" target="_blank">
            
          </a>
        </div>
      </div>
    </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210802/20210802A0DFPP00.html" target="_blank" newsexpo="yw17_1">大V涉嫌盗伐林木罪被刑拘 曾是牡丹江“曹园”事件举报人</a>
          </li>
        <li>
              <a class="q-icons icon-video news_color_2" href="https://new.qq.com/omn/20210803/20210803V01GNI00.html" target="_blank" newsexpo="yw18_1">饭圈乱象专项整治行动：已处置4000多个违规账号</a>
          </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210802/20210802A0EHMG00.html" target="_blank" newsexpo="yw19_1">吴亦凡被刑拘后，五大协会、多家央媒相继发声</a>
          </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210802/20210802V0F7AZ00.html" target="_blank" newsexpo="yw20_1">男子回国擅自脱离隔离点逃至临汾与朋友上网 被警方抓获</a>
          </li>
        <li>
              <a class="" href="https://new.qq.com/omn/20210803/20210803V01JST00.html" target="_blank" newsexpo="yw21_1">超强阵风突袭南宁：路边大树被连根拔起，市民寸步难行</a>
          </li>
  </ul><!--7bb85a3ce8057ac3d64cab4bab31a6bf-->
            </div>
            <!-- /要闻 -->

            <!-- 抗肺炎 -->
            <div id="tab-news-02" class="tab-news" bossexpo="antip_yw" style="display:none;">
              <div id="scaleContainer" style="display:none;" data-src="https://mat1.gtimg.com/rain/apub2019/ddea48edb444.qq_top1x1.svg">
                <a id="antip_charts" style="text-decoration: none;" href="https://news.qq.com//zt2020/page/feiyan.htm" target="_blank" bosszone="antip_click" >
                  <div class="head"></div>
                  <div class="topdataWrap">
                    <div class="antip_timeNum">
                      <div class="bottom">
                        <p class="d"></p>
                      </div>
                    </div>
                    <div class="antip_recentNumber">
                      <div class="antip_icbar confirm">
                        <div class="add">较上日<span>+</span></div>
                        <div class="number">0</div>
                        <div class="text">全国确诊</div>
                      </div>
                      <div class="antip_icbar suspect">
                        <div class="add">较上日<span>+</span></div>
                        <div class="number">0</div>
                        <div class="text">疑似病例</div>
                      </div>
                      <div class="antip_icbar cure">
                        <div class="add">较上日<span>+</span></div>
                        <div class="number">0</div>
                        <div class="text">治愈人数</div>
                      </div>
                      <div class="antip_icbar dead">
                        <div class="add">较上日<span>+</span></div>
                        <div class="number">0</div>
                        <div class="text">死亡人数</div>
                      </div>
                    </div>
                  </div>
                </a>
              </div>
              <ul class="yw-list" bosszone="antip_1">
        <li class="news-top"><a href="https://new.qq.com/omn/20210803/20210803A04DU800.html" target="_blank">国内疫情多地散发，疫苗预防效果如何？张文宏最新分析</a></li>
            <li><a href="https://new.qq.com/omn/20210803/20210803A00ZD700.html" target="_blank">当老年生活遭遇新冠：扬州棋牌室成疫情传播场</a></li>
            <li><a href="https://new.qq.com/omn/20210803/20210803A043DJ00.html" target="_blank">张文宏：未完成接种及接种后未满14天的这些人不适宜出游</a></li>
            <li><a href="https://new.qq.com/omn/20210802/20210802A04WA700.html" target="_blank">国产疫苗对德尔塔毒株有用吗？钟南山给出明确答案</a></li>
            <li><a href="https://new.qq.com/omn/20210803/20210803A04GWV00.html" target="_blank">湖南益阳新增2例无症状感染者，系北京确诊病例舅舅、大伯父</a></li>
            <li><a href="https://new.qq.com/omn/20210803/20210803A01CRE00.html" target="_blank">江苏：所有单位都要配合属地防控工作，确保行政区内无盲区</a></li>
            <li><a href="https://new.qq.com/omn/20210803/20210803A02H3R00.html" target="_blank">最新！又有四地升级，全国现有高中风险区4＋118个</a></li>
            <li><a href="https://new.qq.com/omn/20210803/20210803A01C5300.html" target="_blank">武汉紧急行动：争分夺秒斩断疫情传播</a></li>
                    </ul><ul class="yw-list" bosszone="antip_2">
    <li class="news-pic-txt cf">
      <div class="pic fl">
        <a href="https://new.qq.com/omn/20210803/20210803A05F5C00.html" target="_blank">
                    <img src="//inews.gtimg.com/newsapp_ls/0/13842993997_640330/0" alt="3例确诊、4例无症状，武汉全员核酸检测">
                  </a>
      </div>
      <div class="txt fl">
        <a href="https://new.qq.com/omn/20210803/20210803A05F5C00.html" target="_blank">3例确诊、4例无症状，武汉全员核酸检测</a>
        <div class="info">
          <a href="https://new.qq.com/omn/20210803/20210803A05F5C00.html" target="_blank">
                      </a>
        </div>
      </div>
    </li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A01DVB00.html" target="_blank">四川疾控发布健康提示：这些人员实行14天集中隔离医学观察</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A02XY400.html" target="_blank">上海新增确诊病例为驾驶员，密接52人首次核酸检测阴性</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A01D5J00.html" target="_blank">扬州封控小区只留一个出入口 每户每天一人进出</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A025G000.html" target="_blank">山东烟台新增2例本地确诊：一例经停南京，一例去过扬州</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A01KXM00.html" target="_blank">云南省昨日新增本土新冠肺炎确诊病例2例</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A0262200.html" target="_blank">31省份8月2日新增本土确诊61例，涉及8省份</a></li>
                          </ul><ul class="yw-list" bosszone="antip_3">
    <li class="news-pic-txt cf">
      <div class="pic fl">
        <a href="https://new.qq.com/omn/20210803/20210803A01CUP00.html" target="_blank">
                    <img src="//inews.gtimg.com/newsapp_ls/0/13842113000_640330/0" alt="上海浦东机场连夜展开大规模核酸检测">
                  </a>
      </div>
      <div class="txt fl">
        <a href="https://new.qq.com/omn/20210803/20210803A01CUP00.html" target="_blank">上海浦东机场连夜展开大规模核酸检测</a>
        <div class="info">
          <a href="https://new.qq.com/omn/20210803/20210803A01CUP00.html" target="_blank">
                      </a>
        </div>
      </div>
    </li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A026ZY00.html" target="_blank">哈尔滨机场强化疫情防控 守牢“空中防线”</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A02BLB00.html" target="_blank">北京市卫健委提示市民最大限度减少出京</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A01CVU00.html" target="_blank">海口确诊病例感染毒株初步判定为德尔塔变异株</a></li>
                  <li><a href="https://new.qq.com/omn/20210802/20210802A0EY2D00.html" target="_blank">北京2例确诊1例无症状在张家界轨迹公布，曾观看两场演出</a></li>
                  <li><a href="https://new.qq.com/omn/20210803/20210803A01DMG00.html" target="_blank">武汉新增7例省外关联本地新冠肺炎病例 详情公布</a></li>
    </ul>
            </div>
            <!-- /抗肺炎 -->

            <!-- 地方新闻 -->
            <div id="tab-news-03" class="tab-news undis" bossexpo="bg_dfz">
              <ul class="yw-list" bosszone="dfz_1">
        <li class="news-top"><a href="https://new.qq.com/omn/20200701/20200701A04H7500" target="_blank">北京昨日新增报告3例确诊病例 均在大兴区</a></li>
          <li><a href="https://new.qq.com/omn/20200701/20200701A03LEM00" target="_blank">今明两天北京雷雨频繁 外出需注意防雷避雨</a></li>
          <li><a href="https://new.qq.com/omn/20200701/20200630A0SSK000" target="_blank">新发地周边12个封闭管控小区6月30日起依规解封</a></li>
          <li><a href="https://new.qq.com/omn/20200701/20200630A0SEID00" target="_blank">张文宏：北京疫情只是小范围反弹，中国拒绝第二波疫情</a></li>
          <li><a href="https://new.qq.com/omn/20200701/BJC2020063000998300" target="_blank">北京发布病例详情 多名隔离人员发病不报告</a></li>
          <li><a href="https://new.qq.com/omn/20200701/BJC2020063000799200" target="_blank">北京多人隔离14天后确诊，专家称有两方面原因</a></li>
          <li><a href="https://new.qq.com/omn/20200701/20200630A0A2VC00" target="_blank">北京市银行停业一周？五大行辟谣：仅个别风险区网点暂停</a></li>
          <li><a href="https://new.qq.com/omn/20200701/20200630A0PE5200" target="_blank">北京6月30日有3地疫情风险等级降级</a></li>
                  </ul><ul class="yw-list" bosszone="dfz_2">
    <li class="news-pic-txt cf">
      <div class="pic fl">
        <a href="https://new.qq.com/omn/20200701/20200630A0PPY900" target="_blank">
          <img src="//inews.gtimg.com/newsapp_ls/0/12013918816_640330/0" alt="看“病毒侦探”如何工作：透视北京疫情流调三大焦点">
        </a>
      </div>
      <div class="txt fl">
        <a href="https://new.qq.com/omn/20200701/20200630A0PPY900" target="_blank">看“病毒侦探”如何工作：透视北京疫情流调三大焦点</a>
        <div class="info">
          <a href="https://new.qq.com/omn/20200701/20200630A0PPY900" target="_blank">
            北京日报客户端
          </a>
        </div>
      </div>
    </li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0JTXV00" target="_blank">北京：已经出院的新冠肺炎患者 未发现人传人现象</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0JMSN00" target="_blank">北京：此次疫情重症和危重症患者比例明显偏低</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0JVY300" target="_blank">北京：二级以上医疗机构非急诊全面预约实行常态化机制</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0OLGM00" target="_blank">7月1日起，北京公积金账户余额可直接用来还贷款了！</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0URPY00" target="_blank">北京57家公立医疗机构核酸检测预约电话公布</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0M75V00" target="_blank">北京近4日有37例确诊病例来自集中隔离点</a></li>
                        </ul><ul class="yw-list" bosszone="dfz_3">
    <li class="news-pic-txt cf">
      <div class="pic fl">
        <a href="https://new.qq.com/omn/20200701/20200701A03PSY00" target="_blank">
          <img src="//inews.gtimg.com/newsapp_ls/0/12016212561_640330/0" alt="“圈定”确诊送餐员活动轨迹，他有这些妙招">
        </a>
      </div>
      <div class="txt fl">
        <a href="https://new.qq.com/omn/20200701/20200701A03PSY00" target="_blank">“圈定”确诊送餐员活动轨迹，他有这些妙招</a>
        <div class="info">
          <a href="https://new.qq.com/omn/20200701/20200701A03PSY00" target="_blank">
            北京日报客户端
          </a>
        </div>
      </div>
    </li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0IXXO00" target="_blank">新发地市场一个体经营人员先被诊断为疑似后确诊</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0OF6200" target="_blank">朝阳一诊所因擅自接诊发热患者被停业整顿 当事人被行拘</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A06N0K00" target="_blank">顺义累计采集30余万份样本，结果均为阴性</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A06TMZ00" target="_blank">大兴黄村约谈职能部门，加强企业防疫监管力度</a></li>
                <li><a href="https://new.qq.com/omn/20200701/20200630A0I3XN00" target="_blank">女子与男友吵架 深夜往楼下扔菜刀被控制</a></li>
  </ul><!--991636fbe8c375d613bd50eeee5dadad--><!--[if !IE]>|xGv00|e15d36cd2e0736bada6acb6250ab238f<![endif]-->
            </div>
            <!-- /地方新闻 -->

          </div>
        </div>
      </div>
      <div class="col col-2 fl">

        <!-- 今日话题 -->
        <div class="mod m-topic" bossexpo="bg_jrht">
  <div class="hd cf">
    <h2 class="tit active"><a bosszone="jrht_logo">今日话题</a></h2>
  </div>
  <div class="bd">
    <ul class="news-list">
                  <li class="news-top" bosszone="jrht_1">
          <a href="https://new.qq.com/rain/a/20210803A01OJI00" target="_blank">太空电站：解决人类能源危机的终极出路？</a>
        </li>
                        <li bosszone="jrht_2">
                                                    <a class="cate" href="https://new.qq.com/omn/author/10859191" target="_blank">新华社</a><span class="line">|</span>
                                                        <a href="https://new.qq.com/rain/a/20210803A04BAF00" target="_blank">为什么中国举重选手总是最后才出场？</a>
                                            </li>
                        <li bosszone="jrht_3">
                                                    <a class="cate" href="https://new.qq.com/omn/author/1122" target="_blank">中青报</a><span class="line">|</span>
                                                        <a href="https://new.qq.com/rain/a/20210803A041YE00" target="_blank">90后轮椅女孩杨淑亭：8年帮助1386名村民脱贫致富</a>
                                            </li>
                        <li bosszone="jrht_4">
                                                    <a class="cate" href="https://new.qq.com/omn/author/1795" target="_blank">半月谈</a><span class="line">|</span>
                                                        <a href="https://new.qq.com/rain/a/20210803A04C2W00" target="_blank">告别之战，大家都哭了！郎导，您辛苦了，谢谢！</a>
                                            </li>
                        <li bosszone="jrht_5">
                                                    <a class="cate" href="https://new.qq.com/omn/author/5007264" target="_blank">澎湃</a><span class="line">|</span>
                                                        <a href="https://new.qq.com/omn/20210802/20210802A04MV600.html" target="_blank">亲历苏炳添“封神”一夜，我的双手忍不住颤抖</a>
                                            </li>
                        <li bosszone="jrht_6">
                                                    <a class="cate" href="https://new.qq.com/omn/author/67" target="_blank">36氪</a><span class="line">|</span>
                                                        <a href="https://new.qq.com/rain/a/20210803A04EX900" target="_blank">亏损上万亿的东京奥运会，更让人怀念起08年的北京</a>
                                            </li>
                        <li bosszone="jrht_7">
                                                    <a class="cate" href="https://new.qq.com/omn/author/5107513" target="_blank">较真</a><span class="line">|</span>
                                                        <a href="https://new.qq.com/rain/a/20210802A03U3000" target="_blank">火锅店这几种荤菜千万别点，员工自己都不吃？</a>
                                            </li>
              </ul>
  </div>
</div><!--bc19a7565b160f29639e8a28d56d42d8-->
        <!-- /今日话题 -->

        <!-- 原创 十三邀 -->
        <div class="mod m-yao13" bossexpo="bg_ycsp">
  <div class="hd-2 cf" style="height:50px;">
    <h2 class="tit active">
                        <a href="http://v.qq.com/detail/8/89803.html" target="_blank" bosszone="ycsp_logo">
            <img src="https://mat1.gtimg.com/qqcdn/tupload/1627485253440.png" alt="仅三天">
          </a>
                                                                                                      </h2>
  </div>
  <div class="bd" style="margin-top:-14px;">
    <ul class="news-list">
                                        <li class="news-pic-txt cf" bosszone="ycsp_2">
            <div class="pic video-box click-drag-play fl" bossvv="vv_ycsp">
              <img src="https://mat1.gtimg.com/qqcdn/tupload/1627569350062.jpeg" alt="KTV嗨唱露营地谈心 毛不易陈飞宇灵魂互问：你幸福吗">
              <i class="q-icons icon-play2"></i>
              <div class="txt undis">KTV嗨唱露营地谈心 毛不易陈飞宇灵魂互问：你幸福吗</div>
              <div class="desc undis">y32614qzi15</div>
              <div id="mod-player4" class="mod-player" data-vid="y32614qzi15" data-url="https://v.qq.com/x/cover/mzc00200xurquft/u3262japp02.html" style="display: none;"></div>
              <div class="click-layer"></div>
            </div>
            <div class="txt fl">
              <a href="https://v.qq.com/x/cover/mzc00200xurquft/u3262japp02.html" target="_blank">KTV嗨唱露营地谈心 毛不易陈飞宇灵魂互问：你幸福吗</a>
              <div class="info">

              </div>
            </div>
          </li>
                                          <li bosszone="ycsp_3">
                                        <a href="https://v.qq.com/x/cover/mzc002008rq1ry2/w3261m3xvhr.html" target="_blank">[看点]陈飞宇回忆九岁初次演戏 感慨演员在于塑造生命</a>
                      </li>
                                          <li bosszone="ycsp_4">
                                        <a href="https://v.qq.com/x/cover/mzc00200xurquft/m3261hux6x0.html" target="_blank">[看点]活久见！毛不易陈飞宇合唱大鱼 KTV飙歌停不下来</a>
                      </li>
                                          <li bosszone="ycsp_5">
                                                                            <a class="cate q-icons icon-video" href="https://v.qq.com/detail/8/89877.html" target="_blank">向上之路</a><span class="line">|</span>
                                                                <a href="https://v.qq.com/x/cover/mzc00200lixouil/x32644d64yi.html" target="_blank">探访神秘内蒙“轻骑兵”： 一年跑十万公里</a>
                                                    </li>
                                          <li bosszone="ycsp_6">
                                                                            <a class="cate q-icons icon-video" href="https://v.qq.com/detail/8/89736.html" target="_blank">大美中国路</a><span class="line">|</span>
                                                                <a href="https://v.qq.com/x/cover/mzc00200j8vu7b6/b32658t2840.html" target="_blank">他用一箱油跑1952公里夺冠 如今成汽修王牌</a>
                                                    </li>
                                          <li bosszone="ycsp_7">
                                                                            <a class="cate q-icons icon-video" href="https://v.qq.com/detail/8/89784.html" target="_blank">音乐会</a><span class="line">|</span>
                                                                <a href="https://v.qq.com/x/cover/mzc00200skqqn0n/q326344enxx.html" target="_blank">周笔畅为北漂女孩单独献唱 女孩当场泪崩</a>
                                                    </li>
                      </ul>
  </div>
</div><!--a413775ace5a38270460235cfebc722b-->
        <!-- /原创 十三邀 -->

        <!-- 图话 -->
        <div class="mod m-picture" bossexpo="bg_th">
          <div class="hd-2 cf">
            <h2 class="tit active">
              <a href="https://new.qq.com/ch/photo" target="_blank" bosszone="th_logo">图话</a>
            </h2>
          </div>
          <div class="bd">
            <ul class="news-list">
                    <li class="v-item news-pic-txt cf" bosszone="th_1">
      <div class="pic fl">
        <a href="https://new.qq.com/rain/a/20210730A0GOMZ00" target="_blank">
          <img src="//inews.gtimg.com/newsapp_ls/0/13832527180_640330/0" alt="大撤离中的河南村庄：失散、相聚、重新开始">
        </a>
      </div>
      <div class="txt fl">
        <a href="https://new.qq.com/rain/a/20210730A0GOMZ00" target="_blank">大撤离中的河南村庄：失散、相聚、重新开始</a>
        <div class="info">
          <a href="https://new.qq.com/rain/a/20210730A0GOMZ00" target="_blank">
            中国人的一天 第3946期
          </a>
        </div>
      </div>
    </li>
        <li class="v-item" bosszone="th_2">
                                                <a class="cate q-icons icon-pic" href="https://new.qq.com/omn/author/15329577" target="_blank">萤火计划</a><span class="line">|</span>
                                        <a href="https://new.qq.com/rain/a/20210728A0DQ5J00" target="_blank">焊工、保安、退伍军人，冲上水灾救援一线</a>
                            </li>
  <!--17a3a252b8e0c89b03088b1793f44454-->
                    <li class="v-item" bosszone="th_1">
                                              <a  class="cate q-icons icon-pic" href="https://new.qq.com/omn/author/5505476" target="_blank">谷雨</a><span class="line">|</span>
                                        <a href="https://v.qq.com/x/cover/mzc00200j8vu7b6/b32658t2840.html" target="_blank">他用一箱油跑1952公里夺冠，如今成汽修王牌</a>
                            </li>
  <!--60ac9fe8085541f2b3dd8451268b3eeb-->
                    <li class="v-item" bosszone="th_1">
                                              <a  class="cate q-icons icon-pic" href="http://sports.qq.com/photo/" target="_blank">体坛</a><span class="line">|</span>
                                        <a href="https://fans.sports.qq.com/post.htm?id=1706392292454039783&mid=264#1_allWithElite" target="_blank">【奥运女神】19岁因一张照片惊艳全球！ </a>
                            </li>
  <!--b32ed4e649ba32f8c6a775359d4bd64d-->
                    <li class="v-item" bosszone="th_1">
                                              <a  class="cate q-icons icon-pic" href="https://new.qq.com/ch/ent/" target="_blank">娱乐</a><span class="line">|</span>
                                        <a href="https://new.qq.com/rain/a/20210730A03IJX00" target="_blank">专访王珞丹：期待有一天能演“精致老女人” </a>
                            </li>
  <!--1fb6eabb19a4bee106da41b9f802a242-->
                    <li class="v-item" bosszone="th_1">
                                              <a  class="cate q-icons icon-pic" href="https://new.qq.com/ch/fashion/" target="_blank">时尚圈</a><span class="line">|</span>
                                        <a href="https://new.qq.com/omn/20210730/20210730A08CFE00.html" target="_blank">何为当代摩登家庭？</a>
                            </li>
  <!--a83095ad8c4666126469495cdda603ed-->
            </ul>
          </div>
        </div>
        <!-- /图话 -->

      </div>
      <div class="col col-3 fr">

        <!-- 产品 -->
        <div id="m-product" class="m-product">
  <ul class="list f14">
                                                                                <li class="q-icons prod-1">
                                                <a href="http://news.qq.com/mobile/"  target="_blank" bosszone="cp_1_1_1">新闻APP</a>
                                  <a href="http://sports.qq.com/kbsweb/"  target="_blank" bosszone="cp_1_1_2">体育APP</a>
                                  <a href="https://om.qq.com/userAuth/index"  target="_blank" bosszone="cp_1_1_3">企鹅号</a>
                                  <a href="http://kuaibao.qq.com/"  target="_blank" bosszone="cp_1_1_4">快报</a>
                                  <a href="http://v.qq.com/download.html#pc"  target="_blank" bosszone="cp_1_1_5">视频</a>
                                  <a href="https://browser.qq.com/"  target="_blank" bosszone="cp_1_1_6">浏览器</a>
                                  <a href="http://www.weishi.com/"  target="_blank" bosszone="cp_1_1_7">微视</a>
                                        </li>
                                <li class="q-icons prod-2">
                                                <a href="http://weixin.qq.com/"  target="_blank" bosszone="cp_1_2_1">微信</a>
                                  <a href="https://im.qq.com/index.shtml"  target="_blank" bosszone="cp_1_2_2">QQ</a>
                                  <a href="https://qzone.qq.com/"  target="_blank" bosszone="cp_1_2_3">空间</a>
                                  <a href="https://work.weixin.qq.com/wework_admin/register_wx?from=regopt_tlogo_wxcbar_tengxunwang"  target="_blank" bosszone="cp_1_2_4">企业微信</a>
                                  <a href="https://mail.qq.com/cgi-bin/loginpage"  target="_blank" bosszone="cp_1_2_5">邮箱</a>
                                  <a href="https://cloud.tencent.com/?fromSource=gwzcw.756432.756432.756432"  target="_blank" bosszone="cp_1_2_6">腾讯云</a>
                                  <a href="https://guanjia.qq.com/?ADTAG=news.qqcom"  target="_blank" bosszone="cp_1_2_7">电脑管家</a>
                                  <a href="https://vip.qq.com/"  target="_blank" bosszone="cp_1_2_8">会员</a>
                                        </li>
                                <li class="q-icons prod-3">
                                                <a href="http://lol.qq.com/index.shtml?ADTAG=media.innerenter.qqcom.index_navigation"  target="_blank" bosszone="cp_1_3_1">LOL</a>
                                  <a href="http://dnf.qq.com/?ADTAG=media.innerenter.qqcom.index_navigation"  target="_blank" bosszone="cp_1_3_2">DNF</a>
                                  <a href="http://cf.qq.com/?ADTAG=media.innerenter.qqcom.index_navigation"  target="_blank" bosszone="cp_1_3_3">CF</a>
                                  <a href="http://pvp.qq.com/?ADTAG=media.innerenter.qqcom.index_navigation"  target="_blank" bosszone="cp_1_3_4">王者</a>
                                  <a href="https://gouhuo.qq.com/?ADTAG=QQHOME"  target="_blank" bosszone="cp_1_3_5">单机游戏</a>
                                  <a href="http://huoying.qq.com/act/a20141009landingpage/index.htm?via=45&ADTAG=ied.neiguang&ADTAG=media.innerenter.qqcom.index_navigation"  target="_blank" bosszone="cp_1_3_6">火影OL</a>
                                  <a href="http://wuxia.qq.com/?ADTAG=media.innerenter.qqcom.index_navigation"  target="_blank" bosszone="cp_1_3_7">天刀</a>
                                  <a href="http://iwan.qq.com/index.htm?ADTAG=media.innerenter.qqcom.indexnavigation"  target="_blank" bosszone="cp_1_3_8">爱玩</a>
                                  <a href="http://nz.qq.com/main.shtml?ADTAG=media.innerenter.qqcom.index_navigation"  target="_blank" bosszone="cp_1_3_9">逆战</a>
                                        </li>
                                <li class="q-icons prod-4">
                                                <a href="https://pc.qq.com/"  target="_blank" bosszone="cp_1_4_1">软件</a>
                                  <a href="https://pay.qq.com/"  target="_blank" bosszone="cp_1_4_2">Q币</a>
                                  <a href="https://www.jd.com/?utm_source=qq.com&utm_medium=cpc&utm_campaign=dmp_77&utm_term=dmp_77_11727_d604816f27c2b5e98ae51fd59de8b1c43abfdac_1538472240"  target="_blank" bosszone="cp_1_4_3">京东</a>
                                  <a href="https://map.qq.com/#city=%D6%D0%B9%FA&wd=%D6%D0%B9%FA"  target="_blank" bosszone="cp_1_4_4">腾讯地图</a>
                                  <a href="https://docs.qq.com/"  target="_blank" bosszone="cp_1_4_5">腾讯文档</a>
                                  <a href="https://qian.qq.com/?stat_data=oth87ppcsy00222&ADTAG=SCQD.BD.PC.TXDH1"  target="_blank" bosszone="cp_1_4_6">理财通</a>
                                  <a href="http://www.qq.com/map/" class="more" target="_blank" bosszone="cp_1_4_7">全部</a>
                                        </li>
                </ul>
  <div id="prod-more" class="prod-more">
    <div class="prod-more-btn">
      <div class="q-icons btn-icon">展开</div>
    </div>
    <ul class="list f14">
                        <li class="prod-1">
                                                <a href="https://new.qq.com/omv" target="_blank" bosszone="cp_2_1_6">享看</a>
                                  <a href="http://qq.pinyin.cn/" target="_blank" bosszone="cp_2_1_5">QQ拼音</a>
                                  <a href="http://player.qq.com/" target="_blank" bosszone="cp_2_1_4">QQ影音</a>
                                  <a href="https://pc.qq.com/detail/15/detail_755.html" target="_blank" bosszone="cp_2_1_3">QQ影像</a>
                                  <a href="http://www.weiyun.com/index.html" target="_blank" bosszone="cp_2_1_2">微云</a>
                                  <a href="https://fm.qq.com/" target="_blank" bosszone="cp_2_1_1">企鹅FM</a>
                                        </li>
                                <li class="prod-2">
                                                <a href="http://book.qq.com/?g_f=70085" target="_blank" bosszone="cp_2_2_6">QQ阅读</a>
                                  <a href="https://y.qq.com/" target="_blank" bosszone="cp_2_2_5">QQ音乐</a>
                                  <a href="http://kg.qq.com/" target="_blank" bosszone="cp_2_2_4">全民K歌</a>
                                  <a href="http://z.qzone.com/" target="_blank" bosszone="cp_2_2_3">手机空间</a>
                                  <a href="https://im.qq.com/mobileqq/" target="_blank" bosszone="cp_2_2_2">手机QQ</a>
                                  <a href="https://jiasu.qq.com/?ADTAG=qqcom" target="_blank" bosszone="cp_2_2_1">网游加速器</a>
                                        </li>
                                <li class="prod-3">
                                                <a href="http://speed.qq.com/main.shtml?ADTAG=media.innerenter.qqcom.index_navigation" target="_blank" bosszone="cp_2_3_7">QQ飞车</a>
                                  <a href="http://yxwd.qq.com/?ADTAG=media.innerenter.qqcom.index_navigation" target="_blank" bosszone="cp_2_3_6">英雄</a>
                                  <a href="http://dn.qq.com/cp/a20180904ysjj/index.htm" target="_blank" bosszone="cp_2_3_5">龙之谷</a>
                                  <a href="http://eafifa.qq.com/?ADTAG=media.innerenter.qqcom.index_navigation" target="_blank" bosszone="cp_2_3_4">FIFA</a>
                                  <a href="http://hdl.qq.com/index.shtml?ADTAG=media.innerenter.qqcom.index_navigation" target="_blank" bosszone="cp_2_3_3">魂斗罗</a>
                                  <a href="http://cfm.qq.com/cp/a20180927vacation/index.html" target="_blank" bosszone="cp_2_3_2">CF手游</a>
                                  <a href="http://tlbb.qq.com/main.shtml?ADTAG=media.innerenter.qqcom.index_navigation" target="_blank" bosszone="cp_2_3_1">天龙手游</a>
                                        </li>
                                <li class="prod-4">
                                                <a href="http://xing.qq.com/" target="_blank" bosszone="cp_2_4_7">星钻</a>
                                  <a href="https://888.qq.com/?bc_tag=10161.1.1" target="_blank" bosszone="cp_2_4_6">QQ彩票</a>
                                  <a href="http://cb.qq.com/?attach=200_1000_10090&QQ_from=200_1000_10090" target="_blank" bosszone="cp_2_4_5">彩贝</a>
                                  <a href="http://time.qq.com/?pgv_ref=qqcom" target="_blank" bosszone="cp_2_4_4">时光画轴</a>
                                  <a href="https://tianqi.qq.com/" target="_blank" bosszone="cp_2_4_3">天气</a>
                                  <a href="http://users.qq.com/" target="_blank" bosszone="cp_2_4_2">用户社区</a>
                                  <a href="https://dreamreader.qq.com/" target="_blank" bosszone="cp_2_4_1">海豚智音</a>
                                        </li>
                                                                          </ul>
  </div>
</div><!--987950268e8dbd63939b684bcdb9ae5e--><!--[if !IE]>|xGv00|0319588918e6f0f1297bc805c64d3130<![endif]-->
        <!-- /产品 -->

        <!-- 热门赛事 -->
        <!--include virtual="/ninja/saishi_2018.htm"-->
        <!-- /热门赛事 -->

        <!-- 东京奥运 -->
        <div class="mod m-match olympic_wrap" bossexpo="bg_rmss">
          <div class="hd cf">
              <h2 class="tit active fl">
                  <a href="//2020.qq.com/" target="_blank" bosszone="rmss_logo">东京奥运会</a>
              </h2>
              <div class="fr">
                  <div id="match-info" class="match-info" bosszone="rmss_sc">
                    <div id="scroll_A" class="scroll_A">
                      <div class="inner">
                        <ul class="list cf">
                        </ul>
                      </div>
                    </div>
                    <div class="match_more q-icons">
                      <a href="//2020.qq.com/schedule/index.htm?t=navhot" target="_blank">更多</a>
                    </div>
                  </div>
              </div>
          </div>
          <div class="medal_con">
            <a href="//2020.qq.com/countries/country.html?nocId=26&columnId=130000" target="_blank">
              <div class="medal_logo">
                <img src="//mat1.gtimg.com/qqcdn/qqindex2021/dist/img/olympic_china_logo.png" alt="中国奖牌">
              </div>
              <div class="medal_list">
                <ul class="medal_title">
                  <li>排名</li>
                  <li>金牌</li>
                  <li>银牌</li>
                  <li>铜牌</li>
                  <li>总数</li>
                </ul>
                <ul class="medal_num">
                  <li class="medalRank"></li>
                  <li class="medalGold"></li>
                  <li class="medalSilver"></li>
                  <li class="medalBronze"></li>
                  <li class="medalTotal"></li>
                </ul>
              </div>
            </a>
          </div>
          <div class="bd">
    <ul class="news-list">
                                <li class="video-box click-pop-play" bosszone="rmss_1" bossvv="vv_rmss">
            <img src="//img1.gtimg.com/ninja/2/2021/08/ninja162796834169150.jpg" alt="正直播《赢战东京》：林跃做客 跳水梦之队再冲金">
            <i class="q-icons icon-play"></i>
            <div class="desc undis">1346783200</div>
            <a class="txt" href="https://sports.qq.com/kbsweb/game.htm?mid=100002:20212903" target="_blank">正直播《赢战东京》：林跃做客 跳水梦之队再冲金</a>
                        <div id="mod-player1" class="mod-player" data-vid="1346783200"  data-url="https://sports.qq.com/kbsweb/game.htm?mid=100002:20212903"></div>
            <div class="click-layer"></div>
        </li>
                                            <li bosszone="rmss_2">
            <a href="http://view.inews.qq.com/a/SPO2021080300604400" target="_blank">皮划艇静水男子1000米双人划艇决赛 刘浩/郑鹏飞摘银</a>
            </li>
                                            <li bosszone="rmss_3">
            <a href="https://view.inews.qq.com/a/20210803A01J4Z00" target="_blank">东京奥运射击总结 中国队金牌追平境外奥运会最佳</a>
            </li>
                                            <li bosszone="rmss_4">
            <a href="http://view.inews.qq.com/a/20210803A06XTN00" target="_blank">美国逆转西班牙进四强 卢比奥38＋4杜兰特29分</a>
            </li>
                    </ul>
</div><!--a2678bbd2297ae1cf4f57ccf4f5baab6-->
        </div>
        <!-- /东京奥运 -->

        <!-- 今日热播 -->
        <div class="mod m-todayhot" bossexpo="bg_jrrb">
  <div class="hd-2 cf">
    <h2 class="tit active fl">
      <a href="https://v.qq.com/" target="_blank" bosszone="jrrb_logo">今日热播</a>
    </h2>
  </div>
  <div class="bd">
    <ul id="jrrb_news_1" class="news-list cf">
                                  <li class="video-item fl">
            <div class="pic video-box click-drag-play" bosszone="jrrb_1" bossvv="vv_jrrb">
              <img src="//img1.gtimg.com/ninja/2/2021/08/ninja162794618363356.jpg" alt="中国狂揽金牌！十分钟连夺三金">
              <i class="q-icons icon-play2"></i>
              <div class="txt">中国狂揽金牌！十分钟连夺三金</div>
              <div class="desc undis">x3265d5rgk8</div>
              <div id="mod-player2" class="mod-player" data-vid="x3265d5rgk8" data-url="https://new.qq.com/omn/20210802/20210802V0B2ZB00.html"></div>
              <div class="click-layer"></div>
            </div>
          </li>
                                          <li class="video-item fr">
            <div class="pic video-box click-drag-play" bosszone="jrrb_2" bossvv="vv_jrrb">
              <img src="//img1.gtimg.com/ninja/2/2021/08/ninja162794687630482.jpg" alt="饭圈乱象整治取得阶段性成效">
              <i class="q-icons icon-play2"></i>
              <div class="txt">饭圈乱象整治取得阶段性成效</div>
              <div class="desc undis">p3265qwvqmc</div>
              <div id="mod-player3" class="mod-player" data-vid="p3265qwvqmc" data-url="https://new.qq.com/omn/20210803/20210803V01GNI00.html"></div>
              <div class="click-layer"></div>
            </div>
          </li>
                                                      </ul><ul id="jrrb_news_2" class="news-list">
                    <li class="item" bosszone="jrrb_3">
            <a href="https://new.qq.com/omn/20210802/20210802V0EJG100.html" target="_blank">山西临汾籍擅自脱离郑州隔离点男子为逃犯 已被警方抓获</a>
          </li>
                                                    <li class="item" bosszone="jrrb_4">
            <a href="https://new.qq.com/omn/20210802/20210802V0FHU100.html" target="_blank">云南玉溪山体滑坡瞬间：巨石从山腰翻滚砸落 公路被砸出裂缝</a>
          </li>
                                                    <li class="item" bosszone="jrrb_5">
            <a href="https://new.qq.com/omn/20210802/20210802V070JZ00.html" target="_blank">市民拍西安大唐不夜城疫情前后对比：希望再睹大唐盛景</a>
          </li>
                                                    <li class="item" bosszone="jrrb_6">
            <a href="https://new.qq.com/omn/20210803/20210803V01M7500.html" target="_blank">河南疫情防控形势严峻：乡村连夜调出挖掘机 开始封村断路</a>
          </li>
                                                    <li class="item" bosszone="jrrb_7">
            <a href="https://new.qq.com/omn/20210803/20210803V01IQB00.html" target="_blank">大量印度教徒无视疫情扎堆恒河沐浴：有神的祝福 疫情不会有</a>
          </li>
                  </ul>
  </div>
</div><!--ef752dca453510f91f555499312729e0-->
        <!-- /今日热播 -->
      </div>

    </div>
    <!-- /要闻 -->

    <!-- 视觉焦点 -->
    <div class="layout">
      <div class="index-dom-html structure-imgs" id="visual_focus_20200724"></div>
    </div>
    <!--include virtual="/ninja/visual_focus_20200724.htm"-->
    <!-- /视觉焦点 -->  

    <!-- 广告2 -->
    <div class="layout qq-gg gg-2 cf">
      <div class="col-1 fl">
        <!--NEW_QQCOM_N_Width2_div AD begin...."l=NEW_QQCOM_N_Width2&log=off"--><div id="NEW_QQCOM_N_Width2" style="width:920px;height:90px;" class="l_qq_com"></div><!--NEW_QQCOM_N_Width2 AD end --><!--[if !IE]>|xGv00|6c79a7f4e8d6aec45d089679f71d59ee<![endif]-->
      </div>
      <div class="col-2 fr">
        <!--NEW_QQCOM_N_button1_div AD begin...."l=NEW_QQCOM_N_button1&log=off"--><div id="NEW_QQCOM_N_button1" style="width:440px;height:90px;" class="l_qq_com"></div><!--NEW_QQCOM_N_button1 AD end --><!--[if !IE]>|xGv00|1722760f894cd85a0ed41de85dc28e3b<![endif]-->
      </div>
    </div>
    <!-- /广告2 -->

    <!-- 娱乐/体育/NBA -->
    <div class="layout qq-channel2col channelent channel-num6 cf">

      <div class="col col-2 fl">

        <!-- 娱乐 -->
        <div class="mod-ch">
          <div class="title nst">
            <a class="txt active" href="https://new.qq.com/ch/ent/" target="_blank" bosszone="yule_logo">娱乐</a>
            <span bosszone="yule_dh">
              <a class="txt" href="https://new.qq.com/ch2/tv" target="_blank">电视剧</a>
              <a class="txt" href="https://new.qq.com/ch2/movie" target="_blank">电影</a>
              <a class="txt" href="https://new.qq.com/ch2/music" target="_blank">音乐</a>
            </span>
            <ul class="label" bosszone="yule_om">
                <li><a href="https://new.qq.com/omn/author/32" target="_blank">贵圈</a></li>
  <li><a href="https://new.qq.com/omn/author/26135" target="_blank">懂小姐</a></li>
  <li><a href="https://new.qq.com/omn/author/6853487" target="_blank">认真映画</a></li>
<!--16f576bb002ac16dcbe30d6e528dca11--><!--[if !IE]>|xGv00|f830b6435807e5e6bbb226ae0a5cc0bd<![endif]-->
            </ul>
          </div>
          <div class="bdwrap js_chyh">
			      <div class="index-dom-html structure-text-horizontally structure" id="index_ent_20200724"></div>
            <!--include virtual="/ninja/index_ent_20200724.htm" -->
            <div class="hyh js_hyh" bosszone="yule_more">
              <span class="htxt">换一换</span>
              <ul class="hwrap" id="js_enthyh">
                <li class="hpoint active" data-rel="#js_entbd1"></li>
                <li class="hpoint" data-rel="#js_entbd2"></li>
                <li class="hpoint" data-rel="#js_entbd3"></li>
              </ul>
            </div>
          </div>
        </div>
        <!-- /娱乐 -->

        <!-- 体育 -->
        <div class="mod-ch">
          <div class="title nst">
            <a class="txt active" href="http://sports.qq.com/" target="_blank" bosszone="tiyu_logo">体育</a>
            <span bosszone="tiyu_dh">
              <a class="txt" href="https://sports.qq.com/csocce/csl/" target="_blank">中超</a>
              <a class="txt" href="http://sports.qq.com/cba/" target="_blank">CBA</a>
              <a class="txt" href="http://sports.qq.com/premierleague/" target="_blank">英超</a>
              <a class="txt" href="http://fans.sports.qq.com/#/" target="_blank">社区</a>
            </span>
            <ul class="label" bosszone="tiyu_om">
                <li><a href="https://v.qq.com/detail/8/87659.html" target="_blank">超新星运动会</a></li>
  <li><a href="https://v.qq.com/detail/8/87756.html" target="_blank">女主播大乐透</a></li>
  <li><a href="https://v.qq.com/detail/8/81602.html" target="_blank">有球必应</a></li>
  <li><a href="https://v.qq.com/detail/5/52906.html" target="_blank">中场不安定</a></li>
<!--9bc44bb1ee8fb6a000a940264e51b733--><!--[if !IE]>|xGv00|a85c819f48a9d486de8774f44fbf208e<![endif]-->
            </ul>
          </div>
          <div class="bdwrap js_chyh">
            <div class="index-dom-html structure-text-horizontally structure" loadsport="1" id="index_sports_20200724"></div>
            <!--include virtual="/ninja/index_sports_20200724.htm" -->
            <div class="hyh js_hyh" bosszone="tiyu_more">
              <span class="htxt">换一换</span>
              <ul class="hwrap" id="js_sportshyh">
                <li class="hpoint active" data-rel="#js_sportsbd1"></li>
                <li class="hpoint" data-rel="#js_sportsbd2"></li>
                <li class="hpoint" data-rel="#js_sportsbd3"></li>
              </ul>
            </div>
          </div>
        </div>
        <!-- /体育 -->


      </div>

      <div class="col col-1 fr">
        <div id="mod-recommend" class="mod mod-recommend">
          <i class="line"></i>
          <div class="hd cf">
            <h2 class="tit fl">为你推荐</h2>
            <a class="more-btn fr" href="javascript:;" data-src="https://news.qq.com/" bosszone="wntj_more">点击查看 9 条新内容</a>
            <i class="icon-dot"></i>
          </div>
          <div class="bd">
            <div class="list">
              <div class="index-dom-html structure-text-vertically structure" id="recommend_20201021"></div>
              <!--include virtual="/ninja/recommend_20201021.htm"-->
            </div>
          </div>
        </div>
      </div>

    </div>
    <!-- /娱乐/体育/NBA -->

    <!-- 财经/军事 -->
    <div class="layout channel2col qq-channel2col channel-num6 cf">
      <div class="col col-2 fl">
        <div class="title nst">
          <a class="txt active" href="http://finance.qq.com" target="_blank" bosszone="caijing_logo">财经</a>
          <span bosszone="caijing_dh">
            <a class="txt" href="http://stock.qq.com/" target="_blank">证券</a>
            <a class="txt" href="http://money.qq.com/" target="_blank">理财</a>
          </span>
          <ul class="label" bosszone="caijing_om">
              <li><a href="https://new.qq.com/omn/author/5178949" target="_blank">第一财经</a></li>
  <li><a href="https://new.qq.com/omn/author/5564731" target="_blank">界面新闻</a></li>
  <li><a href="https://new.qq.com/omn/author/5005722" target="_blank">每日经济新闻</a></li>
  <li><a href="https://new.qq.com/omn/author/5373662" target="_blank">财约你</a></li>
<!--d4120556c8f48fa4cc68947280f5af23--><!--[if !IE]>|xGv00|2397e68d03fe1a1234773346456910f5<![endif]-->
          </ul>
        </div>
        <div class="bdwrap js_chyh">
          <div class="bd stockbd cf" id="js_stockbd1" bosszone="caijing_1" bossexpo="bg_caijing_1">
            <div class="bdleft">
              <div class="index-dom-html structure-text structure" id="index_finance1_20200724"></div>
              <!--include virtual="/ninja/index_finance1_20200724.htm" -->
            </div>
            <div class="bdright">
              <div class="index-dom-html" id="index_stock1_zhishu"></div>
              <div class="index-dom-html structure-text structure" id="index_stock1_20200724"></div>
              <!--include virtual="/ninja/index_stock1_20200724.htm" -->
            </div>
          </div>
          <div class="bd cf undis" id="js_stockbd2" bosszone="caijing_2" bossexpo="bg_caijing_2">
            <div class="bdleft">
              <div class="index-dom-html" id="index_finance2_20200724"></div>
              <!--include virtual="/ninja/index_finance2_20200724.htm" -->
            </div>
            <div class="bdright">
              <div class="index-dom-html" id="index_stock2_20200724"></div>
            <!--include virtual="/ninja/index_stock2_20200724.htm" -->
            </div>
          </div>
          <div class="bd cf undis" id="js_stockbd3" bosszone="caijing_3" bossexpo="bg_caijing_3">
            <div class="bdleft">
              <div class="index-dom-html" id="index_finance3_20200724"></div>
              <!--include virtual="/ninja/index_finance3_20200724.htm" -->
            </div>
            <div class="bdright">
              <div class="index-dom-html" id="index_stock3_20200724"></div>
              <!--include virtual="/ninja/index_stock3_20200724.htm" -->
            </div>
          </div>
          <div class="hyh js_hyh" bosszone="caijing_more">
            <span class="htxt">换一换</span>
            <ul class="hwrap" id="js_stockhyh">
              <li class="hpoint active" data-rel="#js_stockbd1"></li>
              <li class="hpoint" data-rel="#js_stockbd2"></li>
              <li class="hpoint" data-rel="#js_stockbd3"></li>
            </ul>
          </div>
        </div>
      </div>
      <div class="col col-1 fr" bossexpo="bg_junshi">
        <div class="title">
          <a class="txt active" href="https://new.qq.com/ch/milite/" target="_blank" data-rel="#js_bdmil" bosszone="junshi_logo">军事</a>
        </div>
        <div bosszone="junshi">
          <div class="index-dom-html structure-text structure" id="index_milite_20200724"></div>
          <!--include virtual="/ninja/index_milite_20200724.htm" -->
        </div>
      </div>
    </div>
    <!-- 财经/军事 -->

    <!-- NBA/大家 -->
    <div class="layout channel2col qq-channel2col channel-num6 cf">
        <div class="col col-2 fl">
            <div class="title nst">
              <a class="txt active" href="http://sports.qq.com/nba/" target="_blank" bosszone="nba_logo">NBA</a>
              <ul class="label" bosszone="nba_om">
                	<li><a href="http://sports.qq.com/nbavideo/" target="_blank">NBA视频</a></li>
	<li><a href="http://sports.qq.com/nbavideo/topsk/" target="_blank">TOP时刻</a></li>
	<li><a href="http://nba.stats.qq.com/stats/" target="_blank">数据中心</a></li>
<!--6d53cccf9c0ee8e250df3d63048c39e4--><!--[if !IE]>|xGv00|47402e3084fcadff3d28e78c0805a35b<![endif]-->
              </ul>
            </div>
            <div class="bdwrap js_chyh">
              <div class="index-dom-html structure-text-horizontally structure" loadsport="1" id="index_sportsnba_20200724"></div>
              <!--include virtual="/ninja/index_sportsnba_20200724.htm" -->
               <div class="hyh js_hyh" bosszone="nba_more">
                <span class="htxt">换一换</span>
                <ul class="hwrap" id="js_nbahyh">
                  <li class="hpoint active" data-rel="#js_nbabd1"></li>
                  <li class="hpoint" data-rel="#js_nbabd2"></li>
                  <li class="hpoint" data-rel="#js_nbabd3"></li>
                </ul>
              </div>
            </div>
        </div>
        <div class="col col-1 fr" bossexpo="bg_dajia">
          <div class="title">
            <!-- <a class="txt active" href="http://dajia.qq.com/" target="_blank" bosszone="dajia_logo">大家</a> -->
            <a class="txt active" href="//new.qq.com/omn/author/5107513" target="_blank" bosszone="dajia_logo">较真</a>
          </div>
          <div bosszone="dajia">
            <div class="index-dom-html structure-text structure" id="index_jiaozhen_2020"></div>
            <!--include virtual="/ninja/index_dajia_2018.htm" -->
            <!--include virtual="/ninja/index_jiaozhen_2020.htm" -->
          </div>
        </div>
      </div>
      <!-- 财经/大家 -->

    <!-- 科技/时尚 -->
    <div class="layout channel2col qq-channel2col channel-num6 cf">
      <div class="col col-2 fl">
        <div class="title nst">
          <a class="txt active" href="https://new.qq.com/ch/tech/" target="_blank" bosszone="keji_logo">科技</a>
          <span bosszone="keji_dh">
            <a class="txt" href="//news.qq.com/newspedia/kepu.htm" target="_blank">科普</a>
          </span>
          <ul class="label" bosszone="keji_om">
              <li><a href="https://new.qq.com/zt/template/?id=TEC2019052000259700" target="_blank">5G资讯</a></li>
<!--a1c4151fc36a592fca14f45269555a08--><!--[if !IE]>|xGv00|0c0f5e155ad8b0bafdc20601737ef7e7<![endif]-->
          </ul>
        </div>
        <div class="bdwrap js_chyh">
          <div class="index-dom-html structure-text-horizontally structure" id="index_tech_20200724"></div>
          <!--include virtual="/ninja/index_tech_20200724.htm" -->
          <div class="hyh js_hyh" bosszone="keji_more">
            <span class="htxt">换一换</span>
            <ul class="hwrap" id="js_techhyh">
              <li class="hpoint active" data-rel="#js_techbd1"></li>
              <li class="hpoint" data-rel="#js_techbd2"></li>
              <li class="hpoint" data-rel="#js_techbd3"></li>
            </ul>
          </div>
         </div>
      </div>
      <div class="col col-1 fr" bossexpo="bg_shishang">
        <div class="title">
          <a class="txt active" href="https://new.qq.com/ch/fashion/" target="_blank" bosszone="shishang_logo">时尚</a>
        </div>
        <div bosszone="shishang">
          <div class="index-dom-html structure-text structure" id="index_fashion_20200724"></div>
          <!--include virtual="/ninja/index_fashion_20200724.htm" -->
        </div>
      </div>
    </div>
    <!-- 科技/时尚 -->

   <!-- 汽车/房产 -->
    <div class="layout channel2col qq-channel2col channel-num6 cf">
      <div class="col col-2 fl">
        <div class="title nst">
          <a class="txt active" href="http://auto.qq.com/" target="_blank" bosszone="qiche_logo">汽车</a>
          <ul class="label" bosszone="qiche_om">
              <li><a href="https://new.qq.com/ch2/hyrd" target="_blank">行业热点</a></li>
  <li><a href="https://auto.qq.com/car_public/index.shtml" target="_blank">车型大全</a></li>
  <li><a href="http://v.qq.com/auto/" target="_blank">精彩视频</a></li>
  <li><a href="http://automall.qq.com/web" target="_blank">汽车商城</a></li>
<!--78bdb5def66102b2e5de894017d41e5d-->
          </ul>
        </div>
        <div class="bdwrap js_chyh">
          <div class="index-dom-html structure-text-horizontally structure" id="index_carauto_20200724"></div>
          <!--include virtual="/ninja/index_carauto_20200724.htm" -->
          <div class="hyh js_hyh" bosszone="qiche_more">
            <span class="htxt">换一换</span>
            <ul class="hwrap" id="js_autohyh">
              <li class="hpoint active" data-rel="#js_autobd1"></li>
              <li class="hpoint" data-rel="#js_autobd2"></li>
              <li class="hpoint" data-rel="#js_autobd3"></li>
            </ul>
          </div>
        </div>
      </div>
      <div class="col col-1 fr" bossexpo="bg_fangchan">
        <div class="title">
          <a class="txt active" href="http://house.qq.com/" target="_blank" bosszone="fangchan_logo">房产</a>
        </div>
        <div bosszone="fangchan">
          <div class="index-dom-html structure-text structure" id="index_house_20200724"></div>
          <!--include virtual="/ninja/index_house_20200724.htm" -->
        </div>
      </div>
    </div>
    <!-- /汽车/房产 -->

    <!-- 综艺影视剧 -->
    <div class="layout qq-videos m40" style="display:none;">
      <div class="title" id="js_vtitle">
        <a class="txt active" href="https://v.qq.com/x/variety/" target="_blank" data-rel="#js_bdzy" bosszone="zongyi_logo">综艺</a>
        <span class="split"></span>
        <a class="txt" href="http://v.qq.com/tv/" target="_blank" data-rel="#js_bdysj" bosszone="zongyi_logo">电视剧</a>
        <span class="split"></span>
        <a class="txt" href="http://v.qq.com/movie/" target="_blank" data-rel="#js_bdmv" bosszone="dianying_logo">电影</a>
        <span class="split"></span>
        <a class="txt" href="https://v.qq.com/child" target="_blank" data-rel="#js_bdchild" bosszone="shaoer_logo">少儿</a>
        <ul class="label" bosszone="zongyi_om">
          <!--include virtual="/ninja/index_omzy_2018.htm" -->
        </ul>
      </div>
      <div class="mainbody" id="js_videosbd">
        <div id="js_bdzy" bossexpo="bg_zongyi">
          <div class="bdwrap">
            <div class="bd-inner cf" id="js_zyCon">

              <!--include virtual="/ninja/index_zongyi_2018.htm"-->
            </div>
          </div>
          <div id="js_zydir" bosszone="zongyi_more">
            <a href="javascript:;" class="prev icon disabled" data-rel="#js_zyCon_0"></a>
            <a href="javascript:;" class="next icon" data-rel="#js_zyCon_1"></a>
          </div>
        </div>
        <div id="js_bdysj" class="undis" bossexpo="bg_dsj">
          <div class="bdwrap">
            <div class="bd-inner cf" id="js_ysjCon">
              <!--include virtual="/ninja/index_ysj_2018.htm"-->
            </div>
          </div>
          <div id="js_ysjdir" bosszone="dsj_more">
            <a href="javascript:;" class="prev icon disabled" data-rel="#js_ysjCon_0"></a>
            <a href="javascript:;" class="next icon" data-rel="#js_ysjCon_1"></a>
          </div>
        </div>
        <div id="js_bdmv" class="undis" bossexpo="bg_dianying">
          <div class="bdwrap">
            <div class="bd-inner cf" id="js_mvCon">
              <!--include virtual="/ninja/index_movie_2018.htm"-->
            </div>
          </div>
          <div id="js_mvdir" bosszone="dianying_more">
            <a href="javascript:;" class="prev icon disabled" data-rel="#js_mvCon_0"></a>
            <a href="javascript:;" class="next icon" data-rel="#js_mvCon_1"></a>
          </div>
        </div>
         <div id="js_bdchild" class="undis" bossexpo="bg_shaoer">
          <div class="bdwrap">
            <div class="bd-inner cf" id="js_childCon">
              <!--include virtual="/ninja/index_child_2018.htm"-->
            </div>
          </div>
          <div id="js_childdir" bosszone="shaoer_more">
            <a href="javascript:;" class="prev icon disabled" data-rel="#js_childCon_0"></a>
            <a href="javascript:;" class="next icon" data-rel="#js_childCon_1"></a>
          </div>
        </div>
        <div class="vplayer">
          <div class="layer"></div>
          <div id="js_videoplayer">

          </div>
        </div>
      </div>
    </div>
    <!-- /综艺影视剧 -->

    <!-- 广告3 -->
    <div class="layout qq-gg gg-3 cf">
      <div class="col-1 fl">
        <!--NEW_QQCOM_N_Width3_div AD begin...."l=NEW_QQCOM_N_Width3&log=off"--><div id="NEW_QQCOM_N_Width3" style="width:920px;height:90px;" class="l_qq_com"></div><!--NEW_QQCOM_N_Width3 AD end --><!--[if !IE]>|xGv00|a54d84501d504c4567ba39a7064f670b<![endif]-->
      </div>
      <div class="col-2 fr">
        <!--NEW_QQCOM_N_button2_div AD begin...."l=NEW_QQCOM_N_button2&log=off"--><div id="NEW_QQCOM_N_button2" style="width:440px;height:90px;" class="l_qq_com"></div><!--NEW_QQCOM_N_button2 AD end --><!--[if !IE]>|xGv00|6743da0daaed589d3944270e6489eda6<![endif]-->
      </div>
    </div>
    <!-- /广告3 -->

    <!-- 军事/历史/文化佛学 -->
    <div class="layout qq-channel3col cf">
      <div class="col col-1">
        <div class="title">
          <a class="txt active" href="https://new.qq.com/ch/edu/" target="_blank" bosszone="jiaoyu_logo">教育</a>
        </div>
        <div bosszone="jiaoyu" bossexpo="bg_jiaoyu">
          <div class="index-dom-html structure-text structure" id="index_education_20200724"></div>
          <!--include virtual="/ninja/index_education_20200724.htm" -->
        </div>
      </div>
      <div class="col col-1">
        <div class="title" id="js_histitle">
          <a class="txt active" href="https://new.qq.com/omn/author/41" target="_blank" data-rel="#js_bdhis" bosszone="lishi_logo">历史</a>
        </div>
        <div class="bdwrap">
          <div class="bd" id="js_bdhis" bosszone="lishi" bossexpo="bg_lishi">
            <div class="index-dom-html structure-text structure" id="index_history_20200724"></div>
            <!--include virtual="/ninja/index_history_20200724.htm" -->
          </div>
        </div>
      </div>
      <div class="col col-1 col-last">
        <div class="title" id="js_title1">
          <a class="txt active" href="https://new.qq.com/ch/cul/" target="_blank" data-rel="#js_bdcul" bosszone="wenhua_logo">文化</a>
          <span class="split"></span>
          <a class="txt" href="https://new.qq.com/ch/cul_ru/" target="_blank" data-rel="#js_bdbud" bosszone="foxue_logo">新国风</a>
        </div>
        <div class="bdwrap">
          <div class="bd" id="js_bdcul" bosszone="wenhua" bossexpo="bg_wenhua">
            <div class="index-dom-html structure-text structure" id="index_culture_20200724"></div>
            <!--include virtual="/ninja/index_culture_20200724.htm" -->
          </div>
          <div class="bd undis" id="js_bdbud" bosszone="foxue" bossexpo="bg_foxue">
            <div class="index-dom-html" id="index_rushidao_20200724"></div>
            <!--include virtual="/ninja/index_rushidao_20200724.htm" -->
          </div>
        </div>
      </div>
    </div>
    <!-- /军事/历史/文化佛学 -->

    <!-- 星座每日运势/游戏动漫/财报 -->
    <div class="layout qq-channel3col cf">
      <div class="col col-1">
        <div class="title" id="js_title2">
          <a class="txt active" href="http://astro.fashion.qq.com/" target="_blank" data-rel="#js_bdastro" bosszone="xingzuo_logo">星座</a>
          <span class="split"></span>
          <a class="txt" href="http://astro.fashion.qq.com/" target="_blank" data-rel="#js_bdfortune" bosszone="yunshi_logo">今日运势</a>
        </div>
        <div class="bdwrap bdwrapast">
          <div class="bd" id="js_bdastro" bosszone="xingzuo" bossexpo="bg_xingzuo">
            <div class="index-dom-html structure-text structure" id="index_astro_20200724"></div>
            <!--include virtual="/ninja/index_astro_20200724.htm" -->
          </div>
          <div class="undis col-astrobd" id="js_bdfortune" bosszone="yunshi" bossexpo="bg_yunshi">
            <div class="astop cf">
              <a class="asleft" href="javascript:;" id="js_aimg" target="_blank">
                <span class="icon Aries" title="白羊座"></span>
                <p class="name">白羊座</p>
              </a>
              <div class="asright">
                <div class="asdesc" id="js_adetail">
                  <div class="desc fortune">
                    <span class="label">今日运势：</span>
                    <div class="fortune-star">
                      <span class="bottom iconastro"></span>
                      <span class="top iconastro"></span>
                    </div>
                  </div>
                  <div class="desc luck">
                    <span class="label">幸运颜色：紫色</span>
                  </div>
                  <div class="desc lucknum">
                    <span class="label">幸运数字：7</span>
                  </div>
                </div>
                <div class="select">
                  <div class="selected iconastro" id="js_aselect">白羊座 03.21-04.19 </div>
                  <ul class="list" id="js_aselectlist">
                    <li class="astroItem" data-value="0">白羊座 03.21-04.19</li>
                    <li class="astroItem" data-value="1">金牛座 04.20-05.20</li>
                    <li class="astroItem active" data-value="2">双子座 05.21-06.21</li>
                    <li class="astroItem" data-value="3">巨蟹座 06.22-07.22</li>
                    <li class="astroItem" data-value="4">狮子座 07.23-08.22</li>
                    <li class="astroItem" data-value="5">处女座 08.23-09.22</li>
                    <li class="astroItem" data-value="6">天秤座 09.23-10.23</li>
                    <li class="astroItem" data-value="7">天蝎座 10.24-11.22</li>
                    <li class="astroItem" data-value="8">射手座 11.23-12.21</li>
                    <li class="astroItem" data-value="9">摩羯座 12.22-01.19</li>
                    <li class="astroItem" data-value="10">水瓶座 01.20-02.18</li>
                    <li class="astroItem" data-value="11">双鱼座 02.19-03.20</li>
                  </ul>
                </div>
              </div>
            </div>
            <p class="astxt" id="js_atxt">今天对于一切的工作都是那么专心致志，隔一段时间就要起来走动一下，才能保证有更高的效率，期待已久的事件总算有了一个交代，虽然不是满分，但也算是好结果。
            </p>
            <ul class="txtArea">
                  <li><a class="" href="https://new.qq.com/omn/20210803/20210803A071I500.html" target="_blank">“羊羊羊”九月份了不得，今晚开始起，运势大增，羊羊得意！</a></li>
                            <li><a class="" href="https://new.qq.com/omn/20210803/20210803A06ZC800.html" target="_blank">再熬三天，属猪人今生最大的贵人就要降临，看看怎么回事？</a></li>
                              </ul>
          </div>
        </div>
      </div>
      <div class="col col-1">
        <div class="title" id="js_title3">
          <a class="txt active" href="https://new.qq.com/ch/games/" target="_blank" data-rel="#js_bdgame" bosszone="youxi_logo">游戏</a>
          <span class="split"></span>
          <a class="txt" href="https://new.qq.com/ch/comic/" target="_blank" data-rel="#js_bdcomic" bosszone="dongman_logo">动漫</a>
        </div>
        <div class="bdwrap">
          <div class="bd" id="js_bdgame" bosszone="youxi" bossexpo="bg_youxi">
            <div class="index-dom-html structure-text structure" id="index_games_20200724"></div>
            <!--include virtual="/ninja/index_games_20200724.htm" -->
          </div>
          <div class="bd undis" id="js_bdcomic" bosszone="dongman" bossexpo="bg_dongman">
            <div class="index-dom-html" id="index_comic_2018test"></div>
            <!--include virtual="/ninja/index_comic_2018test.htm" -->
          </div>
        </div>
      </div>
      <div class="col col-1 col-last col-tencent" bosszone="caibao" bossexpo="bg_caibao">
        <div class="title">
          <a class="txt active" href="https://www.tencent.com/zh-cn/company.html" target="_blank" bosszone="caibao_logo">财报</a>
        </div>
        <div bosszone="caibao">
          <div class="index-dom-html structure-text structure" id="index_caibao_2018_test"></div>
          <!--include virtual="/ninja/index_caibao_2018_test.htm" -->
        </div>
      </div>
    </div>
    <!-- 星座每日运势/游戏动漫/财报 -->

    <!-- 高清组图 -->
    <div class="layout">
      <div class="index-dom-html structure-imgs" id="hd_picture_20200724"></div>
    </div>
    <!--include virtual="/ninja/hd_picture_20200724.htm"-->
    <!-- /高清组图 -->

    <!-- 广告4 -->
    <div class="layout qq-gg gg-4">
      <!--NEW_QQCOM_N_Width4_div AD begin...."l=NEW_QQCOM_N_Width4&log=off"--><div id="NEW_QQCOM_N_Width4" style="width:1400px;height:90px;" class="l_qq_com"></div><!--NEW_QQCOM_N_Width4 AD end --><!--[if !IE]>|xGv00|988d4677a77862e5edbcb3f52aba9377<![endif]-->
    </div>
    <!-- /广告4 -->

    <!--NEW_WWW_RM_RightMove1_div AD begin...."l=NEW_WWW_RM_RightMove1&log=off"--><div id="NEW_WWW_RM_RightMove1" style="width:400px;height:300px;display:none;margin:0 auto;" class="l_qq_com"></div><!--NEW_WWW_RM_RightMove1 AD end --><!--[if !IE]>|xGv00|c020c69143131b5b928166fd08447a05<![endif]-->
    <!--NEW_QQ_Couplet_div AD begin...."l=NEW_QQ_Couplet&log=off"--><div id="NEW_QQ_Couplet" style="width:100px;height:300px;display:none;" class="l_qq_com"></div><!--NEW_QQ_Couplet AD end --><!--[if !IE]>|xGv00|5b0b305532624ef799bf7dc76b9e5338<![endif]-->

    <!-- 版权信息 -->
    <div class="layout qq-footer" bosszone="dibu" bossexpo="bg_dibu">
  <a href="http://www.tencent.com/" target="_blank" rel="nofollow">关于腾讯</a> | <a href="http://www.tencent.com/index_e.shtml"
    target="_blank" rel="nofollow">About Tencent</a> | <a href="http://www.qq.com/contract.shtml" target="_blank"
    rel="nofollow">服务协议</a>
  | <a href="https://privacy.qq.com/" target="_blank" rel="nofollow">隐私政策</a> | <a href="http://open.qq.com/"
    target="_blank" rel="nofollow">开放平台</a><!--  | <a href="http://www.tencentmind.com/" target="_blank" rel="nofollow">广告服务</a> -->
  | <a href="https://www.tencent.com/zh-cn/partnership.html" target="_blank" rel="nofollow">商务洽谈</a> | <a href="http://hr.tencent.com/"
    target="_blank" rel="nofollow">腾讯招聘</a> | <a href="http://gongyi.qq.com/" target="_blank" rel="nofollow">腾讯公益</a>
  | <a href="http://kf.qq.com/" target="_blank" rel="nofollow">客服中心</a> | <a href="http://www.qq.com/map/" target="_blank"
    rel="nofollow">网站导航</a> | <a href="http://dldir1.qq.com/dlomg/qqcom/mini/QQNewsMini5.exe" rel="nofollow">客户端下载</a><!--
  | <a href="http://www.tencent.com/law/mo_law.shtml?/law/copyright.htm" target="_blank" rel="nofollow">版权所有</a>--><br>
  <a href="http://szwljb.sz.gov.cn/" target="_blank" rel="nofollow">深圳举报中心</a> | <a href="http://ga.sz.gov.cn"
    target="_blank" rel="nofollow">深圳公安局</a> | <a href="http://www.qq.com/dzwfggcns.htm" target="_blank" rel="nofollow">抵制违法广告承诺书</a><!--  | <a class="lchot" href="http://www.gdis.cn/admin/newstext_netsun.asp" target="_blank" rel="nofollow">阳光·绿色网络工程</a> -->
  | <a href="http://www.qq.com/copyright.shtml" target="_blank" rel="nofollow">侵权投诉指引</a> | <a href="https://www.qq.com/bjhlwfyflfwgzz.shtml"
    target="_blank" rel="nofollow">北京互联网法院法律服务工作站</a> | <a href="https://gdca.miit.gov.cn/"
    target="_blank" rel="nofollow">广东省通管局</a><br>
  <span><a href="http://www.qq.com/culture.shtml" target="_blank" rel="nofollow">粤网文[2020]3396-195号</a> <a href="http://www.qq.com/internet_licence.htm"
      target="_blank" rel="nofollow">新出网证（粤）字010号</a> <a href="http://www.qq.com/cbst.shtml" target="_blank" rel="nofollow">网络视听许可证1904073号</a>
    增值电信业务经营许可证：<a href="http://beian.miit.gov.cn/" target="_blank" rel="nofollow">粤B2-20090059</a> <a href="http://www.qq.com/icp1.shtml"
      target="_blank" rel="nofollow">B2-20090028</a>
  </span><br>
  <a href="http://www.qq.com/scio.htm" target="_blank" rel="nofollow">新闻信息服务许可证</a> <a href="http://www.qq.com/xwdz.shtml"
    target="_blank" rel="nofollow">粤府新函[2001]87号</a> 违法和不良信息举报电话：0755-83765566-4 <a style="" target="_blank" href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=44030002000001"><span>粤公网安备
      44030002000001号</span></a><br>
  <a href="http://www.qq.com/spyp.htm" target="_blank">互联网药品信息服务资格证书 （粤）—非营业性—2017-0153</a><br>
  <span>Copyright  1998 - </span><span id="currentFullYear">2018</span> <span>Tencent. All Rights Reserved</span>
  <div class="footernew">
    <div class="footernewdiv" style="width:558px">
    <p>
      <span style="width:44px;" class="fl"><a href="http://www.12377.cn/" target="_blank" rel="nofollow"><img width="44" height="44" border="0" alt="中国互联网举报中心" src="//mat1.gtimg.com/www/images/qq2012/buliang.png"></a></span>
      <span style="width:64px;" class="fr"><a class="lcblack" href="http://www.12377.cn/" target="_blank" rel="nofollow">中国互联网<br>
      举报中心</a></span>
    </p>
    <p>
      <span style="width:44px;" class="fl"><a href="http://www.wenming.cn" target="_blank" rel="nofollow"><img width="44" height="42" border="0" alt="中国文明网传播文明" src="//mat1.gtimg.com/www/images/qq2012/wmlogo.gif"></a></span>
      <span style="width:64px;" class="fr"><a class="lcblack" href="http://www.wenming.cn" target="_blank" rel="nofollow">中国文明网<br>传播文明</a></span>
    </p>
    <p style="width:128px;height:52px;border:0;">
      <span style="padding:0;" class="fl"><a href="https://ss.knet.cn/verifyseal.dll?sn=2010051300100001081&ct=df&a=1&pa=337337" target="_blank" rel="nofollow"><img border="0" alt="诚信网站" src="//mat1.gtimg.com/www/images/qq2012/cxrz5.png"></a></span>
    </p>
    <p>
      <span style="width:44px;" class="fl"><a href="http://szcert.ebs.org.cn/6917b6fe-b844-4e12-97c5-85b8d1df30ed" target="_blank" rel="nofollow"><img src="//mat1.gtimg.com/www/images/qq2012/gswj2015.jpg" title="深圳市市场监督管理局企业主体身份公示" alt="深圳市市场监督管理局企业主体身份公示"></a></span>
      <span style="width:64px;" class="fr"><a class="lcblack" href="http://szcert.ebs.org.cn/6917b6fe-b844-4e12-97c5-85b8d1df30ed" target="_blank" rel="nofollow">工商网监<br>电子标识</a></span>
    </p>
    </div>
  </div>
</div>
<script type="text/javascript">
  var currentFullYear = (new Date()).getFullYear();
  document.getElementById('currentFullYear').innerHTML = currentFullYear;
</script>
<!--871eae83c4732085e569c74e6c01bcbe-->
    <!-- /版权信息 -->

    <!-- 电梯 -->
    <div class="elevator" id="elevator">
      <a href="javascript:" class="refresh fix" id="js_refresh" title="刷新" bosszone="shuaxin"><span class="icon"></span><br />刷新</a>
      <a href="https://support.qq.com/products/4312" target="_blank" class="feedback fix" title="用户反馈" bosszone="fankui">用户<br />反馈</a>
      <a href="javascript:void(0)" target="_self" class="backtop" id="js_gotop" title="返回顶部" bosszone="dingbu"><span class="icon"></span></a>
    </div>
    <!-- /电梯 -->

  </div>

  <!-- 视频弹层 -->
  <div id="pop-player" class="pop-player">
  <div class="inner">
    <div class="player-hd">
      <div id="video-pop" class="player-container"></div>
      <div class="tit"></div>
      <a class="close-btn" href="javascript:;">关闭</a>
    </div>
    <div class="player-ft cf">
      <div class="scroll-mod">
        <ul class="player-list cf"></ul>
        <a href="javascript:;" class="q-icons btn btn-left"></a>
        <a href="javascript:;" class="q-icons btn btn-right"></a>
      </div>
    </div>
  </div>
</div>

<div id="pop-player2" class="pop-player pop-player2">
  <div class="inner">
    <div class="player-hd">
      <div id="video-pop2" class="player-container"></div>
      <div class="tit"></div>
      <a class="close-btn" href="javascript:;">关闭</a>
      <div class="hd-info"></div>
    </div>
    <div class="player-ft cf">
      <div class="scroll-mod">
        <ul class="player-list cf"></ul>
      </div>
      <a href="javascript:;" class="q-icons btn btn-left"></a>
      <a href="javascript:;" class="q-icons btn btn-right"></a>
    </div>
  </div>
</div>

<div id="min-player" class="min-player">
  <div class="min-hd cf">
    <h2 class="tit fl"></h2>
    <a class="close-btn fr" href="javascript:;">关闭</a>
  </div>
  <div class="min-bd">
    <div id="video-min" class="player-container"></div>
  </div>
</div><!--ec4544fe058862e423cdc3225e110e49--><!--[if !IE]>|xGv00|6254f87b049c4c938babd0b80a015de3<![endif]-->
  <!-- /视频浮层 -->

  
  <script type="text/javascript">
  //<![CDATA[
  var serverTime = new Date(2021, 08-1, 03, 15, 09, 40);
  //]]>
  </script>
  <script src="//mat1.gtimg.com/www/asset/lib/jquery/jquery/jquery-1.11.1.min.js"></script>
  <script type="text/javascript" src="//joke.qq.com/lucky/jquery.qqscroll.js"></script>
  <script src="//vm.gtimg.cn/tencentvideo/txp/js/txplayer.js" charset="utf-8"></script>
  <script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5d09e4c5.js" charset="utf-8"></script>
  <script src="//mat1.gtimg.com/pingjs/ext2020/configF2017/5e857945.js" charset="utf-8"></script>
  <script src="//mat1.gtimg.com/pingjs/ext2020/dc2017/publicjs/m/ping.js"></script>
	<script>if(typeof(pgvMain) == 'function')pgvMain();</script>
  <script src="//mat1.gtimg.com/qqcdn/qqindex2021/dist/js/qq_c4b4edd7.js" charset="utf-8"></script>

  <script type="text/javascript" src="//imgcache.qq.com/qzone/biz/comm/js/qbs.js"></script>
<script type="text/javascript">
var TIME_BEFORE_LOAD_CRYSTAL = (new Date).getTime();
</script>
<script src="//ra.gtimg.com/web/crystal/v4.7Beta04Build040/crystal-min.js" id="l_qq_com" arguments="{'extension_js_src':'//ra.gtimg.com/web/crystal/v4.7Beta04Build040/crystal_ext-min.js', 'jsProfileOpen':'false', 'mo_page_ratio':'0.01', 'mo_ping_ratio':'0.01', 'mo_ping_script':'//ra.gtimg.com/sc/mo_ping-min.js'}"></script>
<script type="text/javascript">
if(typeof crystal === 'undefined' && Math.random() <= 1) {
  (function() {
    var TIME_AFTER_LOAD_CRYSTAL = (new Date).getTime();
    var img = new Image(1,1);
    img.src = "//dp3.qq.com/qqcom/?adb=1&dm=www&err=1002&blockjs="+(TIME_AFTER_LOAD_CRYSTAL-TIME_BEFORE_LOAD_CRYSTAL);
  })();
}
</script>
<style>.absolute{position:absolute;}</style>
<!--[if !IE]>|xGv00|34ba8056fb38cac38d53027a9f08814a<![endif]-->

  <script>
  // 腾讯分析代码
  var _mtac = {};
  (function() {
      var mta = document.createElement("script");
      mta.src = "//pingjs.qq.com/h5/stats.js?v2.0.2";
      mta.setAttribute("name", "MTAH5");
      mta.setAttribute("sid", "500460529");
      var s = document.getElementsByTagName("script")[0];
      s.parentNode.insertBefore(mta, s);
  })();
  </script>

</body>

</html>`


	tmp, _ := spider.HtmlMinify(html)
	println(tmp)
}
