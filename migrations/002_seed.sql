-- WVP-GB28181 Go版本初始数据
-- 插入默认管理员用户（密码: admin123）
INSERT INTO wvp_user (username, password, nickname, status) 
VALUES ('admin', '$2a$10$your_hashed_password_here', '管理员', 1);

-- 插入默认角色
INSERT INTO wvp_role (name, description) VALUES 
('admin', '系统管理员'),
('operator', '操作员'),
('viewer', '查看者');

-- 分配管理员角色
INSERT INTO wvp_user_role (user_id, role_id) 
SELECT u.id, r.id FROM wvp_user u, wvp_role r 
WHERE u.username = 'admin' AND r.name = 'admin';

-- 插入默认权限
INSERT INTO wvp_permission (name, path, method, description) VALUES 
('设备查看', '/api/device', 'GET', '查看设备列表'),
('设备管理', '/api/device', 'POST', '添加设备'),
('设备编辑', '/api/device/:id', 'PUT', '编辑设备'),
('设备删除', '/api/device/:id', 'DELETE', '删除设备'),
('视频点播', '/api/video/play', 'POST', '开始视频点播'),
('视频停止', '/api/video/stop', 'POST', '停止视频点播'),
('云台控制', '/api/video/ptz', 'POST', '云台控制'),
('录像查看', '/api/record', 'GET', '查看录像列表'),
('平台管理', '/api/platform', 'GET', '查看平台列表'),
('系统设置', '/api/system', 'GET', '查看系统设置');

-- 分配管理员权限
INSERT INTO wvp_role_permission (role_id, permission_id) 
SELECT r.id, p.id FROM wvp_role r, wvp_permission p 
WHERE r.name = 'admin';