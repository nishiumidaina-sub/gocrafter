
CREATE TABLE example_sub (

	id BIGINT NOT NULL AUTO_INCREMENT,

	user_id BIGINT NOT NULL ,

	example_id BIGINT NOT NULL ,

	name VARCHAR(255) NOT NULL ,

	detail VARCHAR(255) DEFAULT NULL ,

	count INT NOT NULL ,

	created_at TIMESTAMP NOT NULL ,

	updated_at TIMESTAMP NOT NULL ,

	PRIMARY KEY(id),
	INDEX(user_id),INDEX(example_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
