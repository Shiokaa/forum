USE forum_db;

-- Rôles
INSERT INTO roles (name) VALUES ('admin'), ('moderator'), ('user');

-- Catégories
INSERT INTO categories (name, description) VALUES
('Technologie', 'Discussions sur les nouveautés tech, le développement et les gadgets.'),
('Culture & Arts', 'Échanges autour de la littérature, du cinéma, de l''histoire et des arts.'),
('Loisirs & Passions', 'Parlez de vos passions : jeux, sport, voyages, cuisine, et bien plus.'),
('Gaming', 'Dédié à l''univers des jeux vidéo, sur consoles, PC et mobiles.'),
('Musique', 'Partagez vos découvertes musicales, discutez de genres et d''artistes.'),
('Vie Pratique', 'Conseils et astuces pour le quotidien, le bricolage, le jardinage.');

-- Utilisateurs
INSERT INTO users (role_id, name, email, password) VALUES
(3, 'Julie Orty', 'Julie@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Jonathan Mira', 'Jonathan@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Alice Dupont', 'alice@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Bob Martin', 'bob@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Charlie Leroy', 'charlie@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Diane Fabre', 'diane@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Éric Lamont', 'eric@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Fiona Girard', 'fiona@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Gael Petit', 'gael@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Hélène Moreau', 'helene@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Igor Lefebvre', 'igor@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8'),
(3, 'Justine Roux', 'justine@example.com', '$2a$12$RI2iB5A8iB.Jb2/C.jJ4a.4JvK.E5uL0zF6qE/I8aB2k4C2b4c6d8');

-- Forums
INSERT INTO forums (category_id, name, description) VALUES
-- Tech
(1, 'Développement Web', 'Frontend, Backend, Fullstack, frameworks et bonnes pratiques.'),
(1, 'Intelligence Artificielle', 'IA, machine learning, éthique et futur de la technologie.'),
(1, 'Smartphones et Mobiles', 'Actualités, tests et débats sur les derniers téléphones.'),
(1, 'Sécurité Informatique', 'Protégez vos données, apprenez les bases du hacking éthique.'),
-- Culture
(2, 'Littérature', 'Romans, poésie, essais, partagez vos lectures et avis.'),
(2, 'Cinéma & Séries TV', 'Discussions et critiques sur les dernières sorties.'),
(2, 'Histoire', 'Des grandes civilisations aux événements plus récents.'),
-- Loisirs
(3, 'Voyages & Aventure', 'Conseils, anecdotes, et photos de vos destinations de rêve.'),
(3, 'Photographie', 'Partagez vos clichés, votre matériel et vos techniques.'),
(3, 'Cuisine & Gastronomie', 'Recettes, astuces de chefs et bonnes adresses.'),
-- Gaming
(4, 'Jeux PC', 'Discussions sur les jeux, les configurations et l''actualité PC.'),
(4, 'Consoles de Salon', 'PlayStation, Xbox, Nintendo : tout sur les consoles.'),
(4, 'Jeux Indépendants', 'Découvrez les pépites et les créations originales.'),
(4, 'Discussions Générales', 'Pour toutes les discussions sur le jeu vidéo qui ne rentrent pas dans les autres forums.'),
-- Musique
(5, 'Rock & Metal', 'Des classiques aux nouveautés, du hard rock au post-rock.'),
(5, 'Musiques Électroniques', 'House, Techno, Trance, Dubstep, ...'),
(5, 'Rap & Hip-Hop', 'Actualités, débats et partage de sons.'),
-- Vie Pratique
(6, 'Bricolage & DIY', 'Rénovation, décoration, vos projets faits maison.'),
(6, 'Jardinage', 'Du potager au jardin d''ornement, échangez vos astuces vertes.');


