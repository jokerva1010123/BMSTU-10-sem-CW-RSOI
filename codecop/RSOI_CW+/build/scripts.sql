CREATE TABLE public.auth
(
    "UserUUID" uuid NOT NULL,
    "Login" text COLLATE pg_catalog."default" NOT NULL,
    "Password" text COLLATE pg_catalog."default" NOT NULL,
    "Role" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT auth_pkey PRIMARY KEY ("UserUUID")
);

INSERT INTO public.auth(
    "UserUUID", "Login", "Password", "Role")
VALUES ('0f7b24d4-c2bb-4c46-8182-ac4748fd96d1', 'Ktlh', md5('password'), 'admin');

CREATE TABLE public.airport
(
    "AirportUUID" uuid NOT NULL,
    "AirportName" text COLLATE pg_catalog."default",
    "City" text COLLATE pg_catalog."default",
    "Description" text COLLATE pg_catalog."default",
    CONSTRAINT airport_pkey PRIMARY KEY ("AirportUUID")
);

INSERT INTO public.airport(
    "AirportUUID", "AirportName", "City", "Description")
VALUES ('28d63850-aada-43cc-8720-15f5269e8088', 'Airport Home', 'Home',
        'This is description just to show something');

CREATE TABLE public.bonus
(
    "UserUUID" uuid NOT NULL,
    "Balance" integer NOT NULL,
    CONSTRAINT bonus_pkey PRIMARY KEY ("UserUUID")
);

CREATE TABLE public.flight
(
    "From" uuid NOT NULL,
    "From_city" text COLLATE pg_catalog."default" NOT NULL,
    "To" uuid NOT NULL,
    "To_city" text COLLATE pg_catalog."default" NOT NULL,
    "Date" date,
    "FlightID" uuid NOT NULL,
    CONSTRAINT flight_pkey PRIMARY KEY ("FlightID")
);

INSERT INTO public.flight(
    "From", "From_city", "To", "To_city", "Date", "FlightID")
VALUES ('28d63850-aada-43cc-8720-15f5269e8088', 'Home',
        '18d63850-aada-43cc-8720-15f5269e8088', 'Not Home', current_date, '5259ae69-ae0e-4dde-a6ed-033027fae15a');

CREATE TABLE public.reports
(
    "ID" SERIAL NOT NULL,
    "UserUUID" uuid NOT NULL,
    "FlightUUID" uuid NOT NULL,
    "TicketUUID" uuid NOT NULL,
    CONSTRAINT reports_pkey PRIMARY KEY ("ID")
);

INSERT INTO public.reports( "UserUUID", "FlightUUID", "TicketUUID")
VALUES ('f933e898-3a35-48bf-9be4-a92b94dfdf38', 'f933e898-3a35-48bf-9be4-a92b94dfdf38', 'f933e898-3a35-48bf-9be4-a92b94dfdf38');

CREATE TABLE public.ticket
(
    "TicketUUID" uuid NOT NULL,
    "FlightUUID" uuid NOT NULL,
    "UserUUID" uuid NOT NULL,
    "Date" date,
    CONSTRAINT ticket_pkey PRIMARY KEY ("TicketUUID")
);
