<template>
  <div id="app">
    <!-- 这里不需要修改了，如果需要缓存的页面在router.index里面加上meta.keepAlive即可,不带的走下面，带的走上面 -->
    <keep-alive>
      <router-view v-if="$route.meta.keepAlive">
      </router-view>
    </keep-alive>
    <router-view v-if="!$route.meta.keepAlive">
    </router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
  created() {
    console.log("env: ", process.env.NODE_ENV)
    if (process.env.NODE_ENV !== "development") {

      var hm = document.createElement("script");
      hm.src = "https://hm.baidu.com/hm.js?f6300972f555296cb7e7d0e2d09bef19";
      var s = document.getElementsByTagName("script")[0];
      s.parentNode.insertBefore(hm, s);
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  color: #2c3e50;
}
</style>
