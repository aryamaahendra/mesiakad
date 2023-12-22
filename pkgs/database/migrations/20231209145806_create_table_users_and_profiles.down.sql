ALTER TABLE
    user_profiles DROP CONSTRAINT IF EXISTS fk_user_profile;

ALTER TABLE
    user_profiles DROP CONSTRAINT IF EXISTS fk_user_pembimbing;

DROP TABLE IF EXISTS user_profiles;

DROP TABLE IF EXISTS users;