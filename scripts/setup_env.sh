#!/bin/zsh

# 设置 SONYFLAKE_MACHINE_ID 环境变量为 "12345"
export SONYFLAKE_MACHINE_ID="12345"

# 定义要添加或更新的导出行
EXPORT_LINE='export SONYFLAKE_MACHINE_ID="12345"'

# 定义配置文件路径
CONFIG_FILE="$HOME/.zshrc"

# 检查配置文件中是否已经存在 SONYFLAKE_MACHINE_ID
if grep -q '^export SONYFLAKE_MACHINE_ID=' "$CONFIG_FILE"; then
    # 使用 sed 替换现有的导出行
    sed -i '' 's/^export SONYFLAKE_MACHINE_ID=.*/export SONYFLAKE_MACHINE_ID="12345"/' "$CONFIG_FILE"
    echo "已更新 $CONFIG_FILE 中的 SONYFLAKE_MACHINE_ID。"
else
    # 如果不存在，则追加导出行
    echo "$EXPORT_LINE" >> "$CONFIG_FILE"
    echo "已将 SONYFLAKE_MACHINE_ID 添加到 $CONFIG_FILE。"
fi

# 使用 ShellCheck 指令指定配置文件路径，帮助 ShellCheck 解析
# shellcheck source=/Users/sybenc/.zshrc
source "$CONFIG_FILE"

echo "SONYFLAKE_MACHINE_ID 已设置为 12345 并应用到当前 Shell 会话中。"