CREATE TABLE schools (
    id integer NOT NULL,
    uuid uuid NOT NULL,
    school_name character varying(255),
    register_code character varying(255),
    street character varying(255),
    town character varying(255),
    postcode character varying(255),
    is_active integer,
    admin_id uuid NOT NULL,
    config uuid NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);

CREATE TABLE configs (
    id integer NOT NULL,
    uuid uuid NOT NULL,
    start_minute integer,
    lesson_length integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
