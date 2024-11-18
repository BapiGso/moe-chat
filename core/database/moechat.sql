CREATE TABLE "chat" (
                        "id"	VARCHAR(255) NOT NULL,
                        "user_id"	VARCHAR(255) NOT NULL,
                        "title"	TEXT NOT NULL,
                        "chat"	TEXT NOT NULL,
                        "share_id"	VARCHAR(255),
                        "archived"	INTEGER NOT NULL,
                        "created_at"	DATETIME NOT NULL,
                        "updated_at"	DATETIME NOT NULL
);

CREATE TABLE "file" (
                        "id"	TEXT NOT NULL,
                        "user_id"	TEXT NOT NULL,
                        "filename"	TEXT NOT NULL,
                        "meta"	TEXT NOT NULL,
                        "created_at"	INTEGER NOT NULL
);

CREATE TABLE "user" (
                        "id"	VARCHAR(255) NOT NULL,
                        "name"	VARCHAR(255) NOT NULL,
                        "email"	VARCHAR(255) NOT NULL,
                        "role"	VARCHAR(255) NOT NULL,
                        "profile_image_url"	TEXT NOT NULL,
                        "api_key"	VARCHAR(255),
                        "created_at"	INTEGER NOT NULL,
                        "updated_at"	INTEGER NOT NULL,
                        "last_active_at"	INTEGER NOT NULL,
                        "settings"	TEXT,
                        "info"	TEXT,
                        "oauth_sub"	TEXT
);