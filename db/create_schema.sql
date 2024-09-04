CREATE SCHEMA IF NOT EXISTS rewild AUTHORIZATION CURRENT_USER;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS rewild.users(
    user_id         UUID DEFAULT gen_random_uuid(),
    first_name      VARCHAR(255) NOT NULL,
    last_name       VARCHAR(255) NOT NULL,
    username        VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    created_ts      TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS rewild.pindrops(
    pindrop_id      UUID DEFAULT gen_random_uuid(),
    longitude       FLOAT NOT NULL,
    latitude        FLOAT NOT NULL,
    created_ts      TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY(pindrop_id)
);

CREATE TABLE IF NOT EXISTS rewild.images(
    image_id            UUID DEFAULT gen_random_uuid(),
    alt_text            TEXT,
    created_ts      TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY(image_id)
);

CREATE TABLE IF NOT EXISTS rewild.timeline_posts(
    timeline_post_id    UUID DEFAULT gen_random_uuid(),
    next_id             UUID NULL,
    prev_id             UUID NULL,
    title               VARCHAR(255) NOT NULL,
    body                TEXT,
    created_ts      TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY(timeline_post_id),
    CONSTRAINT fk_next
        FOREIGN KEY(next_id)
            REFERENCES rewild.timeline_posts(timeline_post_id),
    CONSTRAINT fk_prev
        FOREIGN KEY(prev_id)
            REFERENCES rewild.timeline_posts(timeline_post_id)
);

CREATE TABLE IF NOT EXISTS rewild.timeline_post_images(
    timeline_post_id    UUID NOT NULL,
    image_id            UUID NOT NULL,
    arr_index           INT NOT NULL,
    created_ts      TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY(timeline_post_id),
    CONSTRAINT fk_timeline_post_id
        FOREIGN KEY(timeline_post_id)
            REFERENCES rewild.timeline_posts(timeline_post_id),
    CONSTRAINT fk_image_id
        FOREIGN KEY(image_id)
            REFERENCES rewild.images(image_id)
);

CREATE TABLE IF NOT EXISTS rewild.timelines(
    timeline_id     UUID DEFAULT gen_random_uuid(),
    head_id         UUID NULL,
    tail_id         UUID NULL,
    created_ts      TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY(timeline_id),
    CONSTRAINT fk_head
        FOREIGN KEY(head_id)
            REFERENCES rewild.timeline_posts(timeline_post_id),
    CONSTRAINT fk_tail
        FOREIGN KEY(tail_id)
            REFERENCES rewild.timeline_posts(timeline_post_id)
);

CREATE TABLE IF NOT EXISTS rewild.discussion_board_messages(
    discussion_board_message_id     UUID DEFAULT gen_random_uuid(),
    parent_id                       UUID NULL,
    body                            TEXT,
    author_id                       UUID NULL,
    created_ts      TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY(discussion_board_message_id),
    CONSTRAINT fk_parent
        FOREIGN KEY(parent_id)
            REFERENCES rewild.discussion_board_messages(discussion_board_message_id),
    CONSTRAINT fk_author
        FOREIGN KEY(author_id)
            REFERENCES rewild.users(user_id)
);

CREATE TABLE IF NOT EXISTS rewild.discussion_boards(
    discussion_board_id UUID DEFAULT gen_random_uuid(),
    root_id             UUID NOT NULL,
    created_ts          TIMESTAMP DEFAULT current_timestamp,
    PRIMARY KEY(discussion_board_id),
    CONSTRAINT fk_discussion_board
       FOREIGN KEY(root_id)
           REFERENCES rewild.discussion_board_messages(discussion_board_message_id)
);

CREATE TABLE IF NOT EXISTS rewild.projects(
      project_id              UUID DEFAULT gen_random_uuid(),
      pindrop_id              UUID NOT NULL,
      name                    VARCHAR(255) NOT NULL,
      description             TEXT,
      timeline_id             UUID NOT NULL,
      discussion_board_id     UUID NOT NULL,
      follower_count          INTEGER NOT NULL,
      created_ts      TIMESTAMP DEFAULT current_timestamp,
      PRIMARY KEY(project_id),
      CONSTRAINT fk_pindrop
          FOREIGN KEY(pindrop_id)
              REFERENCES rewild.pindrops(pindrop_id),
      CONSTRAINT fk_timeline
          FOREIGN KEY(timeline_id)
              REFERENCES rewild.timelines(timeline_id),
      CONSTRAINT fk_discussion_board
          FOREIGN KEY(discussion_board_id)
              REFERENCES rewild.discussion_boards(discussion_board_id)
);

CREATE TABLE IF NOT EXISTS rewild.follows(
     user_id     UUID NOT NULL,
     project_id  UUID NOT NULL,
     created_ts      TIMESTAMP DEFAULT current_timestamp,
     PRIMARY KEY(user_id, project_id),
     CONSTRAINT fk_user
         FOREIGN KEY(user_id)
             REFERENCES rewild.users(user_id),
     CONSTRAINT fk_project
         FOREIGN KEY(project_id)
             REFERENCES rewild.projects(project_id)
);

CREATE TABLE IF NOT EXISTS rewild.auths(
     user_id     UUID NOT NULL,
     password    TEXT,
     PRIMARY KEY(user_id),
     CONSTRAINT fk_user
         FOREIGN KEY(user_id)
             REFERENCES rewild.users(user_id)
);