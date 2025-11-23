INSERT INTO tracks (id, name, country, length_km, turns, created_year)
VALUES 
('spa-francorchamps', 'Circuit de Spa-Francorchamps', 'Belgique', 7.004, 19, 1921),
('monza', 'Autodromo Nazionale di Monza', 'Italie', 5.793, 11, 1922);

INSERT INTO sections (track_id, name, description) VALUES
('spa-francorchamps', 'Eau Rouge / Raidillon', 'Enchaînement mythique à très haute vitesse.'),
('spa-francorchamps', 'Blanchimont', 'Virage à gauche ultra rapide.'),
('monza', 'Lesmo', 'Deux virages rapides clés du circuit.'),
('monza', 'Parabolica', 'Dernier virage emblématique de Monza.');

INSERT INTO winners (track_id, year, driver, team) VALUES
('spa-francorchamps', 2024, 'Max Verstappen', 'Red Bull'),
('spa-francorchamps', 2023, 'Max Verstappen', 'Red Bull'),
('monza', 2024, 'Charles Leclerc', 'Ferrari');
