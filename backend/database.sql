CREATE TABLE "Tag"(
  "ID" TEXT NOT NULL,
  "Title" TEXT NOT NULL,
  "CreatedAt" INTEGER NOT NULL,
  PRIMARY KEY("ID"),
  UNIQUE("Title")
);

CREATE TABLE "Comment"(
  "ID" TEXT NOT NULL,
  "UserID" TEXT NOT NULL,
  "PostID" TEXT NOT NULL,
  "Content" TEXT NOT NULL,
  "CreatedAt" INTEGER NOT NULL,
  PRIMARY KEY("ID"),
  CONSTRAINT posts_comments
    FOREIGN KEY ("PostID") REFERENCES "Post" ("ID") ON DELETE No action
      ON UPDATE No action,
  CONSTRAINT users_comments FOREIGN KEY ("UserID") REFERENCES "User" ("ID")
);

CREATE TABLE "Post"(
  "ID" TEXT NOT NULL,
  "UserID" TEXT NOT NULL,
  "Content" TEXT NOT NULL,
  "TagID" TEXT NOT NULL,
  "CreatedAt" INTEGER NOT NULL,
  PRIMARY KEY("ID"),
  CONSTRAINT users_posts
    FOREIGN KEY ("UserID") REFERENCES "User" ("ID") ON DELETE No action
      ON UPDATE No action,
  CONSTRAINT "Category_Post" FOREIGN KEY ("TagID") REFERENCES "Tag" ("ID")
);

CREATE TABLE "User"(
  "ID" TEXT NOT NULL,
  "Email" TEXT NOT NULL,
  "Password" BLOB NOT NULL,
  "Nickname" TEXT NOT NULL,
  "FirstName" TEXT NOT NULL,
  "LastName" TEXT NOT NULL,
  "Gender" TEXT NOT NULL,
  "Age" INTEGER NOT NULL,
  "Color" TEXT NOT NULL,
  "CreatedAt" INTEGER NOT NULL,
  PRIMARY KEY("ID"),
  UNIQUE("Email")
);

CREATE TABLE "Chat"(
  "ID" TEXT NOT NULL,
  "LastMessageTime" INTEGER NOT NULL,
  "CreatedAt" INTEGER NOT NULL,
  PRIMARY KEY("ID")
);

CREATE TABLE "Message"(
  "ID" TEXT NOT NULL,
  "ChatID" TEXT NOT NULL,
  "UserID" TEXT NOT NULL,
  "Content" TEXT NOT NULL,
  "CreatedAt" INTEGER NOT NULL,
  PRIMARY KEY("ID"),
  CONSTRAINT users_messages FOREIGN KEY ("UserID") REFERENCES "User" ("ID"),
  CONSTRAINT "chatTotals_messages" FOREIGN KEY ("ChatID") REFERENCES "Chat" ("ID")
);

CREATE TABLE "Session"(
  "ID" TEXT NOT NULL,
  "UserID" TEXT NOT NULL,
  "CreatedAt" INTEGER NOT NULL,
  PRIMARY KEY("ID"),
  CONSTRAINT "User_Session" FOREIGN KEY ("UserID") REFERENCES "User" ("ID")
);
