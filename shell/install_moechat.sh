#!/bin/sh

# 检查是否存在 wget
if ! command -v wget &> /dev/null; then
    echo "wget 未安装，请先安装 wget。"
    exit 1
fi

# 检查操作系统
OS=$(uname -s)

# 检查系统架构
ARCH=$(uname -m)

# 映射系统架构到 Go 架构
case "$ARCH" in
    x86_64)
        GOARCH="amd64"
        ;;
    arm64|aarch64)
        GOARCH="arm64"
        ;;
    armv7l)
        GOARCH="arm"
        ;;
    mips|mipsel)
        GOARCH="mips"
        ;;
    s390x)
        GOARCH="s390x"
        ;;
    riscv64)
        GOARCH="riscv64"
        ;;
    *)
        echo "不支持的系统架构: $ARCH，请手动下载适用于您系统的 moe-chat 版本。"
        exit 1
        ;;
esac

# 将操作系统名称转换为小写
OS_LOWER=$(echo "$OS" | tr '[:upper:]' '[:lower:]')

# 检查是否支持的操作系统
if [ "$OS" != "Linux" ] && [ "$OS" != "FreeBSD" ] && [ "$OS" != "Darwin" ]; then
    echo "不支持的操作系统: $OS，请手动下载适用于您系统的 moe-chat 版本。"
    exit 1
fi

# 设置下载链接
DOWNLOAD_URL="https://github.com/BapiGso/moe-chat/releases/latest/download/moe-chat_${OS_LOWER}_${GOARCH}"

# 下载 moe-chat
sudo wget "$DOWNLOAD_URL" -O /usr/local/bin/moe-chat

# 检查下载是否成功
if [ $? -ne 0 ]; then
    echo "下载 moe-chat 失败，请检查网络连接或手动下载。"
    exit 1
fi

# 赋予执行权限
sudo chmod +x /usr/local/bin/moe-chat

# 创建工作目录
WORKDIR="/opt/moe-chat"
sudo mkdir -p "$WORKDIR"

# 创建服务文件
if [ "$OS" = "Linux" ]; then
    cat << EOF | sudo tee /etc/systemd/system/moe-chat.service > /dev/null
[Unit]
Description=GoPanel Service
After=network.target

[Service]
Type=simple
User=root
ExecStart=/usr/local/bin/moe-chat -w ${WORKDIR}
Restart=on-failure
WorkingDirectory=${WORKDIR}

[Install]
WantedBy=multi-user.target
EOF

    # 重新加载 systemd 配置
    sudo systemctl daemon-reload

    # 启用并启动 moe-chat 服务
    sudo systemctl enable moe-chat
    sudo systemctl start moe-chat

    # 检查服务状态
    sleep 2
    sudo systemctl status moe-chat
elif [ "$OS" = "FreeBSD" ]; then
    cat << EOF | sudo tee /usr/local/etc/rc.d/moe-chat > /dev/null
#!/bin/sh

# PROVIDE: moe-chat
# REQUIRE: DAEMON NETWORKING
# KEYWORD: shutdown

. /etc/rc.subr

name="moe-chat"
rcvar="\${name}_enable"

load_rc_config \${name}

: \${moe-chat_enable:="NO"}
: \${moe-chat_user:="root"}

command="/usr/local/bin/moe-chat"
command_args="-w ${WORKDIR}"

run_rc_command "\$1"
EOF

    # 赋予 rc.d 脚本执行权限
    sudo chmod +x /usr/local/etc/rc.d/moe-chat

    # 启用并启动 moe-chat 服务
    sudo sysrc moe-chat_enable=YES
    sudo service moe-chat start

    # 检查服务状态
    sleep 2
    sudo service moe-chat status
elif [ "$OS" = "Darwin" ]; then
    # 创建 LaunchDaemon
    PLIST_PATH="/Library/LaunchDaemons/com.moe-chat.service.plist"
    cat << EOF | sudo tee "$PLIST_PATH" > /dev/null
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.moe-chat.service</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/moe-chat</string>
        <string>-w</string>
        <string>${WORKDIR}</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardErrorPath</key>
    <string>/var/log/moe-chat.err</string>
    <key>StandardOutPath</key>
    <string>/var/log/moe-chat.out</string>
    <key>WorkingDirectory</key>
    <string>${WORKDIR}</string>
</dict>
</plist>
EOF

    # 设置权限
    sudo chown root:wheel "$PLIST_PATH"
    sudo chmod 644 "$PLIST_PATH"

    # 加载并启动服务
    sudo launchctl load -w "$PLIST_PATH"

    # 检查服务状态
    sleep 2
    sudo launchctl print system/com.moe-chat.service
else
    echo "不支持的操作系统: $OS，无法启动服务。"
    exit 1
fi