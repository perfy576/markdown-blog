<template>
    <div class="container_body">
        <li>
            <div v-if="item['filepath'] !== '/'" @click="toggle(item['filepath'])">
                {{ get_last_directory_name(item['filepath']) }} ({{ get_file_count(item) }})
            </div>
            <ul v-if="isFolder">
                <Directory class="item" v-for="(child, index) in item['directory']" :key="index" :item="child"
                    :set_open_directory='set_open_directory'></Directory>
            </ul>
        </li>
    </div>
</template>
  
<script>
export default {
    name: "Directory",
    template: "#item-template",
    props: {
        item: Object,
        set_open_directory: {
            type: Function,
            required: true,
        },
    },
    data: function () {
        return {
            isOpen: true
        };
    },
    computed: {
        isFolder: function () {
            return this.item['directory'] && this.item['directory'].length;
        }
    },
    methods: {
        toggle: function (filepath) {
            if (this.isFolder) {
                this.isOpen = !this.isOpen;
            }
            this.set_open_directory(filepath)
        },
        get_last_directory_name(filepath) {
            if (filepath) {
                // console.log("get_last_directory_name:", filepath)
                // 去掉文件名中的路径和文件扩展名
                const parts = filepath.split('/');
                const filename = parts[parts.length - 1];
                return filename
            } else {
                return ""
            }

        },
        get_file_count(item) {

            let counter = 0

            if (item['file'] === undefined) {
                counter = 0
            } else {
                counter = item['file'].length
            }

            for (let index in item['directory']) {
                counter += this.get_file_count(item['directory'][index])
            }

            return counter
        }
    }
}
</script>
  
<style scoped>
.container_body {
    justify-content: center;
    font-size: 12px;
}


.item {
    cursor: pointer;
}

.bold {
    font-weight: bold;
}

ul {
    padding-left: 1em;
    line-height: 1.5em;
    list-style-type: dot;
}
</style>