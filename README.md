# markdown-blog
> 简易的 markdown 笔记/博客

使用 go + vue 实现，将指定目录的 md 文件渲染为html并展示。

带有简单的权限控制。


## 使用

### server
server/config/server.toml 
- chroot 指定了md文件存放的根目录
- authority 为访问密码，密码正确是，先是所有的md文件和目录
- server 配置组下修改server监听ip和port

### nginx 

修改conf目录的conf文件
- 修改域名
- `proxy_pass http://127.0.0.1:10000;` 到server监听的端口
- 修改前端文件部署目录 `root /data/code/markdown-blog/web/dist;`

然后配置到nginx中。

### 权限控制
md 所在目录中，新建 `authority.json` 文件
```json
{
    "show_subdir": false,
    "show_post": false
}
```

可以分别控制该目录的子目录以及该目录下的md文件是否被展示到首页中。但并不影响直接通过链接访问md文件。