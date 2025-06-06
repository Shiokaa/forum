-- Rôles
INSERT INTO roles (name) VALUES ('admin'), ('moderator'), ('user');

-- Utilisateurs
INSERT INTO users (role_id, name, email, password) VALUES
(1, 'Alice Dupont', 'alice@example.com', 'hashed_pw1'),
(2, 'Bob Martin', 'bob@example.com', 'hashed_pw2'),
(3, 'Charlie Leroy', 'charlie@example.com', 'hashed_pw3'),
(3, 'Diane Fabre', 'diane@example.com', 'hashed_pw4'),
(2, 'Éric Lamont', 'eric@example.com', 'hashed_pw5');

-- Catégories
INSERT INTO categories (name, description) VALUES
('Technologie', 'Discussions sur les nouveautés tech'),
('Culture', 'Échanges autour de la culture générale'),
('Loisirs', 'Jeux, sport, voyages, etc.');

-- Forums
INSERT INTO forums (category_id, name, description) VALUES
(1, 'Développement Web', 'Frontend, Backend, Fullstack...'),
(1, 'Intelligence Artificielle', 'IA, machine learning, éthique...'),
(2, 'Littérature', 'Romans, poésie, essais...'),
(3, 'Voyages', 'Conseils, anecdotes, destinations');

-- Topics
INSERT INTO topics (forum_id, user_id, title, status) VALUES
(1, 1, 'Comment apprendre React ?', true),
(2, 3, 'Les risques de l’IA selon vous ?', true),
(3, 2, 'Vos romans préférés ?', true),
(4, 4, 'Voyager seul : pour ou contre ?', true);

-- Messages
INSERT INTO messages (topic_id, user_id, content) VALUES
(1, 2, 'Tu peux commencer avec la doc officielle, elle est bien faite.'),
(1, 3, 'Je recommande aussi des vidéos sur YouTube.'),
(2, 4, 'Je pense que l’IA doit être régulée dès maintenant.'),
(2, 5, 'Les risques existent, mais ne diabolisons pas l’innovation.'),
(3, 1, 'J\'ai adoré "Le Comte de Monte-Cristo", un classique intemporel.'),
(3, 4, 'Je recommande "L\'Étranger" de Camus.'),
(4, 3, 'Voyager seul permet de vraiment se découvrir.'),
(4, 2, 'Mais c’est parfois dangereux, surtout dans certains pays.');

-- Feedbacks
INSERT INTO feedbacks (user_id, message_id, type) VALUES
(1, 1, 'like'),
(3, 1, 'like'),
(2, 3, 'dislike'),
(4, 4, 'like'),
(5, 6, 'like'),
(1, 7, 'like');

-- Réponses aux messages
INSERT INTO message_replies (user_id, reply_to_id, content) VALUES
(1, 1, 'Merci pour la recommandation, je vais tester !'),
(3, 2, 'Oui les vidéos de "OpenClassrooms" sont top aussi.'),
(2, 4, 'Je suis d’accord, mais il faut encadrer les usages.'),
(5, 6, 'Excellent choix, j’ai aussi beaucoup aimé ce livre.'),
(4, 7, 'Totalement d\'accord, j\'ai eu la même expérience.');

