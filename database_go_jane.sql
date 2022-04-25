--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)

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
-- Name: table_name_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.table_name_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.table_name_id_seq OWNER TO postgres;

--
-- Name: table_name_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.table_name_id_seq OWNED BY public.table_name.id;


--
-- Name: tbl_ad; Type: TABLE; Schema: public; Owner: ad_admin
--

CREATE TABLE public.tbl_ad (
    ad_id integer NOT NULL,
    url_id character varying(45),
    campaign_id integer
);


ALTER TABLE public.tbl_ad OWNER TO ad_admin;

--
-- Name: tbl_ad_ad_id_seq; Type: SEQUENCE; Schema: public; Owner: ad_admin
--

CREATE SEQUENCE public.tbl_ad_ad_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_ad_ad_id_seq OWNER TO ad_admin;

--
-- Name: tbl_ad_ad_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ad_admin
--

ALTER SEQUENCE public.tbl_ad_ad_id_seq OWNED BY public.tbl_ad.ad_id;


--
-- Name: tbl_campaign; Type: TABLE; Schema: public; Owner: ad_admin
--

CREATE TABLE public.tbl_campaign (
    campaign_id integer NOT NULL,
    start_campaign timestamp without time zone,
    end_campaign timestamp without time zone,
    max_impressions integer,
    cpm integer,
    impressions integer DEFAULT 0,
    CONSTRAINT tbl_campaign_check CHECK ((end_campaign > start_campaign))
);


ALTER TABLE public.tbl_campaign OWNER TO ad_admin;

--
-- Name: tbl_campaign_campaign_id_seq; Type: SEQUENCE; Schema: public; Owner: ad_admin
--

CREATE SEQUENCE public.tbl_campaign_campaign_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_campaign_campaign_id_seq OWNER TO ad_admin;

--
-- Name: tbl_campaign_campaign_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ad_admin
--

ALTER SEQUENCE public.tbl_campaign_campaign_id_seq OWNED BY public.tbl_campaign.campaign_id;


--
-- Name: tbl_keywords; Type: TABLE; Schema: public; Owner: ad_admin
--

CREATE TABLE public.tbl_keywords (
    campaign_id integer NOT NULL,
    keyword character varying(45) NOT NULL
);


ALTER TABLE public.tbl_keywords OWNER TO ad_admin;

--
-- Name: table_name id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.table_name ALTER COLUMN id SET DEFAULT nextval('public.table_name_id_seq'::regclass);


--
-- Name: tbl_ad ad_id; Type: DEFAULT; Schema: public; Owner: ad_admin
--

ALTER TABLE ONLY public.tbl_ad ALTER COLUMN ad_id SET DEFAULT nextval('public.tbl_ad_ad_id_seq'::regclass);


--
-- Name: tbl_campaign campaign_id; Type: DEFAULT; Schema: public; Owner: ad_admin
--

ALTER TABLE ONLY public.tbl_campaign ALTER COLUMN campaign_id SET DEFAULT nextval('public.tbl_campaign_campaign_id_seq'::regclass);


--
-- Data for Name: table_name; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.table_name (id) FROM stdin;
\.


--
-- Data for Name: tbl_ad; Type: TABLE DATA; Schema: public; Owner: ad_admin
--

COPY public.tbl_ad (ad_id, url_id, campaign_id) FROM stdin;
\.


--
-- Data for Name: tbl_campaign; Type: TABLE DATA; Schema: public; Owner: ad_admin
--

COPY public.tbl_campaign (campaign_id, start_campaign, end_campaign, max_impressions, cpm, impressions) FROM stdin;
\.


--
-- Data for Name: tbl_keywords; Type: TABLE DATA; Schema: public; Owner: ad_admin
--

COPY public.tbl_keywords (campaign_id, keyword) FROM stdin;
\.


--
-- Name: table_name_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.table_name_id_seq', 1, false);


--
-- Name: tbl_ad_ad_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ad_admin
--

SELECT pg_catalog.setval('public.tbl_ad_ad_id_seq', 31, true);


--
-- Name: tbl_campaign_campaign_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ad_admin
--

SELECT pg_catalog.setval('public.tbl_campaign_campaign_id_seq', 51, true);


--
-- Name: tbl_ad tbl_ad_pkey; Type: CONSTRAINT; Schema: public; Owner: ad_admin
--

ALTER TABLE ONLY public.tbl_ad
    ADD CONSTRAINT tbl_ad_pkey PRIMARY KEY (ad_id);


--
-- Name: tbl_campaign tbl_campaign_pkey; Type: CONSTRAINT; Schema: public; Owner: ad_admin
--

ALTER TABLE ONLY public.tbl_campaign
    ADD CONSTRAINT tbl_campaign_pkey PRIMARY KEY (campaign_id);


--
-- Name: tbl_keywords tbl_keywords_pkey; Type: CONSTRAINT; Schema: public; Owner: ad_admin
--

ALTER TABLE ONLY public.tbl_keywords
    ADD CONSTRAINT tbl_keywords_pkey PRIMARY KEY (campaign_id, keyword);


--
-- PostgreSQL database dump complete
--

