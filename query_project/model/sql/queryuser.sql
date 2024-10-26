CREATE TABLE queryuser (
    `id` bigint Not NULL AUTO_INCREMENT,
    `username` varchar(255) Not NULL DEFAULT 1 Comment "登录用户名",
    `passwd` varchar(255) Not NULL DEFAULT "" Comment "登录用户密码",
    `create_time` datetime Not NULL DEFAULT CURRENT_TIMESTAMP Comment "创建时间",
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 Comment '用户登录表';