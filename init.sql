\connect "widgetdb";
CREATE SEQUENCE widget_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."widget" (
    "id" integer DEFAULT nextval('widget_id_seq') NOT NULL,
    "name" character varying(255),
    "category" character varying(255),
    "color" character varying(255),
    CONSTRAINT "widget_pkey" PRIMARY KEY ("id")
) WITH (oids = false);
