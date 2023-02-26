var app = getApp()
const util = require('../../../utils/util.js')

Component({
    lifetimes:{
        ready: function() {
            console.log('show')
            let that = this
            wx.request({
                url: util.baseUrl() + 'wxUser/getBgUrl',
                header: {
                    'token':  wx.getStorageSync('token')
                },
                success(res) {
                    console.log("bg_url" +　res.data.data.bg_url)
                    that.setData({
                        backgroudUrl: res.data.data.bg_url
                    })
                },
                fail(res) {
                    console.log('fail')
                    console(res)
                }
              })
        }
    },
    options: {
      addGlobalClass: true,
    },
    data: {
        backgroudUrl:"",
        avatar: [
            {
                url:'https://ossweb-img.qq.com/images/lol/web201310/skin/big10001.jpg'
            },
            {
                url:'https://ossweb-img.qq.com/images/lol/web201310/skin/big10001.jpg'
            }
        ],
      elements: [
          {
          title: '纪念日',
          name: 'layout',
          color: 'cyan',
          icon: 'newsfill'
        },
        {
          title: '姨妈记录',
          name: 'background',
          color: 'blue',
          icon: 'colorlens'
        }
      ],
    },
  })