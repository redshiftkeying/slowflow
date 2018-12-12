import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import ErrPage404 from '@/views/errpages/Page404'
import ErrPage500 from '@/views/errpages/Page500'
import TreeTable from '@/views/TreeTable'

Vue.use(Router)

export default new Router({
  routes: [
    { path: '/', name: 'HelloWorld', component: HelloWorld },
    { path: '/templet', name: 'TreeTable', component: TreeTable },
    { path: '/404', name: '404', component: ErrPage404 },
    { path: '/500', name: '500', component: ErrPage500 },
    { path: '*', redirect: '/404', hidden: true }
  ]
})
