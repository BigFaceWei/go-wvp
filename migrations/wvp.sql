/*
 Navicat Premium Dump SQL

 Source Server         : 192.168.12.130
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 192.168.12.130:3306
 Source Schema         : wvp

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 11/07/2026 19:43:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 370 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (314, 'p', '888', '/alarm/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (353, 'p', '888', '/alarm/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (312, 'p', '888', '/alarm/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (351, 'p', '888', '/alarm/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (313, 'p', '888', '/alarm/:id/handle', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (352, 'p', '888', '/alarm/:id/handle', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (310, 'p', '888', '/alarm/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (349, 'p', '888', '/alarm/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (311, 'p', '888', '/alarm/statistics', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (350, 'p', '888', '/alarm/statistics', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (9, 'p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (12, 'p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', '888', '/api/enterSyncApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'p', '888', '/api/getApiGroups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'p', '888', '/api/getApiRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', '888', '/api/ignoreApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'p', '888', '/api/setApiRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', '888', '/api/syncApi', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (13, 'p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (183, 'p', '888', '/attachmentCategory/addCategory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (184, 'p', '888', '/attachmentCategory/deleteCategory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (182, 'p', '888', '/attachmentCategory/getCategoryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (25, 'p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (28, 'p', '888', '/authority/getUsersByAuthority', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (27, 'p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (29, 'p', '888', '/authority/setRoleUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (151, 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (150, 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (149, 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (109, 'p', '888', '/autoCode/addFunc', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (100, 'p', '888', '/autoCode/createPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (104, 'p', '888', '/autoCode/createPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (97, 'p', '888', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (120, 'p', '888', '/autoCode/deleteAIWorkflowSession', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (103, 'p', '888', '/autoCode/delPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (98, 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (121, 'p', '888', '/autoCode/dumpAIWorkflowMarkdown', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (119, 'p', '888', '/autoCode/getAIWorkflowSessionDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (118, 'p', '888', '/autoCode/getAIWorkflowSessionList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (95, 'p', '888', '/autoCode/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (91, 'p', '888', '/autoCode/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (92, 'p', '888', '/autoCode/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (102, 'p', '888', '/autoCode/getPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (108, 'p', '888', '/autoCode/getPluginList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (99, 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (94, 'p', '888', '/autoCode/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (101, 'p', '888', '/autoCode/getTemplates', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (105, 'p', '888', '/autoCode/installPlugin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (110, 'p', '888', '/autoCode/mcp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', '888', '/autoCode/mcpList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (114, 'p', '888', '/autoCode/mcpRoutes', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (112, 'p', '888', '/autoCode/mcpStart', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (111, 'p', '888', '/autoCode/mcpStatus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (113, 'p', '888', '/autoCode/mcpStop', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', '888', '/autoCode/mcpTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (93, 'p', '888', '/autoCode/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (106, 'p', '888', '/autoCode/pubPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (107, 'p', '888', '/autoCode/removePlugin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (96, 'p', '888', '/autoCode/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', '888', '/autoCode/saveAIWorkflowSession', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (61, 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (60, 'p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (89, 'p', '888', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (86, 'p', '888', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (88, 'p', '888', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (87, 'p', '888', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (90, 'p', '888', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (294, 'p', '888', '/device', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (333, 'p', '888', '/device', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (296, 'p', '888', '/device/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (335, 'p', '888', '/device/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (293, 'p', '888', '/device/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (332, 'p', '888', '/device/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (295, 'p', '888', '/device/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (334, 'p', '888', '/device/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (298, 'p', '888', '/device/:id/catalog', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (337, 'p', '888', '/device/:id/catalog', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (297, 'p', '888', '/device/:id/channels', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (336, 'p', '888', '/device/:id/channels', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (299, 'p', '888', '/device/:id/status', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (338, 'p', '888', '/device/:id/status', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (292, 'p', '888', '/device/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (331, 'p', '888', '/device/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (144, 'p', '888', '/email/emailTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (145, 'p', '888', '/email/sendEmail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (53, 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (52, 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (56, 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (57, 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (51, 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (58, 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (59, 'p', '888', '/fileUploadAndDownload/importURL', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (54, 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (55, 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (169, 'p', '888', '/info/createInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (170, 'p', '888', '/info/deleteInfo', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (171, 'p', '888', '/info/deleteInfoByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (173, 'p', '888', '/info/findInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (174, 'p', '888', '/info/getInfoList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (172, 'p', '888', '/info/updateInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (62, 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (32, 'p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (34, 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (33, 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (30, 'p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (35, 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (31, 'p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (36, 'p', '888', '/menu/getMenuRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (37, 'p', '888', '/menu/setMenuRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (39, 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (303, 'p', '888', '/platform', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (342, 'p', '888', '/platform', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (305, 'p', '888', '/platform/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (344, 'p', '888', '/platform/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (302, 'p', '888', '/platform/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (341, 'p', '888', '/platform/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (304, 'p', '888', '/platform/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (343, 'p', '888', '/platform/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (306, 'p', '888', '/platform/:id/channels', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (345, 'p', '888', '/platform/:id/channels', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (308, 'p', '888', '/platform/:id/start', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (347, 'p', '888', '/platform/:id/start', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (309, 'p', '888', '/platform/:id/stop', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (348, 'p', '888', '/platform/:id/stop', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (307, 'p', '888', '/platform/:id/sync', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (346, 'p', '888', '/platform/:id/sync', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (300, 'p', '888', '/platform/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (339, 'p', '888', '/platform/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (301, 'p', '888', '/platform/online', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (340, 'p', '888', '/platform/online', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (318, 'p', '888', '/position/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (357, 'p', '888', '/position/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (316, 'p', '888', '/position/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (355, 'p', '888', '/position/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (317, 'p', '888', '/position/:id/track', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (356, 'p', '888', '/position/:id/track', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (315, 'p', '888', '/position/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (354, 'p', '888', '/position/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (327, 'p', '888', '/record/file/download/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (366, 'p', '888', '/record/file/download/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (326, 'p', '888', '/record/file/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (365, 'p', '888', '/record/file/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (328, 'p', '888', '/record/file/preview/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (367, 'p', '888', '/record/file/preview/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (321, 'p', '888', '/record/plan', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (360, 'p', '888', '/record/plan', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (323, 'p', '888', '/record/plan/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (362, 'p', '888', '/record/plan/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (320, 'p', '888', '/record/plan/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (359, 'p', '888', '/record/plan/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (322, 'p', '888', '/record/plan/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (361, 'p', '888', '/record/plan/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (324, 'p', '888', '/record/plan/:id/start', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (363, 'p', '888', '/record/plan/:id/start', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (325, 'p', '888', '/record/plan/:id/stop', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (364, 'p', '888', '/record/plan/:id/stop', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (319, 'p', '888', '/record/plan/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (358, 'p', '888', '/record/plan/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (329, 'p', '888', '/record/playback', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (368, 'p', '888', '/record/playback', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (330, 'p', '888', '/record/playback/stop/:id', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (369, 'p', '888', '/record/playback/stop/:id', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (147, 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (148, 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (146, 'p', '888', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (77, 'p', '888', '/skills/createReference', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (74, 'p', '888', '/skills/createResource', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (71, 'p', '888', '/skills/createScript', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (80, 'p', '888', '/skills/createTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (70, 'p', '888', '/skills/deleteSkill', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (83, 'p', '888', '/skills/getGlobalConstraint', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (78, 'p', '888', '/skills/getReference', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (75, 'p', '888', '/skills/getResource', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (72, 'p', '888', '/skills/getScript', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (68, 'p', '888', '/skills/getSkillDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (67, 'p', '888', '/skills/getSkillList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (81, 'p', '888', '/skills/getTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (66, 'p', '888', '/skills/getTools', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (85, 'p', '888', '/skills/packageSkill', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (84, 'p', '888', '/skills/saveGlobalConstraint', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (79, 'p', '888', '/skills/saveReference', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (76, 'p', '888', '/skills/saveResource', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (73, 'p', '888', '/skills/saveScript', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (69, 'p', '888', '/skills/saveSkill', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (82, 'p', '888', '/skills/saveTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', '888', '/sysApiToken/createApiToken', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', '888', '/sysApiToken/deleteApiToken', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', '888', '/sysApiToken/getApiTokenList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (134, 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (135, 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (137, 'p', '888', '/sysDictionary/exportSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (131, 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (133, 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (136, 'p', '888', '/sysDictionary/importSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (132, 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (124, 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (126, 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (122, 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (129, 'p', '888', '/sysDictionaryDetail/getDictionaryDetailsByParent', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (130, 'p', '888', '/sysDictionaryDetail/getDictionaryPath', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (127, 'p', '888', '/sysDictionaryDetail/getDictionaryTreeList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (128, 'p', '888', '/sysDictionaryDetail/getDictionaryTreeListByType', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (125, 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (123, 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (162, 'p', '888', '/sysError/createSysError', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (163, 'p', '888', '/sysError/deleteSysError', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (164, 'p', '888', '/sysError/deleteSysErrorByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (166, 'p', '888', '/sysError/findSysError', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (167, 'p', '888', '/sysError/getSysErrorList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (168, 'p', '888', '/sysError/getSysErrorSolution', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (165, 'p', '888', '/sysError/updateSysError', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (152, 'p', '888', '/sysExportTemplate/createSysExportTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (153, 'p', '888', '/sysExportTemplate/deleteSysExportTemplate', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (154, 'p', '888', '/sysExportTemplate/deleteSysExportTemplateByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (158, 'p', '888', '/sysExportTemplate/exportExcel', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (159, 'p', '888', '/sysExportTemplate/exportTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (156, 'p', '888', '/sysExportTemplate/findSysExportTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (157, 'p', '888', '/sysExportTemplate/getSysExportTemplateList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (161, 'p', '888', '/sysExportTemplate/importExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (160, 'p', '888', '/sysExportTemplate/previewSQL', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (155, 'p', '888', '/sysExportTemplate/updateSysExportTemplate', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'p', '888', '/sysLoginLog/deleteLoginLog', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', '888', '/sysLoginLog/deleteLoginLogByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', '888', '/sysLoginLog/findLoginLog', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', '888', '/sysLoginLog/getLoginLogList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (140, 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (142, 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (143, 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (138, 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (141, 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (139, 'p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (175, 'p', '888', '/sysParams/createSysParams', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (176, 'p', '888', '/sysParams/deleteSysParams', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (177, 'p', '888', '/sysParams/deleteSysParamsByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (179, 'p', '888', '/sysParams/findSysParams', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (181, 'p', '888', '/sysParams/getSysParam', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (180, 'p', '888', '/sysParams/getSysParamsList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (178, 'p', '888', '/sysParams/updateSysParams', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (65, 'p', '888', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (63, 'p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (64, 'p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (190, 'p', '888', '/sysVersion/deleteSysVersion', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (191, 'p', '888', '/sysVersion/deleteSysVersionByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (187, 'p', '888', '/sysVersion/downloadVersionJson', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (188, 'p', '888', '/sysVersion/exportVersion', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (185, 'p', '888', '/sysVersion/findSysVersion', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (186, 'p', '888', '/sysVersion/getSysVersionList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (189, 'p', '888', '/sysVersion/importVersion', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1, 'p', '888', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (46, 'p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (45, 'p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'p', '888', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (44, 'p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (49, 'p', '888', '/user/resetPassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (43, 'p', '888', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (50, 'p', '888', '/user/setSelfSetting', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (48, 'p', '888', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (47, 'p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', '888', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (193, 'p', '8881', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (196, 'p', '8881', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (198, 'p', '8881', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (195, 'p', '8881', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (194, 'p', '8881', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (199, 'p', '8881', '/api/getApiRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (200, 'p', '8881', '/api/setApiRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (197, 'p', '8881', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (201, 'p', '8881', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (202, 'p', '8881', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (203, 'p', '8881', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (205, 'p', '8881', '/authority/getUsersByAuthority', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (204, 'p', '8881', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (206, 'p', '8881', '/authority/setRoleUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (227, 'p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (226, 'p', '8881', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (233, 'p', '8881', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (234, 'p', '8881', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (231, 'p', '8881', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (232, 'p', '8881', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (235, 'p', '8881', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (223, 'p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (224, 'p', '8881', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (222, 'p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (225, 'p', '8881', '/fileUploadAndDownload/importURL', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (221, 'p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (228, 'p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (209, 'p', '8881', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (211, 'p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (215, 'p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (217, 'p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (210, 'p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (207, 'p', '8881', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (212, 'p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (208, 'p', '8881', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (213, 'p', '8881', '/menu/getMenuRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (214, 'p', '8881', '/menu/setMenuRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (216, 'p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (229, 'p', '8881', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (230, 'p', '8881', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (192, 'p', '8881', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (218, 'p', '8881', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (236, 'p', '8881', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (219, 'p', '8881', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (220, 'p', '8881', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (238, 'p', '9528', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (241, 'p', '9528', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (243, 'p', '9528', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (240, 'p', '9528', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (239, 'p', '9528', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (244, 'p', '9528', '/api/getApiRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (245, 'p', '9528', '/api/setApiRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (242, 'p', '9528', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (246, 'p', '9528', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (247, 'p', '9528', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (248, 'p', '9528', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (250, 'p', '9528', '/authority/getUsersByAuthority', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (249, 'p', '9528', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (251, 'p', '9528', '/authority/setRoleUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (281, 'p', '9528', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (289, 'p', '9528', '/autoCode/deleteAIWorkflowSession', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (290, 'p', '9528', '/autoCode/dumpAIWorkflowMarkdown', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (288, 'p', '9528', '/autoCode/getAIWorkflowSessionDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (287, 'p', '9528', '/autoCode/getAIWorkflowSessionList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (285, 'p', '9528', '/autoCode/mcpRoutes', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (283, 'p', '9528', '/autoCode/mcpStart', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (282, 'p', '9528', '/autoCode/mcpStatus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (284, 'p', '9528', '/autoCode/mcpStop', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (286, 'p', '9528', '/autoCode/saveAIWorkflowSession', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (272, 'p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (271, 'p', '9528', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (279, 'p', '9528', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (277, 'p', '9528', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (278, 'p', '9528', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (276, 'p', '9528', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (280, 'p', '9528', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (268, 'p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (269, 'p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (267, 'p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (270, 'p', '9528', '/fileUploadAndDownload/importURL', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (266, 'p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (273, 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (254, 'p', '9528', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (256, 'p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (260, 'p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (262, 'p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (255, 'p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (252, 'p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (257, 'p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (253, 'p', '9528', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (258, 'p', '9528', '/menu/getMenuRoles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (259, 'p', '9528', '/menu/setMenuRoles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (261, 'p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (274, 'p', '9528', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (275, 'p', '9528', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (237, 'p', '9528', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (263, 'p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (291, 'p', '9528', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (264, 'p', '9528', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (265, 'p', '9528', '/user/setUserAuthority', 'POST', '', '', '');

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (1, '2026-07-11 19:08:17.828', '2026-07-11 19:08:17.828', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmZkZTE5NWUtOTU2ZC00OTVlLTk3ZjQtYWY4YTRjZDVjOTYxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Ik1yLuWlh-a3vCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTc4NDM3MDMwOCwibmJmIjoxNzgzNzY1NTA4fQ.ihy2dN2SZPr-w7VFSokoagqkcBB3331FPEAu6UDS63Q');

-- ----------------------------
-- Table structure for sys_ai_workflow_sessions
-- ----------------------------
DROP TABLE IF EXISTS `sys_ai_workflow_sessions`;
CREATE TABLE `sys_ai_workflow_sessions`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '用户ID',
  `tab` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '会话类型',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '会话标题',
  `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '摘要',
  `conversation_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Dify会话ID',
  `message_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Dify消息ID',
  `current_node_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前选中节点ID',
  `settings` json NULL COMMENT '页面设置',
  `form_data` json NULL COMMENT '表单数据',
  `result_data` json NULL COMMENT '当前展示结果',
  `messages` json NULL COMMENT '会话消息',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_ai_workflow_sessions_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_ai_workflow_sessions_user_id`(`user_id` ASC) USING BTREE,
  INDEX `idx_sys_ai_workflow_sessions_tab`(`tab` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_ai_workflow_sessions
-- ----------------------------

-- ----------------------------
-- Table structure for sys_api_tokens
-- ----------------------------
DROP TABLE IF EXISTS `sys_api_tokens`;
CREATE TABLE `sys_api_tokens`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '用户ID',
  `authority_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `token` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Token',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态',
  `expires_at` datetime(3) NULL DEFAULT NULL COMMENT '过期时间',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_api_tokens_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api_tokens
-- ----------------------------

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 190 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysLoginLog/deleteLoginLog', '删除登录日志', '登录日志', 'DELETE');
INSERT INTO `sys_apis` VALUES (3, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysLoginLog/deleteLoginLogByIds', '批量删除登录日志', '登录日志', 'DELETE');
INSERT INTO `sys_apis` VALUES (4, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysLoginLog/findLoginLog', '根据ID获取登录日志', '登录日志', 'GET');
INSERT INTO `sys_apis` VALUES (5, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysLoginLog/getLoginLogList', '获取登录日志列表', '登录日志', 'GET');
INSERT INTO `sys_apis` VALUES (6, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysApiToken/createApiToken', '签发API Token', 'API Token', 'POST');
INSERT INTO `sys_apis` VALUES (7, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysApiToken/getApiTokenList', '获取API Token列表', 'API Token', 'POST');
INSERT INTO `sys_apis` VALUES (8, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysApiToken/deleteApiToken', '作废API Token', 'API Token', 'POST');
INSERT INTO `sys_apis` VALUES (9, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES (10, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (13, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (14, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES (15, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/user/setSelfSetting', '用户界面配置', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (20, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (27, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/syncApi', '获取待同步API', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (28, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/getApiGroups', '获取路由组', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (29, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/enterSyncApi', '确认同步API', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/ignoreApi', '忽略API', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/getApiRoles', '获取指定API关联角色列表', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (32, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/api/setApiRoles', '全量覆盖API关联角色列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES (37, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (38, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/getUsersByAuthority', '获取角色关联用户ID列表', '角色', 'GET');
INSERT INTO `sys_apis` VALUES (40, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authority/setRoleUsers', '全量覆盖角色关联用户', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (41, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (42, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (45, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (46, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (47, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (48, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (49, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (50, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (51, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (52, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/getMenuRoles', '获取菜单关联角色列表', '菜单', 'GET');
INSERT INTO `sys_apis` VALUES (53, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/menu/setMenuRoles', '全量覆盖菜单关联角色列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (54, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET');
INSERT INTO `sys_apis` VALUES (55, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (56, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (57, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (58, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/upload', '文件上传（建议选择）', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (59, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (60, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (61, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (62, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/fileUploadAndDownload/importURL', '导入URL', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (63, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (64, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (65, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (66, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getTools', '获取技能工具列表', 'skills', 'GET');
INSERT INTO `sys_apis` VALUES (67, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getSkillList', '获取技能列表', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (68, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getSkillDetail', '获取技能详情', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (69, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/saveSkill', '保存技能定义', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (70, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/deleteSkill', '删除技能', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (71, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/createScript', '创建技能脚本', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (72, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getScript', '读取技能脚本', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (73, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/saveScript', '保存技能脚本', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (74, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/createResource', '创建技能资源', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (75, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getResource', '读取技能资源', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (76, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/saveResource', '保存技能资源', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (77, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/createReference', '创建技能参考', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (78, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getReference', '读取技能参考', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (79, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/saveReference', '保存技能参考', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (80, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/createTemplate', '创建技能模板', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (81, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getTemplate', '读取技能模板', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (82, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/saveTemplate', '保存技能模板', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (83, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/getGlobalConstraint', '读取全局约束', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (84, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/saveGlobalConstraint', '保存全局约束', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (85, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/skills/packageSkill', '打包技能', 'skills', 'POST');
INSERT INTO `sys_apis` VALUES (86, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` VALUES (87, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` VALUES (88, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` VALUES (89, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (90, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (91, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (92, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (93, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (94, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (95, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (96, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (97, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/pubPlug', '打包插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (98, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/removePlugin', '卸载插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (99, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getPluginList', '获取已安装插件', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (100, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/mcp', '自动生成 MCP Tool 模板', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (101, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/mcpStatus', '获取 MCP 独立服务状态', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (102, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/mcpStart', '启动 MCP 独立服务', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (103, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/mcpStop', '停用 MCP 独立服务', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (104, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/mcpRoutes', '获取 MCP 路由列表', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (105, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/mcpTest', 'MCP Tool 管理', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (106, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/mcpList', '获取 MCP ToolList', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (107, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/saveAIWorkflowSession', '保存AI需求工作流会话', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (108, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getAIWorkflowSessionList', '获取AI需求工作流会话列表', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (109, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getAIWorkflowSessionDetail', '获取AI需求工作流会话详情', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (110, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/deleteAIWorkflowSession', '删除AI需求工作流会话', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (111, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/dumpAIWorkflowMarkdown', 'AI需求工作流Markdown落盘', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (112, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/createPackage', '配置模板', '模板配置', 'POST');
INSERT INTO `sys_apis` VALUES (113, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getTemplates', '获取模板文件', '模板配置', 'GET');
INSERT INTO `sys_apis` VALUES (114, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getPackage', '获取所有模板', '模板配置', 'POST');
INSERT INTO `sys_apis` VALUES (115, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/delPackage', '删除模板', '模板配置', 'POST');
INSERT INTO `sys_apis` VALUES (116, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (117, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (118, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (119, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (120, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/autoCode/addFunc', '增加模板方法', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (121, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` VALUES (122, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` VALUES (123, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` VALUES (124, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (125, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (126, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/getDictionaryTreeList', '获取字典数列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (127, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/getDictionaryTreeListByType', '根据分类获取字典数列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (128, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/getDictionaryDetailsByParent', '根据父级ID获取字典详情', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (129, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionaryDetail/getDictionaryPath', '获取字典详情的完整路径', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (130, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES (131, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` VALUES (132, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` VALUES (133, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典（建议选择）', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (134, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (135, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionary/importSysDictionary', '导入字典JSON', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES (136, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysDictionary/exportSysDictionary', '导出字典JSON', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (137, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` VALUES (138, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (139, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (140, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (141, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (142, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST');
INSERT INTO `sys_apis` VALUES (143, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (144, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (145, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (146, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/email/sendEmail', '发送邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (147, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (148, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (149, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (150, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/createSysExportTemplate', '新增导出模板', '导出模板', 'POST');
INSERT INTO `sys_apis` VALUES (151, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/deleteSysExportTemplate', '删除导出模板', '导出模板', 'DELETE');
INSERT INTO `sys_apis` VALUES (152, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/deleteSysExportTemplateByIds', '批量删除导出模板', '导出模板', 'DELETE');
INSERT INTO `sys_apis` VALUES (153, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/updateSysExportTemplate', '更新导出模板', '导出模板', 'PUT');
INSERT INTO `sys_apis` VALUES (154, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/findSysExportTemplate', '根据ID获取导出模板', '导出模板', 'GET');
INSERT INTO `sys_apis` VALUES (155, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/getSysExportTemplateList', '获取导出模板列表', '导出模板', 'GET');
INSERT INTO `sys_apis` VALUES (156, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/exportExcel', '导出Excel', '导出模板', 'GET');
INSERT INTO `sys_apis` VALUES (157, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/exportTemplate', '下载模板', '导出模板', 'GET');
INSERT INTO `sys_apis` VALUES (158, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/previewSQL', '预览SQL', '导出模板', 'GET');
INSERT INTO `sys_apis` VALUES (159, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysExportTemplate/importExcel', '导入Excel', '导出模板', 'POST');
INSERT INTO `sys_apis` VALUES (160, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysError/createSysError', '新建错误日志', '错误日志', 'POST');
INSERT INTO `sys_apis` VALUES (161, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysError/deleteSysError', '删除错误日志', '错误日志', 'DELETE');
INSERT INTO `sys_apis` VALUES (162, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysError/deleteSysErrorByIds', '批量删除错误日志', '错误日志', 'DELETE');
INSERT INTO `sys_apis` VALUES (163, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysError/updateSysError', '更新错误日志', '错误日志', 'PUT');
INSERT INTO `sys_apis` VALUES (164, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysError/findSysError', '根据ID获取错误日志', '错误日志', 'GET');
INSERT INTO `sys_apis` VALUES (165, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysError/getSysErrorList', '获取错误日志列表', '错误日志', 'GET');
INSERT INTO `sys_apis` VALUES (166, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysError/getSysErrorSolution', '触发错误处理(异步)', '错误日志', 'GET');
INSERT INTO `sys_apis` VALUES (167, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/info/createInfo', '新建公告', '公告', 'POST');
INSERT INTO `sys_apis` VALUES (168, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/info/deleteInfo', '删除公告', '公告', 'DELETE');
INSERT INTO `sys_apis` VALUES (169, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/info/deleteInfoByIds', '批量删除公告', '公告', 'DELETE');
INSERT INTO `sys_apis` VALUES (170, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/info/updateInfo', '更新公告', '公告', 'PUT');
INSERT INTO `sys_apis` VALUES (171, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/info/findInfo', '根据ID获取公告', '公告', 'GET');
INSERT INTO `sys_apis` VALUES (172, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/info/getInfoList', '获取公告列表', '公告', 'GET');
INSERT INTO `sys_apis` VALUES (173, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysParams/createSysParams', '新建参数', '参数管理', 'POST');
INSERT INTO `sys_apis` VALUES (174, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysParams/deleteSysParams', '删除参数', '参数管理', 'DELETE');
INSERT INTO `sys_apis` VALUES (175, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysParams/deleteSysParamsByIds', '批量删除参数', '参数管理', 'DELETE');
INSERT INTO `sys_apis` VALUES (176, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysParams/updateSysParams', '更新参数', '参数管理', 'PUT');
INSERT INTO `sys_apis` VALUES (177, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysParams/findSysParams', '根据ID获取参数', '参数管理', 'GET');
INSERT INTO `sys_apis` VALUES (178, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysParams/getSysParamsList', '获取参数列表', '参数管理', 'GET');
INSERT INTO `sys_apis` VALUES (179, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysParams/getSysParam', '获取参数列表', '参数管理', 'GET');
INSERT INTO `sys_apis` VALUES (180, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/attachmentCategory/getCategoryList', '分类列表', '媒体库分类', 'GET');
INSERT INTO `sys_apis` VALUES (181, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/attachmentCategory/addCategory', '添加/编辑分类', '媒体库分类', 'POST');
INSERT INTO `sys_apis` VALUES (182, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/attachmentCategory/deleteCategory', '删除分类', '媒体库分类', 'POST');
INSERT INTO `sys_apis` VALUES (183, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysVersion/findSysVersion', '获取单一版本', '版本控制', 'GET');
INSERT INTO `sys_apis` VALUES (184, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysVersion/getSysVersionList', '获取版本列表', '版本控制', 'GET');
INSERT INTO `sys_apis` VALUES (185, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysVersion/downloadVersionJson', '下载版本json', '版本控制', 'GET');
INSERT INTO `sys_apis` VALUES (186, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysVersion/exportVersion', '创建版本', '版本控制', 'POST');
INSERT INTO `sys_apis` VALUES (187, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysVersion/importVersion', '同步版本', '版本控制', 'POST');
INSERT INTO `sys_apis` VALUES (188, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysVersion/deleteSysVersion', '删除版本', '版本控制', 'DELETE');
INSERT INTO `sys_apis` VALUES (189, '2026-07-11 18:24:19.557', '2026-07-11 18:24:19.557', NULL, '/sysVersion/deleteSysVersionByIds', '批量删除版本', '版本控制', 'DELETE');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `authority_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `uni_sys_authorities_authority_id`(`authority_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9529 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2026-07-11 18:24:19.630', '2026-07-11 18:24:20.114', NULL, 888, '普通用户', 0, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2026-07-11 18:24:19.630', '2026-07-11 18:24:20.162', NULL, 8881, '普通用户子角色', 888, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2026-07-11 18:24:19.630', '2026-07-11 18:24:20.138', NULL, 9528, '测试角色', 0, 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns`  (
  `authority_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, 888);
INSERT INTO `sys_authority_menus` VALUES (1, 8881);
INSERT INTO `sys_authority_menus` VALUES (1, 9528);
INSERT INTO `sys_authority_menus` VALUES (2, 888);
INSERT INTO `sys_authority_menus` VALUES (2, 8881);
INSERT INTO `sys_authority_menus` VALUES (2, 9528);
INSERT INTO `sys_authority_menus` VALUES (3, 888);
INSERT INTO `sys_authority_menus` VALUES (3, 8881);
INSERT INTO `sys_authority_menus` VALUES (4, 888);
INSERT INTO `sys_authority_menus` VALUES (4, 8881);
INSERT INTO `sys_authority_menus` VALUES (4, 9528);
INSERT INTO `sys_authority_menus` VALUES (5, 888);
INSERT INTO `sys_authority_menus` VALUES (5, 8881);
INSERT INTO `sys_authority_menus` VALUES (6, 888);
INSERT INTO `sys_authority_menus` VALUES (6, 8881);
INSERT INTO `sys_authority_menus` VALUES (7, 888);
INSERT INTO `sys_authority_menus` VALUES (7, 8881);
INSERT INTO `sys_authority_menus` VALUES (7, 9528);
INSERT INTO `sys_authority_menus` VALUES (8, 888);
INSERT INTO `sys_authority_menus` VALUES (8, 8881);
INSERT INTO `sys_authority_menus` VALUES (9, 888);
INSERT INTO `sys_authority_menus` VALUES (10, 888);
INSERT INTO `sys_authority_menus` VALUES (11, 888);
INSERT INTO `sys_authority_menus` VALUES (12, 888);
INSERT INTO `sys_authority_menus` VALUES (13, 888);
INSERT INTO `sys_authority_menus` VALUES (14, 888);
INSERT INTO `sys_authority_menus` VALUES (15, 888);
INSERT INTO `sys_authority_menus` VALUES (16, 888);
INSERT INTO `sys_authority_menus` VALUES (17, 888);
INSERT INTO `sys_authority_menus` VALUES (18, 888);
INSERT INTO `sys_authority_menus` VALUES (19, 888);
INSERT INTO `sys_authority_menus` VALUES (20, 888);
INSERT INTO `sys_authority_menus` VALUES (21, 888);
INSERT INTO `sys_authority_menus` VALUES (21, 8881);
INSERT INTO `sys_authority_menus` VALUES (22, 888);
INSERT INTO `sys_authority_menus` VALUES (22, 8881);
INSERT INTO `sys_authority_menus` VALUES (23, 888);
INSERT INTO `sys_authority_menus` VALUES (23, 8881);
INSERT INTO `sys_authority_menus` VALUES (24, 888);
INSERT INTO `sys_authority_menus` VALUES (24, 8881);
INSERT INTO `sys_authority_menus` VALUES (25, 888);
INSERT INTO `sys_authority_menus` VALUES (25, 8881);
INSERT INTO `sys_authority_menus` VALUES (26, 888);
INSERT INTO `sys_authority_menus` VALUES (26, 8881);
INSERT INTO `sys_authority_menus` VALUES (27, 888);
INSERT INTO `sys_authority_menus` VALUES (27, 8881);
INSERT INTO `sys_authority_menus` VALUES (28, 888);
INSERT INTO `sys_authority_menus` VALUES (28, 8881);
INSERT INTO `sys_authority_menus` VALUES (29, 888);
INSERT INTO `sys_authority_menus` VALUES (29, 8881);
INSERT INTO `sys_authority_menus` VALUES (30, 888);
INSERT INTO `sys_authority_menus` VALUES (30, 8881);
INSERT INTO `sys_authority_menus` VALUES (31, 888);
INSERT INTO `sys_authority_menus` VALUES (31, 8881);
INSERT INTO `sys_authority_menus` VALUES (32, 888);
INSERT INTO `sys_authority_menus` VALUES (33, 888);
INSERT INTO `sys_authority_menus` VALUES (34, 888);
INSERT INTO `sys_authority_menus` VALUES (35, 888);
INSERT INTO `sys_authority_menus` VALUES (36, 888);
INSERT INTO `sys_authority_menus` VALUES (37, 888);
INSERT INTO `sys_authority_menus` VALUES (38, 888);
INSERT INTO `sys_authority_menus` VALUES (39, 888);
INSERT INTO `sys_authority_menus` VALUES (40, 888);

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表名',
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模块名/插件名',
  `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '前端传入的结构化信息',
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '结构体名称',
  `abbreviation` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '结构体名称缩写',
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务库',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Struct中文名称',
  `templates` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '模板信息',
  `injections` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '注入路径',
  `flag` bigint NULL DEFAULT NULL COMMENT '[0:创建,1:回滚]',
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api表注册内容',
  `menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `export_template_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '导出模板ID',
  `package_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '包ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_histories_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------

-- ----------------------------
-- Table structure for sys_auto_code_packages
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_packages`;
CREATE TABLE `sys_auto_code_packages`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示名',
  `template` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模版',
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名',
  `module` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_packages_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_code_packages
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sys_base_menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_btns_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `menu_level` bigint UNSIGNED NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '高亮菜单',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '是否缓存',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '是否是基础路由（开发中）',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单名',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `close_tab` tinyint(1) NULL DEFAULT NULL COMMENT '自动关闭tab',
  `transition_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由切换动画',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '', 0, 0, '仪表盘', 'odometer', 0, '');
INSERT INTO `sys_base_menus` VALUES (2, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'about', 'about', 0, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0, '');
INSERT INTO `sys_base_menus` VALUES (3, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', 0, 0, '超级管理员', 'user', 0, '');
INSERT INTO `sys_base_menus` VALUES (4, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'person', 'person', 1, 'view/person/person.vue', 4, '', 0, 0, '个人信息', 'message', 0, '');
INSERT INTO `sys_base_menus` VALUES (5, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '', 0, 0, '编程辅助', 'tools', 0, '');
INSERT INTO `sys_base_menus` VALUES (6, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, '', 0, 0, '官方网站', 'customer-gva', 0, '');
INSERT INTO `sys_base_menus` VALUES (7, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'state', 'state', 0, 'view/system/state.vue', 8, '', 0, 0, '服务器状态', 'cloudy', 0, '');
INSERT INTO `sys_base_menus` VALUES (8, '2026-07-11 18:24:19.795', '2026-07-11 18:24:19.795', NULL, 0, 0, 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '', 0, 0, '插件系统', 'cherry', 0, '');
INSERT INTO `sys_base_menus` VALUES (9, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '', 0, 0, '角色管理', 'avatar', 0, '');
INSERT INTO `sys_base_menus` VALUES (10, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '', 1, 0, '菜单管理', 'tickets', 0, '');
INSERT INTO `sys_base_menus` VALUES (11, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '', 1, 0, 'api管理', 'platform', 0, '');
INSERT INTO `sys_base_menus` VALUES (12, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '', 0, 0, '用户管理', 'coordinate', 0, '');
INSERT INTO `sys_base_menus` VALUES (13, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'notebook', 0, '');
INSERT INTO `sys_base_menus` VALUES (14, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, '操作历史', 'pie-chart', 0, '');
INSERT INTO `sys_base_menus` VALUES (15, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'sysParams', 'sysParams', 0, 'view/superAdmin/params/sysParams.vue', 7, '', 0, 0, '参数管理', 'compass', 0, '');
INSERT INTO `sys_base_menus` VALUES (16, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'system', 'system', 0, 'view/systemTools/system/system.vue', 8, '', 0, 0, '系统配置', 'operation', 0, '');
INSERT INTO `sys_base_menus` VALUES (17, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'apiToken', 'apiToken', 0, 'view/systemTools/apiToken/index.vue', 9, '', 0, 0, 'API Token', 'key', 0, '');
INSERT INTO `sys_base_menus` VALUES (18, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'loginLog', 'loginLog', 0, 'view/systemTools/loginLog/index.vue', 10, '', 0, 0, '登录日志', 'monitor', 0, '');
INSERT INTO `sys_base_menus` VALUES (19, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'sysVersion', 'sysVersion', 0, 'view/systemTools/version/version.vue', 11, '', 0, 0, '版本管理', 'server', 0, '');
INSERT INTO `sys_base_menus` VALUES (20, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 3, 'sysError', 'sysError', 0, 'view/systemTools/sysError/sysError.vue', 12, '', 0, 0, '错误日志', 'warn', 0, '');
INSERT INTO `sys_base_menus` VALUES (21, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '模板配置', 'folder', 0, '');
INSERT INTO `sys_base_menus` VALUES (22, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0, '');
INSERT INTO `sys_base_menus` VALUES (23, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 2, '', 0, 0, '自动化代码管理', 'magic-stick', 0, '');
INSERT INTO `sys_base_menus` VALUES (24, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 3, '', 1, 0, '表单生成器', 'magic-stick', 0, '');
INSERT INTO `sys_base_menus` VALUES (25, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'aiWorkflow', 'aiWorkflow', 0, 'view/systemTools/aiWrokflow/index.vue', 4, '', 1, 0, 'AI需求工作流', 'magic-stick', 0, '');
INSERT INTO `sys_base_menus` VALUES (26, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0, '');
INSERT INTO `sys_base_menus` VALUES (27, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'exportTemplate', 'exportTemplate', 0, 'view/systemTools/exportTemplate/exportTemplate.vue', 5, '', 0, 0, '导出模板', 'reading', 0, '');
INSERT INTO `sys_base_menus` VALUES (28, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'mcpTest', 'mcpTest', 0, 'view/systemTools/autoCode/mcpTest.vue', 6, '', 0, 0, 'Mcp Tools管理', 'partly-cloudy', 0, '');
INSERT INTO `sys_base_menus` VALUES (29, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'mcpTool', 'mcpTool', 0, 'view/systemTools/autoCode/mcp.vue', 7, '', 0, 0, 'Mcp Tools模板', 'magnet', 0, '');
INSERT INTO `sys_base_menus` VALUES (30, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'skills', 'skills', 0, 'view/systemTools/skills/index.vue', 8, '', 0, 0, 'Skills管理', 'document', 0, '');
INSERT INTO `sys_base_menus` VALUES (31, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 5, 'picture', 'picture', 0, 'view/systemTools/autoCode/picture.vue', 9, '', 0, 0, 'AI页面绘制', 'picture-filled', 0, '');
INSERT INTO `sys_base_menus` VALUES (32, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 8, 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, '', 0, 0, '插件市场', 'shop', 0, '');
INSERT INTO `sys_base_menus` VALUES (33, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 8, 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, '插件安装', 'box', 0, '');
INSERT INTO `sys_base_menus` VALUES (34, '2026-07-11 18:24:19.804', '2026-07-11 18:24:19.804', NULL, 1, 8, 'pubPlug', 'pubPlug', 0, 'view/systemTools/pubPlug/pubPlug.vue', 3, '', 0, 0, '打包插件', 'files', 0, '');
INSERT INTO `sys_base_menus` VALUES (35, '2026-07-11 18:56:25.000', '2026-07-11 18:56:25.000', NULL, 0, 0, 'gb28181', 'gb28181', 0, 'view/routerHolder.vue', 2, '', 0, 0, '国标GB28181', 'monitor', 0, '');
INSERT INTO `sys_base_menus` VALUES (36, '2026-07-11 18:56:25.000', '2026-07-11 18:56:25.000', NULL, 1, 35, 'device', 'device', 0, 'view/device/index.vue', 1, '', 0, 0, '设备管理', 'monitor', 0, '');
INSERT INTO `sys_base_menus` VALUES (37, '2026-07-11 18:56:25.000', '2026-07-11 18:56:25.000', NULL, 1, 35, 'platform', 'platform', 0, 'view/platform/index.vue', 2, '', 0, 0, '平台级联', 'connection', 0, '');
INSERT INTO `sys_base_menus` VALUES (38, '2026-07-11 18:56:25.000', '2026-07-11 18:56:25.000', NULL, 1, 35, 'alarm', 'alarm', 0, 'view/alarm/index.vue', 3, '', 0, 0, '报警管理', 'warning', 0, '');
INSERT INTO `sys_base_menus` VALUES (39, '2026-07-11 18:56:25.000', '2026-07-11 18:56:25.000', NULL, 1, 35, 'position', 'position', 0, 'view/position/index.vue', 4, '', 0, 0, '移动位置', 'location', 0, '');
INSERT INTO `sys_base_menus` VALUES (40, '2026-07-11 18:56:25.000', '2026-07-11 18:56:25.000', NULL, 1, 35, 'record', 'record', 0, 'view/record/index.vue', 5, '', 0, 0, '录像管理', 'video-camera', 0, '');

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES (888, 888);
INSERT INTO `sys_data_authority_id` VALUES (888, 8881);
INSERT INTO `sys_data_authority_id` VALUES (888, 9528);
INSERT INTO `sys_data_authority_id` VALUES (9528, 8881);
INSERT INTO `sys_data_authority_id` VALUES (9528, 9528);

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父级字典ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2026-07-11 18:24:19.682', '2026-07-11 18:24:19.699', NULL, '性别', 'gender', 1, '性别字典', NULL);
INSERT INTO `sys_dictionaries` VALUES (2, '2026-07-11 18:24:19.682', '2026-07-11 18:24:19.715', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型', NULL);
INSERT INTO `sys_dictionaries` VALUES (3, '2026-07-11 18:24:19.682', '2026-07-11 18:24:19.733', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型', NULL);
INSERT INTO `sys_dictionaries` VALUES (4, '2026-07-11 18:24:19.682', '2026-07-11 18:24:19.749', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型', NULL);
INSERT INTO `sys_dictionaries` VALUES (5, '2026-07-11 18:24:19.682', '2026-07-11 18:24:19.762', NULL, '数据库字符串', 'string', 1, '数据库字符串', NULL);
INSERT INTO `sys_dictionaries` VALUES (6, '2026-07-11 18:24:19.682', '2026-07-11 18:24:19.778', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型', NULL);

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '关联标记',
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父级字典详情ID',
  `level` bigint NULL DEFAULT NULL COMMENT '层级深度',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '层级路径',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2026-07-11 18:24:19.701', '2026-07-11 18:24:19.701', NULL, '男', '1', '', 1, 1, 1, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (2, '2026-07-11 18:24:19.701', '2026-07-11 18:24:19.701', NULL, '女', '2', '', 1, 2, 1, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (3, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'smallint', '1', 'mysql', 1, 1, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (4, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'mediumint', '2', 'mysql', 1, 2, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (5, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'int', '3', 'mysql', 1, 3, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (6, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'bigint', '4', 'mysql', 1, 4, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (7, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'int2', '5', 'pgsql', 1, 5, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (8, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'int4', '6', 'pgsql', 1, 6, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (9, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'int6', '7', 'pgsql', 1, 7, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (10, '2026-07-11 18:24:19.718', '2026-07-11 18:24:19.718', NULL, 'int8', '8', 'pgsql', 1, 8, 2, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (11, '2026-07-11 18:24:19.736', '2026-07-11 18:24:19.736', NULL, 'date', '0', 'mysql', 1, 0, 3, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (12, '2026-07-11 18:24:19.736', '2026-07-11 18:24:19.736', NULL, 'time', '1', 'mysql', 1, 1, 3, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (13, '2026-07-11 18:24:19.736', '2026-07-11 18:24:19.736', NULL, 'year', '2', 'mysql', 1, 2, 3, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (14, '2026-07-11 18:24:19.736', '2026-07-11 18:24:19.736', NULL, 'datetime', '3', 'mysql', 1, 3, 3, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (15, '2026-07-11 18:24:19.736', '2026-07-11 18:24:19.736', NULL, 'timestamp', '5', 'mysql', 1, 5, 3, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (16, '2026-07-11 18:24:19.736', '2026-07-11 18:24:19.736', NULL, 'timestamptz', '6', 'pgsql', 1, 5, 3, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (17, '2026-07-11 18:24:19.750', '2026-07-11 18:24:19.750', NULL, 'float', '0', 'mysql', 1, 0, 4, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (18, '2026-07-11 18:24:19.750', '2026-07-11 18:24:19.750', NULL, 'double', '1', 'mysql', 1, 1, 4, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (19, '2026-07-11 18:24:19.750', '2026-07-11 18:24:19.750', NULL, 'decimal', '2', 'mysql', 1, 2, 4, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (20, '2026-07-11 18:24:19.750', '2026-07-11 18:24:19.750', NULL, 'numeric', '3', 'pgsql', 1, 3, 4, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (21, '2026-07-11 18:24:19.750', '2026-07-11 18:24:19.750', NULL, 'smallserial', '4', 'pgsql', 1, 4, 4, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (22, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'char', '0', 'mysql', 1, 0, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (23, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'varchar', '1', 'mysql', 1, 1, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (24, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'tinyblob', '2', 'mysql', 1, 2, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (25, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'tinytext', '3', 'mysql', 1, 3, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (26, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'text', '4', 'mysql', 1, 4, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (27, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'blob', '5', 'mysql', 1, 5, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (28, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'mediumblob', '6', 'mysql', 1, 6, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (29, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'mediumtext', '7', 'mysql', 1, 7, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (30, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'longblob', '8', 'mysql', 1, 8, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (31, '2026-07-11 18:24:19.763', '2026-07-11 18:24:19.763', NULL, 'longtext', '9', 'mysql', 1, 9, 5, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (32, '2026-07-11 18:24:19.780', '2026-07-11 18:24:19.780', NULL, 'tinyint', '1', 'mysql', 1, 0, 6, NULL, 0, '');
INSERT INTO `sys_dictionary_details` VALUES (33, '2026-07-11 18:24:19.780', '2026-07-11 18:24:19.780', NULL, 'bool', '2', 'pgsql', 1, 0, 6, NULL, 0, '');

-- ----------------------------
-- Table structure for sys_error
-- ----------------------------
DROP TABLE IF EXISTS `sys_error`;
CREATE TABLE `sys_error`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `form` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '错误来源',
  `info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '错误内容',
  `level` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日志等级',
  `solution` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '解决方案',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '未处理' COMMENT '处理状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_error_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_error
-- ----------------------------
INSERT INTO `sys_error` VALUES (1, '2026-07-11 18:25:37.195', '2026-07-11 18:25:37.195', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTemplates (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:45:17)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (2, '2026-07-11 18:25:37.208', '2026-07-11 18:25:37.208', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTableData (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:111:19)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (3, '2026-07-11 18:25:37.534', '2026-07-11 18:25:37.534', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTableData (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:111:19)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (4, '2026-07-11 18:25:37.535', '2026-07-11 18:25:37.535', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTemplates (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:45:17)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (5, '2026-07-11 18:25:47.205', '2026-07-11 18:25:47.205', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getDbFunc (http://192.168.12.130:8080/src/view/systemTools/autoCode/index.vue:613:17)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (6, '2026-07-11 18:25:47.205', '2026-07-11 18:25:47.205', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getPkgs (http://192.168.12.130:8080/src/view/systemTools/autoCode/index.vue:738:17)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (7, '2026-07-11 18:25:50.614', '2026-07-11 18:25:50.614', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTableData (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:111:19)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (8, '2026-07-11 18:25:50.626', '2026-07-11 18:25:50.626', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTableData (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:111:19)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (9, '2026-07-11 18:25:50.631', '2026-07-11 18:25:50.631', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTemplates (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:45:17)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (10, '2026-07-11 18:25:50.652', '2026-07-11 18:25:50.652', NULL, '前端', '错误信息: AxiosError: Request failed with status code 404\nStack: 调用栈: AxiosError: Request failed with status code 404\n    at settle (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1494:14)\n    at XMLHttpRequest.onloadend (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:1801:4)\n    at Axios$1.request (http://192.168.12.130:8080/node_modules/.vite/deps/axios.js?v=0b61d329:2296:37)\n    at async getTemplates (http://192.168.12.130:8080/src/view/systemTools/autoPkg/autoPkg.vue:45:17)', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (11, '2026-07-11 18:26:13.356', '2026-07-11 18:26:13.356', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (12, '2026-07-11 18:56:52.696', '2026-07-11 18:56:52.696', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/device/index.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/device/index.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (13, '2026-07-11 18:56:59.811', '2026-07-11 18:56:59.811', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/platform/index.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/platform/index.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (14, '2026-07-11 18:57:01.455', '2026-07-11 18:57:01.455', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/alarm/index.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/alarm/index.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (15, '2026-07-11 18:57:02.633', '2026-07-11 18:57:02.633', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/position/index.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/position/index.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (16, '2026-07-11 18:57:03.759', '2026-07-11 18:57:03.759', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/record/index.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/record/index.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (17, '2026-07-11 18:57:05.106', '2026-07-11 18:57:05.106', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/device/index.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/device/index.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (18, '2026-07-11 19:07:10.135', '2026-07-11 19:07:10.135', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (19, '2026-07-11 19:08:05.413', '2026-07-11 19:08:05.413', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (20, '2026-07-11 19:11:21.473', '2026-07-11 19:11:21.473', NULL, '前端', '错误信息: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue?t=1783768266631\nStack: 调用栈: TypeError: Failed to fetch dynamically imported module: http://192.168.12.130:8080/src/view/person/person.vue?t=1783768266631', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (21, '2026-07-11 19:20:33.815', '2026-07-11 19:20:33.815', NULL, '后端', 'count devices failed | 错误: Error 1146 (42S02): Table \'wvp.wvp_device\' doesn\'t exist \n 源文件:/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51 \n 调用栈：wvp-go/server/api/v1/gb28181.GetDeviceList\n	/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.CasbinHandler.func4\n	/home/wyd/gb28181/wvp-go/server/middleware/casbin_rbac.go:30\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.JWTAuth.func3\n	/home/wyd/gb28181/wvp-go/server/middleware/jwt.go:69\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.GinRecovery.func2\n	/home/wyd/gb28181/wvp-go/server/middleware/error.go:78\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589\nnet/http.serverHandler.ServeHTTP\n	/usr/local/go/src/net/http/server.go:3301\nnet/http.(*conn).serve\n	/usr/local/go/src/net/http/server.go:2102 \n 最终调用方法:/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51 (GetDeviceList lines 23-72)\n----- 产生日志的方法代码如下 -----\nfunc GetDeviceList(c *gin.Context) {\n	var pageInfo struct {\n		Page     int    `form:\"page\" binding:\"required,min=1\"`\n		PageSize int    `form:\"page_size\" binding:\"required,min=1,max=100\"`\n		DeviceID string `form:\"device_id\"`\n		Name     string `form:\"name\"`\n		Online   *bool  `form:\"online\"`\n	}\n\n	if err := c.ShouldBindQuery(&pageInfo); err != nil {\n		response.Fail(c, response.INVALID_PARAMS, nil)\n		return\n	}\n\n	db := global.GVA_DB.Model(&system.Device{})\n\n	if pageInfo.DeviceID != \"\" {\n		db = db.Where(\"device_id LIKE ?\", \"%\"+pageInfo.DeviceID+\"%\")\n	}\n	if pageInfo.Name != \"\" {\n		db = db.Where(\"name LIKE ?\", \"%\"+pageInfo.Name+\"%\")\n	}\n	if pageInfo.Online != nil {\n		db = db.Where(\"online = ?\", *pageInfo.Online)\n	}\n\n	var total int64\n	if err := db.Count(&total).Error; err != nil {\n		global.GVA_LOG.Error(\"count devices failed\", zap.Error(err))\n		response.Fail(c, response.DB_ERROR, nil)\n		return\n	}\n\n	var devices []system.Device\n	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).\n		Limit(pageInfo.PageSize).\n		Order(\"created_at DESC\").\n		Find(&devices).Error; err != nil {\n		global.GVA_LOG.Error(\"query devices failed\", zap.Error(err))\n		response.Fail(c, response.DB_ERROR, nil)\n		return\n	}\n\n	response.Success(c, gin.H{\n		\"list\":      devices,\n		\"total\":     total,\n		\"page\":      pageInfo.Page,\n		\"page_size\": pageInfo.PageSize,\n	})\n}', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (22, '2026-07-11 19:24:04.878', '2026-07-11 19:24:04.878', NULL, '后端', 'count devices failed | 错误: Error 1054 (42S22): Unknown column \'online\' in \'where clause\' \n 源文件:/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51 \n 调用栈：wvp-go/server/api/v1/gb28181.GetDeviceList\n	/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.CasbinHandler.func4\n	/home/wyd/gb28181/wvp-go/server/middleware/casbin_rbac.go:30\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.JWTAuth.func3\n	/home/wyd/gb28181/wvp-go/server/middleware/jwt.go:69\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.GinRecovery.func2\n	/home/wyd/gb28181/wvp-go/server/middleware/error.go:78\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589\nnet/http.serverHandler.ServeHTTP\n	/usr/local/go/src/net/http/server.go:3301\nnet/http.(*conn).serve\n	/usr/local/go/src/net/http/server.go:2102 \n 最终调用方法:/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51 (GetDeviceList lines 23-72)\n----- 产生日志的方法代码如下 -----\nfunc GetDeviceList(c *gin.Context) {\n	var pageInfo struct {\n		Page     int    `form:\"page\" binding:\"required,min=1\"`\n		PageSize int    `form:\"page_size\" binding:\"required,min=1,max=100\"`\n		DeviceID string `form:\"device_id\"`\n		Name     string `form:\"name\"`\n		Online   *bool  `form:\"online\"`\n	}\n\n	if err := c.ShouldBindQuery(&pageInfo); err != nil {\n		response.Fail(c, response.INVALID_PARAMS, nil)\n		return\n	}\n\n	db := global.GVA_DB.Model(&system.Device{})\n\n	if pageInfo.DeviceID != \"\" {\n		db = db.Where(\"device_id LIKE ?\", \"%\"+pageInfo.DeviceID+\"%\")\n	}\n	if pageInfo.Name != \"\" {\n		db = db.Where(\"name LIKE ?\", \"%\"+pageInfo.Name+\"%\")\n	}\n	if pageInfo.Online != nil {\n		db = db.Where(\"online = ?\", *pageInfo.Online)\n	}\n\n	var total int64\n	if err := db.Count(&total).Error; err != nil {\n		global.GVA_LOG.Error(\"count devices failed\", zap.Error(err))\n		response.Fail(c, response.DB_ERROR, nil)\n		return\n	}\n\n	var devices []system.Device\n	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).\n		Limit(pageInfo.PageSize).\n		Order(\"created_at DESC\").\n		Find(&devices).Error; err != nil {\n		global.GVA_LOG.Error(\"query devices failed\", zap.Error(err))\n		response.Fail(c, response.DB_ERROR, nil)\n		return\n	}\n\n	response.Success(c, gin.H{\n		\"list\":      devices,\n		\"total\":     total,\n		\"page\":      pageInfo.Page,\n		\"page_size\": pageInfo.PageSize,\n	})\n}', 'error', NULL, '未处理');
INSERT INTO `sys_error` VALUES (23, '2026-07-11 19:26:00.690', '2026-07-11 19:26:00.690', NULL, '后端', 'count devices failed | 错误: Error 1054 (42S22): Unknown column \'wvp_device.deleted_at\' in \'where clause\' \n 源文件:/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51 \n 调用栈：wvp-go/server/api/v1/gb28181.GetDeviceList\n	/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.CasbinHandler.func4\n	/home/wyd/gb28181/wvp-go/server/middleware/casbin_rbac.go:30\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.JWTAuth.func3\n	/home/wyd/gb28181/wvp-go/server/middleware/jwt.go:69\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\nwvp-go/server/initialize.Routers.GinRecovery.func2\n	/home/wyd/gb28181/wvp-go/server/middleware/error.go:78\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n	/home/wyd/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589\nnet/http.serverHandler.ServeHTTP\n	/usr/local/go/src/net/http/server.go:3301\nnet/http.(*conn).serve\n	/usr/local/go/src/net/http/server.go:2102 \n 最终调用方法:/home/wyd/gb28181/wvp-go/server/api/v1/gb28181/device.go:51 (GetDeviceList lines 23-72)\n----- 产生日志的方法代码如下 -----\nfunc GetDeviceList(c *gin.Context) {\n	var pageInfo struct {\n		Page     int    `form:\"page\" binding:\"required,min=1\"`\n		PageSize int    `form:\"page_size\" binding:\"required,min=1,max=100\"`\n		DeviceID string `form:\"device_id\"`\n		Name     string `form:\"name\"`\n		Online   *bool  `form:\"online\"`\n	}\n\n	if err := c.ShouldBindQuery(&pageInfo); err != nil {\n		response.Fail(c, response.INVALID_PARAMS, nil)\n		return\n	}\n\n	db := global.GVA_DB.Model(&system.Device{})\n\n	if pageInfo.DeviceID != \"\" {\n		db = db.Where(\"device_id LIKE ?\", \"%\"+pageInfo.DeviceID+\"%\")\n	}\n	if pageInfo.Name != \"\" {\n		db = db.Where(\"name LIKE ?\", \"%\"+pageInfo.Name+\"%\")\n	}\n	if pageInfo.Online != nil {\n		db = db.Where(\"online = ?\", *pageInfo.Online)\n	}\n\n	var total int64\n	if err := db.Count(&total).Error; err != nil {\n		global.GVA_LOG.Error(\"count devices failed\", zap.Error(err))\n		response.Fail(c, response.DB_ERROR, nil)\n		return\n	}\n\n	var devices []system.Device\n	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).\n		Limit(pageInfo.PageSize).\n		Order(\"created_at DESC\").\n		Find(&devices).Error; err != nil {\n		global.GVA_LOG.Error(\"query devices failed\", zap.Error(err))\n		response.Fail(c, response.DB_ERROR, nil)\n		return\n	}\n\n	response.Success(c, gin.H{\n		\"list\":      devices,\n		\"total\":     total,\n		\"page\":      pageInfo.Page,\n		\"page_size\": pageInfo.PageSize,\n	})\n}', 'error', NULL, '未处理');

-- ----------------------------
-- Table structure for sys_export_template_condition
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_template_condition`;
CREATE TABLE `sys_export_template_condition`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_export_template_condition_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_export_template_condition
-- ----------------------------

-- ----------------------------
-- Table structure for sys_export_template_join
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_template_join`;
CREATE TABLE `sys_export_template_join`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联',
  `table` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_export_template_join_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_export_template_join
-- ----------------------------

-- ----------------------------
-- Table structure for sys_export_templates
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_templates`;
CREATE TABLE `sys_export_templates`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `db_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板标识',
  `template_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `sql` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '自定义导出SQL',
  `import_sql` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '自定义导入SQL',
  `limit` bigint NULL DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_export_templates_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_export_templates
-- ----------------------------
INSERT INTO `sys_export_templates` VALUES (1, '2026-07-11 18:24:20.102', '2026-07-11 18:24:20.102', NULL, '', 'api', 'sys_apis', 'api', '{\n\"path\":\"路径\",\n\"method\":\"方法（大写）\",\n\"description\":\"方法介绍\",\n\"api_group\":\"方法分组\"\n}', '', '', NULL, '');

-- ----------------------------
-- Table structure for sys_ignore_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_ignore_apis`;
CREATE TABLE `sys_ignore_apis`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_ignore_apis_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_ignore_apis
-- ----------------------------
INSERT INTO `sys_ignore_apis` VALUES (1, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/swagger/*any', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (2, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/api/freshCasbin', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (3, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/uploads/file/*filepath', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (4, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/health', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (5, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/uploads/file/*filepath', 'HEAD');
INSERT INTO `sys_ignore_apis` VALUES (6, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/autoCode/llmAuto', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (7, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/autoCode/llmAutoSSE', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (8, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/system/reloadSystem', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (9, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/base/login', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (10, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/base/captcha', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (11, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/init/initdb', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (12, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/init/checkdb', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (13, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/info/getInfoDataSource', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (14, '2026-07-11 18:24:19.586', '2026-07-11 18:24:19.586', NULL, '/info/getInfoPublic', 'GET');

-- ----------------------------
-- Table structure for sys_login_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_logs`;
CREATE TABLE `sys_login_logs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '登录状态',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代理',
  `user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_login_logs_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_login_logs
-- ----------------------------
INSERT INTO `sys_login_logs` VALUES (1, '2026-07-11 18:25:08.550', '2026-07-11 18:25:08.550', NULL, 'admin', '127.0.0.1', 1, '登录成功', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', 1);
INSERT INTO `sys_login_logs` VALUES (2, '2026-07-11 19:08:29.695', '2026-07-11 19:08:29.695', NULL, 'admin', '127.0.0.1', 1, '登录成功', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', 1);

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `status` bigint NULL DEFAULT NULL COMMENT '请求状态',
  `latency` bigint NULL DEFAULT NULL COMMENT '延迟',
  `agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
INSERT INTO `sys_operation_records` VALUES (1, '2026-07-11 19:07:30.692', '2026-07-11 19:07:30.692', NULL, '127.0.0.1', 'POST', '/menu/setMenuRoles', 200, 8727339, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', '', '{\"menuId\":35,\"authorityIds\":[888]}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (2, '2026-07-11 19:09:18.191', '2026-07-11 19:09:18.191', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 13473507, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', '', '{\"authorityId\":8881}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (3, '2026-07-11 19:09:32.688', '2026-07-11 19:09:32.688', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 14150256, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', '', '{\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (4, '2026-07-11 19:15:01.786', '2026-07-11 19:15:01.786', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 10943726, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', '', '{\"authorityId\":9528}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (5, '2026-07-11 19:15:05.818', '2026-07-11 19:15:05.818', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 11500356, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', '', '{\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (6, '2026-07-11 19:17:01.494', '2026-07-11 19:17:01.494', NULL, '127.0.0.1', 'GET', '/api/getApiGroups', 200, 2465754, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', '', '{}', '{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"autoCode\":\"代码生成器历史\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"skills\":\"skills\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysExportTemplate\":\"导出模板\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"skills\",\"客户\",\"代码生成器\",\"模板配置\",\"代码生成器历史\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"导出模板\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\"]},\"msg\":\"成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (7, '2026-07-11 19:19:45.247', '2026-07-11 19:19:45.247', NULL, '127.0.0.1', 'GET', '/api/getApiGroups', 200, 1560801, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36 Edg/150.0.0.0', '', '{}', '{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"autoCode\":\"代码生成器历史\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"skills\":\"skills\",\"sysApiToken\":\"API Token\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysError\":\"错误日志\",\"sysExportTemplate\":\"导出模板\",\"sysLoginLog\":\"登录日志\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"登录日志\",\"API Token\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"skills\",\"客户\",\"代码生成器\",\"模板配置\",\"代码生成器历史\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"导出模板\",\"错误日志\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\"]},\"msg\":\"成功\"}', 1);

-- ----------------------------
-- Table structure for sys_params
-- ----------------------------
DROP TABLE IF EXISTS `sys_params`;
CREATE TABLE `sys_params`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数名称',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数键',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数值',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数说明',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_params_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_params
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, 888);
INSERT INTO `sys_user_authority` VALUES (1, 8881);
INSERT INTO `sys_user_authority` VALUES (1, 9528);
INSERT INTO `sys_user_authority` VALUES (2, 888);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `authority_id` bigint UNSIGNED NULL DEFAULT 888 COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  `origin_setting` json NULL COMMENT '配置',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_users_uuid`(`uuid` ASC) USING BTREE,
  INDEX `idx_sys_users_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2026-07-11 18:24:20.053', '2026-07-11 19:15:05.810', NULL, 'ffde195e-956d-495e-97f4-af8a4cd5c961', 'admin', '$2a$10$CGKZmzPU1nCzRirHJpkp3eA8i65iick/8X.esmv330puQPxr/FX1i', 'Mr.奇淼', 'https://qmplusimg.henrongyi.top/gva_header.jpg', 888, '17611111111', '333333333@qq.com', 1, NULL);
INSERT INTO `sys_users` VALUES (2, '2026-07-11 18:24:20.053', '2026-07-11 18:24:20.083', NULL, '02ea0bb9-5aa1-411a-9be7-20f5e6ffe73c', 'a303176530', '$2a$10$iNJFML9mDIRRqQJLECpkv.Rki9dWOc/R5WiNj9kmTMMVO/p8kUh86', '用户1', 'https://qmplusimg.henrongyi.top/1572075907logo.png', 9528, '17611111111', '333333333@qq.com', 1, NULL);

-- ----------------------------
-- Table structure for sys_versions
-- ----------------------------
DROP TABLE IF EXISTS `sys_versions`;
CREATE TABLE `sys_versions`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `version_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本名称',
  `version_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本号',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本描述',
  `version_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '版本数据JSON',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_versions_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_versions
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_alarm
-- ----------------------------
DROP TABLE IF EXISTS `wvp_alarm`;
CREATE TABLE `wvp_alarm`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `channel_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `alarm_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `alarm_type_val` int NULL DEFAULT 0,
  `alarm_priority` tinyint NULL DEFAULT 0,
  `alarm_time` datetime(3) NULL DEFAULT NULL,
  `longitude` decimal(10, 7) NULL DEFAULT 0.0000000,
  `latitude` decimal(10, 7) NULL DEFAULT 0.0000000,
  `description` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `handle_status` tinyint NULL DEFAULT 0,
  `handle_user` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `handle_time` datetime(3) NULL DEFAULT NULL,
  `handle_result` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_device_id`(`device_id` ASC) USING BTREE,
  INDEX `idx_channel_id`(`channel_id` ASC) USING BTREE,
  INDEX `idx_alarm_type`(`alarm_type` ASC) USING BTREE,
  INDEX `idx_alarm_time`(`alarm_time` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_alarm
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_device
-- ----------------------------
DROP TABLE IF EXISTS `wvp_device`;
CREATE TABLE `wvp_device`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `manufacturer` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `model` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `owner` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `civil_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `parental` tinyint NULL DEFAULT 0,
  `parent_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `safety_way` tinyint NULL DEFAULT 0,
  `register_way` tinyint NULL DEFAULT 1,
  `secrecy` tinyint NULL DEFAULT 0,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `transport` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'UDP',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `port` int NULL DEFAULT 5060,
  `online` tinyint(1) NULL DEFAULT 0,
  `register_time` datetime(3) NULL DEFAULT NULL,
  `keepalive_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_device_id`(`device_id` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_device
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_device_channel
-- ----------------------------
DROP TABLE IF EXISTS `wvp_device_channel`;
CREATE TABLE `wvp_device_channel`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `channel_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `parent_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `manufacturer` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `model` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `owner` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `civil_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `parental` tinyint NULL DEFAULT 0,
  `safety_way` tinyint NULL DEFAULT 0,
  `register_way` tinyint NULL DEFAULT 1,
  `secrecy` tinyint NULL DEFAULT 0,
  `status` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'OFF',
  `has_audio` tinyint(1) NULL DEFAULT 0,
  `longitude` decimal(10, 7) NULL DEFAULT 0.0000000,
  `latitude` decimal(10, 7) NULL DEFAULT 0.0000000,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_channel_id`(`channel_id` ASC) USING BTREE,
  INDEX `idx_device_id`(`device_id` ASC) USING BTREE,
  INDEX `idx_parent_id`(`parent_id` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_device_channel
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_device_position
-- ----------------------------
DROP TABLE IF EXISTS `wvp_device_position`;
CREATE TABLE `wvp_device_position`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `channel_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `longitude` decimal(10, 7) NOT NULL DEFAULT 0.0000000,
  `latitude` decimal(10, 7) NOT NULL DEFAULT 0.0000000,
  `altitude` decimal(10, 2) NULL DEFAULT 0.00,
  `speed` decimal(10, 2) NULL DEFAULT 0.00,
  `direction` decimal(10, 2) NULL DEFAULT 0.00,
  `gps_status` tinyint NULL DEFAULT 0,
  `report_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_device_id`(`device_id` ASC) USING BTREE,
  INDEX `idx_channel_id`(`channel_id` ASC) USING BTREE,
  INDEX `idx_report_time`(`report_time` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_device_position
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_media_server
-- ----------------------------
DROP TABLE IF EXISTS `wvp_media_server`;
CREATE TABLE `wvp_media_server`  (
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `hook_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sdp_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `stream_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `http_port` int NULL DEFAULT NULL,
  `http_ssl_port` int NULL DEFAULT NULL,
  `rtmp_port` int NULL DEFAULT NULL,
  `rtmp_ssl_port` int NULL DEFAULT NULL,
  `rtp_proxy_port` int NULL DEFAULT NULL,
  `rtsp_port` int NULL DEFAULT NULL,
  `rtsp_ssl_port` int NULL DEFAULT NULL,
  `flv_port` int NULL DEFAULT NULL,
  `flv_ssl_port` int NULL DEFAULT NULL,
  `mp4_port` int NULL DEFAULT NULL,
  `mp4_ssl_port` int NULL DEFAULT NULL,
  `ws_flv_port` int NULL DEFAULT NULL,
  `ws_flv_ssl_port` int NULL DEFAULT NULL,
  `auto_config` tinyint(1) NULL DEFAULT 0,
  `secret` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'zlm',
  `rtp_enable` tinyint(1) NULL DEFAULT 0,
  `rtp_port_range` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `send_rtp_port_range` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `record_assist_port` int NULL DEFAULT NULL,
  `default_server` tinyint(1) NULL DEFAULT 0,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_media_server
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_platform
-- ----------------------------
DROP TABLE IF EXISTS `wvp_platform`;
CREATE TABLE `wvp_platform`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `platform_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `secret` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `port` int NULL DEFAULT 5060,
  `transport` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'UDP',
  `domain` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `enable` tinyint(1) NULL DEFAULT 1,
  `status` tinyint NULL DEFAULT 0,
  `channel_count` int NULL DEFAULT 0,
  `expires` int NULL DEFAULT 3600,
  `auto_register` tinyint(1) NULL DEFAULT 1,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `register_time` datetime(3) NULL DEFAULT NULL,
  `keepalive_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_platform_id`(`platform_id` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_platform
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_platform_channel
-- ----------------------------
DROP TABLE IF EXISTS `wvp_platform_channel`;
CREATE TABLE `wvp_platform_channel`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `platform_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `channel_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `channel_type` tinyint NULL DEFAULT 1,
  `status` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'OFF',
  `longitude` decimal(10, 7) NULL DEFAULT 0.0000000,
  `latitude` decimal(10, 7) NULL DEFAULT 0.0000000,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_platform_id`(`platform_id` ASC) USING BTREE,
  INDEX `idx_channel_id`(`channel_id` ASC) USING BTREE,
  INDEX `idx_device_id`(`device_id` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_platform_channel
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_record_file
-- ----------------------------
DROP TABLE IF EXISTS `wvp_record_file`;
CREATE TABLE `wvp_record_file`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `channel_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `plan_id` int UNSIGNED NULL DEFAULT 0,
  `file_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `file_path` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `file_size` bigint NULL DEFAULT 0,
  `start_time` datetime(3) NULL DEFAULT NULL,
  `end_time` datetime(3) NULL DEFAULT NULL,
  `duration` int NULL DEFAULT 0,
  `storage_type` tinyint NULL DEFAULT 1,
  `indexes` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_device_id`(`device_id` ASC) USING BTREE,
  INDEX `idx_channel_id`(`channel_id` ASC) USING BTREE,
  INDEX `idx_plan_id`(`plan_id` ASC) USING BTREE,
  INDEX `idx_start_time`(`start_time` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_record_file
-- ----------------------------

-- ----------------------------
-- Table structure for wvp_record_plan
-- ----------------------------
DROP TABLE IF EXISTS `wvp_record_plan`;
CREATE TABLE `wvp_record_plan`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `channel_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `stream_type` tinyint NULL DEFAULT 1,
  `record_type` tinyint NULL DEFAULT 1,
  `enable` tinyint(1) NULL DEFAULT 1,
  `status` tinyint NULL DEFAULT 0,
  `start_time` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `end_time` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `file_days` int NULL DEFAULT 30,
  `storage_path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '',
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_device_id`(`device_id` ASC) USING BTREE,
  INDEX `idx_channel_id`(`channel_id` ASC) USING BTREE,
  INDEX `idx_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of wvp_record_plan
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
