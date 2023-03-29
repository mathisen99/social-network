CREATE TABLE Users (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    dob DATE NOT NULL,
    avatar TEXT,
    nickname TEXT,
    about_me TEXT
);

CREATE TABLE Followers (
    follower_id INTEGER,
    followee_id INTEGER,
    request_status TEXT NOT NULL,
    PRIMARY KEY (follower_id, followee_id),
    FOREIGN KEY (follower_id) REFERENCES Users (id),
    FOREIGN KEY (followee_id) REFERENCES Users (id)
);

CREATE TABLE Profiles (
    user_id INTEGER PRIMARY KEY,
    profile_type TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE Posts (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    image TEXT,
    privacy TEXT NOT NULL,
    creation_date DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE Groups (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    creator_id INTEGER NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES Users (id)
);

CREATE TABLE GroupMembers (
    group_id INTEGER,
    user_id INTEGER,
    membership_status TEXT NOT NULL,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES Groups (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE GroupPosts (
    id INTEGER PRIMARY KEY,
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    image TEXT,
    creation_date DATETIME NOT NULL,
    FOREIGN KEY (group_id) REFERENCES Groups (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE Events (
    id INTEGER PRIMARY KEY,
    group_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    date_time DATETIME NOT NULL,
    FOREIGN KEY (group_id) REFERENCES Groups (id)
);

CREATE TABLE EventResponses (
    event_id INTEGER,
    user_id INTEGER,
    response TEXT NOT NULL,
    PRIMARY KEY (event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES Events (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE PrivateMessages (
    id INTEGER PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    timestamp DATETIME NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES Users (id),
    FOREIGN KEY (receiver_id) REFERENCES Users (id)
);

CREATE TABLE GroupChats (
    id INTEGER PRIMARY KEY,
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    timestamp DATETIME NOT NULL,
    FOREIGN KEY (group_id) REFERENCES Groups (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE Notifications (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    timestamp DATETIME NOT NULL,
    viewed BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id)
);