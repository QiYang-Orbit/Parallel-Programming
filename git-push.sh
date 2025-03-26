#!/bin/bash

# 显示当前状态
echo "当前 Git 状态："
git status

# 询问是否要添加所有更改
read -p "是否添加所有更改？(y/n) " add_all
if [ "$add_all" = "y" ]; then
    git add .
    echo "已添加所有更改"
else
    # 显示更改的文件
    echo "更改的文件："
    git status -s
    read -p "请输入要添加的文件（用空格分隔）：" files
    if [ ! -z "$files" ]; then
        git add $files
        echo "已添加选定的文件"
    fi
fi

# 询问提交信息
read -p "请输入提交信息： " commit_msg
if [ ! -z "$commit_msg" ]; then
    git commit -m "$commit_msg"
    echo "已提交更改"
fi

# 询问是否要推送
read -p "是否推送到远程仓库？(y/n) " push
if [ "$push" = "y" ]; then
    git push
    echo "已推送到远程仓库"
fi

echo "完成！" 