-- Rôles
INSERT INTO roles (name) VALUES ('admin'), ('moderator'), ('user');

-- Utilisateurs
INSERT INTO users (role_id, name, email, password) VALUES
(1, 'Alice Dupont', 'alice@example.com', 'hashed_pw1'),
(2, 'Bob Martin', 'bob@example.com', 'hashed_pw2'),
(3, 'Charlie Leroy', 'charlie@example.com', 'hashed_pw3');

-- Catégories
INSERT INTO categories (name, description) VALUES
('Technologie', 'Discussions sur les nouveautés tech'),
('Culture', 'Échanges autour de la culture générale');

-- Forums
INSERT INTO forums (category_id, name, description) VALUES
(1, 'Développement Web', 'Frontend, Backend, Fullstack...'),
(2, 'Littérature', 'Romans, poésie, essais...');

-- Topics
INSERT INTO topics (forum_id, user_id, title, status) VALUES
(1, 1, 'Comment apprendre React ?', true),
(2, 2, 'Vos romans préférés ?', true);

-- Messages
INSERT INTO messages (topic_id, user_id, content) VALUES
(1, 2, 'Tu peux commencer avec la doc officielle, elle est bien faite.'),
(1, 3, 'Je recommande aussi des vidéos sur YouTube.'),
(2, 1, 'J\'ai adoré "Le Comte de Monte-Cristo", un classique intemporel.');

-- Feedbacks
INSERT INTO feedbacks (user_id, message_id, type) VALUES
(1, 1, 'like'),
(3, 1, 'like'),
(2, 3, 'dislike');

-- Réponses aux messages
INSERT INTO message_replies (user_id, reply_to_id, content) VALUES
(1, 1, 'Merci pour la recommandation, je vais tester !'),
(3, 2, 'Oui les vidéos de "OpenClassrooms" sont top aussi.');
