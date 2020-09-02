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
