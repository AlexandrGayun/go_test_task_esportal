CREATE TABLE IF NOT EXISTS followers(
     id serial primary key,
     follower_id bigint unsigned not null,
     followed_to_id bigint unsigned not null,
     constraint fk_follower foreign key(follower_id) references users(id),
     constraint fk_followedto foreign key(followed_to_id) references users(id)
);
