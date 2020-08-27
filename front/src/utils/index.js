export function getQueryVariable(searchKey) {
  if (!searchKey) {
    return ''
  }
  const queryVars = window.location.search.substring(1).split('&')
  for (let i = 0; i < queryVars.length; i++) {
    const [key, val] = queryVars[i].split('=')
    if (searchKey === key) {
      return val
    }
  }
  return ''
}
