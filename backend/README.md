# DocVault

自托管的云文档管理系统后端

## 特性

- ✅ 单用户、无数据库
- ✅ Markdown 文档 CRUD
- ✅ RESTful API
- ✅ 路径穿越防护
- ✅ Docker 支持

## 快速开始

### 本地运行

```bash
go mod download
go run cmd/server/main.go
```

### Docker 运行

```bash
docker-compose up -d
```

## API 文档

### 创建文档
```http
POST /api/documents
Content-Type: application/json

{
  "name": "hello.md",
  "content": "# Hello World"
}
```

### 获取文档
```http
GET /api/documents/hello.md
```

### 更新文档
```http
PUT /api/documents/hello.md
Content-Type: application/json

{
  "content": "# Updated Content"
}
```

### 删除文档
```http
DELETE /api/documents/hello.md
```

### 列出所有文档
```http
GET /api/documents
```

## 配置

环境变量：
- `PORT`: 服务端口（默认 8080）
- `STORAGE_DIR`: 存储目录（默认 /data/md）

## 架构

```
handler  →  service  →  fs
  ↓          ↓          ↓
HTTP      业务逻辑   文件系统
```

## 未来扩展

- [ ] 版本历史
- [ ] 目录支持
- [ ] Git/S3 存储
- [ ] 多文件类型