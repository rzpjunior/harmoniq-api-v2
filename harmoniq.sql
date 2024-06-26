-- Create the database
CREATE DATABASE IF NOT EXISTS "harmoniq-dev";

-- Use the database
USE "harmoniq-dev";

-- Create artists table
CREATE TABLE IF NOT EXISTS artists (
    artist_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    genre VARCHAR(100),
    country VARCHAR(100),
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create albums table
CREATE TABLE IF NOT EXISTS albums (
    album_id INT AUTO_INCREMENT PRIMARY KEY,
    artist_id INT,
    title VARCHAR(255) NOT NULL,
    release_date DATE,
    genre VARCHAR(100),
    cover_image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (artist_id) REFERENCES artists(artist_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create songs table
CREATE TABLE IF NOT EXISTS songs (
    song_id INT AUTO_INCREMENT PRIMARY KEY,
    album_id INT NULL,
    title VARCHAR(255) NOT NULL,
    duration TIME,
    track_number INT,
    audio_file_url VARCHAR(255),
    image_url VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (album_id) REFERENCES albums(album_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create artist_song relationship table
CREATE TABLE IF NOT EXISTS artist_song (
    artist_id INT,
    song_id INT,
    PRIMARY KEY (artist_id, song_id),
    FOREIGN KEY (artist_id) REFERENCES artists(artist_id),
    FOREIGN KEY (song_id) REFERENCES songs(song_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create user table
CREATE TABLE IF NOT EXISTS user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) DEFAULT '',
    phone VARCHAR(15) DEFAULT '',
    password VARCHAR(250) DEFAULT '',
    name VARCHAR(100) DEFAULT '',
    status TINYINT(1) DEFAULT '0',
    last_login_at DATETIME NULL DEFAULT NULL,
    created_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create favorites table
CREATE TABLE IF NOT EXISTS favorites (
    user_id BIGINT UNSIGNED,
    song_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, song_id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (song_id) REFERENCES songs(song_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create playlists table
CREATE TABLE IF NOT EXISTS playlists (
    playlist_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create playlist_songs relationship table
CREATE TABLE IF NOT EXISTS playlist_songs (
    playlist_id INT,
    song_id INT,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (playlist_id, song_id),
    FOREIGN KEY (playlist_id) REFERENCES playlists(playlist_id),
    FOREIGN KEY (song_id) REFERENCES songs(song_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert artists
INSERT INTO artists (name, genre, country, bio) VALUES 
('Artist One', 'Pop', 'USA', 'Artist One bio.'),
('Artist Two', 'Rock', 'UK', 'Artist Two bio.'),
('Artist Three', 'Jazz', 'Canada', 'Artist Three bio.');

-- Insert albums
INSERT INTO albums (artist_id, title, release_date, genre, cover_image_url) VALUES 
(1, 'Album One', '2021-01-01', 'Pop', 'http://example.com/album1.jpg'),
(2, 'Album Two', '2022-02-02', 'Rock', 'http://example.com/album2.jpg'),
(3, 'Album Three', '2023-03-03', 'Jazz', 'http://example.com/album3.jpg');

-- Insert songs (including singles with image_url)
INSERT INTO songs (album_id, title, duration, track_number, audio_file_url, image_url) VALUES 
(1, 'Song One', '00:03:30', 1, 'http://example.com/song1.mp3', 'http://example.com/song1.jpg'),
(1, 'Song Two', '00:04:00', 2, 'http://example.com/song2.mp3', 'http://example.com/song2.jpg'),
(2, 'Song Three', '00:05:00', 1, 'http://example.com/song3.mp3', 'http://example.com/song3.jpg'),
(2, 'Song Four', '00:03:45', 2, 'http://example.com/song4.mp3', 'http://example.com/song4.jpg'),
(3, 'Song Five', '00:02:50', 1, 'http://example.com/song5.mp3', 'http://example.com/song5.jpg'),
(3, 'Song Six', '00:04:20', 2, 'http://example.com/song6.mp3', 'http://example.com/song6.jpg'),
(NULL, 'Single One', '00:03:15', NULL, 'http://example.com/single1.mp3', 'http://example.com/single1.jpg'),
(NULL, 'Single Two', '00:03:45', NULL, 'http://example.com/single2.mp3', 'http://example.com/single2.jpg');

-- Insert artist-song relationships (including singles)
INSERT INTO artist_song (artist_id, song_id) VALUES 
(1, 1),
(1, 2),
(2, 3),
(2, 4),
(3, 5),
(3, 6),
(1, 7),
(2, 8);

-- Insert users
INSERT INTO user (email, phone, password, name, status, last_login_at, created_at) VALUES 
('user1@example.com', '1234567890', 'hashed_password1', 'User One', 1, '2024-01-01 12:00:00', '2024-01-01 12:00:00'),
('user2@example.com', '0987654321', 'hashed_password2', 'User Two', 1, '2024-01-02 12:00:00', '2024-01-02 12:00:00'),
('user3@example.com', '1122334455', 'hashed_password3', 'User Three', 1, '2024-01-03 12:00:00', '2024-01-03 12:00:00');

-- Insert favorites (including singles)
INSERT INTO favorites (user_id, song_id) VALUES 
(1, 1),
(1, 3),
(2, 2),
(2, 4),
(3, 5),
(3, 6),
(1, 7),
(2, 8);

-- Insert playlists
INSERT INTO playlists (user_id, name, description) VALUES 
(1, 'Chill Vibes', 'Relax and enjoy the music.'),
(2, 'Workout Mix', 'Get pumped with these tracks.'),
(3, 'Jazz Collection', 'All-time favorite Jazz songs.');

-- Insert playlist-songs relationships (including singles)
INSERT INTO playlist_songs (playlist_id, song_id) VALUES 
(1, 1),
(1, 2),
(1, 3),
(2, 3),
(2, 4),
(3, 5),
(3, 6),
(1, 7),
(2, 8);
