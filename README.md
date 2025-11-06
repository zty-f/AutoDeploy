# AutoDeploy

#### 后台运行命令
  ```
  go build -o ./deployScript
  nohup ./deployScript &
  ```

### 服务器部署

```shell
# 如果架构不匹配，需要重新编译：
# 在本地 Mac 上编译 Linux 版本
GOOS=linux GOARCH=amd64 go build -o deploy

# 1. 查看当前服务文件内容
cat /etc/systemd/system/deploy.service

# 2. 编辑服务文件（修复配置）
sudo nano /etc/systemd/system/deploy.service

服务配置如下：
[Unit]
Description=Deploy Service
After=network.target

[Service]
Type=simple
User=root
WorkingDir=/tmp  # 这个是工作目录，需要把代码所需的执行文件放在对应的路径里面
ExecStart=/usr/local/bin/deploy
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target

# 重新加载并启动
sudo systemctl daemon-reload
sudo systemctl restart deploy
sudo systemctl status deploy

# 启动服务
sudo systemctl start deploy

# 停止服务
sudo systemctl stop deploy

# 重启服务
sudo systemctl restart deploy

# 查看服务状态
sudo systemctl status deploy

# 查看实时日志
sudo journalctl -u deploy -f

# 查看最近日志
sudo journalctl -u deploy -n 50

# 设置开机自启（重要！）
sudo systemctl enable deploy

# 取消开机自启
sudo systemctl disable deploy
```