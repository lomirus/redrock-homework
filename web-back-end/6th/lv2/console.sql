SET NAMES utf8;

CREATE DATABASE IF NOT EXISTS school;
USE school;
DROP TABLE IF EXISTS `students`;
DROP TABLE IF EXISTS `classes`;
DROP TABLE IF EXISTS `points`;

CREATE TABLE `students`  (
    `stu_id` BIGINT NOT NULL DEFAULT 0 PRIMARY KEY ,
    `class_id` BIGINT NOT NULL DEFAULT 0,
    `name` VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `password` VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

CREATE TABLE `classes`  (
    `class_id` BIGINT NOT NULL DEFAULT 0 PRIMARY KEY ,
    `subject` VARCHAR(100) NOT NULL,
    `college` VARCHAR(100) NOT NULL,
    `student_number` BIGINT NOT NULL DEFAULT 0
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

CREATE TABLE `points` (
    `subject` VARCHAR (100) NOT NULL,
    `student_id` BIGINT NOT NULL,
    `points` DOUBLE NOT NULL,
    PRIMARY KEY (`subject`, `student_id`)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

ALTER TABLE `students` ADD CONSTRAINT fk_class_id FOREIGN KEY (`class_id`) REFERENCES classes (`class_id`);

INSERT INTO `classes` VALUES (01011906, '通信工程','通信与信息工程学院',27);
INSERT INTO `students` VALUES (2019233333, 01011906,'王小美', 'Wangxiaomei011224');
INSERT INTO `points` VALUES ('MATH',2019233333,3.2);
INSERT INTO `points` VALUES ('SPORT',2019233333,1.0);
INSERT INTO `points` VALUES ('ART',2019233333,3.6);

-- get avg_point:
SELECT AVG(`points`) FROM  `points` where `student_id` = 2019233333;

-- get good_subjects
SELECT `subject` FROM `points`
where `student_id` = 2019233333 and `points` > 3.0;



