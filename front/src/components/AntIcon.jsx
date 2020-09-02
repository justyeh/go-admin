import React from 'react'
import * as Icon from '@ant-design/icons'
export default ({ name, size = 16 }) => {
  return React.createElement(Icon[name], {
    style: { fontSize: size }
  })
}
