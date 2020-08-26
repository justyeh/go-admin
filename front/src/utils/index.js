export function getQueryParams(searchKey) {
  if (!searchKey) {
    return ''
  }

  const queryArray = window.location.search.split('?')[1].split('&')
  for (let i = 0; i < queryArray.length; i++) {
    const [key, val] = queryArray[i].split('=')
    if (searchKey === key) {
      return val
    }
  }
  
  return ''
}
