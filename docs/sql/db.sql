CREATE sequence next_id;

CREATE
    OR REPLACE FUNCTION public.next_id()
    RETURNS bigint
    LANGUAGE plpgsql
AS
$function$
DECLARE
    seq_id bigint;
BEGIN
    SELECT nextval('public.next_id')
    INTO seq_id;
    return seq_id;
END;
$function$
;

CREATE TYPE user_type AS ENUM (
    'ADMIN',
    'USER'
    );

CREATE TABLE public.user
(
    user_id       int8           NOT NULL DEFAULT next_id(),
    fullname      varchar(80)    NOT NULL,
    no_hp         varchar(20)    NOT NULL,
    email         varchar(200)   NOT NULL,
    user_type     user_type      NOT NULL,
    username      varchar(20)    NOT NULL,
    passwd        varchar(200)   NOT NULL,
    photo_id      int8           NOT NULL,
    is_active     bool           NOT NULL,
    last_login_dt timestamptz(0) NOT NULL,
    pass_version  int4           NOT NULL,
    create_by     int8           NOT NULL,
    create_dt     timestamptz(0) NOT NULL,
    update_by     int8           NOT NULL,
    update_dt     timestamptz(0) NOT NULL,
    delete_by     int8           NULL,
    delete_dt     timestamptz(0) NULL,
    CONSTRAINT user_pk PRIMARY KEY (user_id)
);

CREATE TABLE public.property
(
    property_id   int8           NOT NULL DEFAULT next_id(),
    property_name varchar(200)   NOT NULL,
    photo_id      int8           NOT NULL,
    is_active     bool           NOT NULL,
    create_by     int8           NOT NULL,
    create_dt     timestamptz(0) NOT NULL,
    update_by     int8           NOT NULL,
    update_dt     timestamptz(0) NOT NULL,
    delete_by     int8           NULL,
    delete_dt     timestamptz(0) NULL,
    CONSTRAINT property_pk PRIMARY KEY (property_id)
);

CREATE TABLE public.item
(
    item_id     int8           NOT NULL DEFAULT next_id(),
    property_id int8           NOT NULL,
    item_name   varchar(200)   NOT NULL,
    is_active   bool           NOT NULL,
    create_by   int8           NOT NULL,
    create_dt   timestamptz(0) NOT NULL,
    update_by   int8           NOT NULL,
    update_dt   timestamptz(0) NOT NULL,
    delete_by   int8           NULL,
    delete_dt   timestamptz(0) NULL,
    CONSTRAINT item_pk PRIMARY KEY (item_id)
);

CREATE TABLE public.userproperty
(
    userproperty_id int8 NOT NULL DEFAULT next_id(),
    user_id         int8 NOT NULL,
    property_id     int8 NOT NULL,
    CONSTRAINT userproperty_pk PRIMARY KEY (userproperty_id)
);