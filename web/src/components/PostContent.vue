<template>
    <div>
        <div class="container_header">
            <NavBar />
        </div>

        <div class="container_body">
            <div class="left"></div>

            <div class="center">
                <vue-markdown :source="content" class="markdown" :toc="true" v-on:toc-rendered="tocAllRight" toc-id="toc"
                    :toc-anchor-link="false" :toc-last-level=4 :toc-first-level=1></vue-markdown>
            </div>

            <div class="right_placeholder"></div>

            <div class="right">
                <div id="toc"></div>
            </div>
        </div>
    </div>
</template>
  
<script>
// import PostListItem from './PostListItem.vue';
import { post } from '@/utils/net.js'
import { base64_encode, base64_decode } from '@/utils/crypto.js'
import VueMarkdown from 'vue-markdown'
import NavBar from './NavBar'
// import Prism from 'prismjs';
export default {
    name: 'HelloWorld',
    data() {
        return {
            filepath: {},
            content: "",
            content1: "",
        }
    },
    components: { VueMarkdown, NavBar },
    async created() {
        this.content = await this.fetch_file_content()
        this.content1 = '```go\n int main(){\n};'
        this.$nextTick(() => {
            this.loadScript();
        });
    },
    onMounted() {
        console.log("$route.params.post_id:", this.$route.params.post_id)
    },
    methods: {
        async fetch_file_content() {

            var data = {
                'filepath': base64_decode(this.$route.params.post_id),
            }
            const response = await post('/blog/v1/get-markdown-content', data)
            this.content = base64_decode(response['content'])
            return base64_decode(response['content'])



        },
        gen_title_toc_link() {
            return this.$route.params.post_id + "#"
        },
        render_markdown() {
            return `<vue-markdown> this.content </vue-markdown>`
        },
        gen_title_name() {
            const filepath = base64_decode(this.$route.params.post_id)
            // 去掉文件名中的路径和文件扩展名
            const parts = filepath.split('/');
            const filename = parts[parts.length - 1];
            this.filepath = filename
            return filename.replace(/\.md$/, ''); // 去掉 .md 扩展名
        },
        loadScript() {
            this.toc_location()
            let Prism = window.Prism || {};
            Prism.highlightAll();
        },
        tocAllRight: function (tocHtmlStr) {
            // console.log("toc is parsed :", tocHtmlStr);
        },
        allRight: function (htmlStr) {
            // console.log("markdown is parsed !:", htmlStr);
        },
        toc_location() {
            var toc = document.getElementById('toc');
            if (toc) {
                var windowHeight = window.innerHeight;
                var tocHeight = toc.offsetHeight;
                var scrollPosition = window.scrollY;

                var topOffset = (windowHeight - tocHeight) / 2;
                var newPosition = scrollPosition + topOffset;

                toc.style.position = 'fixed';
                toc.style.top = '50%';
                toc.style.transform = 'translateY(-50%)';
            }
        }
    },
}
</script>
  <!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
@media (max-width: 1024px) {
    .container_body {
        display: flex;
        flex-direction: row;
        color: #444;
        font-family: Georgia, Palatino, 'Palatino Linotype', Times, 'Times New Roman', serif;
        font-size: 14px;
        line-height: 1.5em;
        padding: 1em;
        margin: auto;
        width: 100%;
        background: #fefefe;
    }

    .container_header {
        display: block;
        width: 90%;
        margin: auto;
    }

    .left {
        flex: 5%;
    }

    .center {
        display: block;
        width: 90%;
        margin: 0;
        padding: 0;
    }

    .right_placeholder {
        flex: 5%;
    }

    .right {
        flex: 0%;
        display: none;
    }

    #toc {
        display: none;
    }
}

@media (min-width: 1024px) {
    .container_body {
        display: flex;
        flex-direction: row;
        color: #444;
        font-family: Georgia, Palatino, 'Palatino Linotype', Times, 'Times New Roman', serif;
        font-size: 14px;
        line-height: 1.5em;
        padding: 1em;
        margin: auto;
        width: 100%;
        background: #fefefe;
    }

    .container_header {
        display: block;
        left: 20%;
        width: 60%;
        margin: auto;
    }

    .left {
        flex: 20%;
    }

    .center {
        display: block;
        flex: 70%;
        margin: 0;
        padding: 0;
    }

    .right_placeholder {
        display: none;
    }

    .right {
        flex: 20%;
    }
}
</style>
  