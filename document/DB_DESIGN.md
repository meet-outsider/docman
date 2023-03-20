### user 表

| 字段名    | 数据类型 | 描述                      |
| --------- | -------- | -------------------------|
| id        | INT      | 用户ID，主键            |
| username  | VARCHAR  | 用户名                    |
| password  | VARCHAR  | 密码                      |
| email     | VARCHAR  | 电子邮件                  |
| created_at| TIMESTAMP| 创建时间                  |
| updated_at| TIMESTAMP| 更新时间                  |

### role 表

| 字段名    | 数据类型 | 描述                      |
| --------- | -------- | -------------------------|
| id        | INT      | 角色ID，主键            |
| name      | VARCHAR  | 角色名称                  |
| created_at| TIMESTAMP| 创建时间                  |
| updated_at| TIMESTAMP| 更新时间                  |

### role_permission 表

| 字段名      | 数据类型 | 描述                     |
| ----------- | -------- | ------------------------|
| id          | INT      | ID，主键                |
| role_id     | INT      | 角色ID，外键，关联role表|
| path        | VARCHAR  | 文件路径                |
| type        | VARCHAR  | 文件类型                |
| created_at  | TIMESTAMP| 创建时间                |
| updated_at  | TIMESTAMP| 更新时间                |

### user_role 表

| 字段名      | 数据类型 | 描述                     |
| ----------- | -------- | ------------------------|
| id          | INT      | ID，主键                |
| user_id     | INT      | 用户ID，外键，关联user表|
| role_id     | INT      | 角色ID，外键，关联role表|
| created_at  | TIMESTAMP| 创建时间                |
| updated_at  | TIMESTAMP| 更新时间                |

### casbin_rule 表

| 字段名     | 数据类型 | 描述                                             |
| ---------- | -------- | ------------------------------------------------|
| id         | INT      | ID，主键                                        |
| ptype      | VARCHAR  | 权限类型，例如 "p" 表示主题策略，"g" 表示组策略|
| v0         | VARCHAR  | 规则中的第一个字段                               |
| v1         | VARCHAR  | 规则中的第二个字段                               |
| v2         | VARCHAR  | 规则中的第三个字段                               |
| v3         | VARCHAR  | 规则中的第四个字段                               |
| v4         | VARCHAR  | 规则中的第五个字段                               |
| v5         | VARCHAR  | 规则中的第六个字段                               |
