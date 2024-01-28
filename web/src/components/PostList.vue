<template>
  <div>
    <div class="container_header">
      <NavBar :set_open_directory='set_open_directory' />
    </div>

    <div class="container_body">
      <div class="left"></div>
      <ul class="center">
        <li v-for="(file, index) in post_list" :key="index" class="file-item" v-if="is_display(file['filepath'])">
          <h3 class="file-title" @click="navigateToPost(file['filepath'])">{{ getFileName(file["filepath"]) }}</h3>
          <p class="last-modified">修改时间: {{ formatTimestamp(file["modify_time"]) }}</p>
        </li>
      </ul>
      <div class="right_placeholder"></div>

      <ul class="right">
        <Directory class="item" :item="directory_structure" :set_open_directory='set_open_directory'></Directory>
      </ul>
    </div>
  </div>
</template>

<script>
import { post } from '@/utils/net.js'
import { base64_url_encode } from '@/utils/crypto.js'
import Directory from './Directory.vue'
import NavBar from './NavBar'

export default {
  name: 'PostList',
  data() {
    return {
      post_list: [],
      directory_structure: {},
      open_directory: "/",
    }
  },
  components: { Directory, NavBar },
  onActivated() {
    // console.log("onActivated:", this.post_list.length)
  },
  async created() {
    // console.log("创建组件！！:", this.post_list.length)
    if (this.post_list.length === 0) {
      this.directory_structure = await this.fetch_post_list();
    }
  },
  methods: {
    async fetch_post_list() {
      var data = {
        'directory_path': '/',
      }
      const response = await post('/blog/v1/get-directory-structure', data)
      let markdown_list = []
      this.extractMdFiles(response['root'], markdown_list)
      this.post_list = markdown_list
      // response['root']["filepath"] = "/"
      return response['root']
    },
    onUnmounted() {
      // console.log("销毁组件！！")
    },
    getFileName(filepath) {
      // 去掉文件名中的路径和文件扩展名
      const parts = filepath.split('/');
      const filename = parts[parts.length - 1];
      return filename.replace(/\.md$/, ''); // 去掉 .md 扩展名
    },

    formatTimestamp(timestamp) {
      timestamp = Number(timestamp) * 1000
      timestamp = timestamp ? timestamp : null;
      let date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
      let Y = date.getFullYear() + '-';
      let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
      let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
      let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
      let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
      let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
      return Y + M + D + h + m + s;
    },
    extractMdFiles(directory, mdFilesList = []) {
      // 处理当前目录下的文件
      mdFilesList.push(...directory['file'].filter(file => file['filepath'].endsWith('.md')));

      // 递归处理子目录
      for (const subDirectory of directory.directory) {
        this.extractMdFiles(subDirectory, mdFilesList);
      }

      return mdFilesList
    },
    navigateToPost(filepath) {
      const post_id = base64_url_encode(filepath);
      this.$router.push(`/post/${post_id}`);
    },
    is_open(filepath) {
      // 如果 open_directory 未定义或为空，返回 false
      if (!this.open_directory) {
        return false;
      }
      // 检查 filepath 是否是在 open_directory 中或是在 open_directory 的任意子目录中
      const result = this.open_directory === filepath || this.open_directory.startsWith(filepath + '/');
      console.log("is_open: ", this.open_directory, ", ", filepath, " : ", result)
      return result
    },
    is_display(filepath) {
      if (this.open_directory === '/') {
        return true
      }
      // 获取 filepath 的父目录
      const parentDirectory = this.getParentDirectory(filepath);

      // 判断直接父目录是否是 this.open_directory 或者 this.open_directory 是否是 filepath 的子目录
      const result = parentDirectory === this.open_directory || filepath.startsWith(this.open_directory + '/');

      // console.log("is_display: ", this.open_directory, ", ", filepath, ": ", result)
      return result
    },
    set_open_directory(filepath) {
      // console.log("set_open_directory:", filepath, " cur:", this.open_directory, ", 父:", this.getParentDirectory(filepath))

      if (filepath === this.open_directory) {
        filepath = this.getParentDirectory(filepath)
      }
      this.open_directory = filepath
    },
    getParentDirectory(filepath) {
      // 使用 split 将路径字符串切割成数组
      const pathArray = filepath.split('/');

      // 如果数组长度小于等于1，表示没有父目录，直接返回原路径
      if (pathArray.length <= 1) {
        return filepath;
      }

      // 取数组的前面部分，即去掉最后一个元素
      const parentArray = pathArray.slice(0, -1);

      // 使用 join 将数组拼接回字符串，再加上斜杠
      const parentDirectory = parentArray.join('/') + '/';

      return parentDirectory;
    }
  },
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
ul {
  list-style: none;
  padding: 0;
}

.file-item {
  border-bottom: 1px solid #ddd;
  /* 分隔线 */
  padding: 10px;
  cursor: pointer;
}

.file-title {
  margin-bottom: 5px;
}

.last-modified {
  color: #666;
  /* 灰色字体 */
  font-size: 12px;
  text-align: right;
}



@media (max-width: 1024px) {
  .container_header {
    display: block;
    width: 90%;
    margin: auto;
  }

  .container_body {
    display: flex;
    flex-direction: row;
    color: #444;
    font-family: Georgia, Palatino, 'Palatino Linotype', Times, 'Times New Roman', serif;
    font-size: 12px;
    line-height: 1.5em;
    padding: 1em;
    margin: auto;
    width: 100%;
    background: #fefefe;
  }

  .left {
    width: 0px;
    flex: 0px;
  }

  .center {
    flex: 70%;
  }

  .right_placeholder {
    flex: 0px;
  }

  .right {
    display: none;
  }
}

@media (min-width: 1024px) {
  .container_header {
    display: block;
    left: 20%;
    width: 60%;
    margin: auto;
  }

  .container_body {
    display: flex;
    flex-direction: row;
    color: #444;
    font-family: Georgia, Palatino, 'Palatino Linotype', Times, 'Times New Roman', serif;
    font-size: 12px;
    line-height: 1.5em;
    padding: 1em;
    margin: auto;
    width: 100%;
    background: #fefefe;
  }

  .left {
    flex: 20%;
  }

  .center {
    flex: 60%;
  }

  .right_placeholder {
    flex: 0%;
    display: none;
  }

  .right {
    height: 100vh;
    flex: 20%;
  }
}
</style>
