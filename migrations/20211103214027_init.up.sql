CREATE TABLE players (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(32) NOT NULL DEFAULT "",
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE hands (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  game_date DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE half_round_games (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  player_id INT UNSIGNED NOT NULL,
  hand_id INT UNSIGNED NOT NULL,
  game_count INT UNSIGNED NOT NULL DEFAULT 1,
  score INT NOT NULL DEFAULT 0,
  ranking INT UNSIGNED NOT NULL DEFAULT 1,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY fk_player_id (player_id) REFERENCES players(id),
  FOREIGN KEY fk_hand_id (hand_id) REFERENCES hands(id)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
CREATE TABLE players_hands (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  player_id INT UNSIGNED NOT NULL,
  hand_id INT UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY fk_player_id (player_id) REFERENCES players(id),
  FOREIGN KEY fk_hand_id (hand_id) REFERENCES hands(id)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
