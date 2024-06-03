
# EmptyFolderCheck

EmptyFolderCheck 是一个命令行工具，用于扫描指定目录（默认为当前目录）中的空文件夹，并允许用户选择删除这些空文件夹。

## 功能

- 扫描指定目录及其子目录中的空文件夹
- 列出所有找到的空文件夹，并按序号编号
- 允许用户通过输入序号选择删除空文件夹，支持单个、多个以及区间的序号格式

## 安装

1. 确保已安装 [Go 语言](https://golang.org/doc/install) 环境。
2. 克隆或下载此仓库到本地。
3. 在项目根目录下运行以下命令以编译工具：

   ```sh
   go build -o EmptyFolderCheck
   ```

## 使用方法

### 基本用法

在终端运行以下命令，扫描当前目录中的空文件夹：

```sh
./EmptyFolderCheck
```

### 指定目录

也可以通过传递一个目录参数来扫描特定目录：

```sh
./EmptyFolderCheck /path/to/directory
```

### 删除空文件夹

工具会列出所有找到的空文件夹，并按序号编号。输入需要删除的空文件夹的序号，可以使用以下格式：

- 单个序号：例如 `1`
- 多个序号：用逗号分隔，例如 `1,3,5`
- 区间序号：用短横线表示区间，例如 `2-4`

例如：

```sh
Enter the numbers of directories to delete (e.g., 1,2,5-7):
1,3,5-7
```

## 示例

假设当前目录结构如下：

```
.
├── dir1
├── dir2
│   └── subdir1
├── dir3
└── dir4
```

其中 `dir1` 和 `dir3` 是空文件夹。

运行工具：

```sh
./EmptyFolderCheck
```

输出：

```
1: /path/to/current/directory/dir1
2: /path/to/current/directory/dir3
Enter the numbers of directories to delete (e.g., 1,2,5-7):
```

输入 `1,2` 并按下回车键，工具会删除 `dir1` 和 `dir3`。
