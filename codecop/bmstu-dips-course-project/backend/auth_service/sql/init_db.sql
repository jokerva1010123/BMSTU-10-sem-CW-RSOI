--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    uuid uuid NOT NULL,
    login character varying(255) NOT NULL,
    password bytea NOT NULL,
    lastname character varying(255) NOT NULL,
    firstname character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(255) NOT NULL,
    role character varying(50) NOT NULL,
    CONSTRAINT user_role_check CHECK (((role)::text = ANY ((ARRAY['USER'::character varying, 'MODERATOR'::character varying, 'ADMIN'::character varying])::text[])))
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, uuid, login, password, lastname, firstname, email, phone, role) FROM stdin;
1	691ccfe5-6e97-482f-91d8-2cafacf700f6	kovkir	\\x2432622431322458666b4d62557266724a354d564c556831305057564f413970484e5a36474c4a474370436c4a6c39526678756551303730776f2f57	Ковалец	Кирилл	kovaletskir@example.com	+79001004050	ADMIN
2	dd5e398c-2472-46e8-9d82-c5f2165f4283	user	\\x2432622431322438442e4d30756d5630554e53504f33672e754e45394f496138642f78325738636d5470556570454c726637484139626b5651717271	Иванов	Иван	user@example.com	+79002005060	USER
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 2, true);


--
-- Name: user user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: user user_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_login_key UNIQUE (login);


--
-- Name: user user_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_phone_key UNIQUE (phone);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: user user_uuid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_uuid_key UNIQUE (uuid);


--
-- Name: ix_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_user_id ON public."user" USING btree (id);


--
-- PostgreSQL database dump complete
--

