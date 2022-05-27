CREATE TABLE IF NOT EXISTS tweets(
    id serial primary key,
    tweet_text varchar(300) not null,
    posted_by_id bigint unsigned not null,
    publication_date datetime not null default CURRENT_TIMESTAMP,
    constraint fk_tweetpublisher foreign key(posted_by_id) references users(id)
    );