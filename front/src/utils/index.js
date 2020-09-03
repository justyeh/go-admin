// 获取url query参数
export function getQueryVariable(searchKey) {
  if (!searchKey) {
    return ''
  }
  const queryVars = decodeURIComponent(window.location.search).substring(1).split('&')
  for (let i = 0; i < queryVars.length; i++) {
    const [key, val] = queryVars[i].split('=')
    if (searchKey === key) {
      return val
    }
  }
  return ''
}

// 绑定分页
export function bindPage() {
  let current = Number(getQueryVariable('current'))
  if (isNaN(current) || current < 1) {
    current = 1
  }

  let size = Number(getQueryVariable('size'))
  if (![10, 20, 50, 100].includes(size)) {
    size = 10
  }

  return {
    current: current,
    size: size,
    total: 0
  }
}

export function dateFormat(timeStamp, fmt = 'yyyy-MM-dd hh:mm:ss') {
  try {
    const date = new Date(timeStamp * 1000)

    var o = {
      'M+': date.getMonth() + 1, //月份
      'd+': date.getDate(), //日
      'h+': date.getHours(), //小时
      'm+': date.getMinutes(), //分
      's+': date.getSeconds(), //秒
      'q+': Math.floor((date.getMonth() + 3) / 3), //季度
      S: date.getMilliseconds() //毫秒
    }

    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
    for (var k in o) {
      if (new RegExp('(' + k + ')').test(fmt)) {
        fmt = fmt.replace(RegExp.$1, RegExp.$1.length === 1 ? o[k] : ('00' + o[k]).substr(('' + o[k]).length))
      }
    }
  } catch (error) {
    console.error(error)
    fmt = '--'
  }

  return fmt
}
