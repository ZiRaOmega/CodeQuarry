--
-- PostgreSQL database dump
--

-- Dumped from database version 13.15 (Debian 13.15-1.pgdg120+1)
-- Dumped by pg_dump version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)

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
-- Name: verifyemail; Type: TABLE; Schema: public; Owner: codequarry
--

CREATE TABLE public.verifyemail (
    id integer NOT NULL,
    email character varying(50) NOT NULL,
    token character varying(50) NOT NULL,
    validated boolean DEFAULT false
);


ALTER TABLE public.verifyemail OWNER TO codequarry;

--
-- Name: verifyemail_id_seq; Type: SEQUENCE; Schema: public; Owner: codequarry
--

CREATE SEQUENCE public.verifyemail_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.verifyemail_id_seq OWNER TO codequarry;

--
-- Name: verifyemail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: codequarry
--

ALTER SEQUENCE public.verifyemail_id_seq OWNED BY public.verifyemail.id;


--
-- Name: verifyemail id; Type: DEFAULT; Schema: public; Owner: codequarry
--

ALTER TABLE ONLY public.verifyemail ALTER COLUMN id SET DEFAULT nextval('public.verifyemail_id_seq'::regclass);


--
-- Data for Name: verifyemail; Type: TABLE DATA; Schema: public; Owner: codequarry
--

COPY public.verifyemail (id, email, token, validated) FROM stdin;
2	evan.duvivier@icloud.com	YfwIPGS_hmr2f5t1DNNOKMw6iFKbzHM59GaRQ6uMjek=	t
3	polluxagency00@gmail.com	u_5hJ8ttbNM7XFoUPh84qh9uAbH26jZfMk0W2rJwn-Y=	t
4		mHbXRiykc1i9HjkZxQpdXxw4hvGAdNW2VrEabZ4jQ0I=	f
5	svalbardmalfant@gmail.com	_ZEvkeJmn5QA5kJ8uXDu8Zt0sJ0rv2nDTGqWny-3wL4=	t
7	vivien.frebourg@zone01normandie.org	Boj_ZFPsaRju-AgwQRGx_VxNPZVkBX5bIxQcl0BK09M=	f
6	bastien.lagrue@zone01normandie.org	sT2uBaQYaC93O4LYNhwrTBEyIboGTpwBJExuNcp2vMw=	t
8	santoslukombo@gmail.com	yRA_TvlyaQNw895kgm0uXvQNZrmcTUs47RWR4NFV0m8=	f
9	misterzomegaz@gmail.com	1hLo46yLj5DH3iUJ7Ms2mjeO7imw0JihDqhag9LLCzA=	t
\.


--
-- Name: verifyemail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: codequarry
--

SELECT pg_catalog.setval('public.verifyemail_id_seq', 9, true);


--
-- Name: verifyemail verifyemail_email_key; Type: CONSTRAINT; Schema: public; Owner: codequarry
--

ALTER TABLE ONLY public.verifyemail
    ADD CONSTRAINT verifyemail_email_key UNIQUE (email);


--
-- Name: verifyemail verifyemail_pkey; Type: CONSTRAINT; Schema: public; Owner: codequarry
--

ALTER TABLE ONLY public.verifyemail
    ADD CONSTRAINT verifyemail_pkey PRIMARY KEY (id);


--
-- Name: verifyemail verifyemail_token_key; Type: CONSTRAINT; Schema: public; Owner: codequarry
--

ALTER TABLE ONLY public.verifyemail
    ADD CONSTRAINT verifyemail_token_key UNIQUE (token);


--
-- PostgreSQL database dump complete
--

