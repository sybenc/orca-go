#!/bin/zsh

# 定义文本颜色变量
red='\033[31m'
green='\033[32m'
yellow='\033[33m'
blue='\033[34m'
purple='\033[35m'
cyan='\033[36m'
reset='\033[0m'  # 重置颜色
# 定义文本样式变量
bold='\033[1m'
underline='\033[4m'

# 数据库连接信息
db_user="root"
db_name="test"
db_host="localhost"
prefix="[Migrate the database]"

echo -e "${prefix}用户名：${green}${bold}$db_user${reset}"
echo -e "${prefix}数据库名称：${green}${bold}$db_name${reset}"
echo -e "${prefix}数据库主机地址：${green}${bold}$db_host${reset}"

sql_files=(
  './scripts/sql/users.sql'
  './scripts/sql/user_auth.sql'
  './scripts/sql/menu.sql'
  './scripts/sql/roles.sql'
  './scripts/sql/role_menu.sql'
  './scripts/sql/user_role.sql'
)

echo -e "${yellow}${prefix}正在检查数据库 ${green}${db_name}${yellow} 是否存在...${reset}"

# 提示用户输入密码
echo -e "${yellow}${prefix}请输入数据库密码（密码输入过程中不显示）:${reset}"
read -rsp "" db_password

if mysql -u "$db_user" -p"$db_password" -h "$db_host" -e "use $db_name"; then
    echo -e "数据库 ${green}${bold}$db_name${reset} 已存在，继续执行SQL脚本..."
else
    echo -e "${yellow}${bold}数据库 $db_name 不存在，正在创建...${reset}"
    if mysql -u "$db_user" -p"$db_password" -h "$db_host" -e "CREATE DATABASE $db_name DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci"; then
        echo -e "${green}数据库 $db_name 创建成功！${reset}"
    else
        echo -e "${red}${bold}创建数据库 $db_name 失败！退出...${reset}"
        exit 1
    fi
fi

for sql_file in "${sql_files[@]}"
do
  if mysql -u "$db_user" -p"$db_password" -h "$db_host" "$db_name" < "$sql_file"; then
        echo -e "${prefix}${green}$sql_file${reset} 执行成功！"
  else
        echo -e "${prefix}${green}$sql_file${reset} 执行失败！"
        exit 1  # 如果其中一个脚本失败，退出并显示错误
  fi
done

echo -e "${prefix}所有 SQL 脚本执行完毕！"