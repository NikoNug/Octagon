CREATE TABLE users (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Firstname VARCHAR(100) NOT NULL,
    Lastname VARCHAR(100) NOT NULL,
    Username VARCHAR(100) UNIQUE NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Address TEXT,
    City VARCHAR(100)
);

-- CREATE TABLE posts (
--     ID INT AUTO_INCREMENT PRIMARY KEY,
--     Title VARCHAR(255) NOT NULL,
--     Body TEXT NOT NULL
-- );

CREATE TABLE history (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    PostID INT NOT NULL,
    PersonID INT NOT NULL,
    FOREIGN KEY (PostID) REFERENCES posts(PostID) ON DELETE CASCADE,
    FOREIGN KEY (PersonID) REFERENCES users(UserID) ON DELETE CASCADE
);

CREATE TABLE posts (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,                    -- Referencing the user who created the post
    Title VARCHAR(255) NOT NULL,            -- Post title
    Body TEXT NOT NULL,                     -- Post content (can store text, links, etc.)
    ImageURL VARCHAR(255),                  -- URL for image (optional)
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES users(UserID) ON DELETE CASCADE
);

CREATE TABLE comments (
    CommentID INT AUTO_INCREMENT PRIMARY KEY,
    PostID INT NOT NULL,                    -- Referencing the post
    UserID INT NOT NULL,                    -- Referencing the user who made the comment
    Comment TEXT NOT NULL,                  -- The comment text
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (PostID) REFERENCES posts(PostID) ON DELETE CASCADE,
    FOREIGN KEY (UserID) REFERENCES users(UserID) ON DELETE CASCADE
);

CREATE TABLE likes (
    LikeID INT AUTO_INCREMENT PRIMARY KEY,
    PostID INT NOT NULL,                    -- Referencing the post
    UserID INT NOT NULL,                    -- Referencing the user who liked the post
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (PostID) REFERENCES posts(PostID) ON DELETE CASCADE,
    FOREIGN KEY (UserID) REFERENCES users(UserID) ON DELETE CASCADE
);

