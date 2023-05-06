--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

-- Started on 2023-05-03 12:46:52

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
-- TOC entry 214 (class 1259 OID 16404)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id SERIAL NOT NULL,
    customer_name character varying(255) NOT NULL,
    order_date date NOT NULL,
    product character varying(255) NOT NULL,
    quantity integer NOT NULL,
    price numeric NOT NULL
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 3315 (class 0 OID 16404)
-- Dependencies: 214
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, customer_name, order_date, product, quantity, price) FROM stdin;
\.


--
-- TOC entry 3172 (class 2606 OID 16410)
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


-- Completed on 2023-05-03 12:46:52

--
-- PostgreSQL database dump complete
--

