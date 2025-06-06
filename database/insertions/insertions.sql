USE forum_db;

-- Rôles
INSERT INTO roles (name) VALUES
('admin'),
('moderator'),
('member');

-- Utilisateurs
INSERT INTO users (role_id, name, email, password) VALUES
(1, 'Alice Admin', 'alice@forum.com', 'hashed_password_1'),
(2, 'Bob Moderator', 'bob@forum.com', 'hashed_password_2'),
(3, 'Charlie Member', 'charlie@forum.com', 'hashed_password_3'),
(3, 'Diana Member', 'diana@forum.com', 'hashed_password_4');

-- Catégories
INSERT INTO categories (name, description) VALUES
('Général', 'Discussions générales'),
('Tech', 'Discussions technologiques'),
('Jeux', 'Jeux vidéos et discussions liées');

-- Forums
INSERT INTO forums (category_id, name, description) VALUES
(1, 'Présentations', 'Présentez-vous ici !'),
(2, 'Programmation', 'Discutez de code et de dev ici.'),
(3, 'Jeux PC', 'Pour les amateurs de jeux sur PC');

-- Sujets
INSERT INTO topics (forum_id, user_id, title, status) VALUES
(1, 3, 'Salut tout le monde !', TRUE),
(2, 1, 'Comment apprendre Python ?', TRUE),
(3, 4, 'Votre jeu préféré ?', TRUE);

-- Messages
INSERT INTO messages (topic_id, user_id, content) VALUES
(1, 3, 'Je m\'appelle Charlie, ravi d\'être ici.'),
(2, 1, 'Commence par les bases, et pratique beaucoup.'),
(3, 4, 'J\'adore The Witcher 3, et vous ?'),
(3, 3, 'Moi je suis plutôt fan de Skyrim.');

-- Feedbacks
INSERT INTO feedbacks (user_id, message_id, type) VALUES
(1, 1, 'like'),
(2, 2, 'like'),
(3, 3, 'dislike'),
(4, 4, 'like');

-- Réponses aux messages
INSERT INTO message_replies (reply_to_id, content) VALUES
(1, 'Bienvenue Charlie ! Content de te voir ici.'),
(2, 'Merci pour le conseil, je vais commencer avec les bases de Python.'),
(3, 'The Witcher 3 est un excellent choix !'),
(4, 'Skyrim est aussi un classique incontournable.');
