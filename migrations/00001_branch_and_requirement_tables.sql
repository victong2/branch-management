-- +goose Up
-- +goose StatementBegin
CREATE TABLE branches (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    parent_id INTEGER,
    FOREIGN KEY (parent_id) REFERENCES branches(id)
);

CREATE TABLE requirements (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    details TEXT
);

-- Create the branch_requirement junction table
CREATE TABLE branch_requirements (
    branch_id INT NOT NULL,
    requirement_id INT NOT NULL,
    PRIMARY KEY (branch_id, requirement_id),
    CONSTRAINT fk_branch
        FOREIGN KEY(branch_id)
        REFERENCES branches(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_requirement
        FOREIGN KEY(requirement_id)
        REFERENCES requirements(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE branch_requirements;
DROP TABLE branches;
DROP TABLE requirements;
-- +goose StatementEnd
