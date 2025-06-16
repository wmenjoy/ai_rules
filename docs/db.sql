-- Create database
CREATE DATABASE IF NOT EXISTS prompt_template_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE prompt_template_system;

-- Templates table
CREATE TABLE templates (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    usage_count INT UNSIGNED DEFAULT 0,
    is_favorite BOOLEAN DEFAULT FALSE,
    template_type VARCHAR(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Variables table
CREATE TABLE variables (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    template_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    default_value TEXT,
    required BOOLEAN DEFAULT FALSE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Categories table
CREATE TABLE categories (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tags table
CREATE TABLE tags (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Template-Tags relationship table (many-to-many)
CREATE TABLE template_tags (
    template_id BIGINT UNSIGNED NOT NULL,
    tag_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (template_id, tag_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Activities table
CREATE TABLE activities (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    template_id BIGINT UNSIGNED,
    template_title VARCHAR(255) NOT NULL,
    action ENUM('used', 'created', 'edited', 'deleted') NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user VARCHAR(100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes for better performance
CREATE INDEX idx_templates_title ON templates(title);
CREATE INDEX idx_templates_updated_at ON templates(updated_at);
CREATE INDEX idx_templates_usage_count ON templates(usage_count);
CREATE INDEX idx_templates_is_favorite ON templates(is_favorite);
CREATE INDEX idx_templates_template_type ON templates(template_type);
CREATE INDEX idx_variables_template_id ON variables(template_id);
CREATE INDEX idx_variables_name ON variables(name);
CREATE INDEX idx_activities_template_id ON activities(template_id);
CREATE INDEX idx_activities_timestamp ON activities(timestamp);
CREATE INDEX idx_activities_action ON activities(action);

-- Sample data for testing
INSERT INTO categories (name) VALUES 
('客户服务'), 
('内容创作'), 
('营销');

INSERT INTO tags (name) VALUES 
('回复'), 
('创意'), 
('客户');

-- Sample template
INSERT INTO templates (title, description, content, usage_count, is_favorite, template_type) VALUES (
    '客户服务回复模板',
    '用于回复客户服务请求的模板，适用于各种场景的客户查询',
    '尊敬的{{customer_name}}，\n\n感谢您联系我们的客户服务团队。关于您提到的{{issue}}，我们理解这给您带来了不便。\n\n{{solution}}\n\n如果您有任何其他问题，请随时联系我们。\n\n此致，\n{{agent_name}}\n客户支持团队',
    120,
    TRUE,
    'service'
);

-- Sample variables for the template
INSERT INTO variables (template_id, name, description, default_value, required) VALUES 
(1, 'customer_name', '客户姓名', '', FALSE),
(1, 'issue', '客户问题描述', '', TRUE),
(1, 'solution', '问题解决方案', '', TRUE),
(1, 'agent_name', '客服人员姓名', '客服团队', FALSE);

-- Connect template to tags
INSERT INTO template_tags (template_id, tag_id) VALUES 
(1, 1), 
(1, 3);

-- Sample activity
INSERT INTO activities (template_id, template_title, action, user) VALUES 
(1, '客户服务回复模板', 'created', '系统管理员'); 