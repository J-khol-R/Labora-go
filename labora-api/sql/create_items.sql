PGDMP                         {            labora_proyect_1    15.2    15.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16403    labora_proyect_1    DATABASE     �   CREATE DATABASE labora_proyect_1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Spanish_Colombia.1252';
     DROP DATABASE labora_proyect_1;
                postgres    false            �            1259    16404    items    TABLE     �   CREATE TABLE public.items (
    id integer NOT NULL,
    customer_name character varying(255) NOT NULL,
    order_date date NOT NULL,
    product character varying(255) NOT NULL,
    quantity integer NOT NULL,
    price numeric NOT NULL
);
    DROP TABLE public.items;
       public         heap    postgres    false            �          0    16404    items 
   TABLE DATA           X   COPY public.items (id, customer_name, order_date, product, quantity, price) FROM stdin;
    public          postgres    false    214   �       e           2606    16410    items items_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.items DROP CONSTRAINT items_pkey;
       public            postgres    false    214            �      x������ � �     