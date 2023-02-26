const formatTime = date => {
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hour = date.getHours()
    const minute = date.getMinutes()
    const second = date.getSeconds()
  
    return `${[year, month, day].map(formatNumber).join('/')} ${[hour, minute, second].map(formatNumber).join(':')}`
  }
  
  const formatNumber = n => {
    n = n.toString()
    return n[1] ? n : `0${n}`
  }
  
  const baseUrl = n => {
      return 'https://miniprogram.xlxj1314.cn/'
  }
  
  const login = code => {
    wx.checkSession({
        success: (res) => {
          let token = wx.getStorageSync("token")
          console.log("token:" +　token)
          if (!token) {
              _login()
          }
        },
        fail: (res) => {
          _login()
        }
      })
  }


function _login() {
    wx.login({
        success(res) {
            if(res.code) {
                console.log("code:" +　res.code)
                wx.request({
                  url:  baseUrl() + 'wxLogin',
                  method: "GET",
                  data: {
                      code: res.code
                  },
                  success(res) {
                      const token = res.data.data.token
                      wx.setStorageSync("token",token)
                  }
                })
            }else{
                console.log("err code")
            }
        },
        fail(res) {
            console.log(res)
        }
      })
}

  
  module.exports = {
    formatTime,
    login,
    baseUrl
  }
  