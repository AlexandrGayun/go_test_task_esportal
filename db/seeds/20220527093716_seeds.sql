-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
insert into users(username, password, email, last_name, first_name, age) values("tom1", "12345", "tommail@123", "tom_fn", "tom_ln", 20);
insert into users(username, password, email, last_name, first_name, age) values("alex2_followed_by_tom", "12345", "alexmail@123", "alex_fn", "alex_ln", 30);
insert into users(username, password, email, last_name, first_name, age) values("john3_followed_by_tom", "12345", "johnmail@123", "john_fn", "john_ln", 40);
insert into users(username, password, email, last_name, first_name, age) values("vlad4_followed_by_tom", "12345", "vladmail@123l", "vlad_fn", "vlad_ln", 50);
insert into users(username, password, email, last_name, first_name, age) values("sergio_followed_by_alex", "12345", "sergiomail@123", "sergio_fn", "sergio_ln", 60);
insert into users(username, password, email, last_name, first_name, age) values("maria_non_followed", "12345", "mariamail@123", "maria_fn", "maria_ln", 70);
insert into users(username, password, email, last_name, first_name, age) values("ida_followed_by_tom", "12345", "idamail@123", "ida_fn", "ida_ln", 80);

insert into followers(follower_id, followed_to_id) values(1, 2);
insert into followers(follower_id, followed_to_id) values(1, 3);
insert into followers(follower_id, followed_to_id) values(1, 4);
insert into followers(follower_id, followed_to_id) values(2, 5);
insert into followers(follower_id, followed_to_id) values(1, 7);

insert into tweets(posted_by_id, tweet_text, publication_date) values (1, "tom tweet #1", "2020-12-11 19:00:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (1, "tom tweet #2", "2020-12-12 19:00:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (2, "alex tweet #1", "2020-12-12 18:45:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (2, "alex tweet #2", "2020-12-12 18:30:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (3, "john tweet #1", "2020-12-12 18:15:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (3, "john tweet #2", "2020-12-12 18:00:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (3, "john tweet #3", "2020-12-12 17:45:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (4, "vlad tweet #1", "2020-12-12 17:30:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (5, "sergio tweet #1", "2020-12-12 17:15:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (5, "sergio tweet #2", "2020-12-12 17:00:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (6, "maria tweet #1", "2020-12-12 16:45:00");
insert into tweets(posted_by_id, tweet_text, publication_date) values (6, "maria tweet #2", "2020-12-12 16:30:00");
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
