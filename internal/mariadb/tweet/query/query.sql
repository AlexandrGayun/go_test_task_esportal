/* name: GetTweetsOfFollowed :many */
select username, first_name, last_name, tweet_text, publication_date
from followers join tweets
    on followers.followed_to_id = tweets.posted_by_id
join users
    on tweets.posted_by_id = users.id
where follower_id = ?
order by publication_date DESC
limit ? offset ?;
