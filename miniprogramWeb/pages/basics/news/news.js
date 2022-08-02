// pages/basics/news/news.js
Page({
    /**
     * 页面的初始数据
     */
    data: {
        imgList: [
            'https://ossweb-img.qq.com/images/lol/web201310/skin/big10001.jpg',
            'https://ossweb-img.qq.com/images/lol/web201310/skin/big10001.jpg',
            'https://ossweb-img.qq.com/images/lol/web201310/skin/big10001.jpg',
        ],
        modalName:null
    },

    ViewImage(e) {
        wx.previewImage({
            urls: this.data.imgList,
            current: e.currentTarget.dataset.url
        });
    },

    showModal(e) {
        this.setData({
          modalName: e.currentTarget.dataset.target
        })
      },
      hideModal(e) {
        this.setData({
          modalName: null
        })
      },
})