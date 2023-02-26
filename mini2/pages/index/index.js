Page({
    onShow() {

    },
    data: {
      PageCur: 'basics'
    },
    NavChange(e) {
        let that = this
        console.log(e.currentTarget.dataset.cur)
        that.setData({
        PageCur: e.currentTarget.dataset.cur
      })
    },
    onShareAppMessage() {
      return {
        title: 'ColorUI-高颜值的小程序UI组件库',
        imageUrl: '/images/share.jpg',
        path: '/pages/index/index'
      }
    },
  })