-- Topics
INSERT INTO topics (forum_id, user_id, title, status) VALUES
(1, 3, 'Comment bien débuter avec React en 2025 ?', true),
(1, 5, 'Node.js vs Deno : quel futur pour le JavaScript côté serveur ?', true),
(1, 9, 'Le CSS moderne : Grid, Flexbox et les autres', true),
(2, 6, 'Les risques de l’IA selon vous ?', true),
(2, 10, 'Top 5 des outils IA qui ont changé ma vie', true),
(3, 1, 'Quel smartphone choisir pour la photo en 2025 ?', true),
(4, 2, 'Comment sécuriser son réseau Wi-Fi domestique ?', true),
(5, 4, 'Votre film de science-fiction préféré de tous les temps ?', true),
(5, 11, 'La dernière saison de cette série m''a déçu...', true),
(7, 8, 'Quelle période de l''Histoire vous fascine le plus ?', true),
(8, 8, 'Voyager seul en Asie du Sud-Est : mes conseils', true),
(8, 12, 'Quel est votre plus beau souvenir de voyage ?', true),
(10, 7, 'Comment réussir sa pâte à pizza maison ?', true),
(11, 9, 'Ma config PC pour faire tourner les derniers jeux AAA', true),
(11, 5, 'Les perles cachées des soldes Steam', true),
(12, 7, 'La Nintendo Switch 2 : vos attentes ?', true),
(12, 3, 'Exclusivités PlayStation vs Xbox : le débat éternel', true),
(14, 6, 'Le retour en force du vinyle, vous en pensez quoi ?', true),
(15, 10, 'Découverte : un artiste de minimale incroyable !', true),
(17, 4, 'Le meilleur outil pour un projet de bricolage ?', true);


-- Messages
INSERT INTO messages (topic_id, user_id, content) VALUES
(1, 3, 'Salut à tous, je suis un peu perdu avec toutes les ressources disponibles pour apprendre React. Par où commencer ? La doc officielle ? Des cours sur Udemy ?'),
(1, 4, 'La doc officielle est vraiment excellente et toujours à jour. C''est le meilleur point de départ.'),
(1, 9, 'Je plussoie pour la doc. Après, pour la pratique, je te conseille de te lancer directement dans un petit projet perso. C''est comme ça qu''on apprend le mieux !'),
(2, 5, 'Avec la popularité grandissante de Deno, je me demande si Node.js a encore de beaux jours devant lui. Qu''en pensez-vous ?'),
(2, 3, 'L''écosystème de Node est tellement immense que Deno aura du mal à le détrôner à court terme. Mais sa modernité est séduisante.'),
(4, 6, 'Je pense que l’IA doit être régulée dès maintenant pour éviter les dérives.'),
(4, 7, 'Les risques existent, mais ne diabolisons pas l’innovation. L''important est l''éducation.'),
(8, 4, 'Pour moi, rien ne détrônera jamais Blade Runner. L''ambiance, la musique, les thèmes abordés... un chef-d''œuvre.'),
(8, 8, 'Difficile de choisir ! Mais je dirais "Interstellar" pour l''ambition scientifique et l''émotion.'),
(8, 11, 'Je suis plus classique : 2001, l''Odyssée de l''espace. Une expérience visuelle et philosophique unique.'),
(11, 8, 'Je reviens de 3 mois en Asie du Sud-Est, si vous avez des questions pour préparer votre voyage, n''hésitez pas !'),
(11, 12, 'Super ! J''aimerais beaucoup partir au Vietnam. Est-ce que c''est facile de se déplacer dans le pays ?'),
(16, 7, 'J''espère vraiment un écran OLED et plus de puissance pour la prochaine console de Nintendo. Et vous ?'),
(16, 9, 'Plus de puissance, c''est certain ! Et surtout, une meilleure autonomie de la batterie.'),
(17, 3, 'En tant que joueur PC, je dois dire que les exclusivités PlayStation qui arrivent sur Steam, c''est un vrai bonheur !'),
(17, 11, 'C''est vrai, mais ça enlève un peu l''intérêt d''avoir la console au final. Le Game Pass de Xbox est plus direct comme stratégie.');

-- Feedbacks (likes/dislikes)
INSERT INTO feedbacks (user_id, message_id, type) VALUES
(5, 2, 'like'),
(3, 3, 'like'),
(7, 4, 'dislike'),
(4, 7, 'like'),
(11, 8, 'like'),
(4, 9, 'like'),
(12, 12, 'like'),
(7, 14, 'like');

-- Reponses aux messages
INSERT INTO message_replies (user_id, reply_to_id, content) VALUES
(3, 3, 'Merci pour le conseil, je vais me lancer dans un petit portfolio alors !'),
(4, 9, 'Totalement d''accord pour Interstellar, la bande son de Hans Zimmer est magistrale.'),
(8, 12, 'Oui, très facile ! Les bus sont nombreux et peu chers. Tu peux aussi louer un scooter pour plus de liberté.');