#!/bin/sh

# 检查操作系统
OS=$(uname -s)

# 停止并禁用 moe-chat 服务
if [ "$OS" = "Linux" ]; then
    sudo systemctl stop moe-chat
    sudo systemctl disable moe-chat
    sudo rm /etc/systemd/system/moe-chat.service
    sudo systemctl daemon-reload
elif [ "$OS" = "FreeBSD" ]; then
    sudo service moe-chat stop
    sudo sysrc moe-chat_enable=NO
    sudo rm /usr/local/etc/rc.d/moe-chat
elif [ "$OS" = "Darwin" ]; then
    PLIST_PATH="/Library/LaunchDaemons/com.moe-chat.service.plist"
    if [ -f "$PLIST_PATH" ]; then
        sudo launchctl unload -w "$PLIST_PATH"
        sudo rm "$PLIST_PATH"
    else
        echo "找不到 LaunchDaemon 配置文件：$PLIST_PATH"
    fi
else
    echo "不支持的操作系统: $OS，无法停止服务。"
    exit 1
fi

# 删除 moe-chat 二进制文件
sudo rm /usr/local/bin/moe-chat



echo "moe-chat 已成功卸载。"