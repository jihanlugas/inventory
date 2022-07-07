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


CREATE TABLE public.photo
(
    photo_id     int8           NOT NULL DEFAULT next_id(),
    client_name  varchar(200)   NOT NULL,
    server_name  varchar(200)   NOT NULL,
    ext          varchar(5)     NOT NULL,
    photo_path   varchar(200)   NOT NULL,
    photo_size   int8           NOT NULL,
    photo_width  int8           NOT NULL,
    photo_height int8           NOT NULL,
    create_by    int8           NOT NULL,
    create_dt    timestamptz(0) NOT NULL,
    CONSTRAINT photo_pk PRIMARY KEY (photo_id)
);

CREATE TABLE public.photoinc
(
    photoinc_id int8        NOT NULL DEFAULT next_id(),
    ref_table   varchar(50) NOT NULL,
    folder_inc  int8        NOT NULL,
    folder_name varchar(50) NOT NULL,
    running     int8        NOT NULL,
    CONSTRAINT photoinc_pk PRIMARY KEY (photoinc_id)
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
    item_id          int8           NOT NULL DEFAULT next_id(),
    property_id      int8           NOT NULL,
    item_name        varchar(200)   NOT NULL,
    item_description varchar(500)   NOT NULL,
    photo_id         int8           NOT NULL,
    is_active        bool           NOT NULL,
    create_by        int8           NOT NULL,
    create_dt        timestamptz(0) NOT NULL,
    update_by        int8           NOT NULL,
    update_dt        timestamptz(0) NOT NULL,
    delete_by        int8           NULL,
    delete_dt        timestamptz(0) NULL,
    CONSTRAINT item_pk PRIMARY KEY (item_id)
);

CREATE TABLE public.itemvariant
(
    itemvariant_id          int8           NOT NULL DEFAULT next_id(),
    item_id                 int8           NOT NULL,
    itemvariant_name        varchar(200)   NOT NULL,
    itemvariant_description varchar(500)   NOT NULL,
    price                   int8           NOT NULL,
    photo_id                int8           NOT NULL,
    is_active               bool           NOT NULL,
    create_by               int8           NOT NULL,
    create_dt               timestamptz(0) NOT NULL,
    update_by               int8           NOT NULL,
    update_dt               timestamptz(0) NOT NULL,
    delete_by               int8           NULL,
    delete_dt               timestamptz(0) NULL,
    CONSTRAINT itemvariant_pk PRIMARY KEY (itemvariant_id)
);

CREATE TABLE public.userproperty
(
    userproperty_id int8 NOT NULL DEFAULT next_id(),
    user_id         int8 NOT NULL,
    property_id     int8 NOT NULL,
    is_default      bool NOT NULL,
    CONSTRAINT userproperty_pk PRIMARY KEY (userproperty_id)
);