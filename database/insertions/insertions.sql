-- Rôles (inchangé)
INSERT INTO roles (name) VALUES ('admin'), ('moderator'), ('user');

-- Catégories (avec ajouts)
INSERT INTO categories (name, description) VALUES
('Technologie', 'Discussions sur les nouveautés tech, le développement et les gadgets.'),
('Culture', 'Échanges autour de la littérature, du cinéma, de l''histoire et des arts.'),
('Loisirs', 'Parlez de vos passions : jeux, sport, voyages, cuisine, et bien plus.'),
('Gaming', 'Dédié à l''univers des jeux vidéo, sur consoles, PC et mobiles.'),
('Musique', 'Partagez vos découvertes musicales, discutez de genres et d''artistes.');

-- Utilisateurs (avec ajouts)
INSERT INTO users (role_id, name, email, password) VALUES
(1, 'Alice Dupont', 'alice@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(2, 'Bob Martin', 'bob@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Charlie Leroy', 'charlie@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Diane Fabre', 'diane@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(2, 'Éric Lamont', 'eric@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Fiona Girard', 'fiona@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Gael Petit', 'gael@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Hélène Moreau', 'helene@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Igor Lefebvre', 'igor@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Justine Roux', 'justine@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8');

-- Forums (avec ajouts)
-- Catégorie 1: Technologie
INSERT INTO forums (category_id, name, description) VALUES
(1, 'Développement Web', 'Frontend, Backend, Fullstack, frameworks et bonnes pratiques.'),
(1, 'Intelligence Artificielle', 'IA, machine learning, éthique et futur de la technologie.'),
(1, 'Smartphones et Mobiles', 'Actualités, tests et débats sur les derniers téléphones.'),
-- Catégorie 2: Culture
(2, 'Littérature', 'Romans, poésie, essais, partagez vos lectures et avis.'),
(2, 'Cinéma & Séries TV', 'Discussions et critiques sur les dernières sorties.'),
-- Catégorie 3: Loisirs
(3, 'Voyages', 'Conseils, anecdotes, et photos de vos destinations de rêve.'),
(3, 'Photographie', 'Partagez vos clichés, votre matériel et vos techniques.'),
-- Catégorie 4: Gaming
(4, 'Jeux PC', 'Discussions sur les jeux, les configurations et l''actualité PC.'),
(4, 'Consoles de Salon', 'PlayStation, Xbox, Nintendo : tout sur les consoles.'),
-- Catégorie 5: Musique
(5, 'Rock & Metal', 'Des classiques aux nouveautés, du hard rock au post-rock.'),
(5, 'Musiques Électroniques', 'House, Techno, Trance, Dubstep, ...');

-- Topics (avec ajouts, 15 topics au total)
INSERT INTO topics (forum_id, user_id, title, status) VALUES
-- Forum 1: Dév Web
(1, 1, 'Comment bien débuter avec React en 2025 ?', true),
(1, 3, 'Node.js vs Deno : quel futur pour le JavaScript côté serveur ?', true),
(1, 7, 'Le CSS moderne : Grid, Flexbox et les autres', true),
-- Forum 2: IA
(2, 4, 'Les risques de l’IA selon vous ?', true),
(2, 8, 'Top 5 des outils IA qui ont changé ma vie', true),
-- Forum 5: Cinéma
(5, 2, 'Votre film de science-fiction préféré de tous les temps ?', true),
(5, 9, 'La dernière saison de cette série m''a déçu...', true),
-- Forum 6: Voyages
(6, 6, 'Voyager seul en Asie du Sud-Est : mes conseils', true),
(6, 10, 'Quel est votre plus beau souvenir de voyage ?', true),
-- Forum 8: Jeux PC
(8, 7, 'Ma config PC pour faire tourner les derniers jeux AAA', true),
(8, 3, 'Les perles cachées des soldes Steam', true),
-- Forum 9: Consoles
(9, 5, 'La Nintendo Switch 2 : vos attentes ?', true),
(9, 1, 'Exclusivités PlayStation vs Xbox : le débat éternel', true),
-- Forum 10: Rock
(10, 4, 'Le retour en force du vinyle, vous en pensez quoi ?', true),
-- Forum 11: Electro
(11, 8, 'Découverte : un artiste de minimale incroyable !', true);

-- Messages
-- Topic 1: React
INSERT INTO messages (topic_id, user_id, content) VALUES
(1, 1, 'Salut à tous, je suis un peu perdu avec toutes les ressources disponibles pour apprendre React. Par où commencer ? La doc officielle ? Des cours sur Udemy ?'),
(1, 2, 'La doc officielle est vraiment excellente et toujours à jour. C''est le meilleur point de départ.'),
(1, 7, 'Je plussoie pour la doc. Après, pour la pratique, je te conseille de te lancer directement dans un petit projet perso. C''est comme ça qu''on apprend le mieux !'),
-- Topic 2: Node vs Deno
(2, 3, 'Avec la popularité grandissante de Deno, je me demande si Node.js a encore de beaux jours devant lui. Qu''en pensez-vous ?'),
(2, 1, 'L''écosystème de Node est tellement immense que Deno aura du mal à le détrôner à court terme. Mais sa modernité est séduisante.'),
-- Topic 4: Risques IA
(4, 4, 'Je pense que l’IA doit être régulée dès maintenant pour éviter les dérives.'),
(4, 5, 'Les risques existent, mais ne diabolisons pas l’innovation. L''important est l''éducation.'),
-- Topic 6: Film SF
(6, 2, 'Pour moi, rien ne détrônera jamais Blade Runner. L''ambiance, la musique, les thèmes abordés... un chef-d''œuvre.'),
(6, 6, 'Difficile de choisir ! Mais je dirais "Interstellar" pour l''ambition scientifique et l''émotion.'),
(6, 9, 'Je suis plus classique : 2001, l''Odyssée de l''espace. Une expérience visuelle et philosophique unique.'),
-- Topic 8: Voyage Asie
(8, 6, 'Je reviens de 3 mois en Asie du Sud-Est, si vous avez des questions pour préparer votre voyage, n''hésitez pas !'),
(8, 10, 'Super ! J''aimerais beaucoup partir au Vietnam. Est-ce que c''est facile de se déplacer dans le pays ?'),
-- Topic 12: Switch 2
(12, 5, 'J''espère vraiment un écran OLED et plus de puissance pour la prochaine console de Nintendo. Et vous ?'),
(12, 7, 'Plus de puissance, c''est certain ! Et surtout, une meilleure autonomie de la batterie.'),
-- Topic 13: Exclusivités
(13, 1, 'En tant que joueur PC, je dois dire que les exclusivités PlayStation qui arrivent sur Steam, c''est un vrai bonheur !'),
(13, 9, 'C''est vrai, mais ça enlève un peu l''intérêt d''avoir la console au final. Le Game Pass de Xbox est plus direct comme stratégie.');

-- Feedbacks (likes/dislikes)
INSERT INTO feedbacks (user_id, message_id, type) VALUES
(3, 2, 'like'),
(1, 3, 'like'),
(5, 4, 'dislike'),
(2, 7, 'like'),
(9, 8, 'like'),
(2, 9, 'like'),
(10, 12, 'like'),
(5, 14, 'like');

-- Reponses aux messages
INSERT INTO message_replies (user_id, reply_to_id, content) VALUES
(1, 3, 'Merci pour le conseil, je vais me lancer dans un petit portfolio alors !'),
(2, 9, 'Totalement d''accord pour Interstellar, la bande son de Hans Zimmer est magistrale.'),
(6, 10, 'Oui, très facile ! Les bus sont nombreux et peu chers. Tu peux aussi louer un scooter pour plus de liberté.');