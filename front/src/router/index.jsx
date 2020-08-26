import React, { useCallback } from 'react'
import { BrowserRouter, Switch, Route, Redirect, useLocation } from 'react-router-dom'
import Loadable from 'react-loadable'

import Fullpage from '@/layouts/Fullpage'
import Management from '@/layouts/Management'

import routes from './routers'

function LoadableComponent({ layout, component }) {
  const Component = Loadable({
    loader: () => import(`../${component}`),
    loading() {
      return <span />
    }
  })

  const Layout = layout === 'fullpage' ? Fullpage : Management

  return (
    <Layout>
      <Component />
    </Layout>
  )
}

export default function Router() {
  const onEnter = useCallback(({ path, component, layout, title }) => {
    // 设置title
    document.title = title ? 'G-CMS' : 'G-CMS ' + title

    const hasToken = !!window.localStorage.getItem('token')
    const isToLogin = path === '/login'

    if (isToLogin) {
      if (hasToken) {
        return <Redirect to="/" />
      } else {
        return <LoadableComponent layout={layout} component={component} />
      }
    }

    if (!isToLogin) {
      if (hasToken) {
        return <LoadableComponent layout={layout} component={component} />
      } else {
        const { pathname, search } = window.location
        return <Redirect to={`/login?redirect=${encodeURIComponent(pathname + search)}`} />
      }
    }
  }, [])

  return (
    <BrowserRouter>
      <Switch>
        {routes.map((item) => (
          <Route key={item.path} path={item.path} exact={!!item.exact} render={() => onEnter(item)} />
        ))}
      </Switch>
    </BrowserRouter>
  )
}
