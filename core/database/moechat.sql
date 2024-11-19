BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "chat" (
	"id"	VARCHAR(255) NOT NULL,
	"user_id"	VARCHAR(255) NOT NULL,
	"title"	TEXT NOT NULL,
	"share_id"	VARCHAR(255) NOT NULL,
	"archived"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL,
	"updated_at"	DATETIME NOT NULL,
	"chat"	JSON NOT NULL,
	"pinned"	BOOLEAN NOT NULL DEFAULT false,
	"meta"	JSON NOT NULL DEFAULT '{}',
	"folder_id"	TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS "chatidtag" (
	"id"	VARCHAR(255) NOT NULL,
	"tag_name"	VARCHAR(255) NOT NULL,
	"chat_id"	VARCHAR(255) NOT NULL,
	"user_id"	VARCHAR(255) NOT NULL,
	"timestamp"	INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS "config" (
	"id"	INTEGER NOT NULL,
	"data"	JSON NOT NULL,
	"version"	INTEGER NOT NULL,
	"created_at"	DATETIME NOT NULL DEFAULT (CURRENT_TIMESTAMP),
	"updated_at"	DATETIME DEFAULT (CURRENT_TIMESTAMP),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "document" (
	"id"	INTEGER NOT NULL,
	"collection_name"	VARCHAR(255) NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"title"	TEXT NOT NULL,
	"filename"	TEXT NOT NULL,
	"content"	TEXT NOT NULL,
	"user_id"	VARCHAR(255) NOT NULL,
	"timestamp"	INTEGER NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "feedback" (
	"id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"version"	BIGINT NOT NULL,
	"type"	TEXT NOT NULL,
	"data"	JSON NOT NULL,
	"meta"	JSON NOT NULL,
	"snapshot"	JSON NOT NULL,
	"created_at"	BIGINT NOT NULL,
	"updated_at"	BIGINT NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "file" (
	"id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"filename"	TEXT NOT NULL,
	"meta"	JSON NOT NULL,
	"created_at"	INTEGER NOT NULL,
	"hash"	TEXT NOT NULL,
	"data"	JSON NOT NULL,
	"updated_at"	BIGINT NOT NULL,
	"path"	TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS "folder" (
	"id"	TEXT NOT NULL,
	"parent_id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"items"	JSON NOT NULL,
	"meta"	JSON NOT NULL,
	"is_expanded"	BOOLEAN NOT NULL,
	"created_at"	BIGINT NOT NULL,
	"updated_at"	BIGINT NOT NULL,
	PRIMARY KEY("id","user_id")
);
CREATE TABLE IF NOT EXISTS "function" (
	"id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"type"	TEXT NOT NULL,
	"content"	TEXT NOT NULL,
	"meta"	TEXT NOT NULL,
	"created_at"	INTEGER NOT NULL,
	"updated_at"	INTEGER NOT NULL,
	"valves"	TEXT NOT NULL,
	"is_active"	INTEGER NOT NULL,
	"is_global"	INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS "knowledge" (
	"id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"description"	TEXT NOT NULL,
	"data"	JSON NOT NULL,
	"meta"	JSON NOT NULL,
	"created_at"	BIGINT NOT NULL,
	"updated_at"	BIGINT NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "memory" (
	"id"	VARCHAR(255) NOT NULL,
	"user_id"	VARCHAR(255) NOT NULL,
	"content"	TEXT NOT NULL,
	"updated_at"	INTEGER NOT NULL,
	"created_at"	INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS "migratehistory" (
	"id"	INTEGER NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"migrated_at"	DATETIME NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "model" (
	"id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"base_model_id"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"meta"	TEXT NOT NULL,
	"params"	TEXT NOT NULL,
	"created_at"	INTEGER NOT NULL,
	"updated_at"	INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS "prompt" (
	"id"	INTEGER NOT NULL,
	"command"	VARCHAR(255) NOT NULL,
	"user_id"	VARCHAR(255) NOT NULL,
	"title"	TEXT NOT NULL,
	"content"	TEXT NOT NULL,
	"timestamp"	INTEGER NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "tag" (
	"id"	VARCHAR(255) NOT NULL,
	"name"	VARCHAR(255) NOT NULL,
	"user_id"	VARCHAR(255) NOT NULL,
	"meta"	JSON NOT NULL,
	CONSTRAINT "pk_id_user_id" PRIMARY KEY("id","user_id")
);
CREATE TABLE IF NOT EXISTS "tool" (
	"id"	TEXT NOT NULL,
	"user_id"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"content"	TEXT NOT NULL,
	"specs"	TEXT NOT NULL,
	"meta"	TEXT NOT NULL,
	"created_at"	INTEGER NOT NULL,
	"updated_at"	INTEGER NOT NULL,
	"valves"	TEXT
);
CREATE TABLE IF NOT EXISTS "user" (
    "email" VARCHAR(255) NOT NULL,
    "password" BLOB NOT NULL,  -- 使用 BLOB 类型存储密码的二进制数据
    "level" TEXT NOT NULL,   -- 用户级别
    "profile_image_url" TEXT NOT NULL,
    "created_at" INTEGER NOT NULL,
    "updated_at" INTEGER NOT NULL,
    "settings" TEXT NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS "chat_id" ON "chat" (
	"id"
);
CREATE UNIQUE INDEX IF NOT EXISTS "chat_share_id" ON "chat" (
	"share_id"
);
CREATE UNIQUE INDEX IF NOT EXISTS "chatidtag_id" ON "chatidtag" (
	"id"
);
CREATE UNIQUE INDEX IF NOT EXISTS "document_collection_name" ON "document" (
	"collection_name"
);
CREATE UNIQUE INDEX IF NOT EXISTS "document_name" ON "document" (
	"name"
);
CREATE UNIQUE INDEX IF NOT EXISTS "file_id" ON "file" (
	"id"
);
CREATE UNIQUE INDEX IF NOT EXISTS "function_id" ON "function" (
	"id"
);
CREATE UNIQUE INDEX IF NOT EXISTS "memory_id" ON "memory" (
	"id"
);
CREATE UNIQUE INDEX IF NOT EXISTS "model_id" ON "model" (
	"id"
);
CREATE UNIQUE INDEX IF NOT EXISTS "prompt_command" ON "prompt" (
	"command"
);
CREATE UNIQUE INDEX IF NOT EXISTS "tool_id" ON "tool" (
	"id"
);
COMMIT;
