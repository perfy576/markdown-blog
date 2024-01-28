import Vue from 'vue'
import Router from 'vue-router'
// import Index from '@/page/Index'
import PostList from '@/components/PostList'
import PostContent from '@/components/PostContent'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'post list',
      component: PostList,
      meta: { keepAlive: true }
    },
    {
      path: '/post/:post_id',
      name: 'post',
      component: PostContent,
      // meta: { keepAlive: true }
    },
  ]
})
