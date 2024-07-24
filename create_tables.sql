-- Table: challenges
CREATE TABLE challenges (
    id INT PRIMARY KEY AUTO_INCREMENT,
    player_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    won BOOLEAN NOT NULL
);

-- Table: challenge_results
CREATE TABLE challenge_results (
    id INT PRIMARY KEY AUTO_INCREMENT,
    player_id INT NOT NULL,
    won_jackpot BOOLEAN NOT NULL,
    created_at BIGINT NOT NULL
);

-- Table: levels
CREATE TABLE levels (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

-- Table: logs
CREATE TABLE logs (
    id INT PRIMARY KEY AUTO_INCREMENT,
    player_id INT NOT NULL,
    action VARCHAR(255) NOT NULL,
    timestamp BIGINT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

-- Table: payments
CREATE TABLE payments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    method VARCHAR(255) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    details TEXT,
    created_at BIGINT NOT NULL
);

-- Table: players
CREATE TABLE players (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    level VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL
);

-- Table: reservations
CREATE TABLE reservations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    room_id INT NOT NULL,
    date DATE NOT NULL,
    time VARCHAR(255) NOT NULL,
    player_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: rooms
CREATE TABLE rooms (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(255)
);
