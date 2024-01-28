<template>
    <div class="top-navbar">
        <!-- Your navigation content goes here -->
        <ul>
            <li @click="navigator_home()">
                首页
            </li>
        </ul>
        <ul>
            <li>
                密码
            </li>
            <li>
                <input type="password" v-model="authority_password" @input="update_password">
            </li>
        </ul>
    </div>
</template>
  
<script>
export default {
    name: 'NavBar',
    props: {
        item: Object,
        set_open_directory: {
            type: Function,
            required: false,
        },
    },
    data() {
        return {
            authority_password: localStorage.getItem('authority_password')
        }
    },
    methods: {
        navigator_home(path) {
            if (this.$route.path === '/') {
                this.set_open_directory('/')
            } else {
                this.$router.push(`/`);
            }
        },
        update_password() {
            localStorage.setItem('authority_password', this.authority_password);
            // console.log("authority_password", localStorage.getItem('authority_password'))
        }
    }
}
</script>
  
<style scoped>
.top-navbar {
    padding: 10px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

ul {
    list-style: none;
    font-family: Georgia, Palatino, 'Palatino Linotype', Times, 'Times New Roman', serif;
    font-size: 20px;
    list-style: none;
    display: flex;
    padding: 0;
    cursor: pointer;
}

li {
    margin: 0 10px;
}
</style>