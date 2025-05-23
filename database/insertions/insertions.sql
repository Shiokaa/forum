USE forum_db;

-- Insérer des rôles
INSERT INTO roles (name) VALUES
('Admin'),
('Moderator'),
('Member');

-- Insérer des utilisateurs
INSERT INTO users (role_id, name, email, password) VALUES
(1, 'Alice', 'alice.admin@example.com', 'hashed_password1'),
(2, 'Bob', 'bob.mod@example.com', 'hashed_password2'),
(3, 'Charlie', 'charlie.member@example.com', 'hashed_password3'),
(3, 'Diana', 'diana.member@example.com', 'hashed_password4');

-- Insérer des catégories
INSERT INTO categories (name, description) VALUES
('General Discussion', 'Talk about anything and everything.'),
('Tech Talk', 'Discussions related to technology and programming.'),
('Off Topic', 'Anything unrelated to other categories.');

-- Insérer des forums
INSERT INTO forums (category_id, name, description) VALUES
(1, 'Introductions', 'New members introduce themselves.'),
(1, 'Announcements', 'Official forum announcements.'),
(2, 'Programming Languages', 'Discuss various programming languages.'),
(2, 'Hardware', 'Talk about computer hardware and gadgets.'),
(3, 'Random Chat', 'Casual conversation and fun.');

-- Insérer des topics
INSERT INTO topics (forum_id, user_id, title, status) VALUES
(1, 3, 'Hello everyone!', TRUE),
(3, 4, 'What is your favorite programming language?', TRUE),
(4, 2, 'Best CPUs in 2025?', FALSE),
(5, 3, 'Random thoughts', TRUE);

-- Insérer des messages
INSERT INTO messages (topic_id, user_id, content) VALUES
(1, 3, 'Hi all! I am new here and excited to join the community.'),
(1, 1, 'Welcome Charlie! Glad to have you here.'),
(2, 4, 'I love Python because it is so versatile.'),
(2, 2, 'I prefer Rust for its performance and safety.'),
(3, 2, 'The new AMD Ryzen processors are really impressive.'),
(4, 3, 'Sometimes I just like to write random stuff here.');

-- Insérer des feedbacks (likes/dislikes)
INSERT INTO feedbacks (user_id, message_id, type) VALUES
(1, 1, 'like'),
(3, 2, 'like'),
(4, 3, 'like'),
(2, 4, 'dislike'),
(3, 5, 'like'),
(4, 6, 'dislike');

-- Insérer des réponses à des messages
INSERT INTO message_replies (message_id, reply_to_id, content) VALUES
(2, 1, 'Thanks for the warm welcome!'),
(4, 3, 'Rust is great indeed!'),
(6, 5, 'Haha, I feel the same sometimes.');
